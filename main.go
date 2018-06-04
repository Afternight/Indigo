package main

import (

	"github.com/gin-gonic/gin"
	"strconv"
	"encoding/json"
)

const JsonByteStreamHeader = "application/json; charset=utf-8"

const FloatType = "float"
const StringType = "string"
const BinaryType = "binary"

func main(){
	r := gin.Default()

	r.POST("/learn", GetLearningRequest)
	r.POST("/learnIonosphere", GetIonosphereLearningRequest)
	r.POST("/learnAutos",GetAutosLearningRequest)

	r.Run(":3000")
}

func SendResponse(c *gin.Context,resp *LearningResponse){
	bodyBytes, _ := json.Marshal(resp)
	c.Header("Content-Length", strconv.Itoa(len(bodyBytes)))
	c.Data(200,JsonByteStreamHeader,bodyBytes)
}

func GetLearningRequest(c *gin.Context) {
	req := new(JSONLearningRequest)
	handler(req,c)
}

func GetIonosphereLearningRequest(c *gin.Context){
	req := new(IonosphereLearningRequest)
	handler(req,c)
}

func GetAutosLearningRequest(c *gin.Context){
	req := new(AutosLearningRequest)
	handler(req,c)
}

func handler(req LearningRequestInterface, c *gin.Context) {
	req.ReceiveRequest(c)
	resp := ProcessLearningRequest(req)
	SendResponse(c,resp)
}
