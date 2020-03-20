// Copyright 2020 Teserakt AG
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package path

import (
	"path/filepath"
	"runtime"
	"testing"
)

func TestAppResolver(t *testing.T) {
	_, currentFilePath, _, _ := runtime.Caller(1)
	currentDirectory := filepath.Dir(currentFilePath)
	parentDirectory := filepath.Join(currentDirectory, "..")

	resolver, err := NewAppPathResolver(currentFilePath)
	if err != nil {
		t.Fatalf("Failed to create AppPathResolver: %v", err)
	}

	t.Run("ConfigFile returns expected path", func(t *testing.T) {
		path := resolver.ConfigFile("config.yml")
		expectedPath := filepath.Join(parentDirectory, ConfigDir, "config.yml")

		if path != expectedPath {
			t.Errorf("Expected path to be %s, got %s", expectedPath, path)
		}
	})

	t.Run("ConfigDir returns expected path", func(t *testing.T) {
		path := resolver.ConfigDir()
		expectedPath := filepath.Join(parentDirectory, ConfigDir)

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
			{path: "./../test", expectedPath: filepath.Join(parentDirectory, ConfigDir, "./../test")},
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
