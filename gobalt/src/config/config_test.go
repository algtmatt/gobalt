package config

import (
	"gopkg.in/yaml.v2"
	"os"
	"reflect"
	"testing"
)

func getConfigFileLocation() (location string) {
	pwd, _ := os.Getwd()
	config_path := pwd + "/gobalt.yml"
	return config_path
}

func TestNew(t *testing.T) {
	c := New()
	cc := Config{}
	cType := reflect.TypeOf(c)
	if cType != reflect.TypeOf(cc) {
		t.Errorf("Constructor returned incorrect type")
	}
}

func TestConfig_GetConfigFile(t *testing.T) {
	c := Config{ConfigFileLocation: getConfigFileLocation()}
	contents := c.GetConfigFile()
	if len(contents) > 0 {
		// pass
	} else {
		t.Errorf("Zero-byte file retreived")
	}
}

func TestConfig_ParseConfigFile(t *testing.T) {
	c := Config{ConfigFileLocation: getConfigFileLocation()}
	contents := c.GetConfigFile()
	o := Opts{}
	y, err := o.ParseConfigBytes(contents)
	if y.Debug == false && err == nil {
		// pass
	} else {
		t.Errorf("Error in parsing configuration")
	}
}

func TestConfig_ParseConfigFile_MissingOptions(t *testing.T) {
	type partialConfig struct {
		Url string `yaml:"url"`
	}

	partialConfigBytes, _ := yaml.Marshal(&partialConfig{Url: "http://foo.com"})

	o := Opts{}
	_, err := o.ParseConfigBytes(partialConfigBytes)
	if err != nil {
		t.Errorf("Attempt to generate partial config failed")
	}

}

func TestConfig_ParseConfigFile_ExtraOptions(t *testing.T) {
	type partialConfig struct {
		ExtraOption string `yaml:"extra"`
		Url string `yaml:"url"`
	}
	extraConfigBytes, _ := yaml.Marshal(&partialConfig{ExtraOption:"wtfisthisdoinghere", Url:"http://foo.com"})
	o := Opts{}
	_, err := o.ParseConfigBytes(extraConfigBytes)
	if err != nil {
		t.Errorf("Extra option caused issue")
	}
}