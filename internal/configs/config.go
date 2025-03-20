package configs

import "github.com/spf13/viper"

var config *Config

type Option struct {
	configFolders []string
	configFile    string
	configType    string
}
type OptionFunc func(*Option)

func Init(optFuncs ...OptionFunc)error {
	var option *Option = &Option{
		configFolders: getDefaultConfigFolder(),
		configFile:    getDefaultConfigFile(),
		configType:    getDefaultConfigType(),
	}

	for _, optFunc := range optFuncs {
		optFunc(option)
	}

	for _, configFolder := range option.configFolders {
		viper.AddConfigPath(configFolder)
	}

	viper.SetConfigName(option.configFile)
	viper.SetConfigType(option.configType)
	viper.AutomaticEnv()

	config = new(Config)
	var err error = viper.ReadInConfig()
	if err != nil{
		return err
	}

	return viper.Unmarshal(&config)
}

func getDefaultConfigFolder() []string {
	return []string{"./internal/configs/"}
}

func getDefaultConfigFile() string {
	return "config"
}

func getDefaultConfigType() string {
	return "yaml"
}

func WithConfigFolder(configFolders []string) OptionFunc {
	return func(opt *Option) {
		opt.configFolders = configFolders
	}
}

func WithConfigFile(configFile string) OptionFunc {
	return func(opt *Option) {
		opt.configFile = configFile
	}
}

func WithConfigType(configType string) OptionFunc {
	return func(opt *Option) {
		opt.configType = configType
	}
}

func GetConfig() *Config {
	if config == nil {
		config = &Config{}
	}
	return config
}