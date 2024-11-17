package test

import (
	"testing"

	"github.com/verdex-id/verdex-bot/internal/app/service"
	"github.com/verdex-id/verdex-bot/internal/infrastructure/config"
)

func init() {
	config.Load()
}

func TestFetchData(t *testing.T) {
	s := service.NewPDDIKTIService()

	s.GetManyUniversity(1)
}
