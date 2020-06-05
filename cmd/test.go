package cmd

import (
	"github.com/artback/weightedarray"
	"log"
)

type Fruit struct {
	Name   string
	Age    float64
	Weight float64
	Draws  int
}

func main() {
	items := []Fruit{
		Fruit{Name: "Banana", Weight: 1, Age: 5},
		Fruit{Name: "Lemon", Weight: 2, Age: 1},
		Fruit{Name: "Apple", Weight: 4, Age: 2},
		Fruit{Name: "Mellon", Weight: 8, Age: 2},
		Fruit{Name: "Orange", Weight: 10, Age: 10},
		Fruit{Name: "Pineapple", Weight: 7, Age: 10},
	}

	getWeights := func(fruit Fruit) float64 {
		return fruit.Age + fruit.Weight
	}
	var p []float64
	for i, _ := range []int{0, 1, 2, 3, 4, 5} {
		p = append(p, items[i].Weight)
	}

	draw, err := weightedarray.random(items, getWeights)
	if err != nil {
		log.Fatalf("Failed to generate Roulette samples: %v", err)
	}

	for k, _ := range items {
		fmt.Printf("Weight: %.2f, Draws: %d\n", items[k].Weight, items[k].Draws)
	}
}
