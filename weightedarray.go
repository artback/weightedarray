package weightedarray

import (
	"fmt"
	"github.com/gonum/floats"
	"gonum.org/v1/gonum/stat/distuv"
	"math/rand"
	"sort"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}
func getWeights(arr []struct{}, weightFn func(interface{}) float64) ([]float64, error) {
	var weights []float64
	for _, val := range *a {
		weights = append(weights, weightFn(val))
	}
	if weights == nil || len(weights) == 0 {
		return nil, fmt.Errorf("invalid probability weights: %v", weights)
	}
	return weights, nil
}
func random(arr []struct{}, weightFn func(interface{}) float64) (interface{}, error) {
	weights, err := getWeights(a,weightFn)
	if err != nil {
		return nil, err
	}
	cdf := make([]float64, len(weights))
	floats.CumSum(cdf, weights)
	var val = distuv.UnitUniform.Rand() * cdf[len(cdf)-1]
	return sort.Search(len(cdf), func(i int) bool { return cdf[i] > val }), nil
}
