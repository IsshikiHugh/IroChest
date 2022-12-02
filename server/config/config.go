package config

import (
	"io/ioutil"

	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

var filename = "config.yaml" // Config file name.

var Conf Configuration // Config object.

func Init() {
	logrus.Info("Trying to initialize config module...")
	if yamlFile, err := ioutil.ReadFile(filename); err != nil {
		logrus.Fatal(ErrReadYamlFail.Error() + " : " + err.Error())
	} else if err = yaml.Unmarshal(yamlFile, &Conf); err != nil {
		logrus.Fatal(ErrUnmarshalYamlFail.Error() + " : " + err.Error())
	}
	logrus.Info("Initialized!", Conf)
}
