package database

import "os"
import "log"
import "github.com/go-gota/gota/dataframe"

func LoadRatings() dataframe.DataFrame {
	ratingsFile, err := os.Open("data/ratings.csv")

	if err != nil {
		log.Fatal(err)
	}

	defer ratingsFile.Close()

	return dataframe.ReadCSV(ratingsFile)
}


func LoadMovies() dataframe.DataFrame {
	moviesFile, err := os.Open("data/movies.csv")

	if err != nil {
		log.Fatal(err)
	}

	defer moviesFile.Close()

	return dataframe.ReadCSV(moviesFile)
}