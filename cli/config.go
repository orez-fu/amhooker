package cli

import (
	"log"

	"github.com/fatih/structs"
	"github.com/leebenson/conform"
	"github.com/sanity-io/litter"
	"github.com/spf13/viper"
)

type config struct {
	DebugMode       string `mapstructure:"debug" structs:"debug" env:"AMHOOKER_DEBUG_MODE"`
	Port            int    `mapstructure:"port" structs:"port" env:"AMHOOKER_PORT"`
	AlertConfigPath string `mapstructure:"alert_config_path" structs:"alert_config_path" env:"ALERT_CONFIG_PATH"`
	// SensitiveData string `mapstructure:"sensitive_data" structs:"sensitive_data" env:"SENSITIVE_DATA" conform:"redact"`
}

var defaultConfig = &config{
	DebugMode:       "INFO",
	Port:            8866,
	AlertConfigPath: "amhooker/config/config.yaml",
}

var GlobalConfig *config

func readGlobalConfig() {
	// Priority of configuration options
	// 1. CLI Parameters
	// 2. Environment
	// 3. Config File
	// 4. Defaults
	config, err := readConfig()
	if err != nil {
		log.Panic(err)
	}

	GlobalConfig = config
}

func configInit() error {
	cliFlags()
	return bindFlagsAndEnv()
}

// cliFlags defines cli parameters for all config options
func cliFlags() {
	rootCmd.PersistentFlags().String("debug", defaultConfig.DebugMode, "Type of debug mode: INFO | DEBUG | NONE")
	rootCmd.PersistentFlags().Int("port", defaultConfig.Port, "Running application port")
	rootCmd.PersistentFlags().String("alert_config_path", defaultConfig.AlertConfigPath, "Alert Manager Config Path")
}

//bindFlagsAndEnv will assign the environment variables to the cli parameters
func bindFlagsAndEnv() (err error) {
	for _, field := range structs.Fields(&config{}) {
		// get the struct tag values
		key := field.Tag("structs")
		env := field.Tag("env")

		// bind cobra flags to viper
		err = viper.BindPFlag(key, rootCmd.PersistentFlags().Lookup(key))
		if err != nil {
			return err
		}
		err = viper.BindEnv(key, env)
		if err != nil {
			return err
		}
	}
	return nil
}

// print print the config object but remote sensitive data
func (c *config) Print() {
	cp := *c
	_ = conform.Strings(&cp)
	litter.Dump(cp)
}

// string string the config object but remote sensitive data
func (c *config) String() string {
	cp := *c
	_ = conform.Strings(&cp)
	return litter.Sdump(cp)
}

// readConfig a helper to read default from a default config object
func readConfig() (*config, error) {
	// create a map of the default config
	defaultsAsMap := structs.Map(defaultConfig)

	// set defaults
	for key, value := range defaultsAsMap {
		viper.SetDefault(key, value)
	}

	// read config from file
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err == nil {
		log.Println("Read config file: ", viper.ConfigFileUsed())
	}

	// unmarshal config into struct
	c := &config{}
	err := viper.Unmarshal(c)
	if err != nil {
		return nil, err
	}
	return c, nil
}
