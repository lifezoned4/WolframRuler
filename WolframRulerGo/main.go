// Cells project main.go
package main

import (
	"fmt"
	"math/big"
)

const (
	STAR = '*'
	L    = '0'
	S    = 'S'
	SNN  = 'Z'
	N    = 'N'
	NN   = '+'
	B    = 'B'
	T1   = '1'
	T2   = '2'
	T3   = '3'
	T4   = '4'
	T5   = '5'
	T6   = '6'
	T7   = '7'
	T8   = '8'
	T9   = '9'
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

func AutomateSolNB(nb Neighbourhood) rune {
	rules := []Rule{

		Rule{nb: MakeNeighbourhood(STAR, S, L, STAR, STAR), follow: S},
		Rule{nb: MakeNeighbourhood(STAR, STAR, S, L, STAR), follow: L},
		Rule{nb: MakeNeighbourhood(STAR, STAR, S, B, STAR), follow: S},
		Rule{nb: MakeNeighbourhood(STAR, STAR, S, T1, STAR), follow: S},
		Rule{nb: MakeNeighbourhood(STAR, STAR, S, T2, STAR), follow: S},
		Rule{nb: MakeNeighbourhood(STAR, STAR, S, T3, STAR), follow: S},
		Rule{nb: MakeNeighbourhood(STAR, STAR, S, T4, STAR), follow: S},
		Rule{nb: MakeNeighbourhood(STAR, STAR, S, T5, STAR), follow: S},
		Rule{nb: MakeNeighbourhood(STAR, STAR, S, T6, STAR), follow: S},
		Rule{nb: MakeNeighbourhood(STAR, STAR, S, T7, STAR), follow: S},
		Rule{nb: MakeNeighbourhood(STAR, STAR, S, T8, STAR), follow: S},
		Rule{nb: MakeNeighbourhood(STAR, STAR, S, T9, STAR), follow: S},
		Rule{nb: MakeNeighbourhood(STAR, STAR, S, S, STAR), follow: S},
		Rule{nb: MakeNeighbourhood(STAR, B, S, STAR, STAR), follow: L},
		Rule{nb: MakeNeighbourhood(STAR, B, L, STAR, STAR), follow: S},
		Rule{nb: MakeNeighbourhood(STAR, B, L, S, STAR), follow: L},
		Rule{nb: MakeNeighbourhood(STAR, STAR, B, S, STAR), follow: T1},
		Rule{nb: MakeNeighbourhood(STAR, STAR, B, STAR, STAR), follow: B},
		Rule{nb: MakeNeighbourhood(STAR, STAR, T1, L, STAR), follow: T2},
		Rule{nb: MakeNeighbourhood(STAR, STAR, T2, L, STAR), follow: T3},
		Rule{nb: MakeNeighbourhood(STAR, STAR, T3, L, STAR), follow: T4},
		Rule{nb: MakeNeighbourhood(STAR, STAR, T4, L, STAR), follow: T5},
		Rule{nb: MakeNeighbourhood(STAR, STAR, T5, L, STAR), follow: T6},
		Rule{nb: MakeNeighbourhood(STAR, STAR, T6, L, STAR), follow: T7},
		Rule{nb: MakeNeighbourhood(STAR, STAR, T7, L, STAR), follow: T8},
		Rule{nb: MakeNeighbourhood(STAR, STAR, T8, L, STAR), follow: T9},
		Rule{nb: MakeNeighbourhood(STAR, STAR, T9, L, STAR), follow: B},
		Rule{nb: MakeNeighbourhood(STAR, STAR, STAR, STAR, STAR), follow: L}}

	t := STAR
	for _, r := range rules {

		t = r.Eval(nb)
		if t != STAR {
			return t
		}
	}

	return STAR
}

func UsedAutomate(nb Neighbourhood) rune {
	return AutomateSolNB(nb)
	// return AutomateWoNB(nb, 110)
}

func Automate(state []rune) []rune {
	bstate := make([]rune, len(state))

	for i := range state {
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
		state[i] = rt // r.Intn(2)
	}
	state[0] = B

	for i := 0; i < 500; i++ {
		for _, r := range state {
			fmt.Print(string(r))
		}
		fmt.Printf("\n")
		state = Automate(state)
	}

	fmt.Println("END")
}
