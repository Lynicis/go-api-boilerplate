//go:build unit

package path

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Path(t *testing.T) {
	t.Run("should return base path of project", func(t *testing.T) {
		getBasePath := GetProjectBasePath()

		assert.Regexp(t, "/go-rest-api-boilerplate$", getBasePath)
	})
}
