package api

import (
	"errors"

	"github.com/carrot-forest/CarrotaAgent/feishu"
)

func (api *API) GetBotFromBotID(botID string) (*feishu.FeishuBot, error) {
	for _, bot := range api.feishuBot {
		if bot.BotID == botID {
			return bot, nil
		}
	}
	return nil, errors.New("bot not found")
}
