package util

import (
	"github.com/spf13/viper"
)

// Config stores all configuration of the application.
// The values are read by viper from a config file or environment variable.
type Config struct {
	Environment          string `mapstructure:"ENVIRONMENT"`
	HTTPServerAddress    string `mapstructure:"HTTP_SERVER_ADDRESS"`
	TOTPIssuer           string `mapstructure:"TOTP_ISSURE"`
	EmailSenderAddress   string `mapstructure:"EMAIL_SENDER_ADDRESS"`
	EmailSenderPassword  string `mapstructure:"EMAIL_SENDER_PASSWORD"`
	SmtpAuthAddress      string `mapstructure:"SMTP_AUTH_ADDRESS"`
	SmtpServerAddress    string `mapstructure:"SMTP_SERVER_ADDRESS"`
	KmsResourceLocation  string `mapstructure:"KMS_RESOURCE_LOCATION"`
	PassCodeExpirePeriod int    `mapstructure:"PASSCODE_EXPIRE_PERIOD"`
	KeyStringPhrase      string `mapstructure:"KEY_STRING_PHRASE"`
	KeySecretShares      int    `mapstructure:"KEY_SECRET_SHARES"`
	KeyThresholdShares   int    `mapstructure:"KEY_THRESHOLD_SHARES"`
	PrivkeyPath          string `mapstructure:"PRIVATE_KEY_FULL_PATH"`
	PubKeyPath           string `mapstructure:"PUBLICKEY_KEY_FULL_PATH"`
	IssuerBaseURL        string `mapstructure:"ISSUER_BASE_URL"`
	BasicAuthUsername    string `mapstructure:"BASIC_AUTH_USERNAME"`
	BasicAuthPassword    string `mapstructure:"BASIC_AUTH_PASSWORD"`
	RedisHost            string `mapstructure:"REDIS_HOST"`
	RedisPort            string `mapstructure:"REDIS_PORT"`
	RedisPassword        string `mapstructure:"REDIS_PASSWORD"`
	RedisDbname          int    `mapstructure:"REDIS_DBNAME"`
	MongoDBProtocol      string `mapstructure:"MONGODB_PROTOCOL"`
	MongoDBHost          string `mapstructure:"MONGODB_HOST"`
	MongoDBUsername      string `mapstructure:"MONGODB_USERNAME"`
	MongoDBPassword      string `mapstructure:"MONGODB_PASSWORD"`
	MongoDBDbname        string `mapstructure:"MONGODB_DBNAME"`
	MongoDBPort          string `mapstructure:"MONGODB_PORT"`
}

// LoadConfig reads configuration from file or environment variables.
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
