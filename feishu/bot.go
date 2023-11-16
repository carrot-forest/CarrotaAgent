package feishu

import (
	"context"

	"github.com/chyroc/lark"
)

func (bot *FeishuBot) GetBotInfo() *lark.GetBotInfoResp {
	resp, _, err := bot.lark.Bot.GetBotInfo(context.Background(), &lark.GetBotInfoReq{})
	if err != nil {
		panic(err)
	}
	return resp
}
