package algorithms

import (
	"math"
	"sort"
)


type KNN struct {}

type Neighbor struct{
	index int
	distance float64
}

func (k KNN) computeDistance(p1, p2 []float64) float64 {
	if len(p1) != len(p2) {
		panic("Vectors must have the same length")
	}

	var sum float64

	for i := range p1 {
		diff := p1[i] - p2[i]
		sum += diff * diff
	}

	return math.Sqrt(sum)
}

func (knn KNN) GetNeighbors(dataset [][]float64, data []float64, k int) []int {
	var distances []Neighbor
	
	for i, element := range dataset {
		distances = append(distances, Neighbor{
			index: i,
			distance: knn.computeDistance(data, element),
		})
	}

	sort.Slice(distances, func(i, j int) bool {
		return distances[i].distance < distances[j].distance
	})

	neighbors := make([]int, k)
	for i := range k {
		neighbors[i] = distances[i].index
	}

	return neighbors
}

