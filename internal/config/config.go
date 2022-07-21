package config

import (
	"github.com/spf13/viper"
	"os"
	"time"
)

const (
	defaultHTTPPort               = "8000"
	defaultAccessTTL              = time.Minute * 5
	defaultRefreshTTL             = time.Hour * 24 * 14
	defaultShutdownTime           = time.Second * 10
	defaultRWTimeout              = 10 * time.Second
	defaultHTTPMaxHeaderMegabytes = 1
	envLocal                      = "local"
)

type (
	Config struct {
		Environment string
		HTTP        HTTPConfig
		Auth        AuthConfig
		Mongo       MongoConfig
		App         AppConfig
	}

	HTTPConfig struct {
		Port               string        `mapstructure:"port"`
		ReadTimeout        time.Duration `mapstructure:"readTimeout"`
		WriteTimeout       time.Duration `mapstructure:"writeTimeout"`
		MaxHeaderMegabytes int           `mapstructure:"maxHeaderMegaBytes"`
	}

	AuthConfig struct {
		AccessTTL  time.Duration `mapstructure:"accessTokenTTL"`
		RefreshTTL time.Duration `mapstructure:"refreshTokenTTL"`
		SigningKey string        `mapstructure:"key"`
		Issuer     string
	}

	MongoConfig struct {
		Database string `mapstructure:"databaseName"`
		URI      string
	}

	AppConfig struct {
		ShutdownTime time.Duration `mapstructure:"shutdown_time"`
	}
)

func Init(configDir string) (*Config, error) {
	populateDefaults()

	if err := parseConfigFile(configDir, os.Getenv("APP_ENV")); err != nil {
		return nil, err
	}

	var cfg Config
	if err := unmarshal(&cfg); err != nil {
		return nil, err
	}

	setFromEnv(&cfg)

	return &cfg, nil
}

func populateDefaults() {
	viper.SetDefault("auth.refresh_token_ttl", defaultRefreshTTL)
	viper.SetDefault("auth.access_token_ttl", defaultAccessTTL)

	viper.SetDefault("app.shutdown_time", defaultShutdownTime)

	viper.SetDefault("http.port", defaultHTTPPort)
	viper.SetDefault("http.timeouts.read", defaultRWTimeout)
	viper.SetDefault("http.timeouts.write", defaultRWTimeout)
	viper.SetDefault("http.max_header_megabytes", defaultHTTPMaxHeaderMegabytes)
}

func parseConfigFile(configDir, env string) error {
	viper.AddConfigPath(configDir)
	viper.SetConfigName("main")

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	if env == envLocal {
		return nil
	}

	viper.SetConfigName(env)

	return viper.MergeInConfig()
}

func unmarshal(cfg *Config) error {
	if err := viper.UnmarshalKey("http", &cfg.HTTP); err != nil {
		return err
	}
	if err := viper.UnmarshalKey("mongo", &cfg.Mongo); err != nil {
		return err
	}
	if err := viper.UnmarshalKey("auth", &cfg.Auth); err != nil {
		return err
	}

	return nil
}

func setFromEnv(cfg *Config) {
	cfg.Environment = os.Getenv("ENV")
	cfg.Auth.SigningKey = os.Getenv("SECRET")
	cfg.Environment = os.Getenv("ENV")
	cfg.Mongo.URI = os.Getenv("MONGO_URI")

}
