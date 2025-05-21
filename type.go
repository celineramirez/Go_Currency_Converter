package main

type Currency struct {
	Date  int                `json:"date"`
	Base  string             `json:"base"`
	Rates map[string]float64 `json:"rates"`
}