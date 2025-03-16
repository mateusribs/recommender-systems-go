package database

import (
	"fmt"
	"slices"
	"sort"

	"github.com/go-gota/gota/dataframe"
	"github.com/james-bowman/sparse"
)


type UtilityMatrix struct {
	df* dataframe.DataFrame
	UserMapper map[int]int
	MovieMapper map[int]int
	Matrix *sparse.CSR
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

func (u *UtilityMatrix) createSparseMatrix (userIds []int, movieIds []int) *sparse.CSR {
	rows := len(u.UserMapper)
	cols := len(u.MovieMapper)

	data := u.df.Col("rating").Float()

	return sparse.NewCSR(rows, cols, userIds, movieIds, data)
}

func (u *UtilityMatrix) FindSimilarMovies(movie_id int, k int) []int {
	neighbors_ids := make([]int, k + 1)

	movie_ind := u.MovieMapper[movie_id]
	movie_vec := u.Matrix.RowView(movie_ind)

	fmt.Println(movie_ind, movie_vec)

	return neighbors_ids
}


func NewUtilityMatrix(df* dataframe.DataFrame) *UtilityMatrix {

	uMatrix := &UtilityMatrix{df: df}

	uniqueMoviesIds := uMatrix.getUniqueIds("movieId")
	uniqueUsersIds := uMatrix.getUniqueIds("userId")

	uMatrix.UserMapper = uMatrix.setMapper(uniqueUsersIds)
	uMatrix.MovieMapper = uMatrix.setMapper(uniqueMoviesIds)
	uMatrix.Matrix = uMatrix.createSparseMatrix(uniqueUsersIds, uniqueMoviesIds)

	return uMatrix
}