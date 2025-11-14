package structs

import "gitlab.ritsec.cloud/competitions/ists-2023/store/ent"

type AdminInfo struct {
	Users     []*ent.User
	Usernames []string
	UserMap   map[string]*ent.User
	Functions []Function
	Purchases []*ent.Transaction
}
