package model

type Todo struct {
	Id          string `json:"id"`
	UserID      string `json:"userid"`
	Title       string `json:"title"`
	Discription string `json:"discription"`
	Priority    string `json:"priority"`
	Status      string `json:"status"`
}

type Filter map[string]string
