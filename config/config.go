package config

import (
	"io/ioutil"
	"strings"

	"github.com/hashicorp/go-hclog"
	"github.com/nicholasjackson/env"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
)

var configFile = env.String("CONFIG_FILE", false, "config.yaml", "Configuration file location")
var bindAddress = env.String("BIND_ADDRESS", false, "", "Bind address for the Server")
var eatDbHostname = env.String("EAT_DB_HOSTNAME", false, "localhost", "Database host name. Defaults to localhost")
var eatDbPort = env.Int("EAT_DB_PORT", false, 5432, "Database listenting port. Defaults to 5432")
var eatDbDatabase = env.String("EAT_DB_DATABASE", false, "eatables", "Database name. Defaults to eatables")
var logLevel = env.String("LOG_LEVEL", false, "DEBUG", "Log Level, Defaults to INFO")

type configSrc int

const (
	fileConfig configSrc = iota
	envConfig
)

// Config holds the configuration information for eatables and will be unmarshalled
// from YAML.
type Config struct {
	Logger      hclog.Logger
	BindAddress string `yaml:"bind_address"`

	DB struct {
		Hostname string `yaml:"hostname"`
		Port     int    `yaml:"port"`
		Username string
		Password string
		Database string `yaml:"database"`
		SSL      *bool
		MaxConns int `yaml:"max_connections"`
	}

	Log struct {
		Level string `yaml:"level"`
	}
}

// NewConfig creates a new Config
func NewConfig() (*Config, error) {
	env.Parse()
	conf := new(Config)
	if configFile != nil && *configFile != "" {
		if err := conf.parse(*configFile); err != nil {
			return nil, err
		}
	} else {
		conf.parseEnv()
	}
	if err := conf.validate(); err != nil {
		return nil, err
	}
	conf.Logger = hclog.New(&hclog.LoggerOptions{
		Name:  "eatables",
		Level: hclog.LevelFromString(conf.Log.Level),
	})
	return conf, nil
}

// parse parses configuration yaml file and constructs the config structure
func (conf *Config) parse(file string) error {
	f, err := ioutil.ReadFile(file)
	if err != nil {
		hclog.Default().Error("Error reading configuration", "file", file, "error", err.Error())
		return errors.Wrapf(err, "Error reading %s", file)
	}

	if err := yaml.Unmarshal(f, conf); err != nil {
		hclog.Default().Error("Error parsing configuration", "file", file, "error", err.Error())
		return errors.Wrapf(err, "Error parsing %s", file)
	}
	return nil
}

// ParseEnv parses the environment variable starting with RISTAT
// and constructs the config structure
func (conf *Config) parseEnv() {
	conf.BindAddress = *bindAddress
	conf.DB.Hostname = *eatDbHostname
	conf.DB.Port = *eatDbPort
	conf.DB.Database = *eatDbDatabase
	conf.Log.Level = *logLevel
}

func (conf *Config) validate() error {
	validations := []struct {
		validator    func() bool
		missedConfig string
	}{
		{func() bool { return conf.BindAddress != "" }, "bind_address / BIND_ADDRESS"},
		{func() bool { return conf.DB.Hostname != "" }, "db.hostname / EAT_DB_HOSTNAME"},
		{func() bool { return conf.DB.Database != "" }, "db.database / EAT_DB_DATABASE"},
	}

	missing := []string{}
	for _, v := range validations {
		if v.validator() {
			continue
		}
		missing = append(missing, v.missedConfig)
	}

	if len(missing) > 0 {
		return errors.New("Missing required values: " + strings.Join(missing, ", "))
	}
	return nil
}
