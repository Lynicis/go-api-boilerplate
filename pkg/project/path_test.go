//go:build unit

package project

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Path(t *testing.T) {
	t.Run("should return base path of project", func(t *testing.T) {
		rootDirectory := GetRootDirectory()

		assert.Regexp(t, "/go-api-boilerplate$", rootDirectory)
	})
}
