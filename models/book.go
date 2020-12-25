package models

import ()

type Book struct {
	ID     uint   `json:"id" grom:"primary_key"`
	Title  string `json:"title"`
	Author string `json:"author"`
}
