// Package is a simpler simple threshold neuron aka perceptron.
// A perceptron takes several binary inputs, x1,x2,…x1,x2,…,
// and produces a single binary output
package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

const (
	rate              = 0.2
	trainingThreshold = 98
	maxEpochs         = 10
)

// Real numbers expressing the importance of the respective inputs to the output
type weights []float64

// epoch is a single pass through the entire training set
// and testing the verification set.
type epoch struct {
	id      int
	results []Output
	success float64
}

func main() {
	w, t := setup()
	train(t, w)
}

func setup() (weights, TrainingSet) {
	w := initWeights(2)
	t := getTrainingData()
	return w, t
}

func train(t TrainingSet, w weights) {
	epochId := 0
	trained := false
	epochLimit := false

	for !trained && !epochLimit {
		e := epoch{id: epochId, results: []Output{}, success: 0.0}
		w, e = runEpoch(t, w, e)
		// Check if training has achieved satisfactory score.
		trained = e.success > trainingThreshold
		epochLimit = e.id >= maxEpochs
		epochId++
	}
}

func runEpoch(t TrainingSet, w weights, e epoch) (weights, epoch) {
	displayEpochHeader(e.id)

	// Run training - iterate over each case in the training data.
	for _, c := range t {
		w, e.results = runTest(c, w, e.results)
	}

	e.success = score(e.results)

	displayEpochFooter(e.id, e.success)

	return w, e
}

func score(r []Output) float64 {
	return (sum(r) / float64(len(r))) * 100
}

// sum returns the total of all added results.
func sum(v []Output) float64 {
	s := 0.0
	for _, j := range v {
		s += float64(j)
	}
	return s
}

// Run the test, and adjust the weights according to the results.
//func runTest(test TrainingCase, w weights) (weights, []float64) {
func runTest(test TrainingCase, w weights, epochResults []Output) (weights, []Output) {
	displayTestHeader(test.Input, w)

	neuronOutput := neuronActivationFunction(test.Input, w)
	activated := test.Expect == neuronOutput

	if !activated {
		w = adjustWeights(w, test.Expect, neuronOutput)
	}
	activationScore := Output(0.0)
	if activated {
		activationScore = 1.0
	}
	epochResults = append(epochResults, activationScore)
	fmt.Println(">>", epochResults)
	displayTestFooter(test.Input, neuronOutput, activated) //, epochResults)

	return w, epochResults
}

func adjustWeights(w weights, expect, neuronOutput Output) weights {
	for i, _ := range w {
		w[i] = w[i] + (rate * float64(expect-neuronOutput))
	}
	return w
}

func neuronActivationFunction(in Inputs, w weights) Output {
	sum := in[0]*w[0] + in[1]*w[1]

	activationThreshold := 1.0 //2.0
	if sum > activationThreshold {
		return 1.0
	} else {
		return 0.0
	}
}

/*
func saveTestResult(epochResults []float64, isActivated bool) []float64 {
	activated := 0.0
	if isActivated {
		activated = 1.0
	}
	epochResults = append(epochResults, activated)
	return epochResults
}
*/
func random() float64 {
	rand.Seed(time.Now().UnixNano())
	r := rand.Float64() * 2
	return r
}

func initWeights(max int) weights {
	w := make(weights, max)
	for i := 0; i < max; i++ {
		rndWeight := random()
		w[i] = rndWeight
	}
	return w
}

type TrainingSet []TrainingCase

type TrainingCase struct {
	Input  Inputs
	Expect Output
}

type Inputs []float64
type Output float64

func getTrainingData() TrainingSet {
	return []TrainingCase{
		TrainingCase{Input: Inputs{0.0, 0.0}, Expect: 0.0},
		TrainingCase{Input: Inputs{0.0, 1.0}, Expect: 1.0},
		TrainingCase{Input: Inputs{1.0, 0.0}, Expect: 1.0},
		TrainingCase{Input: Inputs{1.0, 1.0}, Expect: 1.0},
	}
}

func displayEpochHeader(epochId int) {
	fmt.Println("")
	fmt.Printf("##### Epoch #%d #####\n", epochId)
	fmt.Println("-----------------------------------------------")
}

func displayTestHeader(inputs, weights []float64) {
	fmt.Printf("Inputs: \t%+v\n", joinFloats(inputs, "\t"))
	fmt.Printf("Weights:\t%+v\n", joinFloats(weights, "\t"))
}

func displayTestFooter(testInputs Inputs, neuronOutput Output, isActivated bool) {
	//func displayTestFooter(testInputs Inputs, neuronOutput Output, isActivated bool, r []float64) {
	fmt.Printf("Result: \t#%+v : #%+v\n", neuronOutput, isActivated)
	//fmt.Printf("--> %+v\n", r)
	fmt.Println("-----------------------------------------------")
}

func displayEpochFooter(epochId int, epochSuccess float64) {
	fmt.Printf("##### Epoch #%d : #%.1f #####\n", epochId, epochSuccess)
	fmt.Println("-----------------------------------------------")
}

func joinFloats(f []float64, d string) string {
	csv := ""
	for i, j := range f {
		csv += strconv.FormatFloat(j, 'f', 1, 64)
		if i < len(f) {
			csv += d
		}
	}
	return csv
}
