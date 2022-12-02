package config

import "errors"

// Define the errors.
var ErrReadYamlFail = errors.New("ERR_READ_YAML_FAIL")
var ErrUnmarshalYamlFail = errors.New("ERR_UNMARSHAL_YAML_FAIL")

type Configuration struct {
	OriginAccessKeys string `yaml:"access"`
	Port             string `yaml:"port"`
}
