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

// Package path provides basic helpers to resolve path from the executable location
//
// It answers some questions with regards to config files:
//
//  1. How do we locate the actual configuration file (independent of whether
// 	viper is used)?
//  2. How do we locate files specified as relative paths in the configuration file.
//
// The deployment layout should look like this:
//
//    /opt/e4/bin/binary
//    /opt/e4/configs/projectconf.yaml
//    /opt/e4/configs/sslcert.pem
//    /opt/e4/share/e4/projectconf.yaml.template
//
// and so on. /opt/e4 is an arbitrary (but reasonable) prefix that users may
// decide to change, e.g. to /usr/local. As such:
//
//  1. The config file should be located at ../configs/name.yaml relative to
// 	the binary.
//  2. Additional configuration data is likely located in the configs directory
// 	too. So relative paths for such loads should be relative to the config
// 	file location.
//
// In development, we have paths that look like this:
//
// 	GITREPO/bin/binary
// 	GITREPO/configs/projectconf.yaml
//
// etc. So this logic works both for development and for production scenarios.
package path

import (
	"path/filepath"
)

// ConfigDir defines the default directory name holding the configuration files
var ConfigDir = "configs"

// ConfigDirResolver defines method for resolving application configuration path
type ConfigDirResolver interface {
	ConfigDir() string
	ConfigRelativePath(relPath string) string
}

// AppPathResolver represents the state of an application path lookup for future use
type AppPathResolver struct {
	binarypath         string
	absolutePrefixPath string
}

// NewAppPathResolver returns a new instance of the AppPathResolver
// binaryPath is the path to the current executable, usually argv[0].
func NewAppPathResolver(binaryPath string) (*AppPathResolver, error) {
	dir, err := filepath.Abs(filepath.Dir(binaryPath))
	if err != nil {
		return nil, err
	}

	return &AppPathResolver{
		binarypath:         binaryPath,
		absolutePrefixPath: dir,
	}, nil
}

// ConfigFile returns the path to the config file, given confFilename as a config file argument
func (a *AppPathResolver) ConfigFile(confFilename string) string {
	return filepath.Join(a.ConfigDir(), confFilename)
}

// ConfigDir returns the path to the config file directory, given confFilename as a config file argument
func (a *AppPathResolver) ConfigDir() string {
	return filepath.Join(a.absolutePrefixPath, "..", ConfigDir)
}

// BinaryFile returns the path to the binary itself, in case this is ever useful
func (a *AppPathResolver) BinaryFile() string {
	return a.binarypath
}

// ConfigRelativePath resolves a relative filepath from the config file. If the filepath is
// absolute then it is returned unchanged. This is suitable to be called
// for all file resolutions
func (a *AppPathResolver) ConfigRelativePath(relPath string) string {
	if filepath.IsAbs(relPath) {
		return relPath
	}
	return filepath.Join(a.ConfigDir(), relPath)
}
