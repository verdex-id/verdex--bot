package event

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func (h *DiscordEventHandler) Ready(s *discordgo.Session, _ *discordgo.Ready) {
	s.UpdateStatusComplex(discordgo.UpdateStatusData{
		Activities: []*discordgo.Activity{
			{
				Name: "source.verdex.id",
				Type: discordgo.ActivityTypeCompeting,
				URL:  "https://source.verdex.id",
			},
		},
	})
	log.Println("Discord bot online as " + s.State.User.Username)

	// s.GuildOnboardingEdit("", &discordgo.GuildOnboarding{
	// 	Prompts: &[]discordgo.GuildOnboardingPrompt{
	// 		{
	// 			ID:       "university",
	// 			Type:     discordgo.GuildOnboardingPromptTypeDropdown,
	// 			Required: false,
	// 			Options: []discordgo.GuildOnboardingPromptOption{
	// 				{}
	// 			},
	// 		},
	// 	},
	// })
}
