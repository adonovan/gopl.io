package eval

import (
	"fmt"
	"math"
)

// Expr 은 산술 표현식
type Expr interface {
	Eval(env Env) float64
}

// Var 는 변수를 식별
type Var string

func (v Var) Eval(env Env) float64 {
	return env[v]
}

// literal 수치형 상수
type literal float64

func (l literal) Eval(_ Env) float64 {
	return float64(l)
}

// unary 는 단항 연산자 표현식
type unary struct {
	op rune
	x  Expr
}

func (u unary) Eval(env Env) float64 {
	switch u.op {
	case '+':
		return +u.x.Eval(env)
	case '-':
		return -u.x.Eval(env)
	}
	panic(fmt.Sprintf("unsupported unary operator: %q", u.op))
}

// binary 는 이항 연산자 표현식
type binary struct {
	op   rune
	x, y Expr
}

func (b binary) Eval(env Env) float64 {
	switch b.op {
	case '+':
		return b.x.Eval(env) + b.y.Eval(env)
	case '-':
		return b.x.Eval(env) - b.y.Eval(env)
	case '*':
		return b.x.Eval(env) * b.y.Eval(env)
	case '/':
		return b.x.Eval(env) / b.y.Eval(env)
	}
	panic(fmt.Sprintf("unsupported binary operator: %q", b.op))
}

// call 음 함수 호출 표현식
type call struct {
	fn   string
	args []Expr
}

func (c call) Eval(env Env) float64 {
	switch c.fn {
	case "pow":
		return math.Pow(c.args[0].Eval(env), c.args[1].Eval(env))
	case "sin":
		return math.Sin(c.args[0].Eval(env))
	case "sqrt":
		return math.Sqrt(c.args[0].Eval(env))
	}
	panic(fmt.Sprintf("unsupported function operator: %q", c.fn))
}

// Env 는 변수명을 값과 매핑한다
type Env map[Var]float64
