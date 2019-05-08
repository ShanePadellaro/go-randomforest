package main

import (
	"bufio"
	"fmt"
	"go-test/treeEnsemble"
	"os"

	"github.com/kniren/gota/dataframe"
	"github.com/kniren/gota/series"
)

func main() {

	// df := dataframe.New(
	// 	series.New([]string{"b", "a", "c", "d"}, series.String, "COL.1"),
	// 	series.New([]int{1, 2, 3, 4}, series.Int, "COL.2"),
	// 	series.New([]float64{3.0, 4.0, 5.0, 6.0}, series.Float, "COL.3"),
	// )

	// randomForest := treeEnsemble.New(df, df, 5, 2, 2)

	// fmt.Println(df)
	// fmt.Println(randomForest)

	// fmt.Println(df.Subset(1))

	// mean := func(s series.Series) series.Series {
	// 	floats := s.Float()
	// 	sum := 0.0
	// 	for _, f := range floats {
	// 		sum += f
	// 	}
	// 	return series.Floats(sum / float64(len(floats)))
	// }

	// col := df.Select([]int{1})
	// meanValues := col.Capply(mean)

	// fmt.Println(meanValues)
	// fmt.Println(meanValues.Elem(0, 0))

	file, err := os.Open("Train.csv")
	if err != nil {
		panic("bad bad")
	}
	reader := bufio.NewReader(file)

	df := dataframe.ReadCSV(reader)
	y := df.Select([]string{"SalePrice"})
	x := df.Select([]string{"YearMade", "MachineHoursCurrentMeter"})
	// fmt.Println(df.Subset([]int{1}))
	fmt.Println(x.Describe())
	fmt.Println(y.Describe())

	noNan := func(s series.Series) series.Series {
		for i, isNan := range s.IsNaN() {
			if isNan {
				s.Set([]int{i}, series.Floats([]float64{-1}))
			}
		}
		return s
	}

	x.Capply(noNan)
	y.Capply(noNan)
	fmt.Println(x.Describe())
	fmt.Println(y.Describe())

	randomForest := treeEnsemble.New(x, y, 1, 1000, 5)
	fmt.Println(randomForest)

}
