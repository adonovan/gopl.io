package eval

type Expr interface {
	Eval(env Env) float64
	Check(vars map[Var]bool) error
	String() string
}

type Var string

type literal float64

type unary struct {
	op rune
	x  Expr
}

type binary struct {
	op   rune
	x, y Expr
}

type call struct {
	fn   string
	args []Expr
}

type postUnary struct {
	op rune
	x  Expr
}
