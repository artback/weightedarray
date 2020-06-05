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

func (values Values) Random() interface{} {
	sort.Slice(values, func(i, j int) bool {
		return values[i].GetWeight() < values[j].GetWeight()
	})
	totals := make([]int, len(values))
	runningTotal := 0
	for i, v := range values {
		runningTotal += int(v.GetWeight())
		totals[i] = runningTotal
	}
	r := rand.Intn(runningTotal) + 1
	i := sort.SearchInts(totals, r)
	return values[i]

}
