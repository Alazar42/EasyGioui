package components

import (
	"easygioui"
	"fmt"
	"strconv"
	"strings"
)

type Calculator struct {
	accumulator   float64
	currentNumber string
	operation     string
	display       string
}

func NewCalculator() *Calculator {
	return &Calculator{
		accumulator:   0,
		currentNumber: "0",
		operation:     "",
		display:       "0",
	}
}

// AppendDigit adds a digit to the current number
func (c *Calculator) AppendDigit(digit string) {
	if c.currentNumber == "0" && digit != "." {
		c.currentNumber = digit
	} else if digit == "." && strings.Contains(c.currentNumber, ".") {
		return // Don't allow multiple decimals
	} else {
		c.currentNumber += digit
	}
	easygioui.SetText("display", c.currentNumber)
}

// Digit0 through Digit9 handlers
func (c *Calculator) Digit0()  { c.AppendDigit("0") }
func (c *Calculator) Digit1()  { c.AppendDigit("1") }
func (c *Calculator) Digit2()  { c.AppendDigit("2") }
func (c *Calculator) Digit3()  { c.AppendDigit("3") }
func (c *Calculator) Digit4()  { c.AppendDigit("4") }
func (c *Calculator) Digit5()  { c.AppendDigit("5") }
func (c *Calculator) Digit6()  { c.AppendDigit("6") }
func (c *Calculator) Digit7()  { c.AppendDigit("7") }
func (c *Calculator) Digit8()  { c.AppendDigit("8") }
func (c *Calculator) Digit9()  { c.AppendDigit("9") }
func (c *Calculator) Decimal() { c.AppendDigit(".") }

// SetOperation stores current number and sets the operation
func (c *Calculator) SetOperation(op string) {
	if c.currentNumber == "" {
		return
	}
	if c.operation != "" && c.currentNumber != "" {
		c.Calculate() // Chain operations
	} else {
		num, _ := strconv.ParseFloat(c.currentNumber, 64)
		c.accumulator = num
	}
	c.operation = op
	c.currentNumber = ""
	easygioui.SetText("display", op)
}

// Operation handlers
func (c *Calculator) OpAdd() { c.SetOperation("+") }
func (c *Calculator) OpSub() { c.SetOperation("-") }
func (c *Calculator) OpMul() { c.SetOperation("*") }
func (c *Calculator) OpDiv() { c.SetOperation("/") }

// Calculate performs the stored operation
func (c *Calculator) Calculate() {
	if c.operation == "" || c.currentNumber == "" {
		return
	}
	num, _ := strconv.ParseFloat(c.currentNumber, 64)
	result := c.accumulator

	switch c.operation {
	case "+":
		result = c.accumulator + num
	case "-":
		result = c.accumulator - num
	case "*":
		result = c.accumulator * num
	case "/":
		if num != 0 {
			result = c.accumulator / num
		} else {
			easygioui.SetText("display", "Error: Div/0")
			c.operation = ""
			c.currentNumber = ""
			c.accumulator = 0
			return
		}
	}

	// Format result to remove unnecessary decimals
	resultStr := fmt.Sprintf("%g", result)
	easygioui.SetText("display", resultStr)
	c.accumulator = result
	c.currentNumber = ""
	c.operation = ""
}

// Equal handler
func (c *Calculator) Equal() {
	c.Calculate()
}

// Clear resets the calculator
func (c *Calculator) Clear() {
	c.accumulator = 0
	c.currentNumber = "0"
	c.operation = ""
	c.display = "0"
	easygioui.SetText("display", "0")
}
