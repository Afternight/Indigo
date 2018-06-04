package main

import (
	"github.com/gin-gonic/gin"
	"errors"
	"encoding/json"
)

//region Ionosphere Requests

type IonosphereLearningRequest struct {
	BaseLearningRequest
	Targets []IonospherePoint
	Data []IonospherePoint
}

func (g *IonosphereLearningRequest) ReceiveRequest (c *gin.Context) error {
	jsonDecodeErr := json.NewDecoder(c.Request.Body).Decode(&g)

	if jsonDecodeErr != nil {
		return jsonDecodeErr
	}
	return nil
}

func (g *IonosphereLearningRequest) GetBaseRequest() LearningRequestInterface {
	return g
}

func (g *IonosphereLearningRequest) GetTargetsLength() int {
	return len(g.Targets)
}

func (g *IonosphereLearningRequest) GetTargetByIndex(i int) DataPoint {
	return &g.Targets[i]
}

func (g *IonosphereLearningRequest) GetDataLength() int {
	return len(g.Data)
}

func (g *IonosphereLearningRequest) GetDataByIndex(i int) DataPoint {
	return &g.Data[i]
}

//endregion

//region IonospherePoint

type IonospherePoint struct {
	A01 float64
	A02 float64
	A03 float64
	A04 float64
	A05 float64
	A06 float64
	A07 float64
	A08 float64
	A09 float64
	A10 float64
	A11 float64
	A12 float64
	A13 float64
	A14 float64
	A15 float64
	A16 float64
	A17 float64
	A18 float64
	A19 float64
	A20 float64
	A21 float64
	A22 float64
	A23 float64
	A24 float64
	A25 float64
	A26 float64
	A27 float64
	A28 float64
	A29 float64
	A30 float64
	A31 float64
	A32 float64
	A33 float64
	A34 float64
	Class string
}

func (g *IonospherePoint) GetValue(toGet string) Attribute {
	if toGet == "A01" {
		return createFloatAttribute(g.A01)
	} else if toGet == "A02" {
		return createFloatAttribute(g.A02)
	} else if toGet == "A03" {
		return createFloatAttribute(g.A03)
	} else if toGet == "A04" {
		return createFloatAttribute(g.A04)
	} else if toGet == "A05" {
		return createFloatAttribute(g.A05)
	} else if toGet == "A06" {
		return createFloatAttribute(g.A06)
	} else if toGet == "A07" {
		return createFloatAttribute(g.A07)
	} else if toGet == "A08" {
		return createFloatAttribute(g.A08)
	} else if toGet == "A09" {
		return createFloatAttribute(g.A09)
	} else if toGet == "A10" {
		return createFloatAttribute(g.A10)
	} else if toGet == "A11" {
		return createFloatAttribute(g.A11)
	} else if toGet == "A12" {
		return createFloatAttribute(g.A12)
	} else if toGet == "A13" {
		return createFloatAttribute(g.A13)
	} else if toGet == "A14" {
		return createFloatAttribute(g.A14)
	} else if toGet == "A15" {
		return createFloatAttribute(g.A15)
	} else if toGet == "A16" {
		return createFloatAttribute(g.A16)
	} else if toGet == "A17" {
		return createFloatAttribute(g.A17)
	} else if toGet == "A18" {
		return createFloatAttribute(g.A18)
	} else if toGet == "A19" {
		return createFloatAttribute(g.A19)
	} else if toGet == "A20" {
		return createFloatAttribute(g.A20)
	} else if toGet == "A21" {
		return createFloatAttribute(g.A21)
	} else if toGet == "A22" {
		return createFloatAttribute(g.A22)
	} else if toGet == "A23" {
		return createFloatAttribute(g.A23)
	} else if toGet == "A24" {
		return createFloatAttribute(g.A24)
	} else if toGet == "A25" {
		return createFloatAttribute(g.A25)
	} else if toGet == "A26" {
		return createFloatAttribute(g.A26)
	} else if toGet == "A27" {
		return createFloatAttribute(g.A27)
	} else if toGet == "A28" {
		return createFloatAttribute(g.A28)
	} else if toGet == "A29" {
		return createFloatAttribute(g.A29)
	} else if toGet == "A30" {
		return createFloatAttribute(g.A30)
	} else if toGet == "A31" {
		return createFloatAttribute(g.A31)
	} else if toGet == "A32" {
		return createFloatAttribute(g.A32)
	} else if toGet == "A33" {
		return createFloatAttribute(g.A33)
	} else if toGet == "A34" {
		return createFloatAttribute(g.A34)
	} else if toGet == "Class" {
		return createBinaryAttribute(g.Class)
	}

	return Attribute{}
}

