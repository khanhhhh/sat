package message

import "math"

// Message :
type Message float64

// FromInt :
func FromInt(a int, b int) (ratOut Message) {
	ratOut = Message(float64(a) / float64(b))
	return ratOut
}

// FromFloat :
func FromFloat(floatIn float64) (ratOut Message) {
	ratOut = Message(floatIn)
	return ratOut
}

// ToFloat :
func ToFloat(ratIn Message) (floatOut float64) {
	floatOut = float64(ratIn)
	return floatOut
}

// IsNaN :
func IsNaN(ratIn Message) (IsNaN bool) {
	IsNaN = math.IsNaN(float64(ratIn))
	return IsNaN
}

// Sign :
func Sign(ratIn Message) (signOut int) {
	if ratIn > 0 {
		signOut = +1
	}
	if ratIn < 0 {
		signOut = -1
	}
	if ratIn == 0 {
		signOut = 0
	}
	return signOut
}

// Abs :
func Abs(ratIn Message) (ratOut Message) {
	if Sign(ratIn) == 1 {
		ratOut = +ratIn
	} else {
		ratOut = -ratIn
	}
	return ratOut
}

// Cmp :
func Cmp(ratIn1 Message, ratIn2 Message) (signOut int) {
	if ratIn1 > ratIn2 {
		signOut = +1
	}
	if ratIn1 < ratIn2 {
		signOut = -1
	}
	if ratIn1 == ratIn2 {
		signOut = 0
	}
	return signOut
}

// Add :
func Add(ratIn1 Message, ratIn2 Message) (ratOut Message) {
	ratOut = ratIn1 + ratIn2
	return ratOut
}

// Sub :
func Sub(ratIn1 Message, ratIn2 Message) (ratOut Message) {
	ratOut = ratIn1 - ratIn2
	return ratOut
}

// Mul :
func Mul(ratIn1 Message, ratIn2 Message) (ratOut Message) {
	ratOut = ratIn1 * ratIn2
	return ratOut
}

// Div :
func Div(ratIn1 Message, ratIn2 Message) (ratOut Message) {
	ratOut = ratIn1 / ratIn2
	return ratOut
}
