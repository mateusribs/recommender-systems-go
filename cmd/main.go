package main

import (
	"fmt"
	"github.com/mateusribs/recommender-systems-go/internal/database"
)


func main() {
    fmt.Println("Olá, Golang!")

	ratings := database.LoadRatings()
	movies := database.LoadMovies().Select([]int{0, 1}).Maps()

	n_ratings := ratings.Nrow()
	n_users := len(ratings.GroupBy("userId").GetGroups())
	n_movies := len(ratings.GroupBy("movieId").GetGroups())

	fmt.Printf("Ratings total: %d, Movies total: %d, Users total: %d\n", n_ratings, n_movies, n_users)

	
	uMatrix := database.NewUtilityMatrix(&ratings)

	recommendations := uMatrix.FindSimilarMovies(1, 10, movies)

	fmt.Println(recommendations)
}