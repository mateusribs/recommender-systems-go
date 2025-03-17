package database

import (
	"slices"
	"sort"

	"github.com/go-gota/gota/dataframe"
	"github.com/mateusribs/recommender-systems-go/internal/algorithms"
)


type UtilityMatrix struct {
	df* dataframe.DataFrame
	UserMapper map[int]int
	UserMapperInv map[int]int
	MovieMapper map[int]int
	MovieMapperInv map[int]int
	Matrix [][]float64
}


func (u *UtilityMatrix) getUniqueIds (col string) []int {
	ids, err := u.df.Col(col).Int()

	sort.Ints(ids)

	if err != nil {
		panic("Error during Series transformation")
	}

	return slices.Compact(ids)
}


func (u *UtilityMatrix) setMapper (ids []int) map[int]int {
	N := len(ids)
	mapper := make(map[int]int)

	for i := 0; i < N; i++ {
		k := ids[i]
		mapper[k] = i
	}

	return mapper
}

func (u *UtilityMatrix) setMapperInv (ids []int) map[int]int {
	N := len(ids)
	mapper := make(map[int]int)

	for i := 0; i < N; i++ {
		k := ids[i]
		mapper[i] = k
	}

	return mapper
}

func (u *UtilityMatrix) createSparseMatrix (userIds []int, movieIds []int) [][]float64 {
	rows := len(movieIds)
	cols := len(userIds)

	matrix := make([][]float64, rows)

	for i := range matrix {
		matrix[i] = make([]float64, cols)
		for j := range matrix[i] {
			matrix[i][j] = 0
		}
	}

	data := u.df.Select([]string{"movieId", "userId", "rating"}).Maps()
	
	for _, row := range data {
		userId, _ := row["userId"].(int)
		movieId, _ := row["movieId"].(int)
		n := u.UserMapper[userId]
		m := u.MovieMapper[movieId]
		if rating, ok := row["rating"].(float64); ok {
			matrix[m][n] = rating
		}
	}

	return matrix
}

func (u *UtilityMatrix) FindSimilarMovies(movie_id int, k int, movies []map[string]interface{}) []string {
	knn := algorithms.KNN{}

	movie_ind := u.MovieMapper[movie_id]
	movie_vec := u.Matrix[movie_ind]

	neighbors := knn.GetNeighbors(u.Matrix, movie_vec, k + 1)
	neighbors = neighbors[1:]

	moviesTitles := make(map[int]string)

	for _, mapper := range movies {
		
		movieId, _ := mapper["movieId"].(int)
		movieTitle, _ := mapper["title"].(string)
		moviesTitles[movieId] = movieTitle
	}

	neighborsTitles := make([]string, len(neighbors))

	for i, movieId := range neighbors {
		id := u.MovieMapperInv[movieId]
		neighborsTitles[i] = moviesTitles[id]
	}

	return neighborsTitles
}


func NewUtilityMatrix(data *dataframe.DataFrame) *UtilityMatrix {

	uMatrix := &UtilityMatrix{df: data}

	uniqueMoviesIds := uMatrix.getUniqueIds("movieId")
	uniqueUsersIds := uMatrix.getUniqueIds("userId")

	uMatrix.UserMapper = uMatrix.setMapper(uniqueUsersIds)
	uMatrix.MovieMapper = uMatrix.setMapper(uniqueMoviesIds)
	uMatrix.UserMapperInv = uMatrix.setMapperInv(uniqueUsersIds)
	uMatrix.MovieMapperInv = uMatrix.setMapperInv(uniqueMoviesIds)
	uMatrix.Matrix = uMatrix.createSparseMatrix(uniqueUsersIds, uniqueMoviesIds)

	return uMatrix
}