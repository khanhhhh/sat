package message

import "math"

// Message :
// Store value x as (1-x) to reduce float vanishing
type Message float64

// FromInt :
func FromInt(a int, b int) (ratOut Message) {
	ratOut = Message(1 - float64(a)/float64(b))
	return ratOut
}

// FromFloat :
func FromFloat(floatIn float64) (ratOut Message) {
	ratOut = Message(1 - floatIn)
	return ratOut
}

// ToFloat :
func (rat Message) ToFloat() (floatOut float64) {
	floatOut = float64(1 - rat)
	return floatOut
}

// IsNaN :
func (rat Message) IsNaN() (IsNaN bool) {
	IsNaN = math.IsNaN(float64(1 - rat))
	return IsNaN
}

// Sign :
func (rat Message) Sign() (signOut int) {
	if 1-rat > 0 {
		signOut = +1
	}
	if 1-rat < 0 {
		signOut = -1
	}
	if 1-rat == 0 {
		signOut = 0
	}
	return signOut
}

// Abs :
func (rat Message) Abs() (ratOut Message) {
	if (1 - rat).Sign() == 1 {
		ratOut = +rat
	} else {
		//ratOut = 1 - (-(1 - ratIn))
		ratOut = 2 - rat
	}
	return ratOut
}

// Cmp :
func (rat Message) Cmp(ratIn Message) (signOut int) {
	if 1-rat > 1-ratIn {
		signOut = +1
	}
	if 1-rat < 1-ratIn {
		signOut = -1
	}
	if 1-rat == 1-ratIn {
		signOut = 0
	}
	return signOut
}

// Add :
func (rat Message) Add(ratIn Message) (ratOut Message) {
	//ratOut = 1 - ((1 - rat) + (1 - ratIn))
	ratOut = -1 + rat + ratIn
	return ratOut
}

// Sub :
func (rat Message) Sub(ratIn Message) (ratOut Message) {
	//ratOut = 1 - ((1 - rat) - (1 - ratIn))
	ratOut = 1 + rat - ratIn
	return ratOut
}

// Mul :
func (rat Message) Mul(ratIn Message) (ratOut Message) {
	//ratOut = 1 - ((1 - rat) * (1 - ratIn))
	ratOut = rat + ratIn - rat*ratIn
	return ratOut
}

// Div :
func (rat Message) Div(ratIn Message) (ratOut Message) {
	//ratOut = 1 - ((1 - rat) / (1 - ratIn))
	ratOut = (rat - ratIn) / (1 - ratIn)
	return ratOut
}
