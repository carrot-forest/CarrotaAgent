package feishu

import (
	"context"

	"github.com/chyroc/lark"
)

func (bot *FeishuBot) SendRawMessage(receiveIDType string, receiveID string, msgType string, msg string) {
	content := convertRawMessageToContent(msg, msgType)
	_, _, err := bot.lark.Message.SendRawMessage(context.Background(), &lark.SendRawMessageReq{
		ReceiveIDType: lark.IDType(receiveIDType),
		ReceiveID:     receiveID,
		MsgType:       lark.MsgType(msgType),
		Content:       content,
	})

	if err != nil {
		panic(err)
	}
}

func (bot *FeishuBot) SendRawMessageToGroup(groupID string, msgType string, msg string) {
	bot.SendRawMessage("chat_id", groupID, msgType, msg)
}

func (bot *FeishuBot) SendRawMessageToUser(userID string, msgType string, msg string) {
	bot.SendRawMessage("open_id", userID, msgType, msg)
}

func (bot *FeishuBot) ReplyRawMessage(msgID string, msgType string, msg string) {
	content := convertRawMessageToContent(msg, msgType)
	_, _, err := bot.lark.Message.ReplyRawMessage(context.Background(), &lark.ReplyRawMessageReq{
		MessageID: msgID,
		MsgType:   lark.MsgType(msgType),
		Content:   content,
	})
	if err != nil {
		panic(err)
	}
}

// func (bot *FeishuBot) GetMessageFile(msgID string) {
// 	res, response, err := bot.Lark.Message.GetMessageFile(context.Background(), &lark.GetMessageFileReq{
// 		MessageID: msgID,
// 		File
// 	})
// 	if err != nil {
// 		panic(err)
// 	}
// 	res.Items
// }

// emoji_list: https://open.feishu.cn/document/server-docs/im-v1/message-reaction/emojis-introduce
func (bot *FeishuBot) CreateMessageReaction(msgID string, emojiType string) *lark.CreateMessageReactionResp {
	resp, _, err := bot.lark.Message.CreateMessageReaction(context.Background(), &lark.CreateMessageReactionReq{
		MessageID: msgID,
		ReactionType: &lark.CreateMessageReactionReqReactionType{
			EmojiType: emojiType,
		},
	})
	if err != nil {
		panic(err)
	}
	return resp
}

func (bot *FeishuBot) DeleteMessageReaction(msgID string, reactionID string) {
	_, _, err := bot.lark.Message.DeleteMessageReaction(context.Background(), &lark.DeleteMessageReactionReq{
		MessageID:  msgID,
		ReactionID: reactionID,
	})
	if err != nil {
		panic(err)
	}
}
