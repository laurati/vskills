package appliances

import "errors"

// The main interface used to describe appliances
type Appliance interface {
	Start()
	GetPurpose() string
}

// Appliances type
const (
	STOVE = iota
	FRIDGE
	MICROWAVE
)

func CreateAppliance(myType int) (Appliance, error) {
	switch myType {
	case STOVE:
		return new(Stove), nil
	case FRIDGE:
		return new(Fridge), nil
	case MICROWAVE:
		return new(Microwave), nil

	default:
		return nil, errors.New("Invalid Appliance Type")
	}
}
