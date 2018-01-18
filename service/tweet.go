package service

import "github.com/makki0205/reply-push/model"

var Tweet = tweettimpl{}

type tweettimpl struct {
}

func (self *tweettimpl) Set(tweet string) model.Tweet {
	var t = model.Tweet{
		Body: tweet,
		Done: false,
	}
	db.Create(&t)
	return t
}
func (self *tweettimpl) GetFirst() model.Tweet {
	var tweet model.Tweet

	db.Where("done = ?", false).First(&tweet)
	tweet.Done = true
	db.Update(&tweet)

	return tweet
}

func (self *tweettimpl) RolBack(id uint) {
	var tweet model.Tweet
	db.First(&tweet, id)
	tweet.Done = false
	db.Update(&tweet)

}
