package reader

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"io/ioutil"
	"os"
	"reflect"
)

type Reader struct {
	configStruct interface{}
	settings     Settings
}

type Settings struct {
	ConfigFile string
	EnvPrefix  string
}

func New(configStruct interface{}, settings Settings) Reader {
	return Reader{
		configStruct: configStruct,
		settings:     settings,
	}
}

func (r Reader) Read() error {
	if r.settings.ConfigFile != "" {
		fmt.Printf("try to apply config from file %s...\n", r.settings.ConfigFile)
		if err := r.SetFromFile(); err != nil {
			return fmt.Errorf("can't apply config from file: %w", err)
		}
	}
	if r.settings.EnvPrefix != "" {
		fmt.Printf("try to apply config from environment...\n")
		if err := r.SetFromEnv(); err != nil {
			return fmt.Errorf("can't apply envvars to config:%w", err)
		}
	}
	return nil
}

// Method adds and replace config fields from file.
func (r Reader) SetFromFile() error {
	if r.settings.ConfigFile == "" {
		return nil
	}
	f, err := os.Open(r.settings.ConfigFile)
	if err != nil {
		return fmt.Errorf("can't open config file: %w", err)
	}
	defer f.Close()
	l, err := ioutil.ReadAll(f)
	if err != nil {
		return fmt.Errorf("can't read content of the config file : %w", err)
	}
	_, err = toml.Decode(string(l), r.configStruct)
	if err != nil {
		return fmt.Errorf("can't parse config file : %w", err)
	}
	return nil
}

// Method adds and replace config fields from env.
func (r Reader) SetFromEnv() error {
	if r.settings.EnvPrefix == "" {
		return nil
	}
	return getEnvVar(reflect.ValueOf(r.configStruct), reflect.TypeOf(r.configStruct), -1, r.settings.EnvPrefix)
}
