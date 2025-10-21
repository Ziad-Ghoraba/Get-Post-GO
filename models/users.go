package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

type User struct {
	Id      int
	Name    string
	Age     int
	Address Address
}

var (
	usersDir = "users" // Folder to store all user files
	nextId   = 1
)

func init() {
	os.MkdirAll(usersDir, os.ModePerm)
}

func GetUserById(id int) (User, error) {
	filename := filepath.Join(usersDir, fmt.Sprintf("%d.json", id))

	data, err := os.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return User{}, fmt.Errorf("user with id %d not found", id)
		}
		return User{}, err
	}

	var u User
	err = json.Unmarshal(data, &u)
	if err != nil {
		return User{}, err
	}

	return u, nil
}

func AddUser(u User) (User, error) {
	if u.Id != 0 {
		return User{}, errors.New("User id must not include id or it must be set to zero")
	}
	u.Id = nextId
	nextId++
	filename := filepath.Join(usersDir, fmt.Sprintf("%d.json", u.Id))
	// Save user to file
	data, err := json.MarshalIndent(u, "", "  ")
	if err != nil {
		return User{}, err
	}

	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		return User{}, err
	}

	return u, nil
}
