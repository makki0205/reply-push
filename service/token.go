package service

import (
	"github.com/makki0205/reply-push/cash"
	"github.com/makki0205/reply-push/model"
)

var Token = tokenimpl{}

type tokenimpl struct {
}

func (self *tokenimpl) Exists(token string) bool {
	tokens := cash.GetTokent()
	for _, value := range tokens {
		if token == value.Token {
			return true
		}
	}
	return false
}

func (self *tokenimpl) Set(token string) {
	db.Create(&model.Token{Token: token})
}
