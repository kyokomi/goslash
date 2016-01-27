package lgtm

import (
	"strings"

	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/kyokomi/goslash/goslash"
	"github.com/kyokomi/goslash/plugins"
)

const lgtmURL = "http://lgtm.in/g"

type plugin struct {
	httpClient *http.Client
}

func New(httpClient *http.Client) plugins.Plugin {
	return &plugin{
		httpClient: httpClient,
	}
}

func (p *plugin) Do(req goslash.SlashCommandRequest) goslash.SlashCommandMessage {
	_, args := req.CmdArgs()

	sendMessage, ok := p.getLGTMImageURL(p.buildRandomURL(strings.Join(args, " ")))
	if ok {
		return goslash.NewInChannelMessage(sendMessage)
	} else {
		return goslash.NewMessage(sendMessage)
	}
}

func (p *plugin) buildRandomURL(message string) string {
	randomURL := lgtmURL
	args := strings.Fields(message)
	if len(args) == 2 {
		randomURL += "/" + args[1]
	}
	return randomURL
}

func (p *plugin) getLGTMImageURL(lgtmURL string) (string, bool) {
	res, err := p.httpClient.Get(lgtmURL)
	if err != nil {
		return err.Error(), false
	}
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromResponse(res)
	if err != nil {
		return err.Error(), false
	}

	text, exists := doc.Find("#imageUrl").Attr("value")
	if !exists {
		return lgtmURL + ": not exists", false
	}

	return text, true
}

var _ plugins.Plugin = (*plugin)(nil)
