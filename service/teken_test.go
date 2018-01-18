package service

import (
	"fmt"
	"testing"

	"github.com/makki0205/reply-push/cash"
)

func TestTokenimpl_Exists(t *testing.T) {
	Token.Set("hoge01")
	fmt.Println(cash.GetTokent())
	cash.ReloadTokent()
	fmt.Println(cash.GetTokent())

}
