package main

import (
	"github.com/gin-gonic/gin"
	"encoding/json"
)

//region JSON requests / responses

type JSONLearningRequest struct {
	BaseLearningRequest
	Targets []JSONPoint
	Data []JSONPoint
}

func (g *JSONLearningRequest) ReceiveRequest(c *gin.Context) error {
	jsonDecodeErr := json.NewDecoder(c.Request.Body).Decode(&g)

	if jsonDecodeErr != nil {
		return jsonDecodeErr
	}
	return nil
}

func (g *JSONLearningRequest) GetBaseRequest() LearningRequestInterface {
	return g
}

func (g *JSONLearningRequest) GetTargetsLength() int {
	return len(g.Targets)
}

func (g *JSONLearningRequest) GetTargetByIndex(i int) DataPoint {
	return &g.Targets[i]
}

func (g *JSONLearningRequest) GetDataLength() int {
	return len(g.Data)
}

func (g *JSONLearningRequest) GetDataByIndex(i int) DataPoint {
	return &g.Data[i]
}

//endregion

//region GenericJSON Point

//json format of a datapoint to be compatible with the object map input
type JSONPoint struct {
	Values map[string]Attribute //map of the name of the attribute to the attribute value itself
}

func (g *JSONPoint) GetValue(toGet string) Attribute {
	return g.Values[toGet]
}

func (g *JSONPoint) SetValue(name string, toSet Attribute) error {
	g.Values[name] = toSet
	return nil
}

func (g *JSONPoint) GetAttributeValues(target string) map[string]Attribute {
	var dup map[string]Attribute
	dup = g.Values
	delete(dup,target)
	return dup
}

//endregion


