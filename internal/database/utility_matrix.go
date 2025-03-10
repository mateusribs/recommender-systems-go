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
	userMapper map[int]int
	movieMapper map[int]int
}


func (u UtilityMatrix) getUniqueIds (col string) []int {
	ids, err := u.df.Col(col).Int()

	sort.Ints(ids)

	if err != nil {
		panic("Error during Series transformation")
	}

	return slices.Compact(ids)
}


func (u UtilityMatrix) setMapper (ids []int) map[int]int {
	N := len(ids)
	mapper := make(map[int]int)

	for i := 0; i < N; i++ {
		k := ids[i]
		mapper[k] = i
	}

	return mapper
}

func (u UtilityMatrix) createSparseMatrix (userIds []int, movieIds []int) {
	rows := len(u.userMapper)
	cols := len(u.movieMapper)

	data := u.df.Col("rating").Float()

	csrMatrix := sparse.NewCSR(rows, cols, userIds, movieIds, data)

	fmt.Println(csrMatrix)

}


func NewUtilityMatrix(df* dataframe.DataFrame) UtilityMatrix {

	uMatrix := UtilityMatrix{df: df}

	uniqueMoviesIds := uMatrix.getUniqueIds("movieId")
	uniqueUsersIds := uMatrix.getUniqueIds("userId")

	uMatrix.userMapper = uMatrix.setMapper(uniqueUsersIds)
	uMatrix.movieMapper = uMatrix.setMapper(uniqueMoviesIds)

	uMatrix.createSparseMatrix(uniqueUsersIds, uniqueMoviesIds)

	return uMatrix
}