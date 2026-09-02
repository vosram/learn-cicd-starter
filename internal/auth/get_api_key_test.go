package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := map[string]struct {
		includeHeader bool
		headerVal     string
		expectedErr   bool
		expectedVal   string
	}{
		"Valid API Key":       {includeHeader: true, headerVal: "ApiKey 4bba243dc0d5cc18dd9b059a104bc4cf", expectedErr: false, expectedVal: "4bba243dc0d5cc18dd9b059a104bc4ab"},
		"No Auth Header":      {includeHeader: false, headerVal: "", expectedErr: true, expectedVal: ""},
		"Invalid Auth Header": {includeHeader: true, headerVal: "ApiKey", expectedErr: true, expectedVal: ""},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			h := http.Header{}
			if tc.includeHeader {
				h.Add("Authorization", tc.headerVal)
			}
			apiKey, err := GetAPIKey(h)
			if tc.expectedErr && err == nil {
				t.Fatal("Expected Error but no error occurred")
			}
			if apiKey != tc.expectedVal {
				t.Fatalf("Expected %v, but got %s", tc.expectedVal, apiKey)
			}
		})
	}
}
