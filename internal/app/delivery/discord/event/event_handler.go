package event

import (
	"github.com/bwmarrin/discordgo"
	"github.com/verdex-id/verdex-bot/internal/app/service"
)

type DiscordEventHandler struct {
	PDDIKTIService *service.PDDIKTIService
}

func NewDiscordEventHandler(pddiktiService *service.PDDIKTIService) *DiscordEventHandler {
	return &DiscordEventHandler{
		PDDIKTIService: pddiktiService,
	}
}

func (h *DiscordEventHandler) IsAdmin(s *discordgo.Session, user *discordgo.User, channelID string) bool {
	perms, _ := s.State.UserChannelPermissions(user.ID, channelID)

	return perms&discordgo.PermissionAdministrator != 0
}
