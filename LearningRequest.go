package main

import (
	"errors"
	"github.com/jupp0r/go-priority-queue"
	"math"
	"fmt"
	"reflect"
)

func ProcessLearningRequest(req LearningRequestInterface) *LearningResponse {
	//create the response object and assign the origin request to it for debug later if need be
	response := new(LearningResponse)
	response.OriginRequest = req.GetBaseRequest()
	targetAttribute := req.GetTargetAttribute()

	measuredValues := 0
	var totalErr float64
	totalErr = 0

	//iterate through the given "targets" to determine the target value for each of the data points
	for i := 0; i < req.GetTargetsLength(); i++ {
		val := req.GetTargetByIndex(i)
		//find the k nearest Neighbours (k from the learning request)
		neighbours, nErr := FindNearestNeighbours(val,req)
		if nErr != nil {
			fmt.Println(nErr)
			continue //if something fails on the target point, skip it and work on the others
		}

		//determine the target value from the neighbour points
		//that is, we predict what the target value is here
		predictedAttribute := DeterminePoint(val,neighbours,req)

		//Copy the value of the predicted point from the target array
		//We copy the value in order to preserve our debug statement of the original request
		//before we change its values to reflect our prediction
		predictedPoint := reflect.New(reflect.ValueOf(val).Elem().Type()).Interface().(DataPoint)
		for key,val := range val.GetAttributeValues(targetAttribute) {
			predictedPoint.SetValue(key,val)
		}

		//set the value on the copied datapoint according to what we predicted
		predictedPoint.SetValue(targetAttribute,predictedAttribute)

		//if there was already a value assigned to the target attribute on the target point
		//we assign it to our total so we can calculate the mean error later
		expectedAttribute := val.GetValue(targetAttribute)
		if expectedAttribute.Value != nil {
			measuredValues++
		}

		//Create the completed point structure
		//Attaches nessasary information regarding the prediction, as well as the prediction itself
		completedPoint := new(CompletedPoint)
		completedPoint.PredictedAttribute = predictedAttribute
		completedPoint.FoundNeighbours = neighbours
		completedPoint.PointValue = predictedPoint
		completedPoint.ExpectedAttribute = expectedAttribute
		completedPoint.Error = 0 //we set the error as 0 to start with so if the expected is nil we still have something

		//if expected existed and its of float type, assign the point relative error based on its
		//percentage error, that is how incorrect was the prediction
		if completedPoint.PointValue.GetValue(targetAttribute).Type == FloatType && expectedAttribute.Value != nil {
			//Error = (|PredictedValue - ExpectedValue| / ExpectedValue) * 100
			//We multiply by 100 to show it as a percentage
			//this also allows reading on a per point basis
			completedPoint.Error = (math.Abs(completedPoint.PredictedAttribute.Value.(float64) -completedPoint.ExpectedAttribute.Value.(float64)) / completedPoint.ExpectedAttribute.Value.(float64)) * 100
		} else if completedPoint.PointValue.GetValue(targetAttribute).Type != FloatType && expectedAttribute.Value != nil {
			// We enter here if it wasn't a numeric prediction
			//we assign a simple 1 or 0 based on the accuracy here
			if completedPoint.PredictedAttribute.Value != completedPoint.ExpectedAttribute.Value {
				completedPoint.Error = 100
			} else {
				completedPoint.Error = 0
			}
		}

		//add the calculated error to the total error amount for mean calculation later
		totalErr = totalErr + completedPoint.Error

		response.CompletedPoints = append(response.CompletedPoints, *completedPoint)
	}

	//if we had measured values, then we calculate the mean of the errors and assign it to the root response object
	if measuredValues > 0 {
		response.Error = (totalErr / float64(measuredValues * 100)) * 100
	}

	//return the response to be processed by whatever inputted it
	return response
}

