package cash

import "github.com/makki0205/reply-push/model"

var token = loadToken()

func loadToken() []model.Token {
	var a []model.Token
	db.Find(&a)
	return a
}
func ReloadTokent() {
	token = loadToken()
}
func GetTokent() []model.Token {
	return token
}
