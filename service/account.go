package service

import (
	"github.com/makki0205/reply-push/cash"
	"github.com/makki0205/reply-push/model"
)

var Account = accountimpl{}

type accountimpl struct {
}

func (self *accountimpl) GetAll(token string) []model.Account {
	return cash.GetAccount()
}
