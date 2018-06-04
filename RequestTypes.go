package main

import (
	"github.com/gin-gonic/gin"
)

type BaseLearningRequest struct {
	K int //value for K in kNN learning algo
	TargetAttribute string //name of the target attribute
	Comparers map[string]Comparer //map of name of type to its valid comparer
}

func (g *BaseLearningRequest) GetTargetAttribute() string {
	return g.TargetAttribute
}

func (g *BaseLearningRequest) GetK() int {
	return g.K
}

func (g *BaseLearningRequest) GetComparers() map[string]Comparer {
	return g.Comparers
}

//The general structure needed for learning request interface implementations
//type LearningRequest struct {
//	BaseLearningRequest
//	Targets []DataPoint //points that need target value determined
//	Data []DataPoint //the points to be used to train the dataset
//}

type LearningRequestInterface interface {
	ReceiveRequest(*gin.Context) error
	GetBaseRequest() LearningRequestInterface

	GetTargetAttribute() string
	GetK() int
	GetComparers() map[string]Comparer

	GetTargetsLength() int
	GetTargetByIndex(int) DataPoint
	GetDataLength() int
	GetDataByIndex(int) DataPoint
}

type LearningResponse struct {
	OriginRequest LearningRequestInterface //Original request
	CompletedPoints []CompletedPoint //Array of completed points
	Error float64 //the Mean Squared Error if numeric, percentage correct if not of the predictions
}

type CompletedPoint struct {
	PredictedAttribute Attribute //The attribute that was predicted
	ExpectedAttribute Attribute //The given value of the attribute (null if not given)
	Error float64 //the percentage error between the predicted and expected attributes
	PointValue DataPoint //The full value of the point that had its target attribute predicted
	FoundNeighbours []DataPoint //the neighbours found (array is k length)
}