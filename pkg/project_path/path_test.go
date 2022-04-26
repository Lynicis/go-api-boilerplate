//go:build unit

package project_path

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Path(t *testing.T) {
	t.Run("should return base path of project", func(t *testing.T) {
		getBasePath := GetRootDirectory()

		assert.Regexp(t, "/go-rest-api-boilerplate$", getBasePath)
	})
}
