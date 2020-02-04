package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	eval "gopl.io/ch7/_beomyonglee/26.ex_7.14"
)

const assignment_error = 2

func main() {
	exitCode := 0
	stdin := bufio.NewScanner(os.Stdin)
	fmt.Printf("Expression: ")
	stdin.Scan()
	exprStr := stdin.Text()
	fmt.Printf("Variables (<var>=<val>, eg: x=3): ")
	stdin.Scan()
	envStr := stdin.Text()
	if stdin.Err() != nil {
		fmt.Fprintln(os.Stderr, stdin.Err())
		os.Exit(1)
	}

	env := eval.Env{}
	assignments := strings.Fields(envStr)
	for _, a := range assignments {
		fields := strings.Split(a, "=")
		if len(fields) != 2 {
			fmt.Fprintf(os.Stderr, "bad assignment: %s\n", a)
			exitCode = assignment_error
		}
		ident, valStr := fields[0], fields[1]
		val, err := strconv.ParseFloat(valStr, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "bad value for %s, using zero: %s\n", ident, err)
			exitCode = assignment_error
		}
		env[eval.Var(ident)] = val
	}

	expr, err := eval.Parse(exprStr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "bad expression: %s\n", err)
		os.Exit(1)
	}
	fmt.Println(expr.Eval(env))
	os.Exit(exitCode)
}
