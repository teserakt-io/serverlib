// Copyright 2019 Teserakt AG
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

package config_test

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/teserakt-io/serverlib/config"
	"github.com/teserakt-io/serverlib/path"
)

func ExampleLoader_Load() {
	log.SetFlags(0)

	pathResolver, err := path.NewAppPathResolver(os.Args[0])
	if err != nil {
		log.Fatalf("failed to create path resolver: %v", err)
	}

	// generate a dummy config file
	if err := os.MkdirAll(pathResolver.ConfigDir(), 0700); err != nil {
		log.Fatalf("failed to create config dir: %v", err)
	}

	configData := `
url_config_override: http://fromFile
count: 12
`
	if err := ioutil.WriteFile(filepath.Join(pathResolver.ConfigDir(), "config.yaml"), []byte(configData), 0600); err != nil {
		log.Fatalf("failed to write config: %v", err)
	}

	// and set some env
	os.Setenv("URL_ENV_OVERRIDE", "http://127.0.0.1:1234")

	// "config" is the name of the file to be loaded.
	// viper will search in pathResolver.ConfigDir() for any "config.yml", "config.yaml", "config.json"...
	loader := config.NewViperLoader("config", pathResolver)

	var url1, url2, url3 string
	var count int

	fields := []config.ViperCfgField{
		config.ViperCfgField{&url1, "url_default", config.ViperString, "http://localhost:8080", "URL_DEFAULT"},
		config.ViperCfgField{&url2, "url_env_override", config.ViperString, "http://localhost:8080", "URL_ENV_OVERRIDE"},
		config.ViperCfgField{&url3, "url_config_override", config.ViperString, "http://localhost:8080", "URL_CONFIG_OVERRIDE"},
		config.ViperCfgField{&count, "count", config.ViperInt, 0, ""},
	}

	if err := loader.Load(fields); err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	log.Printf("Url (default): %s\nUrl (env override): %s\nUrl (config override): %s\nCount: %d\n", url1, url2, url3, count)
}
