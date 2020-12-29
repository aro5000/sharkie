package main

import (
	"testing"
)

func TestCompare (t *testing.T) {
	// Test out of bounds number
	expectedValues := []int{200,300,400,500}
	value := 202


	// This should not be successful, because 200 should be specified, not 202
	ok := compare(expectedValues, value)

	if ok {
		t.Errorf("%d should not match values in %v", value, expectedValues)
	}

}


func TestCleanup (t *testing.T) {
	// Test to make sure values get reset
	TDATA.Url = "example.com"
	TDATA.Expected = 200

	cleanup()

	if TDATA.Url != "" {
		t.Errorf("TDATA.Url should be empty but is set to: %v", TDATA.Url)
	}

	if TDATA.Expected != 0 {
		t.Errorf("TDATA.Expected should be 0, but is set to: %v", TDATA.Expected)
	}
}


func TestParse (t *testing.T) {
	testCases := []struct {
		Name           string
		Servers        []string
		Url            string
		ExpectedPort   string
		ExpectedHost   string
		ExpectedPath   string
	} {
		{"localhost port parsing", []string{}, "http://localhost:8080", "8080", "localhost", "/"},
		{"http check", []string{"127.0.0.1","127.0.0.2"}, "http://example.com", "80", "example.com", "/"},
		{"https check", []string{"127.0.0.1","127.0.0.2"}, "https://example.com", "443", "example.com", "/"},
		{"path check", []string{"127.0.0.1","127.0.0.2"}, "https://example.com/example/path", "443", "example.com", "/example/path"},
		{"default host check", []string{}, "127.0.0.1", "80", "127.0.0.1", "/"},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			TDATA.Url = tc.Url
			tc.Servers = parse(tc.Servers)

			if tc.ExpectedPort != TDATA.Port {
				t.Errorf("TDATA.Port is: %v Expected: %v", TDATA.Port, tc.ExpectedPort)
			}

			if tc.ExpectedHost != TDATA.Host {
				t.Errorf("TDATA.Host is: %v Expected: %v", TDATA.Host, tc.ExpectedHost)
			}
			if tc.ExpectedPath != TDATA.Path {
				t.Errorf("TDATA.Path is: %v Expected: %v", TDATA.Path, tc.ExpectedPath)
			}
		})
	}

}