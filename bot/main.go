package bot

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"gitlab.ritsec.cloud/competitions/ists-2023/store/config"
	"gitlab.ritsec.cloud/competitions/ists-2023/store/data"
)

var (
	Session         *discordgo.Session
	Commands        []*discordgo.ApplicationCommand
	CommandHandlers map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate) = make(map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate))

	minvalue float64 = 0
)

func init() {
	session, err := discordgo.New("Bot " + config.BotToken)
	if err != nil {
		panic(err)
	}

	Session = session
}

func Init() {
	err := Session.Open()
	if err != nil {
		panic(err)
	}

	Commands = []*discordgo.ApplicationCommand{
		{
			Name:        "add",
			Description: "add balance to a user",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Name:        "user",
					Description: "the user to add balance to",
					Type:        discordgo.ApplicationCommandOptionString,
					Required:    true,
					Choices: []*discordgo.ApplicationCommandOptionChoice{
						{
							Name:  "Team01",
							Value: "team01",
						},
						{
							Name:  "Team02",
							Value: "team02",
						},
						{
							Name:  "Team03",
							Value: "team03",
						},
						{
							Name:  "Team04",
							Value: "team04",
						},
						{
							Name:  "Team05",
							Value: "team05",
						},
						{
							Name:  "Team06",
							Value: "team06",
						},
						{
							Name:  "Team07",
							Value: "team07",
						},
						{
							Name:  "Team08",
							Value: "team08",
						},
						{
							Name:  "Team09",
							Value: "team09",
						},
						{
							Name:  "Team10",
							Value: "team10",
						},
						{
							Name:  "Team11",
							Value: "team11",
						},
						{
							Name:  "Team12",
							Value: "team12",
						},
						{
							Name:  "Team13",
							Value: "team13",
						},
						{
							Name:  "Team14",
							Value: "team14",
						},
						{
							Name:  "Team15",
							Value: "team15",
						},
						{
							Name:  "Black",
							Value: "black",
						},
						{
							Name:  "Admin",
							Value: "admin",
						},
						{
							Name:  "Red",
							Value: "red",
						},
						{
							Name:  "White",
							Value: "white",
						},
					},
				},
				{
					Name:        "amount",
					Description: "the amount to add",
					Type:        discordgo.ApplicationCommandOptionInteger,
					Required:    true,
					MinValue:    &minvalue,
				},
			},
		},
		{
			Name:        "balance",
			Description: "get the balance of a user",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Name:        "user",
					Description: "the user to get the balance of",
					Type:        discordgo.ApplicationCommandOptionString,
                    Required:    true,
					Choices: []*discordgo.ApplicationCommandOptionChoice{
						{
							Name:  "Team01",
							Value: "team01",
						},
						{
							Name:  "Team02",
							Value: "team02",
						},
						{
							Name:  "Team03",
							Value: "team03",
						},
						{
							Name:  "Team04",
							Value: "team04",
						},
						{
							Name:  "Team05",
							Value: "team05",
						},
						{
							Name:  "Team06",
							Value: "team06",
						},
						{
							Name:  "Team07",
							Value: "team07",
						},
						{
							Name:  "Team08",
							Value: "team08",
						},
						{
							Name:  "Team09",
							Value: "team09",
						},
						{
							Name:  "Team10",
							Value: "team10",
						},
						{
							Name:  "Team11",
							Value: "team11",
						},
						{
							Name:  "Team12",
							Value: "team12",
						},
						{
							Name:  "Team13",
							Value: "team13",
						},
						{
							Name:  "Team14",
							Value: "team14",
						},
						{
							Name:  "Team15",
							Value: "team15",
						},
						{
							Name:  "Black",
							Value: "black",
						},
						{
							Name:  "Admin",
							Value: "admin",
						},
						{
							Name:  "Red",
							Value: "red",
						},
						{
							Name:  "White",
							Value: "white",
						},
					},
				},
			},
		},
	}

	CommandHandlers["add"] = func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		_, err := data.User.IncrementBalance(i.ApplicationCommandData().Options[0].StringValue(), int(i.ApplicationCommandData().Options[1].IntValue()))
		if err != nil {
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: err.Error(),
				},
			})
			return
		}
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: fmt.Sprintf("Added %v points to %v", i.ApplicationCommandData().Options[1].IntValue(), i.ApplicationCommandData().Options[0].StringValue()),
			},
		})
	}

	CommandHandlers["balance"] = func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		user, err := data.User.Get(i.ApplicationCommandData().Options[0].StringValue())
		if err != nil {
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: err.Error(),
				},
			})
			return
		}
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: fmt.Sprintf("%v has %v points", i.ApplicationCommandData().Options[0].StringValue(), user.Balance),
			},
		})
	}

	_, err = Session.ApplicationCommandBulkOverwrite(AppID, GuildID, Commands)
	if err != nil {
		panic(err)
	}

	Session.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := CommandHandlers[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}
	})
}

func Close() {
	Session.Close()
}
