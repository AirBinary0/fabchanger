package config

import (
	"errors"
	"flag"
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	General
	Mode   string
	Join   string
	Output string
	Input  string
	Merge  string
}

type General struct {
	Identity          string
	MSPId             string
	Channel           string
	ConnectionProfile string
	ConfigTxPath      string
	OrgToJoinMSP      string
	MyOrg             string
}

func GetConfig() (*Config, error) {

	var Configuration *Config
	mode := flag.String("mode", "", "operating mode")
	input := flag.String("f", "", "input file name")
	merge := flag.String("merge", "", "file to merge with")
	output := flag.String("o", "", "output file name")
	join := flag.String("join", "", "join org or orderer")
	flag.Parse()

	if *mode == "" {
		return nil, errors.New("please specify --mode")
	}

	// read config
	viper.SetConfigName("config")
	viper.AddConfigPath("./config")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		return nil, errors.New(fmt.Sprintf("unable to read config file, %s", err))
	}

	err := viper.Unmarshal(&Configuration)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("unable to decode into struct, %v", err))
	}
	Configuration.Mode = *mode
	Configuration.Input = *input
	Configuration.Output = *output
	Configuration.Join = *join
	Configuration.Merge = *merge
	return Configuration, nil
}
