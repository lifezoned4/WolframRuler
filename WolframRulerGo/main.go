// Cells project main.go
package main

import (
	"fmt"
	"math/big"
)

const (
	L    = '0'
	S    = 'S'
	STAR = '*'
)

type Rule struct {
	nb     Neighbourhood
	follow rune
}

type Neighbourhood struct {
	minusTwo rune
	minusOne rune
	self     rune
	plusOne  rune
	plusTwo  rune
}

func GetNeighbourhood(state []rune, at int) Neighbourhood {
	var nb = new(Neighbourhood)

	nb.self = state[at]
	if at < 1 {
		nb.minusTwo = state[len(state)-2]
		nb.minusOne = state[len(state)-1]
	} else if at < 2 {
		nb.minusTwo = state[len(state)-1]
		nb.minusOne = state[at-1]
	} else {
		nb.minusTwo = state[at-2]
		nb.minusOne = state[at-1]
	}

	if at >= len(state)-2 {
		nb.plusOne = state[0]
		nb.plusTwo = state[1]
	} else if at >= len(state)-1 {
		nb.plusOne = state[at+1]
		nb.plusTwo = state[0]
	} else {
		nb.plusOne = state[at+1]
		nb.plusTwo = state[at+2]
	}

	return *nb
}

func SubWoRulerNB(nb Neighbourhood, num *big.Int) rune {
	var bit0 rune = L
	if num.Bit(0) == 1 {
		bit0 = S
	}

	var bit1 rune = L
	if num.Bit(1) == 1 {
		bit1 = S
	}

	var bit2 rune = L
	if num.Bit(2) == 1 {
		bit2 = S
	}

	if nb.minusOne == bit2 && nb.self == bit1 && nb.plusOne == bit0 {
		return S
	} else {
		return STAR
	}
}

func AutomateWoNB(nb Neighbourhood, rule int) rune {
	rulenum := big.NewInt(int64(rule))

	var tnum rune = 0

	for i := 0; i < 8; i++ {
		tnum = SubWoRulerNB(nb, big.NewInt(int64(i)))
		if tnum != STAR {
			if rulenum.Bit(i) == 1 {
				return S
			} else {
				return L
			}
		}
	}
	return STAR
}

func MakeNeighbourhood(minusTwo rune, minusOne rune, self rune, plusOne rune, plusTwo rune) Neighbourhood {
	return Neighbourhood{minusTwo: minusTwo,
		minusOne: minusOne,
		self:     self,
		plusOne:  plusOne,
		plusTwo:  plusTwo}
}

func (r *Rule) Eval(nb Neighbourhood) rune {
	if (r.nb.minusTwo == nb.minusTwo || r.nb.minusTwo == STAR) &&
		(r.nb.minusOne == nb.minusOne || r.nb.minusOne == STAR) &&
		(r.nb.self == nb.self || r.nb.self == STAR) &&
		(r.nb.plusOne == nb.plusOne || r.nb.plusOne == STAR) &&
		(r.nb.plusTwo == nb.plusTwo || r.nb.plusTwo == STAR) {
		return r.follow
	}
	return STAR
}

func UsedAutomate(nb Neighbourhood) rune {
	// return AutomateSolNB(nb)
	return AutomateWoNB(nb, 30)
}

func Automate(state []rune) []rune {
	bstate := make([]rune, len(state))

	for i, _ := range state {
		nb := GetNeighbourhood(state, i)
		bstate[i] = UsedAutomate(nb)
	}

	return bstate
}

func main() {
	state := make([]rune, 100)

	for i := range state {
		rt := L
		if i == len(state)/2 {
			rt = S
		}
		state[i] = rt
	}

	for i := 0; i < 50; i++ {
		for _, r := range state {
			fmt.Print(string(r))
		}
		fmt.Printf("\r\n")
		state = Automate(state)
	}

	fmt.Println("END")
}
