package tttt

import (
	"fmt"
	"testing"
)

func TestQueryExample(t *testing.T) {
	u := Userinfo{
		Userid: 1,
		Phone:  1,
		Email:  "123@qq.com",
		Passwd: "1",
	}
	res, err := u.QueryExample()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}
