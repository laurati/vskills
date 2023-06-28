package appliances

// define a fridge struct, the struct contain a string representing the type name
type Microwave struct {
	typeName string
}

func (fr *Microwave) Start() {
	fr.typeName = "Microwave"
}

func (fr *Microwave) GetPurpose() string {
	return "I am a " + fr.typeName + " I warm stuff!"
}
