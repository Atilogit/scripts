package main

import (
	"fmt"
	"math"
	"reflect"
	"strings"
)

// MODIFY HERE

// xor: !=
const exprString = "y"

func expr(P, Q bool) bool {
	return imply(imply(!P, Q) && imply(!P, !Q), P)
}

// DONT MODIFY PAST THIS LINE

func imply(a, b bool) bool {
	if !a {
		return true
	}
	return b
}

func main() {
	v := reflect.ValueOf(expr)
	fmt.Printf("\\begin{tabular}{ %s|c }\n", strings.Repeat("c|", v.Type().NumIn()))

	names := make([]string, v.Type().NumIn())
	for i := range names {
		names[i] = string(rune(int('a') + i))
	}
	fmt.Printf("  %s & $%s$ \\\\\n  \\hline\n", strings.Join(names, " & "), exprString)

	rows := int64(math.Pow(2, float64(v.Type().NumIn())))
	for i := int64(0); i < rows; i++ {
		res, bin := call(v, i)

		strs := make([]string, len(bin))
		for s := range strs {
			strs[s] = string(bin[s])
		}

		fmt.Printf("  %s & %d \\\\\n", strings.Join(strs, " & "), b2i(res))
	}
	fmt.Printf("\\end{tabular}\n")
}

func b2i(b bool) int {
	if b {
		return 1
	}
	return 0
}

func call(v reflect.Value, in int64) (bool, string) {
	bin := fmt.Sprintf("%064b", in)
	input := make([]reflect.Value, v.Type().NumIn())
	for i := range input {
		input[i] = reflect.ValueOf(bin[len(bin)-1-i] == '1')
	}
	res := v.Call(input)

	return res[0].Bool(), bin[64-v.Type().NumIn():]
}
