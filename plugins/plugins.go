package plugins

import (
	"fmt"
	"net/http"

	"github.com/kyokomi/goslash/goslash"
)

type SlashCommand interface {
	AddPlugin(cmd string, plugin Plugin)
	Execute(req goslash.SlashCommandRequest) string
}

type slashCommand struct {
	slash   *goslash.SlashCommandService
	plugins map[string]Plugin
}

func (c *slashCommand) AddPlugin(cmd string, plugin Plugin) {
	c.plugins[cmd] = plugin
}

func (c *slashCommand) Execute(req goslash.SlashCommandRequest) string {
	cmd, _ := req.CmdArgs()
	p, ok := c.plugins[cmd]
	if !ok {
		return fmt.Sprintf("%s command not found", cmd)
	}

	resp, err := c.slash.Reply(req, p.Do(req))
	if err != nil {
		return err.Error()
	}
	resp.Body.Close()

	return ""
}

func New(client *http.Client, plugins map[string]Plugin) SlashCommand {
	return &slashCommand{
		slash:   goslash.New(client).SlashCommand,
		plugins: plugins,
	}
}

type Plugin interface {
	Do(req goslash.SlashCommandRequest) goslash.SlashCommandMessage
}
