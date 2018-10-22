package common

import "time"

type DataResourceContext struct {
	Server    Server     `yaml:"server"`
	DBConfigs []DBConfig `yaml:"dbconfigs"`
}

type Server struct {
	Name    string `yaml:"name"`
	Version string `yaml:"version"`
}

type DBConfig struct {
	Name   string `yaml:"name"`
	Config Config `yaml:"config"`
}

type Config struct {
	Mode            bool          `yaml:"mode"`
	Driver          string        `yaml:"driver"`
	Host            string        `yaml:"host"`
	Port            int           `yaml:"port"`
	UserName        string        `yaml:"username"`
	Password        string        `yaml:"password"`
	DataBaseName    string        `yaml:"databasename"`
	ConnMaxLifetime time.Duration `yaml:"lifetime"`
	MaxOpenNum      int           `yaml:"max-open-num"`
	MaxIdleNum      int           `yaml:"max-idle-num"`
}
