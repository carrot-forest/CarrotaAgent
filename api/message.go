package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/carrot-forest/CarrotaAgent/carrota"
)

type APIMessageSendRequest struct {
	Agent     string   `json:"agent"`
	GroupID   string   `json:"group_id"`
	UserID    string   `json:"user_id"`
	MessageID string   `json:"message_id"`
	Message   []string `json:"message"`
}

type APIMessageSendResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data string `json:"data"`
}

func (api *API) AddAPIMessageSend(path string) {
	http.HandleFunc(fmt.Sprintf("%s%s", api.basepath, path), func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading request body", http.StatusBadRequest)
			return
		}

		var requestData APIMessageSendRequest
		if err := json.Unmarshal(body, &requestData); err != nil {
			http.Error(w, "Error decoding JSON", http.StatusBadRequest)
			return
		}

		fmt.Printf("[api]message send reqdata=%+v\n", requestData)

		response := APIMessageSendResponse{}

		bot, err := api.GetBotFromBotID(requestData.Agent)
		if err != nil {
			response = APIMessageSendResponse{
				Code: 400,
				Msg:  err.Error(),
				Data: "failed",
			}
		} else if bot.Platform == carrota.PlatformTypeFeishu {
			if requestData.MessageID != "" {
				bot.ReplyRawMessage(requestData.MessageID, "text", requestData.Message[0])
				bot.CreateMessageReaction(requestData.MessageID, "MeMeMe")
			} else if requestData.GroupID != "" {
				bot.SendRawMessageToGroup(requestData.GroupID, "text", requestData.Message[0])
			} else if requestData.UserID != "" {
				bot.SendRawMessageToUser(requestData.UserID, "text", requestData.Message[0])
			}
			response = APIMessageSendResponse{
				Code: 200,
				Msg:  "",
				Data: "ok",
			}
		} else {
			response = APIMessageSendResponse{
				Code: 503,
				Msg:  errors.New("platform not supported").Error(),
				Data: "failed",
			}
		}

		jsonResponse, err := json.Marshal(response)
		if err != nil {
			http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")

		w.Write(jsonResponse)
	})
}
