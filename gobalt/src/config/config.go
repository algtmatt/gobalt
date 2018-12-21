package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	Opts
	ConfigFileLocation string
	config             map[string]string
}

type Opts struct {
	Url   string `yaml:"url"`
	Port  int    `yaml:"port"`
	Debug bool   `yaml:"debug"`

	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

func New() (config Config) {
	return Config{
		ConfigFileLocation: "/etc/gobalt.yml",
	}
}

func (c *Config) GenerateConfig() Opts {
	f := c.GetConfigFile()
	o, err := c.ParseConfigBytes(f)
	if err != nil {
		fmt.Println("Error parsing configuration file")
	}
	return o
}

func (c *Config) GetConfigFile() (contents []byte) {
	configFile, err := ioutil.ReadFile(c.ConfigFileLocation) //TODO Better paths?
	if err != nil {
		fmt.Printf("Could not read file at: %s\n", c.ConfigFileLocation)
		panic(err)
	} else {
		return configFile
	}
}

func (o Opts) ParseConfigBytes(fileContents []byte) (Opts, error) {
	err := yaml.Unmarshal(fileContents, &o)
	if err != nil {
		panic(err)
	}
	return o, err
}
