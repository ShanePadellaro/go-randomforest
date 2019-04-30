package decisionTree

import (
	"github.com/kniren/gota/dataframe"
)

type DecisionTree struct {
	x, y    dataframe.DataFrame
	minLeaf int
}

func New(x dataframe.DataFrame, y dataframe.DataFrame, minLeaf int) DecisionTree {

	return DecisionTree{x: x, y: y, minLeaf: minLeaf}
}
