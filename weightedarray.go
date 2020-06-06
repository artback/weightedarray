package weightedarray

import (
	"math/rand"
	"sort"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

type Value interface {
	GetWeight() float64
}
type Values []Value

func (values Values) getRandomIndex() int {
	totals := make([]int, len(values))
	runningTotal := 0
	for i, v := range values {
		runningTotal += int(v.GetWeight())
		totals[i] = runningTotal
	}
	r := rand.Intn(runningTotal) + 1
	return sort.SearchInts(totals, r)
}
func (values Values) Sort() {
	sort.Slice(values, func(i, j int) bool {
		return values[i].GetWeight() < values[j].GetWeight()
	})
}

func (values Values) Random() interface{} {
	values.Sort()
	return values[values.getRandomIndex()]
}

func (values Values) Probability(value Value) float64 {
	var totalWeight float64
	for _, v := range values {
		totalWeight += v.GetWeight()
	}
	return value.GetWeight() / totalWeight
}
