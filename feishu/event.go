package feishu

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	plugincenter "github.com/carrot-forest/CarrotaAgent/plugin-center"
	"github.com/chyroc/lark"
)

func (bot *FeishuBot) AddHandlerP2PChatCreate() {
	bot.lark.EventCallback.HandlerEventV1P2PChatCreate(
		func(ctx context.Context, cli *lark.Lark, schema string, header *lark.EventHeaderV1, event *lark.EventV1P2PChatCreate) (string, error) {
			fmt.Println("got event:", event.ChatID, event.User.Name)
			_, _, err := cli.Message.Send().ToChatID(event.ChatID).SendText(ctx, "hello! "+event.User.Name)
			return "", err
		})
}

func (bot *FeishuBot) AddHandlerIMMessageReceive() {
	bot.lark.EventCallback.HandlerEventV2IMMessageReceiveV1(
		func(ctx context.Context, cli *lark.Lark, schema string, header *lark.EventHeaderV2, event *lark.EventV2IMMessageReceiveV1) (string, error) {
			_, err := lark.UnwrapMessageContent(event.Message.MessageType, event.Message.Content)
			if err != nil {
				return "", err
			}
			fmt.Printf("[feishu]event im.message.receive event=%+v\n", event)

			switch event.Message.MessageType {
			case lark.MsgTypeText:
				bot.CreateMessageReaction(event.Message.MessageID, string(lark.EmojiOnIt))
				reqData := plugincenter.ForwardMessageRequest{}
				t, _ := strconv.ParseInt(event.Message.CreateTime, 10, 64)
				reqData = plugincenter.ForwardMessageRequest{
					Agent:       bot.BotID,
					UserID:      event.Sender.SenderID.OpenID,
					UserName:    bot.GetUser(event.Sender.SenderID.OpenID).User.Name,
					Time:        t,
					IsMentioned: bot.isMentionedBot(event.Message.Mentions),
					MessageID:   event.Message.MessageID,
					Message:     convertContentToRawMessageAndDeleteAtUserText(event.Message.Content, "text"),
				}
				if event.Message.ChatType == "group" {
					reqData.GroupID = event.Message.ChatID
					reqData.GroupName = string(event.Message.ChatType)
				}
				bot.pluginCenter.ForwardMessage(reqData)
			// case lark.MsgTypeFile:
			// 	bot.CreateMessageReaction(event.Message.MessageID, string(lark.EmojiOnIt))
			// case lark.MsgTypeImage:
			// 	bot.CreateMessageReaction(event.Message.MessageID, string(lark.EmojiOnIt))
			default:
				bot.CreateMessageReaction(event.Message.MessageID, string(lark.EmojiVRHeadSet))
			}

			return "", err
		})
}

func (bot *FeishuBot) AddHandlerChatMemberBotAdded() {
	bot.lark.EventCallback.HandlerEventV2IMChatMemberBotAddedV1(
		func(ctx context.Context, cli *lark.Lark, schema string, header *lark.EventHeaderV2, event *lark.EventV2IMChatMemberBotAddedV1) (string, error) {
			fmt.Println("got event:", event.ChatID, event.Name)
			return "", nil
		})
}

func (bot *FeishuBot) Runhandler(callbackPath string) {
	http.HandleFunc(callbackPath, func(w http.ResponseWriter, r *http.Request) {
		bot.lark.EventCallback.ListenCallback(r.Context(), r.Body, w)
	})
}
