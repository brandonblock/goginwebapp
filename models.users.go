package main

import "errors"

type user struct {
	Username string `json:"username"`
	Password string `json:"-"`
}

var userList = []user{
	user{Username: "HopFrog", Password: "1234"},
	user{Username: "GoldBug", Password: "1234"},
	user{Username: "TheSphinx", Password: "1234"},
}

func registerNewUser(username, password string) (*user, error) {
	return nil, errors.New("placeholder error")
}

func isUsernameAvailable(username string) bool {
	return false
}
