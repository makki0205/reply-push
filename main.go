package main

import "github.com/makki0205/reply-push/router"

func main() {
	r := router.GetRouter()
	r.Run(":8000")
}
