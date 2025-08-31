package main

import (
	"fmt"
	"strings"
)

var b *strings.Builder
var hex = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "A", "B", "C", "D", "E", "F"}
var bases = []string{
	"nil",
	"one",
	"binary",
	"ternary",
	"quarternary",
	"quinary",
	"senary",
	"septenary",
	"octal",
	"nonary",
	"decimal",
	"undecimal",
	"duodecimal",
	"tridecimal",
	"tetradecimal",
	"pentadecimal",
	"hexadecimal",
}

func fprintf(format string, a ...any) {
	_, _ = fmt.Fprintf(b, format, a...)
}

func main() {
	b = new(strings.Builder)

	for i := 2; i <= 256; i++ {
		emit(i)
	}
	fmt.Println(b.String())
}

func emit(i int) {
	name := fmt.Sprintf("base-%d", i)
	if i < len(bases)-1 {
		name = bases[i]
	}

	fprintf("{\n")
	fprintf("\t\"name\":\"Base%d\",\n", i)
	fprintf("\t\"docs\":\"Base%d represents a single %s placeholder.\\n\\nNOTE: All 𝑡𝑖𝑛𝑦 placeholder lexemes are hexadecimal.\",\n", i, name)
	fprintf("\t\"nameSet\":\"Base%dDigits\",\n", i)
	fprintf("\t\"set\": [\n\t\t")

	var ii = 0
	var iii = 0
	for j := 0; j < i; j++ {
		if i <= 16 {
			fprintf("\"%s\"", hex[ii])
		} else {
			fprintf("\"%s%s\"", hex[iii], hex[ii])
		}
		if j < i-1 {
			fprintf(",")
		}

		ii++
		if ii >= 16 {
			iii++
			ii = 0
		}
		if iii >= 16 {
			iii = 0
		}
	}

	fprintf("\n\t]\n")
	fprintf("},\n")
}
