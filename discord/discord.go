package discord

import (
	"fmt"

	"github.com/gtuk/discordwebhook"
	"gitlab.ritsec.cloud/competitions/ists-2023/store/config"
)

func Send(teamname string, item string, instructions string) {
	content := fmt.Sprintf("<@&1285064061582835756>\n\n**%s** has purchased **%s**\n**Instructions**:\n%s", teamname, item, instructions)

	message := discordwebhook.Message{
		Content: &content,
	}

	_ = discordwebhook.SendMessage(config.Webhook, message)
}
