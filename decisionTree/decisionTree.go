package decisionTree

import (
	"github.com/kniren/gota/dataframe"
)

type DecisionTree struct {
	x, y          dataframe.DataFrame
	minLeaf, n, c int
	idxs          []int
}

func New(x dataframe.DataFrame, y dataframe.DataFrame, minLeaf int, idxs []int) DecisionTree {
	if idxs == nil {
		idxs = make([]int, 0)
		for i := 0; i < y.Nrow(); i++ {
			idxs = append(idxs, i)
		}
	}

	tree := DecisionTree{x: x, y: y, minLeaf: minLeaf, idxs: idxs}
	tree.c = len(idxs)
	tree.n = x.Ncol()

	return tree
}

func (tree *DecisionTree) Predict() dataframe.DataFrame {
	return dataframe.DataFrame{}
}
