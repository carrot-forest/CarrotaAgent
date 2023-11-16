package feishu

// func TestMessageReply(t *testing.T) {
// 	carrota.LoadConfig("../config.yaml")
// 	bot := NewFeishuBot().WithAppInfo(carrota.C.Feishu.AppID, carrota.C.Feishu.AppSecret).WithCallBackInfo(carrota.C.Feishu.EncryptKey, carrota.C.Feishu.VerificationToken)
// 	bot.Build()
// 	bot.lark.EventCallback.HandlerEventV2IMMessageReceiveV1(
// 		func(ctx context.Context, cli *lark.Lark, schema string, header *lark.EventHeaderV2, event *lark.EventV2IMMessageReceiveV1) (string, error) {
// 			_, err := lark.UnwrapMessageContent(event.Message.MessageType, event.Message.Content)
// 			if err != nil {
// 				return "", err
// 			}
// 			bot.ReplyRawMessage(event.Message.MessageID, "text", "hello\nworld")
// 			// replyMsg, reply, err := cli.Message.Reply(event.Message.MessageID).SendText(ctx, fmt.Sprintf("got text: %s", content.Text.Text))
// 			// fmt.Println(replyMsg, reply, err.Error())
// 			return "", err
// 		})
// 	bot.Runhandler(carrota.C.Feishu.CallbackAPI)
// }
