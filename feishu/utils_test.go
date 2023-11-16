package feishu

import (
	"fmt"
	"testing"
)

func TestConvertTextToRawMessage(t *testing.T) {
	message := "@_user_1 你为什么不开心"
	messageRaw := "你为什么不开心"
	msg := convertRawMessageToText(message)
	msg_ := convertContentToRawMessageAndDeleteAtUserText(msg, "text")
	fmt.Println(messageRaw, msg_)
	if msg_ != messageRaw {
		t.Fail()
	}
}
