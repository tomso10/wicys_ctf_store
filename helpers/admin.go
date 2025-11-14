package helpers

import (
	"gitlab.ritsec.cloud/competitions/ists-2023/store/data"
	"gitlab.ritsec.cloud/competitions/ists-2023/store/ent"
	"gitlab.ritsec.cloud/competitions/ists-2023/store/structs"
)

func NewAdminInfo() *structs.AdminInfo {
	adminInfo := &structs.AdminInfo{}

	adminInfo.Users, _ = data.User.GetAll()
	adminInfo.UserMap = make(map[string]*ent.User)
	adminInfo.Usernames = []string{}
	adminInfo.Purchases, _ = data.Transaction.GetAll()
	adminInfo.Functions = []structs.Function{
		{
			Name: "Add Balance",
			ID:   "ADD",
		},
		{
			Name: "Subtract Balance",
			ID:   "SUB",
		},
		{
			Name: "Set Balance",
			ID:   "SET",
		},
	}

	for _, user := range adminInfo.Users {
		adminInfo.Usernames = append(adminInfo.Usernames, user.Username)
		adminInfo.UserMap[user.Username] = user
	}

	return adminInfo
}
