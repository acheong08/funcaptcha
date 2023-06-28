package funcaptcha

import (
	"strings"
	"testing"
)

func TestGetToken(t *testing.T) {
	token, _ := GetOpenAITokenV1()
	if !strings.Contains(token, "sup=") {
		t.Errorf("Token does not contain 'sup='")
	}
	if !strings.Contains(token, "rid=") {
		t.Errorf("Token does not contain 'rid='")
	}
}
