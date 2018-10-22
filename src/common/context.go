package common

import (
	"gopkg.in/yaml.v2"
	"hdg.com/tools/src/common-util"
	"os"
)

var ConfigurationContext *DataResourceContext

func init() {
	context := new(DataResourceContext)
	logger.Infof("init context ... ")
	configPath:=os.Getenv(ConfigEnvName)
	if configPath == ""{
		configPath = defaultConfigFilePath
	}
	err := context.Parse(configPath)
	if err != nil {
		panic(err)
	}
	logger.Infof("server name:%s version:%s",context.Server.Name,context.Server.Version)
	ConfigurationContext = context
}

type Configuration interface {
	Parse(path string) error
}

func (context *DataResourceContext) Parse(path string) error {
	reader := common_util.NewBufferFileReader(path)
	data, err := reader.Read()
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(data, context)
	if err != nil {
		return err
	}
	return nil
}
