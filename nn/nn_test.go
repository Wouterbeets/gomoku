package nn

import (
	"testing"
)

func TestNewNet(test *testing.T) {
	var tests = []struct {
		inputNeurons  int
		hiddenNeurons int
		hiddenLayers  int
		outputNeurons int
	}{
		{
			inputNeurons:  2,
			hiddenNeurons: 10,
			hiddenLayers:  2,
			outputNeurons: 1,
		},
		{
			inputNeurons:  19 * 19,
			hiddenNeurons: 30,
			hiddenLayers:  2,
			outputNeurons: 1,
		},
	}
	for _, t := range tests {
		n := NewNet(t.inputNeurons, t.hiddenNeurons, t.hiddenLayers, t.outputNeurons)
		_ = n
	}

}
