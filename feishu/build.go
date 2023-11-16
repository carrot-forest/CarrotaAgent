package feishu

import (
	"github.com/carrot-forest/CarrotaAgent/carrota"
	plugincenter "github.com/carrot-forest/CarrotaAgent/plugin-center"
	"github.com/chyroc/lark"
)

type FeishuBot struct {
	Platform     carrota.PlatformType
	BotName      string
	BotID        string
	identifyMeta struct {
		appID             string
		appSecret         string
		encryptKey        string
		verificationToken string
	}
	lark     *lark.Lark
	larkMeta struct {
		openID string
	}
	pluginCenter *plugincenter.PluginCenter
}

func NewFeishuBot(name string, id string) *FeishuBot {
	bot := &FeishuBot{
		Platform: carrota.PlatformTypeFeishu,
		BotName:  name,
		BotID:    id,
	}
	return bot
}

func (bot *FeishuBot) WithAppInfo(appID string, appSecret string) *FeishuBot {
	bot.identifyMeta.appID = appID
	bot.identifyMeta.appSecret = appSecret
	return bot
}

func (bot *FeishuBot) WithCallBackInfo(encryptKey string, verificationToken string) *FeishuBot {
	bot.identifyMeta.encryptKey = encryptKey
	bot.identifyMeta.verificationToken = verificationToken
	return bot
}

func (bot *FeishuBot) WithPluginCenter(pc *plugincenter.PluginCenter) *FeishuBot {
	bot.pluginCenter = pc
	return bot
}

func (bot *FeishuBot) Build() {
	bot.lark = lark.New(
		lark.WithAppCredential(bot.identifyMeta.appID, bot.identifyMeta.appSecret),
		lark.WithEventCallbackVerify(bot.identifyMeta.encryptKey, bot.identifyMeta.verificationToken),
		lark.WithNonBlockingCallback(true),
	)
	botInfo := bot.GetBotInfo()
	bot.larkMeta.openID = botInfo.OpenID
}
