// WolframRulerGo project main.go
package main

import (
	"fmt"
	"math/rand"
)

type WolframState struct {
	state bool
}

func (w WolframState) String() string {
	if w.state {
		return "X"
	} else {
		return "0"
	}
}

func (w *WolframState) setState(set bool) {
	w.state = set
}

func (w *WolframState) nextState(nb []WolframState) {
	// fmt.Print(nb)
	if !nb[0].state && !nb[1].state && !nb[2].state {
		w.state = false
	} else if !nb[0].state && nb[1].state && !nb[2].state {
		w.state = true
	}
}

const lineLen = 40

var r = rand.New(rand.NewSource(99))

func FlipCoin() bool {
	if r.Intn(3) == 1 {
		fmt.Printf("FILP\n")
		return true
	} else {
		fmt.Printf("FLOP\n")
		return false
	}
}

func main() {
	fmt.Println("Hello WolframRuler!")
	var statesT0 []WolframState = make([]WolframState, lineLen)
	var statesT1 []WolframState = make([]WolframState, lineLen)

	for i, _ := range statesT0 {
		statesT0[i].setState(FlipCoin())
		statesT1[i].setState(false)
	}

	for _, e := range statesT0 {
		fmt.Print(e)
	}
	fmt.Printf("\n")

	for i := 0; i < 100; i++ {
		for index, _ := range statesT0 {
			// fmt.Printf("%v\n", index)
			var nb = make([]WolframState, 3)
			if (index > 1) && (index < (lineLen - 1)) {
				copy(nb, statesT0[index-1:index+1])
				statesT1[index].nextState(nb)
			}
		}

		for _, e := range statesT1 {
			fmt.Print(e)
		}

		fmt.Printf("\n")

		statesT0 = statesT1
		statesT1 = make([]WolframState, lineLen)
	}
}
