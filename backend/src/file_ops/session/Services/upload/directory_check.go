package upload

import (
	"archive/tar"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"oggcloudserver/src/db"
	"oggcloudserver/src/file_ops/file"
	"path"
	"strings"

	"github.com/google/uuid"
)

const STORAGE_DIR_NAME = "Storage"
const PREVIEW_DIR_NAME = "Preview"
const CHECKSUM_FILENAME = "checksum.json"

type directory struct {
	filenames           []string
	checksumFileContent *checksumFileStructure
	name                string
	path                string
}

func contains[K comparable](arr []K, val K) bool {
	for _, v := range arr {
		if v == val {
			return true
		}
	}
	return false
}

type checksumFileStructure struct {
	Preview bool
	Files   []struct {
		Checksum string
		Filetype string
		Filename string
	}
}

func createFileInstancesDB(data *checksumFileStructure, session *Session) error {
	for _, descriptor := range data.Files {

		lx := strings.Split(descriptor.Filename, ".")
		dbname := fmt.Sprintf("%s.%s", lx[0], lx[1])
		u := file.File{
			ID:         uuid.New(),
			FileName:   dbname,
			SessionID:  session.ID,
			UserID:     session.UserID,
			IsPreview:  data.Preview,
			FileType:   &descriptor.Filetype,
			Checksum:   &descriptor.Checksum,
			HasPreview: !data.Preview,
		}
		res := db.DB.Create(&u)
		if res.Error != nil {
			return fmt.Errorf("error occured while saving to db")
		}
	}
	return nil
}

func checkDirectoryValidity(r io.ReadSeeker, session *Session) error {
	dirs, err := determineDirectories(r)
	if err != nil {
		return fmt.Errorf("error occured while determining directories:\n\t%w", err)
	}
	fields := []string{STORAGE_DIR_NAME, PREVIEW_DIR_NAME}
	if err = doDirFieldCheck(dirs, fields); err != nil {
		return err
	}

	storage_dir := dirs[STORAGE_DIR_NAME]
	preview_dir := dirs[PREVIEW_DIR_NAME]

	if len(preview_dir.filenames) != len(storage_dir.filenames) {
		return fmt.Errorf("storage_dir and preview_dir have non-identical lengths")
	}
	if !contains(storage_dir.filenames, CHECKSUM_FILENAME) {
		return fmt.Errorf("storage directory doesn't contain checksum file")
	}
	if !contains(preview_dir.filenames, CHECKSUM_FILENAME) {
		return fmt.Errorf("preview directory doesn't contain checksum file")
	}

	if err = createFileInstancesDB(dirs[STORAGE_DIR_NAME].checksumFileContent, session); err != nil {
		return fmt.Errorf("error occured occured while creating db instances for %s\n\t%w", STORAGE_DIR_NAME, err)
	}
	if err = createFileInstancesDB(dirs[PREVIEW_DIR_NAME].checksumFileContent, session); err != nil {
		return fmt.Errorf("error occured occured while creating db instances for %s\n\t%w", PREVIEW_DIR_NAME, err)
	}

	return nil
}

func doDirFieldCheck(m map[string]*directory, fields []string) error {
	for _, field := range fields {
		if _, s := m[field]; !s {
			return fmt.Errorf("field %s doesn't exist", field)
		}
	}
	return nil
}

func determineDirectories(r io.ReadSeeker) (map[string]*directory, error) {
	gzipReader, err := gzip.NewReader(r)
	if err != nil {
		return nil, fmt.Errorf("error occured while creating new gzip reader:\n\t%w", err)
	}
	defer gzipReader.Close()
	tarReader := tar.NewReader(gzipReader)
	dirs := make(map[string]*directory)
	for {
		header, err := tarReader.Next()
		if err == io.EOF {
			log.SetPrefix("INFO: ")
			log.Println("archive end reached")
			break
		} else if err != nil {
			return nil, fmt.Errorf("error occured while reading the next entry in tar reader:\n\t%v", err)
		}
		dir := directory{}

		cleanPath := path.Clean(header.Name)
		if cleanPath == "." {
			continue
		}
		if header.Typeflag == tar.TypeDir {
			dir.name = cleanPath
			dir.path = header.Name
			dirs[cleanPath] = &dir
		} else if header.Typeflag == tar.TypeReg {
			fdir := path.Dir(cleanPath)
			lx := strings.Split(cleanPath, "/")
			if fdir != "." {
				dirobj, exists := func() (*directory, bool) {

					dirname := lx[len(lx)-2]
					dirobj, s := dirs[dirname]
					if dirobj.name != fdir {
						s = false
					}
					return dirobj, s
				}()
				if !exists {
					return nil, fmt.Errorf("orphaned file with path:\n\t%s", cleanPath)
				}
				if lx[len(lx)-1] == CHECKSUM_FILENAME {
					buf, err := io.ReadAll(tarReader)
					if err != nil {
						return nil, fmt.Errorf("error while reading from tar reader:\n\t%w", err)
					}
					var unmarshaledJson checksumFileStructure
					if err := json.Unmarshal(buf, &unmarshaledJson); err != nil {
						return nil, fmt.Errorf("error while unmarshaling json:\n\t%w", err)
					}
					dirobj.checksumFileContent = &unmarshaledJson
				}
				dirobj.filenames = append(dirobj.filenames, lx[len(lx)-1])
			}
		}
	}
	return dirs, nil

}
