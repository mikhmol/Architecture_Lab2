package lab2_test

import (
	"fmt"

	lab2 "github.com/mikhmol/Architecture_Lab2"
	. "gopkg.in/check.v1"
)

type PrefixToPostfixSuite struct{}

var _ = Suite(&PrefixToPostfixSuite{})

func (s *PrefixToPostfixSuite) TestPrefixToPostfix(c *C) {
	simpleExpression := "+ 1 2"
	expectedResultSimple := "1 2 +"
	resultSimple, err := lab2.PrefixToPostfix(simpleExpression)
	c.Assert(err, IsNil)
	c.Assert(resultSimple, Equals, expectedResultSimple)

	complexExpression := "* - + 3 2 1 / 1 2"
	expectedResultComplex := "3 2 + 1 - 1 2 / *"
	resultComplex, err := lab2.PrefixToPostfix(complexExpression)
	c.Assert(err, IsNil)
	c.Assert(resultComplex, Equals, expectedResultComplex)

	emptyExpression := ""
	expectedErrorEmpty := "invalid prefix expression"
	_, errEmpty := lab2.PrefixToPostfix(emptyExpression)
	c.Assert(errEmpty, NotNil)
	c.Assert(errEmpty.Error(), Equals, expectedErrorEmpty)

	invalidOperatorExpression := "+ 5 $ 3"
	expectedErrorInvalidOperator := "invalid prefix expression"
	_, errInvalidOperator := lab2.PrefixToPostfix(invalidOperatorExpression)
	c.Assert(errInvalidOperator, NotNil)
	c.Assert(errInvalidOperator.Error(), Equals, expectedErrorInvalidOperator)
}

func ExamplePrefixToPostfix() {
	res, _ := lab2.PrefixToPostfix("+ 2 2")
	fmt.Println(res)

	// Output:
	// 2 2 +
}
