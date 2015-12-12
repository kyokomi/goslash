package time

import (
	"time"

	"github.com/kyokomi/goslash/goslash"
	"github.com/kyokomi/goslash/plugins"
)

type plugin struct {
}

func New() plugins.Plugin {
	return &plugin{}
}

func (p *plugin) Do(_ goslash.SlashCommandRequest) goslash.SlashCommandMessage {
	return goslash.NewInChannelMessage(
		time.Now().Format(time.RFC3339),
	)
}
