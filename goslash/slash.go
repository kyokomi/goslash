package goslash

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strings"
)

type SlashCommandRequest struct {
	Token       string `json:"token"`
	TeamID      string `json:"team_id"`
	TeamDomain  string `json:"team_domain"`
	ChannelID   string `json:"channel_id"`
	ChannelName string `json:"channel_name"`
	UserID      string `json:"user_id"`
	UserName    string `json:"user_name"`
	Command     string `json:"command"`
	Text        string `json:"text"`
	ResponseURL string `json:"response_url"`
}

func (r *SlashCommandRequest) CmdArgs() (cmd string, args []string) {
	fields := strings.Fields(r.Text)
	if len(fields) == 0 {
		return "", []string{}
	}
	cmd, args = fields[0], fields[1:]
	return cmd, args
}

type SlashCommandMessage struct {
	ResponseType string       `json:"response_type,omitempty"`
	Text         string       `json:"text"`
	Attachments  []Attachment `json:"attachments,omitempty"`
}

func NewMessage(text string) SlashCommandMessage {
	return SlashCommandMessage{
		Text:        text,
		Attachments: []Attachment{},
	}
}

func NewInChannelMessage(text string) SlashCommandMessage {
	return SlashCommandMessage{
		ResponseType: "in_channel",
		Text:         text,
		Attachments:  []Attachment{},
	}
}

func ParseFormSlashCommandRequest(r *http.Request) (SlashCommandRequest, error) {
	if err := r.ParseForm(); err != nil {
		return SlashCommandRequest{}, err
	}
	return SlashCommandRequest{
		Token:       r.PostForm.Get("token"),
		TeamID:      r.PostForm.Get("team_id"),
		TeamDomain:  r.PostForm.Get("team_domain"),
		ChannelID:   r.PostForm.Get("channel_id"),
		ChannelName: r.PostForm.Get("channel_name"),
		UserID:      r.PostForm.Get("user_id"),
		UserName:    r.PostForm.Get("user_name"),
		Command:     r.PostForm.Get("command"),
		Text:        r.PostForm.Get("text"),
		ResponseURL: r.PostForm.Get("response_url"),
	}, nil
}

type SlashCommandService struct {
	client *Client
}

func (c *SlashCommandService) Reply(req SlashCommandRequest, msg SlashCommandMessage) (*http.Response, error) {
	var jsonData bytes.Buffer
	if err := json.NewEncoder(&jsonData).Encode(&msg); err != nil {
		return nil, err
	}

	return c.client.Post(
		req.ResponseURL,
		"application/json; charset=utf-8",
		&jsonData,
	)
}
