package lab2

import (
	"bytes"
	"strings"

	. "gopkg.in/check.v1"
)

func (s *ComputeHandler) TestComputeHandler(c *C) {
	b := bytes.NewBuffer(make([]byte, 0))

	handler := ComputeHandler{
		Input:  strings.NewReader("- 6 2"),
		Output: b,
	}
	err := handler.Compute()

	c.Assert(err, Equals, nil)
	c.Assert(b.String(), Equals, "6 2 -")
}

func (s *ComputeHandler) TestComputeHandlerHard(c *C) {
	b := bytes.NewBuffer(make([]byte, 0))

	handler := ComputeHandler{
		Input:  strings.NewReader("/ 5 * + 9 8 3"),
		Output: b,
	}
	err := handler.Compute()

	c.Assert(err, Equals, nil)
	c.Assert(b.String(), Equals, "5 9 8 + 3 * /")
}

func (s *ComputeHandler) TestComputeHandlerErrorExceedOperands(c *C) {
	b := bytes.NewBuffer(make([]byte, 0))

	handler := ComputeHandler{
		Input:  strings.NewReader("+ 5 5 5 5"),
		Output: b,
	}
	err := handler.Compute()

	c.Assert(err, NotNil)
}

func (s *ComputeHandler) TestComputeHandlerErrorExceedOperators(c *C) {
	b := bytes.NewBuffer(make([]byte, 0))

	handler := ComputeHandler{
		Input:  strings.NewReader("- - - - 1000"),
		Output: b,
	}
	err := handler.Compute()

	c.Assert(err, NotNil)
}
