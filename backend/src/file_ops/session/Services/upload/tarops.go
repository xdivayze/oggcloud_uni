package upload

import (
	"archive/tar"
	"bufio"
	"compress/gzip"
	"errors"
	"fmt"
	"io"
	"log"
	"oggcloudserver/src/db"
	"oggcloudserver/src/file_ops/file"
	"oggcloudserver/src/oggcrypto"
	"os"
	"path"
	"strings"
)

var PreviewExemptions = []string{} // I think it's cool that preview checksum is shown as storage checksum's preview

var ErrTooEarlyToBeAPreview = errors.New("no owner for preview to be associated with an owner")

func extractFile(tarReader *tar.Reader, header *tar.Header, previewMode bool, belongsTo *file.File) error {
	fw := strings.Split(header.Name, "/")
	fparts := strings.Split(fw[len(fw)-1], ".")

	outfilename := fmt.Sprintf("%s.%s", fparts[0], fparts[1])
	outFilePath := fmt.Sprintf("%s/%s", currentWorkingPath, outfilename)
	outFile, err := os.Create(outFilePath)
	if err != nil {
		return fmt.Errorf("error occured while creating file at path %s:\n\t%w", outFilePath, err)
	}
	defer outFile.Close()
	bufr := bufio.NewReader(tarReader)
	bufw := bufio.NewWriter(outFile)
	if _, err = io.Copy(bufw, bufr); err != nil {
		return fmt.Errorf("error occured while writing from reader to file:\n\t%w", err)
	}
	if err = bufw.Flush(); err != nil {
		return fmt.Errorf("error occured while flushing buffered writer:\n\t%w", err)
	}
	var fileObj file.File
	if res := db.DB.Where("file_name = ?", outfilename).Where("is_preview = ?", previewMode).Find(&fileObj); res.Error != nil {
		return fmt.Errorf("error occured while searching in database:\n\t%w", err)
	}
	size := header.FileInfo().Size()
	fileObj.Size = size
	if _, err = outFile.Seek(0, io.SeekStart); err != nil {
		return fmt.Errorf("error occured while seeking:\n\t%w", err)
	}
	sum, err := oggcrypto.CalculateSHA256sum(outFile)
	if err != nil {
		return fmt.Errorf("error occured while calculating checksum:\n\t%w", err)
	}
	if sum != *fileObj.Checksum {
		return fmt.Errorf("file checksum doesn't match") //TODO add a cleanup function
	}

	if previewMode {
		if belongsTo == nil {
			return fmt.Errorf("error: parent file object is nil: %w", ErrTooEarlyToBeAPreview) 
		}
		belongsTo.PreviewID = &(fileObj.ID)
		belongsTo.Preview = &fileObj
		belongsTo.HasPreview = true
		fileObj.IsPreview = true
		res := db.DB.Save(belongsTo)
		if res.Error != nil {
			return fmt.Errorf("error occured while updating the owner file instance:\n\t%w", res.Error)
		}

	}
	if res := db.DB.Save(&fileObj); res.Error != nil {
		return fmt.Errorf("error occured while saving to db:\n\t%w", res.Error)
	}
	return nil
}

var currentWorkingPath string

// TODO add concurrency back <3
func extractTarGz(r io.Reader, session *Session) error {

	f, err := setupTar(r)
	if err != nil {
		return err
	}
	defer f.Close()

	checksum, err := oggcrypto.CalculateSHA256sum(f)
	if err != nil {
		return fmt.Errorf("error occured while calculating sha256sum:\n\t%w", err)
	}
	if _, err = f.Seek(0, io.SeekStart); err != nil {
		return fmt.Errorf("error while seeking to start:\n\t%w", err)
	}

	if checksum != session.SessionChecksum {
		return fmt.Errorf("checksum no match")
	}

	if err := checkDirectoryValidity(f, session); err != nil { //ensure storage and preview directories exist and are valid
		return err
	}

	if _, err = f.Seek(0, io.SeekStart); err != nil {
		return fmt.Errorf("error while seeking to start:\n\t%w", err)
	}

	gzipReader, err := gzip.NewReader(f)
	if err != nil {
		return fmt.Errorf("error occurred while creating new gzip reader:\n\t%w", err)
	}
	defer gzipReader.Close()
	tarReader := tar.NewReader(gzipReader)
	for {
		header, err := tarReader.Next()
		if err == io.EOF {
			log.SetPrefix("INFO: ")
			log.Println("archive end reached")
			break
		} else if err != nil {
			return fmt.Errorf("error occured while reading the next entry in tar reader:\n\t%v", err)
		}
		cp := path.Clean(header.Name)
		lx := strings.Split(cp, "/")
		if cp == "." {
			continue
		}
		if header.Typeflag == tar.TypeDir && (lx[len(lx)-1] == STORAGE_DIR_NAME || lx[len(lx)-1] == PREVIEW_DIR_NAME) {
			currentWorkingPath = fmt.Sprintf("%s/%s", DirectorySession, cp)
			if err = os.MkdirAll(currentWorkingPath, 0777); err != nil {
				return fmt.Errorf("error occured creating path at %s :\n\t%w", currentWorkingPath, err)
			}
		}
		if header.Typeflag == tar.TypeReg && (lx[len(lx)-1] != CHECKSUM_FILENAME) {
			previewLoadingMode := false
			var owningFileRef *file.File
			owningFileRef = nil
			if lx[len(lx)-2] == PREVIEW_DIR_NAME { //see if preview directory is being exported
				previewLoadingMode = true
				var err error
				owningFileRef, err = session.FindOwnedFileWithName(lx[len(lx)-1])
				if err != nil {
					return fmt.Errorf("error occured while finding owned file with name %s:\n\t%w", lx[len(lx)-1], err)
				}
				if owningFileRef == nil {
					return fmt.Errorf("error occured while trying to find owner file")
				}

			}
			err = extractFile(tarReader, header, previewLoadingMode, owningFileRef)
			if err != nil {
				return fmt.Errorf("error occured while extracting file:\n\t%w", err)
			}
		}
	}
	return nil
}

func setupTar(r io.Reader) (*os.File, error) {
	buffer := make([]byte, 1024*4)
	f, err := os.CreateTemp("", "compressedtar*.tar.gz")
	if err != nil {
		return nil, fmt.Errorf("error creating temporary file at path %s:\n\t%w", f.Name(), err)
	}

	defer os.Remove(f.Name())

	if _, err = io.CopyBuffer(f, r, buffer); err != nil {
		return nil, fmt.Errorf("error occured while copying file to new buffer:\n\t%w", err)
	}

	if err = f.Sync(); err != nil {
		return nil, fmt.Errorf("error trying to sync file:\n\t%w", err)
	}

	f.Close()

	f, err = os.Open(f.Name())
	if err != nil {
		return nil, fmt.Errorf("error occured while opening file:\n\t%w", err)
	}
	return f, nil
}
