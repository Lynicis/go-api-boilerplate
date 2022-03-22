//go:build unit

package directory

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Directory(t *testing.T) {
	t.Run("should create new directory instance and return directory instance", func(t *testing.T) {
		getDirectory, _ := os.Getwd()
		basePath := fmt.Sprintf("../../%s", getDirectory)

		testDirectoryInstance := NewDirectory(basePath)

		expectedDirectoryInstance := &directory{
			BasePath: basePath,
		}

		assert.Equal(t, expectedDirectoryInstance, testDirectoryInstance)
	})

	t.Run("should call base path method return base path from directory from directory instance", func(t *testing.T) {
		directoryInstance := createTestDirectoryInstance()

		getDirectory, _ := os.Getwd()
		expectedBasePathFromDirectoryInstance := fmt.Sprintf("../../%s", getDirectory)

		callProjectPathMethod := directoryInstance.GetProjectBasePath()
		assert.Equal(t, expectedBasePathFromDirectoryInstance, callProjectPathMethod)
	})
}

func createTestDirectoryInstance() Directory {
	getDirectory, _ := os.Getwd()
	basePath := fmt.Sprintf("../../%s", getDirectory)
	return NewDirectory(basePath)
}
