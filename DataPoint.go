package main

type DataPoint interface {
	GetValue(attributeName string) Attribute //gets an attribute based on the name of the attribute given
	SetValue(attributeName string, valToSet Attribute) error //sets the attribute value, throws an error if they aren't the right types
	GetAttributeValues(targetName string) map[string]Attribute //a map of the name of the attribute to the attribute itself, excluding the target
}

type Attribute struct {
	Value interface{}
	Type string //the type can either be int or anything else. If its anything else we need to use its comparer, that is the method in which we can compare this value to others of the same class
}

func createFloatAttribute(input float64) Attribute {
	return createAttribute(input,FloatType)
}

func createBinaryAttribute(input string) Attribute {
	return createAttribute(input,BinaryType)
}

func createAttribute(input interface{}, inputType string) Attribute {
	a := new(Attribute)
	a.Value = input
	a.Type = inputType
	return *a
}



