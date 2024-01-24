package store

import (
	"errors"
	"math/rand"
	"time"

	"github.com/todoService/model"
)

const (
	userID   = "userid"
	title    = "title"
	desc     = "discription"
	priority = "priority"
	status   = "status"
)

type Error struct {
	stage   string
	error   error
	message string
}

func GenerateID() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(90000) + 10000
}

func QueryGenerator(filter model.Filter) (string, error) {
	query := "select * from todo where "

	if val, exists := filter[userID]; exists {
		query = query + userID + "=" + val
	} else if val, exists := filter[title]; exists {
		query = query + title + "=" + val
	} else if val, exists := filter[desc]; exists {
		query = query + desc + "=" + val
	} else if val, exists := filter[priority]; exists {
		query = query + priority + "=" + val
	} else if val, exists := filter[status]; exists {
		query = query + status + "=" + val
	} else {
		return "", errors.New("Invalid Filter")
	}

	return query, nil
}
