package main

type Currency struct {
	Date  string             `json:"date"`
	Base  string             `json:"base"`
	Rates map[string]string  `json:"rates"`
}
