package main

import (
	"fmt"
	"math"
	"reflect"
	"strings"
)

// MODIFY HERE

// xor: !=
const exprString = "a \\lxor b"

func expr(a, b bool) bool {
	return a != b
}

// DONT MODIFY PAST THIS LINE

func main() {
	v := reflect.ValueOf(expr)
	fmt.Printf("\\begin{tabular}{ |%s|c| }\n  \\hline\n", strings.Repeat("c|", v.Type().NumIn()))

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
	fmt.Printf("  \\hline\n\\end{tabular}\n")
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
