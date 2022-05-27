package models

import (
	"errors"
	"fmt"
)

type User struct {
	// By default, each field is initialised to its zero value
	ID         int    // 0
	FirstName  string // empty
	SecondName string // empty
}

// Variable block
var (
	// Slice containing pointers to users.
	users  []*User
	nextID = 1 // Implicit declaration as int (colon not needed at package level)
)

func GetUsers() []*User {
	return users
}

func AddUser(u User) (User, error) {
	if u.ID != 0 {
		return User{}, errors.New("cannot include a user with an existing ID")
	}
	u.ID = nextID
	nextID++
	users = append(users, &u)
	return u, nil
}

func userIndex(id int) (int, error) {
	for i, u := range users {
		if u.ID == id {
			return i, nil
		}
	}

	return 0, fmt.Errorf("User with ID '%v' not found", id)
}

func GetUser(id int) (User, error) {
	index, err := userIndex(id)
	if err != nil {
		return User{}, err
	}

	return *users[index], nil
}

func UpdateUser(u User) (User, error) {
	index, err := userIndex(u.ID)
	if err != nil {
		return User{}, err
	}

	users[index] = &u
	return u, nil
}

func RemoveUser(id int) error {
	index, err := userIndex(id)
	if err != nil {
		return err
	}
	// Removing using a splicing approach.
	users = append(users[:index], users[index+1:]...)
	return nil
}
