package lab2

import (
	"bytes"
	"io"
)

// ComputeHandler should be constructed with input io.Reader and output io.Writer.
// Its Compute() method should read the expression from input and write the computed result to the output.
type ComputeHandler struct {
	// TODO: Add necessary fields.
	Input  io.Reader
	Output io.Writer
}

func (ch *ComputeHandler) Compute() error {
	input, err := io.ReadAll(ch.Input)

	if err != nil {
		return err
	}

	trimInp := bytes.Trim(input, "\x00")
	res, err := PrefixToPostfix(string(trimInp))
	if err != nil {
		return err
	}

	ch.Output.Write([]byte(res))
	return nil
}
