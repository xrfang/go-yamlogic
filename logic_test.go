package yamlogic

import (
	"fmt"
	"testing"
)

const (
	EXPRESSION = `
---
and:
- item1
- or: [item2, item3]
- not: [~extra]
`
	EXPMULTI = `
---
all_of:
- item1
- 2_of:
  - item2
  - item3
  - item4
`
)

func TestLoad(t *testing.T) {
	_, err := Parse(EXPRESSION)
	if err != nil {
		t.Logf("Parse() failed: %v", err)
		t.Fail()
		return
	}
	t.Log("Parse() successful")
}

func TestEvalTrue(t *testing.T) {
	x, _ := Parse(EXPRESSION)
	if !x.Eval([]string{"item1", "item3"}) {
		t.Log("expected true, got false.")
		t.Fail()
		return
	}
	t.Log("evaluation successful")
}

func TestEvalFail1(t *testing.T) {
	x, _ := Parse(EXPRESSION)
	if x.Eval([]string{"item2"}) {
		t.Log("expected false, got true.")
		t.Fail()
		return
	}
	t.Log("evaluation successful")
}

func TestEvalFail2(t *testing.T) {
	x, _ := Parse(EXPRESSION)
	if x.Eval([]string{"item1", "item2", "extra_item"}) {
		t.Log("expected false, got true.")
		t.Fail()
		return
	}
	t.Log("evaluation successful")
}

func TestMultiSelection(t *testing.T) {
	x, _ := Parse(EXPMULTI)
	if !x.Eval([]string{"item1", "item2", "item4"}) {
		t.Log("expected false, got true.")
		t.Fail()
		return
	}
	t.Log("evaluation successful")
}

func ExampleLoad() {
	ex, _ := Parse(`
---
and:
- item1
- or: [item2, item3]
- not: [~extra]`)
	fmt.Println(ex.String())
	//Output:
	//all_of:
	//- item1
	//- any_of:
	//   - item2
	//   - item3
	//- none_of:
	//   - ~extra
}
func ExampleExpression_Eval() {
	ex, err := Parse(`---
and:
- item1
- or: [item2, item3]
- not: [~extra]`)
	if err != nil {
		fmt.Println("invalid expression:", err)
		return
	}
	fmt.Println(ex.Eval([]string{"item2", "item3", "extra_item"}))
	//Output: false
}
