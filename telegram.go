package gomisc

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

// SendMessage sends a message to telegram chat on behalf of a bot
func SendMessage(botID string, chatID string, message string) (*http.Response, error) {
	return SendMessageFormatted(botID, chatID, message, "text")
}

// SendHTMLMessage sends an html formatted message to telegram chat on behalf of a bot
func SendHTMLMessage(botID string, chatID string, html string) (*http.Response, error) {
	return SendMessageFormatted(botID, chatID, html, "html")
}

// SendMarkdownMessage sends an markdown formatted message to telegram chat on behalf of a bot
func SendMarkdownMessage(botID string, chatID string, markdown string) (*http.Response, error) {
	return SendMessageFormatted(botID, chatID, markdown, "markdown")
}

// SendMessageFormatted sends a message to telegram chat on behalf of a bot.
//
// Message can be formatted as "text", "html", "markdown".
func SendMessageFormatted(botID string, chatID string, message string, format string) (*http.Response, error) {
	form := url.Values{}
	form.Add("chat_id", chatID)
	form.Add("text", message)
	if format == "text" {
		// default format is used
	} else if format == "html" || format == "markdown" {
		form.Add("parse_mode", format)
	} else {
		return nil, fmt.Errorf("Format '%s' is not recognized", format)
	}
	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", botID)
	return http.Post(url,
		"application/x-www-form-urlencoded",
		strings.NewReader(form.Encode()))
}
