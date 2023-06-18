package database

import (
	"sort"
)

func (dbs *DBStruct) getUsers() []User {
	users := []User{}
	for _, user := range dbs.Users {
		users = append(users, user)
	}
	sort.Slice(users, func(i, j int) bool { return users[i].ID < users[j].ID })
	return users
}

func (dbs *DBStruct) getUserByEmail(email string) (User, bool) {
	users := dbs.getUsers()
	for _, user := range users {
		if user.Email == email {
			return user, true
		}
	}
	return User{}, false
}

func (dbs *DBStruct) isRevoked(token string) bool {
	_, exists := dbs.RevokedTokens[token]
	return exists
}
