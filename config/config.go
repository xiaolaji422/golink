package config

import (
	"fmt"

	"github.com/pelletier/go-toml"
)

var Conf = conf{}

type conf struct {
	*toml.Tree
}

func init() {
	config, err := toml.LoadFile("./config/config.toml")
	if err != nil {
		fmt.Println("error on conf init:", err.Error())
	}
	Conf = conf{
		config,
	}
}
