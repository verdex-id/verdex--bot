package event

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/spf13/viper"
)

func (h *DiscordEventHandler) MemberAdd(s *discordgo.Session, member *discordgo.GuildMemberAdd) {
	channelID := viper.GetString("discord.setting.welcome.channel_id")
	welcomeMessage := viper.GetString("discord.setting.welcome.message")

	embed := &discordgo.MessageEmbed{
		Author: &discordgo.MessageEmbedAuthor{
			Name:    member.User.Username,
			IconURL: member.User.AvatarURL("64x64"),
		},
		Description: welcomeMessage,
		Color:       viper.GetInt("discord.setting.embed.color.default"),
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "Start Chat",
				Value:  fmt.Sprintf("<#%s>", viper.GetString("discord.setting.chat.channel_id")),
				Inline: true,
			},
			{
				Name:   "Verdex Source",
				Value:  "Digital Sources, visit [**source.verdex.id**](https://source.verdex.id)",
				Inline: true,
			},
		},
	}

	s.ChannelMessageSendComplex(channelID, &discordgo.MessageSend{
		Content: member.Mention(),
		Embed:   embed,
	})
}
