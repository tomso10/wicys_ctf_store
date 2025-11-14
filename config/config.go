package config

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"gitlab.ritsec.cloud/competitions/ists-2023/store/data"
)

var (
	Hostname string
	JWTKey   string
	Webhook  string
	Timeout  int
	BotToken string
)

func init() {
	viper.SetConfigFile("databases/config.yml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	UpdateConfigs()

	viper.OnConfigChange(
		func(e fsnotify.Event) {
			if e.Op == fsnotify.Write {
				UpdateConfigs()
			}
		},
	)

	viper.WatchConfig()
}

func UpdateConfigs() {
	JWTKey = viper.GetString("key")
	Hostname = viper.GetString("hostname")
	Webhook = viper.GetString("webhook")
	Timeout = viper.GetInt("timeout")
	BotToken = viper.GetString("bot")
	viper.UnmarshalKey("items", &data.Items)
	viper.UnmarshalKey("users", &data.Users)
	viper.UnmarshalKey("token_info", &data.TokenInfo)
	viper.UnmarshalKey("tokens", &data.Tokens)
	data.UpdateUsers()
	data.UpdateItems()
	data.UpdateTokens()
}
