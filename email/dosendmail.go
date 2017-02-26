package main

import (
	"fmt"
	"libofm"
)

func main() {
	mycontent := "Dear my golang"

	email := libofm.NewEmail("seafooler@hust.edu.cn", "test golang email", mycontent)

	err := libofm.SendEmail(email)

	fmt.Println(err)
}
