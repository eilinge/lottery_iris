// Copyright 2019 The KubeSphere Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

package config

import (
	"flag"
	"fmt"
	"os"
	"sync"

	"github.com/koding/multiconfig"
	"github.com/prometheus/common/log"

	"imooc.com/lottery/constants"
)

type Config struct {
	Mysql struct {
		Host     string `json:"host" default:"47.101.187.227"`
		Port     string `json:"port" default:"3306"`
		User     string `json:"user" default:"root"`
		Password string `json:"password" default:"password"`
		Database string `json:"database" default:"lottery"`
		Disable  bool   `json:"disable" default:"false"`
		LogMode  bool   `json:"logmode" default:"true"`
	}

	Redis struct {
		Host      string `default:"47.101.187.227"`
		Port      string `default:"6379"`
		User      string `default:""`
		Pwd       string `default:""`
		IsRunning bool   `default:"true"`
	}

	App struct {
		Host string `default:"localhost"`
		Port string `default:"8080"`
	}
}

var instance *Config

var once sync.Once

func GetInstance() *Config {
	once.Do(func() {
		instance = &Config{}
	})
	return instance
}

type GrpcConfig struct {
	ShowErrorCause bool `default:"false"` // show grpc error cause to frontend
}

type LogConfig struct {
	Level string `default:"debug"` // debug, info, warn, error, fatal
}

func (c *Config) PrintUsage() {
	fmt.Fprintf(os.Stdout, "Usage of %s:\n", os.Args[0])
	flag.PrintDefaults()
	fmt.Fprint(os.Stdout, "\nSupported environment variables:\n")
	e := newLoader(constants.ServiceName)
	e.PrintEnvs(new(Config))
	fmt.Println("")
}

func (c *Config) GetFlagSet() *flag.FlagSet {
	flag.CommandLine.Usage = c.PrintUsage
	return flag.CommandLine
}

func (c *Config) ParseFlag() {
	c.GetFlagSet().Parse(os.Args[1:])
}

func (c *Config) LoadConf() *Config {
	c.ParseFlag()
	config := instance

	m := &multiconfig.DefaultLoader{}
	m.Loader = multiconfig.MultiLoader(newLoader(constants.ServiceName))
	m.Validator = multiconfig.MultiValidator(
		&multiconfig.RequiredValidator{},
	)
	err := m.Load(config)
	if err != nil {
		panic(err)
	}

	log.Info(nil, "LoadConf: %+v", config)

	return config
}
