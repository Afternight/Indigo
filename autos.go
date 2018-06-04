package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"encoding/json"
)

//region Autos requests

type AutosLearningRequest struct {
	BaseLearningRequest
	Targets []AutosPoint
	Data []AutosPoint
}

func (g *AutosLearningRequest) ReceiveRequest (c *gin.Context) error {
	jsonDecodeErr := json.NewDecoder(c.Request.Body).Decode(&g)

	if jsonDecodeErr != nil {
		return jsonDecodeErr
	}
	return nil
}

func (g *AutosLearningRequest) GetBaseRequest() LearningRequestInterface {
	return g
}

func (g *AutosLearningRequest) GetTargetsLength() int {
	return len(g.Targets)
}

func (g *AutosLearningRequest) GetTargetByIndex(i int) DataPoint {
	return &g.Targets[i]
}

func (g *AutosLearningRequest) GetDataLength() int {
	return len(g.Data)
}

func (g *AutosLearningRequest) GetDataByIndex(i int) DataPoint {
	return &g.Data[i]
}

//endregion

//region Autos Point

type AutosPoint struct {
	NormalizedLosses float64
	Make string
	FuelType string
	Aspiration string
	NumOfDoors string
	BodyStyle string
	DriveWheels string
	EngineLocation string
	WheelBase float64
	Length float64
	Width float64
	Height float64
	CurbWeight float64
	EngineType string
	NumOfCylinders string
	EngineSize float64
	FuelSystem string
	Bore float64
	Stroke float64
	CompressionRatio float64
	Horsepower float64
	PeakRpm float64
	CityMpg float64
	HighwayMpg float64
	Price float64
	Symboling float64
}

func (g *AutosPoint) GetValue(toGet string) Attribute {
	if toGet == "NormalizedLosses" {
		return createFloatAttribute(g.NormalizedLosses)
	} else if toGet == "Make" {
		return createAttribute(g.Make,"Make")
	} else if toGet == "FuelType" {
		return createAttribute(g.FuelType,"FuelType")
	} else if toGet == "Aspiration" {
		return createAttribute(g.Aspiration,"Aspiration")
	} else if toGet == "NumOfDoors" {
		return createAttribute(g.NumOfDoors,"NumOfDoors")
	} else if toGet == "BodyStyle" {
		return createAttribute(g.BodyStyle,"BodyStyle")
	} else if toGet == "DriveWheels" {
		return createAttribute(g.DriveWheels,"DriveWheels")
	} else if toGet == "EngineLocation" {
		return createAttribute(g.EngineLocation,"EngineLocation")
	} else if toGet == "WheelBase" {
		return createFloatAttribute(g.WheelBase)
	} else if toGet == "Length" {
		return createFloatAttribute(g.Length)
	} else if toGet == "Width" {
		return createFloatAttribute(g.Width)
	} else if toGet == "Height" {
		return createFloatAttribute(g.Height)
	} else if toGet == "CurbWeight" {
		return createFloatAttribute(g.CurbWeight)
	} else if toGet == "EngineType" {
		return createAttribute(g.EngineType,"EngineType")
	} else if toGet == "NumOfCylinders" {
		return createAttribute(g.NumOfCylinders,"NumOfCylinders")
	} else if toGet == "EngineSize" {
		return createFloatAttribute(g.EngineSize)
	} else if toGet == "FuelSystem" {
		return createAttribute(g.FuelSystem,"FuelSystem")
	} else if toGet == "Bore" {
		return createFloatAttribute(g.Bore)
	} else if toGet == "Stroke" {
		return createFloatAttribute(g.Stroke)
	} else if toGet == "CompressionRatio" {
		return createFloatAttribute(g.CompressionRatio)
	} else if toGet == "Horsepower" {
		return createFloatAttribute(g.Horsepower)
	} else if toGet == "PeakRpm" {
		return createFloatAttribute(g.PeakRpm)
	} else if toGet == "CityMpg" {
		return createFloatAttribute(g.CityMpg)
	} else if toGet == "HighwayMpg" {
		return createFloatAttribute(g.HighwayMpg)
	} else if toGet == "Price" {
		return createFloatAttribute(g.Price)
	} else if toGet == "Symboling" {
		return createFloatAttribute(g.Symboling)
	}

	return Attribute{}
}

