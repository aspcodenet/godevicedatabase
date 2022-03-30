package main

type Device struct {
	Name          string `json:"name"`
	WeightInGrams int    `json:"weight"`
	Manufacturer  string
}
