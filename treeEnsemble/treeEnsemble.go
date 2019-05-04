package treeEnsemble

import (
	"math/rand"
	"time"

	"go-test/decisionTree"

	"github.com/kniren/gota/dataframe"
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

	tree := decisionTree.New(ensemble.x.Subset(rndIndexes), ensemble.y.Subset(rndIndexes), ensemble.minLeaf, nil)
	return &tree
}

func (ensemble *TreeEnsemble) Predict() dataframe.DataFrame {
	predictions := make([]dataframe.DataFrame, 0)

	for _, tree := range ensemble.trees {
		predictions = append(predictions, tree.Predict())

	}

	for i := 0; i < ensemble.y.Nrow(); i++ {

	}
	return dataframe.DataFrame{}
}

func shuffle(indexes []int) {
	d := len(indexes)
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(d, func(i, j int) { indexes[i], indexes[j] = indexes[j], indexes[i] })
}
