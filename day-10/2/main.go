package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	z3 "github.com/aclements/go-z3/z3"
)

// solveMachine models a machine as an integer-linear program: each button is
// a non-negative integer variable, every counter contributes an equality
// constraint requiring the weighted sum of button presses to reach its target,
// and we minimize the total number of presses. The Go bindings lack the
// Optimize API, so we binary search the objective by repeatedly asserting
// “total <= mid” with Push/Pop and checking satisfiability.
func solveMachine(ctx *z3.Context, target []int, buttons [][]int) (int, error) {
	intSort := ctx.IntSort()
	zero := ctx.FromInt(0, intSort).(z3.Int)
	solver := z3.NewSolver(ctx)

	vars := make([]z3.Int, len(buttons))
	for i := range buttons {
		v := ctx.FreshConst("btn", intSort).(z3.Int)
		solver.Assert(v.GE(zero))
		vars[i] = v
	}

	// Counter constraints: each counter sees the sum of every button
	// variable that touches it (times the multiplicity) and must equal
	// the demanded target.
	for counterIdx, want := range target {
		sum := zero
		for btnIdx, btn := range buttons {
			coeff := 0
			for _, idx := range btn {
				if idx == counterIdx {
					coeff++
				}
			}
			if coeff == 0 {
				continue
			}
			term := vars[btnIdx]
			if coeff > 1 {
				term = term.Mul(ctx.FromInt(int64(coeff), intSort).(z3.Int))
			}
			sum = sum.Add(term)
		}
		targetVal := ctx.FromInt(int64(want), intSort).(z3.Int)
		solver.Assert(sum.Eq(targetVal))
	}

	// Total presses = sum of all button variables.
	total := zero
	for _, v := range vars {
		total = total.Add(v)
	}

	if sat, err := solver.Check(); err != nil {
		return 0, err
	} else if !sat {
		return 0, fmt.Errorf("machine is unsatisfiable")
	}

	upper := 0
	for _, t := range target {
		upper += t
	}
	low, high := 0, upper
	best := upper
	// Binary search the optimal objective by temporarily constraining the
	// sum of presses and asking Z3 if a solution exists under that bound.
	for low <= high {
		mid := (low + high) / 2
		solver.Push()
		bound := total.LE(ctx.FromInt(int64(mid), intSort).(z3.Int))
		solver.Assert(bound)
		sat, err := solver.Check()
		solver.Pop()
		if err != nil {
			return 0, err
		}
		if sat {
			best = mid
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return best, nil
}

func main() {
	ctx := z3.NewContext(nil)

	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var sum uint64
	for scanner.Scan() {
		var buttons [][]int
		line := scanner.Text()
		for i, s := range strings.Split(line, " ") {
			if i == 0 {
				// skip targets for part two
				continue
			}
			if i == len(strings.Split(line, " "))-1 {
				s = s[1 : len(s)-1]
				parts := strings.Split(s, ",")
				targets := make([]int, 0, len(parts))
				for _, part := range parts {
					target, _ := strconv.Atoi(part)
					targets = append(targets, target)
				}
				sol, err := solveMachine(ctx, targets, buttons)
				if err != nil {
					log.Fatal(err)
				}
				sum += uint64(sol)
				continue
			}
			// construct the buttons array
			s = s[1 : len(s)-1]
			parts := strings.Split(s, ",")
			btn := []int{}
			for _, part := range parts {
				x, _ := strconv.Atoi(part)
				btn = append(btn, x)
			}
			buttons = append(buttons, btn)
		}
	}
	fmt.Println(sum)
}
