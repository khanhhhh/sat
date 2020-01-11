package surveydecimation

import (
	"errors"

	"github.com/khanhhhh/sat/guesser/surveydecimation/message"
	"github.com/khanhhhh/sat/instance"
)

type surveyPropagationGraph struct {
	piMap  map[edge][3]message.Message // variable -> clause [converge to 1]
	etaMap map[edge]message.Message    // clause -> variable [converge to 1]
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

var iter = true

// iterateSurveyPropagationGraph :
// Iterate clauseA Survey Propagation Graph
func iterateSurveyPropagationGraph(ins instance.Instance, graphIn *surveyPropagationGraph, smooth float64) (absoluteEtaChange float64, graphOut *surveyPropagationGraph, err error) {
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
					sum := triplet[0].Add(triplet[1]).Add(triplet[2])
					eta = eta.Mul(triplet[0].Div(sum))
				}
				// detect nan : if sum triplet == 0 => eta = NaN
				if eta.IsNaN() {
					return 0, graphIn, errors.New("eta: NaN")
				}
			}
			if eta.Sub(graphIn.etaMap[edge]).Abs().ToFloat() > absoluteEtaChange {
				absoluteEtaChange = eta.Sub(graphIn.etaMap[edge]).Abs().ToFloat()
			}
			graphOut.etaMap[edge] = eta
		}
		// pi
		{
			var oneMessage = message.FromInt(1, 1)
			var productAgree = message.FromInt(1, 1)
			var productDisagree = message.FromInt(1, 1)
			for _, clauseB := range clauseAgree(ins, edge) {
				productAgree = productAgree.Mul(oneMessage.Sub(graphIn.etaMap[newEdge(variableI, clauseB)]))
			}
			for _, clauseB := range clauseDisagree(ins, edge) {
				productDisagree = productDisagree.Mul(oneMessage.Sub(graphIn.etaMap[newEdge(variableI, clauseB)]))
			}
			var triplet [3]message.Message
			smoothConst := message.FromFloat(smooth)
			// detect zero, negative
			if productAgree.Sign() != 1 && productDisagree.Sign() != 1 {
				return 0, graphIn, errors.New("triplet: non-positive")
			}
			triplet[0] = oneMessage.Sub(smoothConst.Mul(productDisagree)).Mul(productAgree)
			triplet[1] = oneMessage.Sub(smoothConst.Mul(productAgree)).Mul(productDisagree)
			triplet[2] = smoothConst.Mul(productAgree).Mul(productDisagree)
			// detect nan
			if triplet[0].IsNaN() || triplet[1].IsNaN() || triplet[2].IsNaN() {
				return 0, graphIn, errors.New("triplet: NaN")
			}
			graphOut.piMap[edge] = triplet
		}
	}
	return absoluteEtaChange, graphOut, nil
}
