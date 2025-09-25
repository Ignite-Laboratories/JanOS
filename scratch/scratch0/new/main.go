package main

import (
	"fmt"
)

func main() {
	MyFunc(&MyNeuron{})
}

type MyNeuron struct {
}

func (n MyNeuron) Impulse() {

}

func (n MyNeuron) Reveal() any {
	return n
}

func MyFunc(operand any) {
	switch operand.(type) {
	case Neuron:
		fmt.Println("here")
	}
}

type Neuron interface {
	Impulse()
	Reveal() any
}
