package main

import (
	"fmt"
	"go-test/treeEnsemble"

	"github.com/kniren/gota/dataframe"
	"github.com/kniren/gota/series"
)

func main() {

	df := dataframe.New(
		series.New([]string{"b", "a", "c", "d"}, series.String, "COL.1"),
		series.New([]int{1, 2, 3, 4}, series.Int, "COL.2"),
		series.New([]float64{3.0, 4.0, 5.0, 6.0}, series.Float, "COL.3"),
	)

	randomForest := treeEnsemble.New(df, df, 5, 2, 2)

	fmt.Println(df)
	fmt.Println(randomForest)

	fmt.Println(df.Subset(1))

}
