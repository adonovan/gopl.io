package eval2

import (
	"fmt"
	"strings"
)

// Var 는 string 이다. 이걸 체크하면
// vars 라는 맵에 이게 있는지를 확인하는 것일텐데 이 구현은
// 그냥 true 로 만들어줄 뿐이다.
func (v Var) Check(vars map[Var]bool) error {
	vars[v] = true
	return nil
}

//  이건 의미없음
func (literal) Check(vars map[Var]bool) error {
	return nil
}

// unary Check 는 뭘까?
// 첫째 u.op 는 연산자 문자 하나인데 + 또는 - 인지를 체크한다
// 그다음에는 Expr 을 체크한다.
func (u unary) Check(vars map[Var]bool) error {
	if !strings.ContainsRune("+-", u.op) {
		return fmt.Errorf("unexpected unary op %q", u.op)
	}
	return u.x.Check(vars)
}

// binary 도 유사하다. 다믄 Expr 가 두 개 있으니 둘 다 체크해준다.
func (b binary) Check(vars map[Var]bool) error {
	if !strings.ContainsRune("+-*/", b.op) {
		return fmt.Errorf("unexpected binary op %q", b.op)
	}
	if err := b.x.Check(vars); err != nil {
		return err
	}
	return b.y.Check(vars)
}

// call 은 일단 pow, sin sqrt 중 하나인지 확인하고
// 그다음에는 피연산자 arity 개수가 맞는 지 확인한다.
// 마지막으로 피연산자 각각, Expr 을 체크해준다.
func (c call) Check(vars map[Var]bool) error {
	arity, ok := numParams[c.fn]
	if !ok {
		return fmt.Errorf("unknown function %q", c.fn)
	}
	if len(c.args) != arity {
		return fmt.Errorf("call to %s has %d args, want %d",
			c.fn, len(c.args), arity)

	}
	for _, arg := range c.args {
		if err := arg.Check(vars); err != nil {
			return err
		}
	}
	return nil
}

var numParams = map[string]int{
	"pow":  2,
	"sin":  1,
	"sqrt": 1,
}
