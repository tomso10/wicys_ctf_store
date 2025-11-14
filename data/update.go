package data

import "log"

func UpdateUsers() {
	for _, user := range Users {
		_, err := User.Update(user.Name, user.Password, user.Team)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func UpdateItems() {
	for _, item := range Items {
		_, err := Item.Update(item.Name, item.Description, item.Image, item.Price)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func UpdateTokens() {
	for _, item := range Tokens {
		_, err := Token.Update(item.Value, item.Type)
		if err != nil {
			log.Fatal(err)
		}
	}
}
