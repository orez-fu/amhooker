package models

type TelegramAlertConfig struct {
	Enabled  bool   `yaml:"enabled"`
	BotToken string `yaml:"botToken"`
}

type AlertConfig struct {
	TimeZone          string              `yaml:"timeZone"`
	TimeoutFormat     string              `yaml:"timeOutputFormat"`
	SplitMessageBytes int                 `yaml:"splitMessageBytes"`
	TelegramConfig    TelegramAlertConfig `yaml:"telegram"`

	// TelegramToken       string `yaml:"telegramToken"`
	// SplitToken          string `yaml:"splitToken"`
	// SendOnly            bool   `yaml:"sendOny"`
	// DisableNotification bool   `yaml:"disableNotification"`
}
