package surveydecimation

import (
	"github.com/khanhhhh/sat/guesser/surveydecimation/message"
	"github.com/khanhhhh/sat/instance"
)

// surveyDecimation :
// inference max bias variable from a Survey Propagation Graph
func surveyDecimation(ins instance.Instance, graphIn *surveyPropagationGraph, smooth float64) (nonTrivialCover bool, maxBiasVariable instance.Variable, maxBiasValue bool) {
	var maxBias = message.FromInt(0, 1)
	// select maxBias over all variables
	for variable := range ins.VariableMap() {
		// calculate mu
		var mu [3]message.Message
		{
			var oneMessage = message.FromInt(1, 1)
			var productPositive = message.FromInt(1, 1)
			var productNegative = message.FromInt(1, 1)
			for _, clause := range clausePositive(ins, variable) {
				productPositive = productPositive.Mul(oneMessage.Sub(graphIn.etaMap[newEdge(variable, clause)]))
			}
			for _, clause := range clauseNegative(ins, variable) {
				productNegative = productNegative.Mul(oneMessage.Sub(graphIn.etaMap[newEdge(variable, clause)]))
			}
			smoothConst := message.FromFloat(smooth)
			mu[0] = oneMessage.Sub(smoothConst.Mul(productNegative)).Mul(productPositive)
			mu[1] = oneMessage.Sub(smoothConst.Mul(productPositive)).Mul(productNegative)
			mu[2] = smoothConst.Mul(productPositive).Mul(productNegative)
		}
		// normalize
		{
			sum := mu[0].Add(mu[1]).Add(mu[2])
			if sum.Sign() == 1 {
				mu[0] = mu[0].Div(sum)
				mu[1] = mu[1].Div(sum)
				mu[2] = mu[2].Div(sum)
			}
		}
		// select maxBias
		{
			bias := mu[1].Sub(mu[0]).Abs()
			if bias.Cmp(maxBias) == 1 {
				maxBias = bias
				maxBiasVariable = variable
				maxBiasValue = mu[1].Cmp(mu[0]) == 1
			}
		}
	}
	// detect trivial cover
	if maxBias.Sign() == 0 {
		nonTrivialCover = false
	} else {
		nonTrivialCover = true
	}
	return nonTrivialCover, maxBiasVariable, maxBiasValue
}
