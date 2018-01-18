package service

import (
	"github.com/ChimeraCoder/anaconda"
	"github.com/makki0205/reply-push/model"
)

func getTwitterApi(a model.Account) *anaconda.TwitterApi {
	anaconda.SetConsumerKey(a.CustomerKey)
	anaconda.SetConsumerSecret(a.CustomerSecret)
	api := anaconda.NewTwitterApi(a.AccessToken, a.AccessSecret)
	return api
}

func TweetAPI(a model.Account, body string) error {
	api := getTwitterApi(a)
	_, err := api.PostTweet(body, nil)
	if err != nil {
		return err
	}
	return nil
}