func DeterminePoint(Origin DataPoint, Neighbours []DataPoint, req LearningRequestInterface) Attribute {
	//foreach neighbour assert the votes of each to determine what the target attribute should be
	//if the type is a numeric one we could add averages etc here
	targetAttrName := req.GetTargetAttribute()

	if Origin.GetValue(targetAttrName).Type != FloatType {
		prMap := make(map[Attribute]int)

		for _,val := range Neighbours {
			prMap[val.GetValue(targetAttrName)]++
		}

		best := -1
		var bestAttr Attribute
		for key,val := range prMap {
			if val > best {
				best = val
				bestAttr = key
			}
		}
		return bestAttr
	} else {
		//take the average
		var total float64
		for _,val := range Neighbours {
			total = total + val.GetValue(targetAttrName).Value.(float64)
		}

		var a Attribute
		a.Value = total / float64(len(Neighbours))
		a.Type = FloatType
		return a
	}

}

func FindNearestNeighbours(MainPoint DataPoint, req LearningRequestInterface) ([]DataPoint, error) {
	//create a new priority queue so that neighbours can be ranked by distance
	pq := pq.New()
	var nullArray []DataPoint
	//for each data point in the training data set we calculate the distance and add it to the priority queue
	//the priority is equal to the distance of the point

	for j:= 0; j < req.GetDataLength(); j++ {
		val := req.GetDataByIndex(j)
		currDistance, distErr := Distance(MainPoint,val,req) // find the distance

		//if there is an error in the distance calculation
		if distErr != nil {
			//we assert that the calculation for this point is corrupted, therefore return an error and disregard
			//the previous distance calculations
			return nullArray, distErr
		}

		//insert into pq with the priority of the distance
		pq.Insert(val,currDistance)
		//fmt.Println(pq.Len())
	}

	//we now pop the appropriate amount of points off the priority queue
	var kNeighbours []DataPoint
	for i := 1; i <= req.GetK(); i++ {
		obj, pErr := pq.Pop()

		//if we fail to pop
		if pErr != nil {
			//if we have failed to pop, we need to assert a corrupt learning and disregard predictions for this target
			return nullArray,pErr
		}

		//add the neighbours to our return value
		kNeighbours = append(kNeighbours,obj.(DataPoint))
	}

	//return the response, this is the only exit that doesn't yield an error
	return kNeighbours,nil
}

//takes in two data points and evaluates the distance using the given comparers and basic numerics
func Distance(Origin DataPoint, Neighbour DataPoint, req LearningRequestInterface) (float64,error){

	//Get the attributes to be used to calculate the distance
	//the GetAttributeValues function does not return a set that includes the target value
	neighbourAttributes := Neighbour.GetAttributeValues(req.GetTargetAttribute())

	var distance float64

	//Iterate through each attribute and calculate the distance (similarity) using the appropriate formula
	for key,val := range(Origin.GetAttributeValues(req.GetTargetAttribute())){
		//get the neigbours attribute using the dictionary
		//this is to allow us to simply iterate through the origin (to be predicted) points attributes
		//and find the appropriate neighbour attributes
		currNeighbour := neighbourAttributes[key]

		//if somehow the attributes have different types we cannot compare them
		if val.Type != currNeighbour.Type {
			return -1, errors.New("Attributes have incompatible types")
		}

		if val.Type == FloatType {
			//if we know the types are integers or are numeric values, we can simply use the euclidean distance function
			distance = distance + NumericDistance(val.Value.(float64),currNeighbour.Value.(float64))
		} else {
			//if the types arethe same and aren't a float, then we are going to need to use a comparer to determine the distance
			distance = distance + ComparerDistance(val.Value.(string),val.Value.(string),req.GetComparers()[val.Type])
		}
	}

	//as the distance functions are defined to return euclidean distances
	//we need to root the distance
	return math.Sqrt(distance), nil
}

func NumericDistance(origin float64, destination float64) float64 {
	//return the euclidean distance
	//we can substitute other distance functions in here if need be
	baseDiff := origin - destination
	//this is essentially (origin - destination)^2
	return baseDiff * baseDiff
}

func ComparerDistance(origin string, destination string, comparer Comparer) float64 {
	//we simply use the encodings defined in the comparer to convert the type
	//into a numeric one that we can then use the euclidean distance on
	return NumericDistance(comparer.Encoding[origin],comparer.Encoding[destination])
}
