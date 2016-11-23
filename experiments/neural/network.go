package main

import (
	"fmt"
	"math/rand"
	// "time"
)

type Network struct {
	layers [][]float64
	b_nums []float64
}

func (n *Network) generateLayers(count_layers, count_neurons int) {
	for i := 0; i < count_layers; i++ {
		layer := make([]float64, count_neurons)
		for j := 0; j < count_neurons; j++ {
			layer[j] = rand.Float64()*2 - 1
		}
		n.layers = append(n.layers, layer)
	}

	for i := 0; i < count_layers; i++ {
		n.b_nums = append(n.b_nums, rand.Float64()*2-1)
	}
}

func (n *Network) think(inputs []float64) {
	tmp := make([]float64, len(inputs))
	for layer := range n.layers {
		for i := range inputs {
			tmp[i] = n.layers[layer][i] * inputs[i]
		}
	}

	fmt.Println(tmp)
	// for i := range inputs {
	// 	fmt.Println(i, inputs[i])
	// }
}

func newNetwork() Network {
	// rand.Seed(time.Now().UnixNano())
	rand.Seed(1)
	net := Network{}
	net.generateLayers(2, 2)
	return net
}
