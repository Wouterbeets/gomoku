package main

import (
	"fmt"
)

type Net struct {
	Layers []*layer
}

func NewNet(inputNeurons, hiddenNeurons, hiddenLayers, outputNeurons int) *Net {
	n := &Net{
		Layers: make([]*layer, hiddenLayers+2),
	}
	n.Layers[0] = newLayer(nil, inputNeurons, hiddenNeurons)
	for i := 1; i < hiddenLayers+1; i++ {
		n.Layers[i] = newLayer(n.Layers[i-1], hiddenNeurons, hiddenNeurons)
	}
	n.Layers[hiddenLayers+1] = newLayer(n.Layers[hiddenLayers], outputNeurons, 0)
	return n
}

func (n *Net) String() string {
	str := ""
	for lNum, layer := range n.Layers {
		str += fmt.Sprintln("layer", lNum, "holds\n", layer)
	}
	return str
}

type layer struct {
	neurons []*neuron
}

func newLayer(prevL *layer, numNeurons, numNextL int) *layer {
	l := &layer{
		neurons: make([]*neuron, numNeurons),
	}
	for i := 0; i < numNeurons; i++ {
		out := make([]chan float64, numNextL)
		if prevL != nil {
			inp := make([]chan float64, len(prevL.neurons))
			for k, v := range prevL.neurons {
				inp[k] = v.output[i]
			}
			l.neurons[i] = newNeuron(inp, out)
		}
		l.neurons[i] = newNeuron(nil, out)
	}
	return l
}
func (l *layer) String() string {
	str := ""
	for nNum, neur := range l.neurons {
		str += fmt.Sprintln("\tneur", nNum, "holds", neur)
	}
	return str
}

func (n *neuron) String() string {
	str := ""
	for k, v := range n.weight {
		str += fmt.Sprintln("\t\tweight aplied to neur", k, "is", v)
	}
	return str
}

type neuron struct {
	input  []chan float64
	output []chan float64
	weight []float64
	bias   float64
	id     int
}

func newNeuron(input []chan float64, output []chan float64) *neuron {
	n := &neuron{
		input:  input,
		output: output,
		weight: make([]float64, len(input)),
	}
	return n
}

func main() {
	n := NewNet(30, 30, 2, 1)
	fmt.Println(n)
}

//type con struct {
//
//}

//func newCon(shallow *layer, deep *layer) *connections
