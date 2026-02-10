package bot

import (
	"testing"

	"github.com/arya237/foodPilot/internal/config"
	"github.com/arya237/foodPilot/internal/infrastructure/bot"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSend(t *testing.T) {
	cfg, err := config.New()
	require.NoError(t, err)

	bot, err := bot.New(cfg.BaleBot)
	require.NoError(t, err)

	sender := NewSender(bot)

	const (
		chatID  = "398725378"
		message = "test"
	)
	err = sender.Send(chatID, message)
	assert.NoError(t, err)
}
