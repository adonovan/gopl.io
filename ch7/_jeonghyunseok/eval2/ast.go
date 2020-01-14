package eval2

// Expr 라는 인터페이스는
// 메쏘드로 Eval 과 Check 를 가지고 있다
type Expr interface {
	// Eval 은 Expr 의 값을 환경변수에서 가져온다.
	Eval(env Env) float64
	// Expr 의 에러를 리포팅 해주고, Var 를 set 해준다.
	Check(vars map[Var]bool) error
}

// Var 는 string 이다. 변수를 정의한다.
type Var string

// literal 은 상수 숫자값이다.
type literal float64

// unary 는 unary operator 를 의미한다. -x 와 같은 경우의 -
type unary struct {
	op rune //'+', '-' 같은 문자이다.
	x  Expr // 이 문자에 대한 실제 동작을 여기 정의할 수 있겠다.
}

// binary 는 binary operator 를 의미한다.
type binary struct {
	op   rune
	x, y Expr
}

type call struct {
	fn   string // 함수명이다. pow, sin, sqrt 같은거
	args []Expr
}
