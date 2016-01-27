package suddendeath

import (
	"unicode/utf8"

	"github.com/kyokomi/goslash/goslash"
	"github.com/kyokomi/goslash/plugins"
)

type plugin struct {
}

func NewPlugin() plugins.Plugin {
	return &plugin{}
}

func (p *plugin) Do(req goslash.SlashCommandRequest) goslash.SlashCommandMessage {
	_, args := req.CmdArgs()

	size := utf8.RuneCountInString(args[0])
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
	reMessage += "＞　" + args[0] + "　＜"
	reMessage += "\n"
	reMessage += "￣Y" + fotter + "￣"

	return goslash.NewMessage(reMessage)
}

var _ plugins.Plugin = (*plugin)(nil)
