package model

type DoModify struct {
	Tag     string `json:"tag"`
	Data    string `json:"value"`
	Operate string `json:"operate"`
	Leaf    string `json:"leaf"`
}
