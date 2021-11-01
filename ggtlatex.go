package main

import (
	"fmt"
	"math"
)

type GGTStep struct {
	a, b, q, r, x, y int
}

func main() {
	a, b := 0, 0
	fmt.Print("a, b = ")
	fmt.Scan(&a, &b)

	steps := ggt(a, b, []GGTStep{})
	extendGgt(steps)

	fmt.Println("\n")

	fmt.Println("Euklidischer Algorithmus:")
	fmt.Println("\\begin{enumerate}[leftmargin=1.5cm]")
	for _, s := range steps {
		fmt.Printf("\t\\item[$%d$] $= %d * %d + %d$\n", s.a, s.q, s.b, s.r)
	}
	fmt.Println("\\end{enumerate}")
	fmt.Println("Erweiterung:")
	fmt.Println("\\begin{enumerate}[leftmargin=1.5cm]")
	for _, s := range steps {
		fmt.Printf("\t\\item[$%d = $] %d * %d + %d * %d\n", steps[len(steps)-1].b, s.a, s.x, s.b, s.y)
	}
	fmt.Println("\\end{enumerate}")

	fmt.Printf("ggT$(%d, %d) = %d * %d + %d * %d = %d$\n\n", a, b, steps[0].x, a, steps[0].y, b, steps[len(steps)-1].b)
}

func ggt(a, b int, steps []GGTStep) []GGTStep {
	af, bf := float64(a), float64(b)
	q := int(math.Floor(af / bf))
	step := GGTStep{
		a: a,
		b: b,
		q: q,
		r: a - q*b,
		x: 0,
		y: 1,
	}
	steps = append(steps, step)

	if step.r == 0 {
		return steps
	}
	return ggt(step.b, step.r, steps)
}

func extendGgt(steps []GGTStep) {
	for i := len(steps) - 2; i >= 0; i-- {
		steps[i].x = steps[i+1].y
		steps[i].y = steps[i+1].x - steps[i].q*steps[i+1].y
	}
}
