package feishu

import (
	"regexp"
	"strings"

	"github.com/chyroc/lark"
)

func deleteAtUserText(text string) string {
	pattern := `@_user_\d+ `
	re := regexp.MustCompile(pattern)
	return re.ReplaceAllString(text, "")
}

func convertRawMessageToText(rawMessage string) string {
	rawMessage = strings.Replace(rawMessage, "\n", "\\n", -1)
	return "{\"text\":\"" + rawMessage + "\"}"
}

func convertRawMessageToContent(msg string, msgType string) string {
	if msgType == "text" {
		return convertRawMessageToText(msg)
	}
	return msg
}

func convertTextToRawMessage(text string) string {
	text = strings.Replace(text, "\\n", "\n", -1)
	return text[9 : len(text)-2]
}

func convertContentToRawMessageAndDeleteAtUserText(content string, msgType string) string {
	if msgType == "text" {
		return deleteAtUserText(convertTextToRawMessage(content))
	}
	return ""
}

func (bot *FeishuBot) isMentionedBot(mentions []*lark.EventV2IMMessageReceiveV1MessageMention) bool {
	for _, v := range mentions {
		if v.ID.OpenID == bot.larkMeta.openID {
			return true
		}
	}
	return false
}
