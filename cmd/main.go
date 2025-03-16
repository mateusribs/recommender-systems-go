package main

import (
	"fmt"
	"github.com/mateusribs/recommender-systems-go/internal/algorithms"
	"github.com/mateusribs/recommender-systems-go/internal/database"
)


func main() {
    fmt.Println("Olá, Golang!")

	ratings := database.LoadRatings()
	// movies := database.LoadMovies()

	n_ratings := ratings.Nrow()
	n_users := len(ratings.GroupBy("userId").GetGroups())
	n_movies := len(ratings.GroupBy("movieId").GetGroups())

	fmt.Printf("Ratings total: %d, Movies total: %d, Users total: %d\n", n_ratings, n_movies, n_users)

	
	uMatrix := database.NewUtilityMatrix(&ratings)

	fmt.Println(uMatrix.FindSimilarMovies(15, 10))

	dataset := [][]float64{
		{2.7810836,2.550537003},
		{1.465489372,2.362125076},
		{3.396561688,4.400293529},
		{1.38807019,1.850220317},
		{3.06407232,3.005305973},
		{7.627531214,2.759262235},
		{5.332441248,2.088626775},
		{6.922596716,1.77106367},
		{8.675418651,-0.242068655},
		{7.673756466,3.508563011},
	}

	knn := algorithms.KNN{}

	fmt.Println(knn.GetNeighbors(dataset, dataset[0], 7))


}