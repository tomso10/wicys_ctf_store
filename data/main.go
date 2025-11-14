package data

import (
	"context"
	"log"

	_ "github.com/mattn/go-sqlite3"
	"gitlab.ritsec.cloud/competitions/ists-2023/store/ent"
	"gitlab.ritsec.cloud/competitions/ists-2023/store/structs"
)

var (
	Client      *ent.Client
	Item        *item_s
	User        *user_s
	Transaction *transaction_s
	Token       *token_s

	TokenInfo structs.TokenInfo
	Users     []structs.Account
	Items     []structs.Item
	Tokens    []structs.Token

	Ctx context.Context = context.Background()
)

func init() {
	client, err := ent.Open("sqlite3", "file:databases/data.sqlite?_loc=auto&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf(err.Error())
	}

	Client = client

	// Run the auto migration tool.
	if err := Client.Schema.Create(Ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	Item = &item_s{Client: Client, Ctx: Ctx}
	User = &user_s{Client: Client, Ctx: Ctx}
	Transaction = &transaction_s{Client: Client, Ctx: Ctx}
	Token = &token_s{Client: Client, Ctx: Ctx}
}

func Close() {
	Client.Close()
}
