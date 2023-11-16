package plugincenter

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type ForwardMessageRequest struct {
	Agent       string `json:"agent"`
	GroupID     string `json:"group_id"`
	GroupName   string `json:"group_name"`
	UserID      string `json:"user_id"`
	UserName    string `json:"user_name"`
	MessageID   string `json:"message_id"`
	IsMentioned bool   `json:"is_mention"`
	Time        int64  `json:"time"`
	Message     string `json:"message"`
}

type ForwardMessageResponse struct {
	IsReply bool     `json:"is_reply"`
	Message []string `json:"message"`
}

func (pc *PluginCenter) ForwardMessage(reqData ForwardMessageRequest) {
	url := fmt.Sprintf("http://%s:%s%s%s", pc.meta.IP, pc.meta.Port, pc.meta.API.BasePath, pc.meta.API.MessageReportPath)

	jsonData, err := json.Marshal(reqData)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	fmt.Printf("[plugin-center]forward message reqdata=%+v\n", reqData)

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	fmt.Printf("[plugin-center]forward message response=%s\n", string(body))
}
