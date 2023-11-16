package feishu

import (
	"context"

	"github.com/chyroc/lark"
)

func (bot *FeishuBot) GetUserList() []*lark.GetUserListRespItem {
	res, _, err := bot.lark.Contact.GetUserList(context.Background(), &lark.GetUserListReq{
		DepartmentID: "0",
	})
	if err != nil {
		panic(err)
	}
	// fmt.Println("req-id:", response.RequestID)
	// fmt.Println("res:", res)
	userList := res.Items
	for res.PageToken != "" {
		res, _, err = bot.lark.Contact.GetUserList(context.Background(), &lark.GetUserListReq{
			DepartmentID: "0",
			PageToken:    &res.PageToken,
		})
		if err != nil {
			panic(err)
		}
		userList = append(userList, res.Items...)
	}
	return userList
}

func (bot *FeishuBot) GetUser(openID string) *lark.GetUserResp {
	res, _, err := bot.lark.Contact.GetUser(context.Background(), &lark.GetUserReq{
		UserID:     openID,
		UserIDType: lark.IDTypePtr(lark.IDTypeOpenID),
	})
	if err != nil {
		panic(err)
	}
	return res
}
