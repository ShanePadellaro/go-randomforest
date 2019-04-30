package main

import (
	"fmt"
	"math/rand"
	"time"

	"go-test/decisionTree"

	"github.com/kniren/gota/dataframe"
	"github.com/kniren/gota/series"
)

type TreeEnsemble struct {
	x, y                             dataframe.DataFrame
	numberTress, sampleSize, minLeaf int
	trees                            []*decisionTree.DecisionTree
}

func New(x dataframe.DataFrame, y dataframe.DataFrame, numberTress int, sampleSize int, minLeaf int) *TreeEnsemble {
	ensemble := new(TreeEnsemble)
	ensemble.y = y
	ensemble.x = x
	ensemble.numberTress = numberTress
	ensemble.sampleSize = sampleSize
	ensemble.minLeaf = minLeaf
	ensemble.trees = make([]*decisionTree.DecisionTree, 0)

	for i := 0; i < numberTress; i++ {
		ensemble.trees = append(ensemble.trees, ensemble.CreateTree())
	}

	return ensemble
}

func (ensemble *TreeEnsemble) CreateTree() *decisionTree.DecisionTree {
	indexes := make([]int, 0)
	rows := ensemble.y.Nrow()
	for j := 0; j < rows; j++ {
		indexes = append(indexes, j)
	}
	shuffle(indexes)
	rndIndexes := indexes[:ensemble.sampleSize]

	tree := decisionTree.New(ensemble.x.Subset(rndIndexes), ensemble.y.Subset(rndIndexes), ensemble.minLeaf)
	return &tree
}

func shuffle(indexes []int) {
	d := len(indexes)
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(d, func(i, j int) { indexes[i], indexes[j] = indexes[j], indexes[i] })
}

func main() {

	df := dataframe.New(
		series.New([]string{"b", "a", "c", "d"}, series.String, "COL.1"),
		series.New([]int{1, 2, 3, 4}, series.Int, "COL.2"),
		series.New([]float64{3.0, 4.0, 5.0, 6.0}, series.Float, "COL.3"),
	)

	randomForest := New(df, df, 5, 2, 2)

	fmt.Println(df)
	fmt.Println(randomForest)

}
