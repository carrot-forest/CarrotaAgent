package api

import "github.com/carrot-forest/CarrotaAgent/feishu"

type API struct {
	basepath  string
	feishuBot []*feishu.FeishuBot
}

func NewAPI() *API {
	api := &API{}
	return api
}

func (api *API) WithBasePath(basepath string) *API {
	api.basepath = basepath
	return api
}

func (api *API) WithFeishuBot(bot *feishu.FeishuBot) *API {
	api.feishuBot = append(api.feishuBot, bot)
	return api
}
