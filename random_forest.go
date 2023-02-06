package main

import (
	"fmt"

	"github.com/sjwhitworth/golearn/base"
	"github.com/sjwhitworth/golearn/datasets"
	"github.com/sjwhitworth/golearn/evaluation"
	"github.com/sjwhitworth/golearn/tree"
)

func main() {
	iris, err := datasets.LoadIrisDataset()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Load the iris dataset into the Instances structure
	irisInstances := base.NewDenseInstances()
	irisInstances.AddAttr("sepal_length", base.AttributeContinuous)
	irisInstances.AddAttr("sepal_width", base.AttributeContinuous)
	irisInstances.AddAttr("petal_length", base.AttributeContinuous)
	irisInstances.AddAttr("petal_width", base.AttributeContinuous)
	irisInstances.AddClassAttribute("class", base.AttributeClass, []string{"setosa", "versicolor", "virginica"})

	for i := 0; i < len(iris); i++ {
		irisInstances.AddInstance(base.NewDenseInstance(iris[i][:4], iris[i][4:]))
	}

	// Create a Random Forest model
	rf := tree.NewRandomForest(10, 4)

	// Train the model on the iris dataset
	rf.Fit(irisInstances)

	// Evaluate the model using cross-validation
	cv, err := evaluation.GenerateCrossFoldValidationConfusionMatrix(irisInstances, rf, 10)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(evaluation.GetSummary(cv))
}
