package users

import "fmt"

type User struct {
	Sub               string `json:"sub"`
	Name              string `json:"name"`
	GivenName         string `json:"given_name"`
	FamilyName        string `json:"family_name"`
	PreferredUsername string `json:"preferred_username"`
	Email             string `json:"email"`
	Picture           string `json:"picture"`
}

func Auth(login, password, mfa string) (bool, User, error) {
	if login == "omar" && password == "password" {
		return true, GetAllUsers()[0], nil
	}
	return false, User{}, fmt.Errorf("Invalid login or passwordo")
}

func GetAllUsers() []User {
	return []User{
		{
			Sub:               "9-9-9-9",
			Name:              "omar",
			GivenName:         "Omar",
			FamilyName:        "hm",
			PreferredUsername: "omar",
			Email:             "omar@domain.inv",
		},
	}
}
