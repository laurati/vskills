package appliances

// define a fridge struct, the struct contain a string representing the type name
type Fridge struct {
	typeName string
}

func (fr *Fridge) Start() {
	fr.typeName = "Fridge"
}

func (fr *Fridge) GetPurpose() string {
	return "I am a " + fr.typeName + " I cool stuff down!"
}
