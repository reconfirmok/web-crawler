package main

import (
	"net/url"
	"reflect"
	"testing"
)

func TestGetURLsFromHTML(t *testing.T) {
	tests := []struct {
		name      string
		inputURL  string
		inputBody string
		expected  []string
	}{
		{
			name:      "URLs from absloute path",
			inputURL:  "https://crawler-test.com",
			inputBody: `<html><body><a href="https://crawler-test.com"><span>Boot.dev</span></a></body></html>`,
			expected:  []string{"https://crawler-test.com"},
		},
		{
			name:      "URLs from relative path",
			inputURL:  "https://crawler-test.com",
			inputBody: `<html><body><a href="/test/path"><span>Boot.dev</span></a></body></html>`,
			expected:  []string{"https://crawler-test.com/test/path"},
		},
		{
			name:      "no URLs",
			inputURL:  "https://crawler-test.com",
			inputBody: `<html><body><p>Test Title</p></body></html>`,
			expected:  nil,
		},
		{
			name:     "absloute and relative URLs",
			inputURL: "https://crawler-test.com",
			inputBody: `
<html>
	<body>
		<a href="/path/one">
			<span>Boot.dev</span>
		</a>
		<a href="https://other.com/path/one">
			<span>Boot.dev</span>
		</a>
	</body>
</html>
`,
			expected: []string{"https://crawler-test.com/path/one", "https://other.com/path/one"},
		},
		{
			name:     "bad HTML",
			inputURL: "https://crawler-test.com",
			inputBody: `
<html body>
	<a href="path/one">
		<span>Boot.dev</span>
	</a>
</html body>
`,
			expected: []string{"https://crawler-test.com/path/one"},
		},
		{
			name:     "invalid href URL",
			inputURL: "https://crawler-test.com",
			inputBody: `
<html>
	<body>
		<a href=":\\invalidURL">
			<span>Boot.dev</span>
		</a>
	</body>
</html>
`,
			expected: nil,
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			baseURL, err := url.Parse(tc.inputURL)
			if err != nil {
				t.Errorf("Test %v - '%s' FAILED: Error parsing input URL: %v", i, tc.name, err)
				return
			}

			actual, err := getURLsFromHTML(tc.inputBody, baseURL)
			if err != nil {
				t.Errorf("Test %v - '%s' FAILED: unexpected error: %v", i, tc.name, err)
				return
			}

			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Test %v - '%s' FAILED: expected %v, got %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}
