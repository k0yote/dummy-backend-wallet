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
	KeyThresholdShares   int    `mapstructure:"KEY_THREDHOLD_SHARES"`
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
