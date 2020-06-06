package main

import (
	"github.com/artback/weightedarray"
	"log"
)

type Fruit struct {
	Name   string
	Weight float64
	Draws  int
}

func (f Fruit) GetWeight() float64 {
	return f.Weight
}

func main() {
	items := weightedarray.Values{
		&Fruit{Name: "Banana", Weight: 1},
		&Fruit{Name: "Lemon", Weight: 2},
		&Fruit{Name: "Apple", Weight: 4},
		&Fruit{Name: "Mellon", Weight: 8},
		&Fruit{Name: "Orange", Weight: 10},
		&Fruit{Name: "Pineapple", Weight: 7},
	}
	log.Print(items.Probability(items[4]))
}
