package googleimage

import (
	"log"
	"math/rand"
	"strings"
	"time"

	"github.com/kyokomi/goslash/goslash"
	"github.com/kyokomi/goslash/plugins"
	"github.com/kyokomi/slackbot/plugins/googleimage"
)

type plugin struct {
	rd     *rand.Rand
	client googleimage.GoogleImageAPIClient
	cx     string
	apiKey string
}

func NewPlugin(client googleimage.GoogleImageAPIClient) plugins.Plugin {
	return &plugin{
		rd:     rand.New(rand.NewSource(time.Now().UnixNano())),
		client: client,
	}
}

func (p *plugin) Do(req goslash.SlashCommandRequest) goslash.SlashCommandMessage {
	_, args := req.CmdArgs()
	message := strings.Join(args, " ")

	query := strings.Replace(strings.TrimLeft(message, "image me"), "image me", "", 1)

	links, err := p.client.GetImageLinks(query)
	if err != nil {
		log.Println(err)
		goslash.NewMessage(err.Error())
	}

	idx := int(p.rd.Int() % len(links))
	return goslash.NewInChannelMessage(links[idx])
}

var _ plugins.Plugin = (*plugin)(nil)
