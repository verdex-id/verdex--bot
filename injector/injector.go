//go:build wireinject
// +build wireinject

package injector

import (
	"github.com/google/wire"
	delivery "github.com/verdex-id/verdex-bot/internal/app/delivery/discord"
	"github.com/verdex-id/verdex-bot/internal/app/delivery/discord/event"
)

func InjectDeliveryDiscord() *delivery.DiscordDelivery {
	wire.Build(
		event.NewDiscordEventHandler,
		delivery.NewDiscordDelivery,
	)

	return nil
}
