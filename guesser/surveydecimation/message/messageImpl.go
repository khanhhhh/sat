package message

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
	return float64(ratIn)
}

// Abs :
func Abs(ratIn Message) (ratOut Message) {
	if ratIn >= 0 {
		ratOut = +ratIn
	} else {
		ratOut = -ratIn
	}
	return ratOut
}

// Sign :
func Sign(rat Message) (signOut int) {
	if rat > 0 {
		signOut = +1
	}
	if rat == 0 {
		signOut = 0
	}
	if rat < 0 {
		signOut = -1
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
