package event

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/spf13/viper"
)

func (h *DiscordEventHandler) QuestionUniversityCommand(s *discordgo.Session, msg *discordgo.MessageCreate) {
	if msg.Content != "!question university" {
		return
	}

	if !h.IsAdmin(s, msg.Author, msg.ChannelID) {
		s.ChannelMessageSendComplex(msg.ChannelID, &discordgo.MessageSend{
			Content:   "Kamu tidak memiliki akses untuk perintah ini!",
			Reference: msg.Reference(),
			Flags:     discordgo.MessageFlagsEphemeral,
		})
		return
	}

	universities, err := h.PDDIKTIService.GetManyUniversity(1)
	if err != nil {
		s.ChannelMessageSendComplex(msg.ChannelID, &discordgo.MessageSend{
			Content:   "Error: " + err.Error(),
			Reference: msg.Reference(),
			Flags:     discordgo.MessageFlagsEphemeral,
		})
		return
	}

	embed := &discordgo.MessageEmbed{
		Title:       "Kamu sedang berkuliah dimana?",
		Description: fmt.Sprintf("Yuk pilih Universitas kamu dibawah ini, siapa tau jadi ada kenalan teman satu kampus 😊\nTotal ada %d dari %d halaman Universitas nih yang bisa kamu pilih dibawah ini. Kalau ngga ada bisa lapor ke kami ya 🤗", universities.TotalItems, universities.TotalPages),
		Color:       viper.GetInt("discord.setting.embed.color.default"),
	}

	// embedPagination := &discordgo.MessageEmbed{}

	selectMenuOption := []discordgo.SelectMenuOption{}
	for _, university := range universities.Data {
		selectMenuOption = append(selectMenuOption, discordgo.SelectMenuOption{
			Label:       university.NamaPT,
			Description: fmt.Sprintf("%s, %s", university.KabKotaPT, university.ProvinsiPT),
			Value:       university.NamaPT,
		})
	}

	_, err = s.ChannelMessageSendComplex(msg.ChannelID, &discordgo.MessageSend{
		Embeds: []*discordgo.MessageEmbed{
			embed,
		},
		Components: []discordgo.MessageComponent{
			discordgo.ActionsRow{
				Components: []discordgo.MessageComponent{
					discordgo.SelectMenu{
						CustomID: "universities",
						Options:  selectMenuOption,
					},
				},
			},
		},
	})

	if err != nil {
		log.Fatal(err)
	}
}
