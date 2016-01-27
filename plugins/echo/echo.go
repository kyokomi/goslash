package echo

import (
	"strings"

	"github.com/kyokomi/goslash/goslash"
	"github.com/kyokomi/goslash/plugins"
)

type plugin struct {
}

func New() plugins.Plugin {
	return &plugin{}
}

func (p *plugin) Do(req goslash.SlashCommandRequest) goslash.SlashCommandMessage {
	_, args := req.CmdArgs()
	return goslash.NewInChannelMessage(strings.Join(args, ""))
}
