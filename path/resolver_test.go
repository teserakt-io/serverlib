package path

import (
	"path/filepath"
	"runtime"
	"testing"
)

func TestAppResolver(t *testing.T) {
	_, currentFilePath, _, _ := runtime.Caller(1)
	currentDirectory := filepath.Dir(currentFilePath)

	resolver, err := NewAppPathResolver(currentFilePath)
	if err != nil {
		t.Fatalf("Failed to create AppPathResolver: %v", err)
	}

	t.Run("ConfigFile returns expected path", func(t *testing.T) {
		path := resolver.ConfigFile("config.yml")
		expectedPath := filepath.Join(currentDirectory, ConfigDir, "config.yml")

		if path != expectedPath {
			t.Errorf("Expected path to be %s, got %s", expectedPath, path)
		}
	})

	t.Run("ConfigDir returns expected path", func(t *testing.T) {
		path := resolver.ConfigDir()
		expectedPath := filepath.Join(currentDirectory, ConfigDir)

		if path != expectedPath {
			t.Errorf("Expected path to be %s, got %s", expectedPath, path)
		}
	})

	t.Run("BinaryFile returns expected path", func(t *testing.T) {
		path := resolver.BinaryFile()

		if path != currentFilePath {
			t.Errorf("Expected path to be %s, got %s", currentFilePath, path)
		}
	})

	t.Run("ConfigRelativePath returns expected path", func(t *testing.T) {
		testCases := []struct {
			path         string
			expectedPath string
		}{
			{path: "./../test", expectedPath: filepath.Join(currentDirectory, ConfigDir, "./../test")},
			{path: "/test/something", expectedPath: "/test/something"},
		}

		for _, testCase := range testCases {
			path := resolver.ConfigRelativePath(testCase.path)

			if path != testCase.expectedPath {
				t.Errorf("Expected path to be %s, got %s", testCase.expectedPath, path)
			}
		}
	})
}
