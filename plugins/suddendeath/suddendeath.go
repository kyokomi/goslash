package suddendeath

import (
	"unicode/utf8"

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

	text := "突然の" + args[0]
	size := utf8.RuneCountInString(text)
	header := ""
	for i := 0; i < size+2; i++ {
		header += "人"
	}

	fotter := ""
	for i := 0; i < size; i++ {
		fotter += "^Y"
	}

	reMessage := "＿" + header + "＿"
	reMessage += "\n"
	reMessage += "＞　" + text + "　＜"
	reMessage += "\n"
	reMessage += "￣Y" + fotter + "￣"

	return goslash.NewInChannelMessage(reMessage)
}

var _ plugins.Plugin = (*plugin)(nil)
