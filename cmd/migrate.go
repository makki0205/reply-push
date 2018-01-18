package main

import "github.com/makki0205/reply-push/model"

func main() {
	db := model.GetDBConn()

	db.DropTableIfExists(&model.Token{})
	db.DropTableIfExists(&model.Account{})
	db.DropTableIfExists(&model.Tweet{})

	db.AutoMigrate(&model.Token{})
	db.AutoMigrate(&model.Account{})
	db.AutoMigrate(&model.Tweet{})
}
