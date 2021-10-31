package main

import (
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
)

func main() {
	userId := 10
	username, err := QueryUserName(userId)
	if err == sql.ErrNoRows {
		fmt.Println(err)
	}
	fmt.Println(username)
}
func QueryUserName(id int) (string, error) {
	userName := ""
	err := db.QueryRow("select user_name from users where id=?", id).Scan(&userName)
	return userName, errors.Wrap(err, "error when get user name")
}
