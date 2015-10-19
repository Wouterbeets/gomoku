package nn

import (
	"fmt"
	"math/rand"
)

type neuron struct {
	input  []chan float64
	output []chan float64
	weight []float64
	bias   float64
	id     int
}

func (n *neuron) mergeInputChannels() chan float64 {
	ch := make(chan float64)
	for v, k := range n.input {
		go func()
	}
}

func (n *neuron) activate() {
	 TODO: somewhere here i need to use an activation function
	lenInp := len(n.input)
	lenOutp := len(n.output)
		c := n.mergeInputChannels()
	for {
		resp := float64(0)
		for i := 0; i < lenInp; i++ {
			resp += <-c
		}
		for i := 0; i < lenOut; i++ {
			n.output[i] <- resp
		}
	}
}

func (n *neuron) String() string {
	str := ""
	for k, v := range n.weight {
		str += fmt.Sprintln("\t\tweight aplied to neur", k, "is", v)
	}
	for k, v := range n.input {
		str += fmt.Sprintln("\t\tchannel", k, "is", v)
	}
	return str
}

func newNeuron(input []chan float64, output []chan float64) *neuron {
	n := &neuron{
		input:  input,
		output: output,
		weight: make([]float64, len(input), len(input)),
	}
	for k, _ := range n.weight {
		n.weight[k] = rand.NormFloat64()
	}
	return n
}
