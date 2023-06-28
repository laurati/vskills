package appliances

// define a fridge struct, the struct contain a string representing the type name
type Stove struct {
	typeName string
}

func (fr *Stove) Start() {
	fr.typeName = "Stove"
}

func (fr *Stove) GetPurpose() string {
	return "I am a " + fr.typeName + " I bake stuff!"
}
