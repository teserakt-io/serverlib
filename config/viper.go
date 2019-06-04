package config

import (
	"fmt"

	"github.com/spf13/viper"

	"gitlab.com/teserakt/serverlib/path"
)

// Loader defines a service able to load configuration
type Loader interface {
	Load([]ViperCfgField) error
}

// viperConfigLoader implements config.Loader
type viperConfigLoader struct {
	v              *viper.Viper
	configResolver path.ConfigDirResolver
}

// NewViperLoader creates a new configuration loader using Viper
// It will attempt to load file identified by configName (without extension)
// in pathResolver.ConfigDir()
//
// Example:
//
//    import (
//        "fmt"
//        "gitlab.com/teserakt/serverlib/path"
//     )
//
//     loader := NewViperLoader("config", path.NewAppPathResolver())
//
//     var url string
//     var count int
//
//     fields := []ViperCfgField{
// 	        ViperCfgField{&url, "url", ViperString, "http://localhost:8080", "URL"},
// 	        ViperCfgField{&count, "count", ViperInt, 0, ""},
//     }
//
//     if err := loader.Load(fields); err != nil {
//	        fmt.Fatalf("Failed to load configuration: %v", err)
//     }
//
//     fmt.Printf("Url: %s, count: %d\n", url, count)
//
func NewViperLoader(configName string, configResolver path.ConfigDirResolver) Loader {
	v := viper.New()
	v.SetConfigName(configName)
	v.AddConfigPath(configResolver.ConfigDir())

	return &viperConfigLoader{
		v:              v,
		configResolver: configResolver,
	}
}

// ViperType allow to instruct viper how to cast the loaded values
type ViperType int

const (
	// ViperInt defines a viper type for an int
	ViperInt ViperType = iota
	// ViperString defines a viper type for a string
	ViperString
	// ViperStringSlice defines a viper type for a []string
	ViperStringSlice
	// ViperBool defines a viper type for a bool
	ViperBool
	// ViperDBType defines a viper type for a DBType
	ViperDBType
	// ViperDBSecureConnection defines a viper type for a DBSecureConnectionType
	ViperDBSecureConnection
)

// ViperCfgField defines a struct to instruct viper what and how to load configuration data.
type ViperCfgField struct {
	// Target must be a pointer to a variable which will hold the loaded value
	Target interface{}
	// KeyName is the name which must be found in the configuration file
	KeyName string
	// CfgType must be one of the ViperType, telling viper how to cast the value
	CfgType ViperType
	// DefaultValue is the value to be set on the Target when it can't be found in the configuration file
	DefaultValue interface{}
	// EnvMapping is the name of the environment variable to look for, which will replace any defined value in the configuration file
	EnvMapping string
}

// Load configure viper and read the configuration, attempting to populate the Target of every given ViperCfgFields.
// For each given fields, tt will first try to read it from a configuration file,
// then fallback to env variable if provided, at last use the default value when nothing else matched.
func (loader *viperConfigLoader) Load(fields []ViperCfgField) error {
	for _, field := range fields {
		loader.v.SetDefault(field.KeyName, field.DefaultValue)

		if field.EnvMapping != "" {
			loader.v.BindEnv(field.KeyName, field.EnvMapping)
		}
	}

	if err := loader.v.ReadInConfig(); err != nil {
		return err
	}

	for _, field := range fields {
		switch field.CfgType {
		case ViperInt:
			v := field.Target.(*int)
			*v = loader.v.GetInt(field.KeyName)
		case ViperString:
			v := field.Target.(*string)
			*v = loader.v.GetString(field.KeyName)
		case ViperStringSlice:
			v := field.Target.(*[]string)
			*v = loader.v.GetStringSlice(field.KeyName)
		case ViperBool:
			v := field.Target.(*bool)
			value := loader.v.GetBool(field.KeyName)
			*v = value
		case ViperDBType:
			v := field.Target.(*DBType)
			*v = DBType(loader.v.GetString(field.KeyName))
		case ViperDBSecureConnection:
			v := field.Target.(*DBSecureConnectionType)
			*v = DBSecureConnectionType(loader.v.GetString(field.KeyName))
		default:
			return fmt.Errorf("unsupported configuration type %v for field %v", field.CfgType, field.KeyName)
		}
	}

	return nil
}
