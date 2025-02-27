package upload_test

import (
	"oggcloudserver/src/user/testing_material"
	"testing"
)

func TestDBIntegrity(t *testing.T) {

	testing_material.TestDBIntegrity(t)
}

func TestDataHandling(t *testing.T) {
	testing_material.TestDataHandling(t)

}