func (g *IonospherePoint) SetValue(name string, toSet Attribute) error {
	if name == "A01" {
		g.A01 = toSet.Value.(float64)
	} else if name == "A02" {
		g.A02 = toSet.Value.(float64)
	} else if name == "A03" {
		g.A03 = toSet.Value.(float64)
	} else if name == "A04" {
		g.A04 = toSet.Value.(float64)
	} else if name == "A05" {
		g.A05 = toSet.Value.(float64)
	} else if name == "A06" {
		g.A06 = toSet.Value.(float64)
	} else if name == "A07" {
		g.A07 = toSet.Value.(float64)
	} else if name == "A08" {
		g.A08 = toSet.Value.(float64)
	} else if name == "A09" {
		g.A09 = toSet.Value.(float64)
	} else if name == "A10" {
		g.A10 = toSet.Value.(float64)
	} else if name == "A11" {
		g.A11 = toSet.Value.(float64)
	} else if name == "A12" {
		g.A12 = toSet.Value.(float64)
	} else if name == "A13" {
		g.A13 = toSet.Value.(float64)
	} else if name == "A14" {
		g.A14 = toSet.Value.(float64)
	} else if name == "A15" {
		g.A15 = toSet.Value.(float64)
	} else if name == "A16" {
		g.A16 = toSet.Value.(float64)
	} else if name == "A17" {
		g.A17 = toSet.Value.(float64)
	} else if name == "A18" {
		g.A18 = toSet.Value.(float64)
	} else if name == "A19" {
		g.A19 = toSet.Value.(float64)
	} else if name == "A20" {
		g.A20 = toSet.Value.(float64)
	} else if name == "A21" {
		g.A21 = toSet.Value.(float64)
	} else if name == "A22" {
		g.A22 = toSet.Value.(float64)
	} else if name == "A23" {
		g.A23 = toSet.Value.(float64)
	} else if name == "A24" {
		g.A24 = toSet.Value.(float64)
	} else if name == "A25" {
		g.A25 = toSet.Value.(float64)
	} else if name == "A26" {
		g.A26 = toSet.Value.(float64)
	} else if name == "A27" {
		g.A27 = toSet.Value.(float64)
	} else if name == "A28" {
		g.A28 = toSet.Value.(float64)
	} else if name == "A29" {
		g.A29 = toSet.Value.(float64)
	} else if name == "A30" {
		g.A30 = toSet.Value.(float64)
	} else if name == "A31" {
		g.A31 = toSet.Value.(float64)
	} else if name == "A32" {
		g.A32 = toSet.Value.(float64)
	} else if name == "A33" {
		g.A33 = toSet.Value.(float64)
	} else if name == "A34" {
		g.A34 = toSet.Value.(float64)
	} else if name == "Class" {
		g.Class = toSet.Value.(string)
	} else {
		return errors.New("No match")
	}
	return nil
}

func (g *IonospherePoint) GetAttributeValues(target string) map[string]Attribute {
	newMap := make(map[string]Attribute)

	if target != "A01" {
		newMap["A01"] = createFloatAttribute(g.A01)
	}

	if target != "A02" {
		newMap["A02"] = createFloatAttribute(g.A02)
	}

	if target != "A03" {
		newMap["A03"] = createFloatAttribute(g.A03)
	}

	if target != "A04" {
		newMap["A04"] = createFloatAttribute(g.A04)
	}

	if target != "A05" {
		newMap["A05"] = createFloatAttribute(g.A05)
	}

	if target != "A06" {
		newMap["A06"] = createFloatAttribute(g.A06)
	}

	if target != "A07" {
		newMap["A07"] = createFloatAttribute(g.A07)
	}

	if target != "A08" {
		newMap["A08"] = createFloatAttribute(g.A08)
	}

	if target != "A09" {
		newMap["A09"] = createFloatAttribute(g.A09)
	}

	if target != "A10" {
		newMap["A10"] = createFloatAttribute(g.A10)
	}

	if target != "A11" {
		newMap["A11"] = createFloatAttribute(g.A11)
	}

	if target != "A12" {
		newMap["A12"] = createFloatAttribute(g.A12)
	}

	if target != "A13" {
		newMap["A13"] = createFloatAttribute(g.A13)
	}

	if target != "A14" {
		newMap["A14"] = createFloatAttribute(g.A14)
	}

	if target != "A15" {
		newMap["A15"] = createFloatAttribute(g.A15)
	}

	if target != "A16" {
		newMap["A16"] = createFloatAttribute(g.A16)
	}

	if target != "A17" {
		newMap["A17"] = createFloatAttribute(g.A17)
	}

	if target != "A18" {
		newMap["A18"] = createFloatAttribute(g.A18)
	}

	if target != "A19" {
		newMap["A19"] = createFloatAttribute(g.A19)
	}

	if target != "A20" {
		newMap["A20"] = createFloatAttribute(g.A20)
	}

	if target != "A21" {
		newMap["A21"] = createFloatAttribute(g.A21)
	}

	if target != "A22" {
		newMap["A22"] = createFloatAttribute(g.A22)
	}

	if target != "A23" {
		newMap["A23"] = createFloatAttribute(g.A23)
	}

	if target != "A24" {
		newMap["A24"] = createFloatAttribute(g.A24)
	}

	if target != "A25" {
		newMap["A25"] = createFloatAttribute(g.A25)
	}

	if target != "A26" {
		newMap["A26"] = createFloatAttribute(g.A26)
	}

	if target != "A27" {
		newMap["A27"] = createFloatAttribute(g.A27)
	}

	if target != "A28" {
		newMap["A28"] = createFloatAttribute(g.A28)
	}

	if target != "A29" {
		newMap["A29"] = createFloatAttribute(g.A29)
	}

	if target != "A30" {
		newMap["A30"] = createFloatAttribute(g.A30)
	}

	if target != "A31" {
		newMap["A31"] = createFloatAttribute(g.A31)
	}

	if target != "A32" {
		newMap["A32"] = createFloatAttribute(g.A32)
	}

	if target != "A33" {
		newMap["A33"] = createFloatAttribute(g.A33)
	}

	if target != "A34" {
		newMap["A34"] = createFloatAttribute(g.A34)
	}

	if target != "Class" {
		newMap["Class"] = createBinaryAttribute(g.Class)
	}

	return newMap
}

//endregion
