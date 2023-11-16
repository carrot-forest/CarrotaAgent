package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/carrot-forest/CarrotaAgent/api"
	"github.com/carrot-forest/CarrotaAgent/carrota"
	"github.com/carrot-forest/CarrotaAgent/feishu"
	plugincenter "github.com/carrot-forest/CarrotaAgent/plugin-center"
)

func main() {
	carrota.LoadConfig("config.yaml")

	pluginCenter := plugincenter.NewPluginCenter().
		WithAddress(carrota.C.PluginCenter.IP, carrota.C.PluginCenter.Port).
		WithBasePath(carrota.C.PluginCenter.API.BasePath).
		WithMessageReport(carrota.C.PluginCenter.API.MessageReport)
	pluginCenter.Build()

	api := api.NewAPI().WithBasePath(carrota.C.API.BasePath)

	for _, f := range carrota.C.Feishu {
		fmt.Printf("add bot %s [%s]\n", f.BotName, f.BotID)

		feishuBot := feishu.NewFeishuBot(f.BotName, f.BotID).
			WithAppInfo(f.AppID, f.AppSecret).
			WithCallBackInfo(f.EncryptKey, f.VerificationToken).
			WithPluginCenter(pluginCenter)
		feishuBot.Build()

		feishuBot.AddHandlerP2PChatCreate()
		feishuBot.AddHandlerIMMessageReceive()
		feishuBot.AddHandlerChatMemberBotAdded()
		feishuBot.Runhandler(f.CallbackAPI)

		api = api.WithFeishuBot(feishuBot)
	}

	api.AddAPIMessageSend(carrota.C.API.MessageSendAPI)

	fmt.Println("start server ...")
	log.Fatal(http.ListenAndServe(":"+carrota.C.Port, nil))
}
