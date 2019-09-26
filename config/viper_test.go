package config

import (
	"path/filepath"
	"reflect"
	"runtime"
	"testing"

	"github.com/teserakt-io/serverlib/path"
)

type testResolver struct {
	configDir string
}

var _ path.ConfigDirResolver = &testResolver{}

func (t *testResolver) ConfigDir() string {
	return t.configDir
}

func (t *testResolver) ConfigRelativePath(relPath string) string {
	return filepath.Join(t.configDir, relPath)
}

// getRootDir retrieve project root directory path from current test file
func getRootDir() string {
	_, filename, _, _ := runtime.Caller(0)

	return filepath.Join(filepath.Dir(filename), "..")
}

func TestViper(t *testing.T) {
	resolver := &testResolver{
		configDir: filepath.Join(getRootDir(), filepath.Join("test", "data")),
	}

	loader := NewViperLoader("_viper.config", resolver)

	type testConfig struct {
		TestString                    string
		TestInt                       int
		TestStringSlice               []string
		TestBool                      bool
		TestDbTypePostgress           DBType
		TestDbTypeSQLite              DBType
		TestDBSecureCnxTypeEnabled    DBSecureConnectionType
		TestDBSecureCnxTypeSelfSigned DBSecureConnectionType
		TestDBSecureCnxTypeInsecure   DBSecureConnectionType
		TestViperPath                 string
	}

	var cfg testConfig

	fields := []ViperCfgField{
		ViperCfgField{&cfg.TestString, "test-string", ViperString, "", ""},
		ViperCfgField{&cfg.TestInt, "test-int", ViperInt, 0, ""},
		ViperCfgField{&cfg.TestStringSlice, "test-stringslice", ViperStringSlice, []string{}, ""},
		ViperCfgField{&cfg.TestBool, "test-bool", ViperBool, false, ""},
		ViperCfgField{&cfg.TestDbTypePostgress, "test-dbtype-postgres", ViperDBType, DBTypeEmpty, ""},
		ViperCfgField{&cfg.TestDbTypeSQLite, "test-dbtype-sqlite3", ViperDBType, DBTypeEmpty, ""},
		ViperCfgField{
			&cfg.TestDBSecureCnxTypeEnabled,
			"test-dbsecurecnxtype-enabled",
			ViperDBSecureConnection,
			DBSecureConnectionEmpty,
			"",
		},
		ViperCfgField{
			&cfg.TestDBSecureCnxTypeSelfSigned,
			"test-dbsecurecnxtype-selfsigned",
			ViperDBSecureConnection,
			DBSecureConnectionEmpty,
			"",
		},
		ViperCfgField{
			&cfg.TestDBSecureCnxTypeInsecure,
			"test-dbsecurecnxtype-insecure",
			ViperDBSecureConnection,
			DBSecureConnectionEmpty,
			"",
		},
		ViperCfgField{&cfg.TestViperPath, "test-path", ViperRelativePath, "", ""},
	}

	if err := loader.Load(fields); err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	expectedCfg := testConfig{
		TestString:                    "str",
		TestInt:                       1,
		TestStringSlice:               []string{"str1", "str2"},
		TestBool:                      true,
		TestDbTypePostgress:           DBTypePostgres,
		TestDbTypeSQLite:              DBTypeSQLite,
		TestDBSecureCnxTypeEnabled:    DBSecureConnectionEnabled,
		TestDBSecureCnxTypeInsecure:   DBSecureConnectionInsecure,
		TestDBSecureCnxTypeSelfSigned: DBSecureConnectionSelfSigned,
		TestViperPath:                 resolver.ConfigRelativePath("../test/path"),
	}

	if reflect.DeepEqual(cfg, expectedCfg) == false {
		t.Errorf("Expected config to be %#v, got %#v", expectedCfg, cfg)
	}
}