func (g *AutosPoint) SetValue(name string, toSet Attribute) error {
	if name == "NormalizedLosses" {
		g.NormalizedLosses = toSet.Value.(float64)
	} else if name == "Make" {
		g.Make = toSet.Value.(string)
	} else if name == "FuelType" {
		g.FuelType = toSet.Value.(string)
	} else if name == "Aspiration" {
		g.Aspiration = toSet.Value.(string)
	} else if name == "NumOfDoors" {
		g.NumOfDoors = toSet.Value.(string)
	} else if name == "BodyStyle" {
		g.BodyStyle = toSet.Value.(string)
	} else if name == "DriveWheels" {
		g.DriveWheels = toSet.Value.(string)
	} else if name == "EngineLocation" {
		g.EngineLocation = toSet.Value.(string)
	} else if name == "WheelBase" {
		g.WheelBase = toSet.Value.(float64)
	} else if name == "Length" {
		g.Length = toSet.Value.(float64)
	} else if name == "Width" {
		g.Width = toSet.Value.(float64)
	} else if name == "Height" {
		g.Height = toSet.Value.(float64)
	} else if name == "CurbWeight" {
		g.CurbWeight = toSet.Value.(float64)
	} else if name == "EngineType" {
		g.EngineType = toSet.Value.(string)
	} else if name == "NumOfCylinders" {
		g.NumOfCylinders = toSet.Value.(string)
	} else if name == "EngineSize" {
		g.EngineSize = toSet.Value.(float64)
	} else if name == "FuelSystem" {
		g.FuelSystem = toSet.Value.(string)
	} else if name == "Bore" {
		g.Bore = toSet.Value.(float64)
	} else if name == "Stroke" {
		g.Stroke = toSet.Value.(float64)
	} else if name == "CompressionRatio" {
		g.CompressionRatio = toSet.Value.(float64)
	} else if name == "Horsepower" {
		g.Horsepower = toSet.Value.(float64)
	} else if name == "PeakRpm" {
		g.PeakRpm = toSet.Value.(float64)
	} else if name == "CityMpg" {
		g.CityMpg = toSet.Value.(float64)
	} else if name == "HighwayMpg" {
		g.HighwayMpg = toSet.Value.(float64)
	} else if name == "Price" {
		g.Price = toSet.Value.(float64)
	} else if name == "Symboling" {
		g.Symboling = toSet.Value.(float64)
	} else {
		return errors.New("No match")
	}
	return nil
}

func (g *AutosPoint) GetAttributeValues(target string) map[string]Attribute {
	newMap := make(map[string]Attribute)

	if target != "NormalizedLosses" {
		newMap["NormalizedLosses"] = createFloatAttribute(g.NormalizedLosses)
	}

	if target != "Make" {
		newMap["Make"] = createAttribute(g.Make,"Make")
	}

	if target != "FuelType" {
		newMap["FuelType"] = createAttribute(g.FuelType,"FuelType")
	}

	if target != "Aspiration" {
		newMap["Aspiration"] = createAttribute(g.Aspiration,"Aspiration")
	}

	if target != "NumOfDoors" {
		newMap["NumOfDoors"] = createAttribute(g.NumOfDoors,"NumOfDoors")
	}

	if target != "BodyStyle" {
		newMap["BodyStyle"] = createAttribute(g.BodyStyle,"BodyStyle")
	}

	if target != "DriveWheels" {
		newMap["DriveWheels"] = createAttribute(g.DriveWheels,"DriveWheels")
	}

	if target != "EngineLocation" {
		newMap["EngineLocation"] = createAttribute(g.EngineLocation,"EngineLocation")
	}

	if target != "WheelBase" {
		newMap["WheelBase"] = createFloatAttribute(g.WheelBase)
	}

	if target != "Length" {
		newMap["Length"] = createFloatAttribute(g.Length)
	}

	if target != "Width" {
		newMap["Width"] = createFloatAttribute(g.Width)
	}

	if target != "Height" {
		newMap["Height"] = createFloatAttribute(g.Height)
	}

	if target != "CurbWeight" {
		newMap["CurbWeight"] = createFloatAttribute(g.CurbWeight)
	}

	if target != "EngineType" {
		newMap["EngineType"] = createAttribute(g.EngineType,"EngineType")
	}

	if target != "NumOfCylinders" {
		newMap["NumOfCylinders"] = createAttribute(g.NumOfCylinders,"NumOfCylinders")
	}

	if target != "EngineSize" {
		newMap["EngineSize"] = createFloatAttribute(g.EngineSize)
	}

	if target != "FuelSystem" {
		newMap["FuelSystem"] = createAttribute(g.FuelSystem,"FuelSystem")
	}

	if target != "Bore" {
		newMap["Bore"] = createFloatAttribute(g.Bore)
	}

	if target != "Stroke" {
		newMap["Stroke"] = createFloatAttribute(g.Stroke)
	}

	if target != "CompressionRatio" {
		newMap["CompressionRatio"] = createFloatAttribute(g.CompressionRatio)
	}

	if target != "Horsepower" {
		newMap["Horsepower"] = createFloatAttribute(g.Horsepower)
	}

	if target != "PeakRpm" {
		newMap["PeakRpm"] = createFloatAttribute(g.PeakRpm)
	}

	if target != "CityMpg" {
		newMap["CityMpg"] = createFloatAttribute(g.CityMpg)
	}

	if target != "HighwayMpg" {
		newMap["HighwayMpg"] = createFloatAttribute(g.HighwayMpg)
	}

	if target != "Price" {
		newMap["Price"] = createFloatAttribute(g.Price)
	}

	if target != "Symboling" {
		newMap["Symboling"] = createFloatAttribute(g.Symboling)
	}

	return newMap
}

//endregion
