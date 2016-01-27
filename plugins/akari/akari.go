package akari

import (
	"fmt"
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
	message := strings.Join(args, " ")

	return goslash.NewInChannelMessage(
		fmt.Sprintf("わぁい%s あかり%s大好き", message, message),
	)
}

var _ plugins.Plugin = (*plugin)(nil)
