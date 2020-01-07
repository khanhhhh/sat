package unitpropagation

import "github.com/khanhhhh/sat/instance"

// Guess :
func Guess(ins instance.Instance) (converged bool, variableOut instance.Variable, valueOut bool) {
	converged = false
	for variable := range ins.VariableMap() {
		if len(ins.VariableMap()[variable]) == 1 {
			converged = true
			variableOut = variable
			for _, value := range ins.VariableMap()[variable] {
				valueOut = value
			}
			break
		}
	}
	return converged, variableOut, valueOut
}
