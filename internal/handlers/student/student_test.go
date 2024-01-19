package env

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"net/http"
    "net/http/httptest"
    "github.com/golang/mock/gomock"
    "github.com/google/uuid"
)

func TestGetEnv(t *testing.T) {
	t.Run("Should Return environment default", func(t *testing.T) {
		defaultValue := "GOLANG"
		environment := "PROGRAM"
		assert.Equal(t, GetEnv(environment, defaultValue), defaultValue)
	})

	t.Run("Should Return environment default", func(t *testing.T) {
		defaultValue := ""
		environment := "HOME"
		assert.NotEmpty(t, GetEnv(environment, defaultValue))
	})
}