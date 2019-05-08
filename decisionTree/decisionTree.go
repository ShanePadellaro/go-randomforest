package decisionTree

import (
	"math"

	"github.com/kniren/gota/dataframe"
	"github.com/kniren/gota/series"
)

type DecisionTree struct {
	x, y          dataframe.DataFrame
	minLeaf, n, c int
	idxs          []int
	val, Score    float64
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

	ss := y.Subset(idxs)
	column := ss.Select([]int{0})

	mean := func(s series.Series) series.Series {
		floats := s.Float()
		sum := 0.0
		for _, f := range floats {
			sum += f
		}
		return series.Floats(sum / float64(len(floats)))
	}

	tree.val = column.Capply(mean).Elem(0, 0).Float()
	tree.Score = math.Inf(1)
	tree.findVarSplit()
	return tree
}

func (tree *DecisionTree) Predict() dataframe.DataFrame {
	return dataframe.DataFrame{}
}

func (tree *DecisionTree) findVarSplit() {

}
