package surveydecimation

import (
	"github.com/khanhhhh/sat/guesser/surveydecimation/message"
	"github.com/khanhhhh/sat/instance"
)

type surveyPropagationGraph struct {
	piMap  map[edge][3]message.Message // variable -> clause
	etaMap map[edge]message.Message    // clause -> variable
}

// makeSurveyPropagationGraph :
// Make an empty Survey Propagation Graph
func makeSurveyPropagationGraph(ins instance.Instance) (graph *surveyPropagationGraph) {
	graph = &surveyPropagationGraph{
		make(map[edge][3]message.Message),
		make(map[edge]message.Message),
	}
	for _, edge := range allEdges(ins) {
		graph.piMap[edge] = [3]message.Message{
			message.FromInt(1, 2),
			message.FromInt(1, 2),
			message.FromInt(1, 2),
		}
		graph.etaMap[edge] = message.FromInt(1, 2)
	}
	return graph
}

// iterateSurveyPropagationGraph :
// Iterate clauseA Survey Propagation Graph
func iterateSurveyPropagationGraph(ins instance.Instance, graphIn *surveyPropagationGraph, smooth float64) (absoluteEtaChange float64, graphOut *surveyPropagationGraph) {
	// initialize etaChange to 0
	absoluteEtaChange = 0
	// make empty graphOut
	graphOut = &surveyPropagationGraph{
		make(map[edge][3]message.Message),
		make(map[edge]message.Message),
	}
	// calculate graphOut value for all edges
	for _, edge := range allEdges(ins) {
		variableI := edge.variable
		clauseA := edge.clause
		// eta
		{
			var eta = message.FromInt(1, 1)
			for variableJ := range ins.ClauseMap()[clauseA] {
				if variableJ != variableI {
					triplet := graphIn.piMap[newEdge(variableJ, clauseA)]
					sum := message.Add(message.Add(triplet[0], triplet[1]), triplet[2])
					eta = message.Mul(eta, message.Div(triplet[0], sum))
				}
			}
			// detect nan : if sum triplet == 0 => eta = NaN
			if message.IsNaN(eta) {
				panic("eta: NaN")
			}
			if message.ToFloat(message.Abs(message.Sub(eta, graphIn.etaMap[edge]))) > absoluteEtaChange {
				absoluteEtaChange = message.ToFloat(message.Abs(message.Sub(eta, graphIn.etaMap[edge])))
			}
			graphOut.etaMap[edge] = eta
		}
		// pi
		{
			var productAgree = message.FromInt(1, 1)
			var productDisagree = message.FromInt(1, 1)
			for _, clauseB := range clauseAgree(ins, edge) {
				productAgree = message.Mul(
					productAgree,
					message.Sub(message.FromInt(1, 1), graphIn.etaMap[newEdge(variableI, clauseB)]),
				)
			}
			for _, clauseB := range clauseDisagree(ins, edge) {
				productDisagree = message.Mul(
					productDisagree,
					message.Sub(message.FromInt(1, 1), graphIn.etaMap[newEdge(variableI, clauseB)]),
				)
			}
			var triplet [3]message.Message
			smoothConst := message.FromFloat(smooth)
			// detect zero
			if message.Sign(productAgree) == 0 && message.Sign(productDisagree) == 0 {
				panic("triplet: Zero")
			} else {
				triplet[0] = message.Mul(
					productAgree,
					message.Sub(message.FromInt(1, 1), message.Mul(smoothConst, productDisagree)),
				)
				triplet[1] = message.Mul(
					productDisagree,
					message.Sub(message.FromInt(1, 1), message.Mul(smoothConst, productAgree)),
				)
				triplet[2] = message.Mul(
					message.FromFloat(smooth),
					message.Mul(productAgree, productDisagree),
				)
			}
			// detect nan
			if message.IsNaN(triplet[0]) || message.IsNaN(triplet[1]) || message.IsNaN(triplet[2]) {
				panic("triplet: NaN")
			}
			graphOut.piMap[edge] = triplet
		}
	}
	return absoluteEtaChange, graphOut
}
