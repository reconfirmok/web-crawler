package main

import "testing"

func TestGetHeadingFromHTML(t *testing.T) {
	tests := []struct {
		name      string
		inputBody string
		expected  string
	}{
		{
			name:      "Get h1 content",
			inputBody: "<html><body><h1>Test Title</h1></body></html>",
			expected:  "Test Title",
		},
		{
			name:      "Get h2 content",
			inputBody: "<html><body><h2>Fallback Title</h2></body></html>",
			expected:  "Fallback Title",
		},
		{
			name:      "No heading",
			inputBody: "<html><body><p></p></body></html>",
			expected:  "",
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := getHeadingFromHTML(tc.inputBody)
			if err != nil {
				t.Errorf("Test %v - '%s' FAILED: unexpected error: %v", i, tc.name, err)
				return
			}
			if actual != tc.expected {
				t.Errorf("Test %v - '%s' FAILED: expected %q, got %q", i, tc.name, tc.expected, actual)
			}
		})
	}
}

func TestGetFirstParagraphFromHTMLMainPriority(t *testing.T) {
	tests := []struct {
		name      string
		inputBody string
		expected  string
	}{
		{
			name:      "Get <p> from <main>",
			inputBody: `<html><body><p>Outside paragraph.</p><main><p>Main paragraph.</p></main></body></html>`,
			expected:  "Main paragraph.",
		},
		{
			name:      "Get first <p>",
			inputBody: "<html><body><p>First paragraph outside main.</p><p>Second paragraph outside main.</p></body></html>",
			expected:  "First paragraph outside main.",
		},
		{
			name:      "Empty <p>",
			inputBody: `<html><body></body></html>`,
			expected:  "",
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := getFirstParagraphFromHTML(tc.inputBody)
			if err != nil {
				t.Errorf("Test %v - '%s' FAILED: unexpected error: %v", i, tc.name, err)
				return
			}
			if actual != tc.expected {
				t.Errorf("Test %v - '%s' FAILED: expected %q, got %q", i, tc.name, tc.expected, actual)
			}
		})
	}
}
