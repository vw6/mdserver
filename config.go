package main
import (
	"io/ioutil"

	"yaml-v2"
)
// Config - dmserver.yaml
type Config struct {
	Listen string `yaml:"listen"`
}

func readConfig(ConfigName string) (x *Config, err error) {
	var file []byte
	if file, err = ioutil.ReadFile(ConfigName); err != nil {
		return nil, err
	}
	x = new(Config)
	if err = yaml.Unmarshal(file, x); err != nil {
		return nil, err
	}
	return x, nil
}
