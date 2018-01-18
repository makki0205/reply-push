package cash

import "github.com/makki0205/reply-push/model"

var account = loadAccount()

func loadAccount() []model.Account {
	var a []model.Account
	db.Find(&a)
	return a
}
func ReloadAccount() {
	account = loadAccount()
}
func GetAccount() []model.Account {
	return account
}
