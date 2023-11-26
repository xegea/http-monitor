package httprequest

import (
	"strings"
	"testing"
)

func Test_doRequest(t *testing.T) {
	t.Run("Invalid request", func(t *testing.T) {
		expected := "invalid request"
		_, err := Get("test@test.com")
		if !strings.Contains(err.Error(), expected) {
			t.Errorf("expected %s error but got %v", expected, err)
		}
	})
}
