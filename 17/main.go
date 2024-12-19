package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	// registerA, registerB, registerC, program := readInput("test.txt")
	executeProgram(readInput("mini.txt"))
}

func readInput(filename string) (int, int, int, []int) {
	content, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	parts := strings.Split(string(content), "\n")

	registerAString := strings.Split(parts[0], ": ")[1]
	registerBString := strings.Split(parts[1], ": ")[1]
	registerCString := strings.Split(parts[2], ": ")[1]
	//* skip empty line
	programString := strings.Split(parts[4], ": ")[1]

	registerA, _ := strconv.Atoi(registerAString)
	registerB, _ := strconv.Atoi(registerBString)
	registerC, _ := strconv.Atoi(registerCString)

	program := []int{}
	programParts := strings.Split(programString, ",")
	for _, part := range programParts {
		partInt, _ := strconv.Atoi(part)
		program = append(program, partInt)
	}

	return registerA, registerB, registerC, program
}

func executeProgram(registerA, registerB, registerC int, program []int) {
	outputs := []int{}
	for i := 0; i < len(program); i++ {
		instruction := program[i]
		i++
		operand := program[i]

		a, b, c, idx, outs := executeInstruction(instruction, operand, registerA, registerB, registerC, i)
		registerA = a
		registerB = b
		registerC = c
		i = idx
		outputs = append(outputs, outs...)
	}

	fmt.Println("A", registerA)
	fmt.Println("B", registerB)
	fmt.Println("C", registerC)
	fmt.Println("outputs", outputs)
}

func executeInstruction(instruction int, operand int, registerA, registerB, registerC int, i int) (int, int, int, int, []int) {
	outputs := []int{}
	switch instruction {
	case 0:
		comboOperand := getComboOperand(operand, registerA, registerB, registerC)
		registerA = adv(registerA, comboOperand)
	case 1:
		registerB = bxl(registerA, operand)
	case 2:
		comboOperand := getComboOperand(operand, registerA, registerB, registerC)
		registerB = bst(comboOperand)
	case 3:
		return registerA, registerB, registerC, jnz(registerA, operand, i), outputs
	case 4:
		registerB = bxc(registerB, registerC)
	case 5:
		comboOperand := getComboOperand(operand, registerA, registerB, registerC)
		outputs = append(outputs, bst(comboOperand))

		return registerA, registerB, registerC, i, outputs
	case 6:
		comboOperand := getComboOperand(operand, registerA, registerB, registerC)
		registerB = adv(registerA, comboOperand)
	case 7:
		comboOperand := getComboOperand(operand, registerA, registerB, registerC)
		registerC = adv(registerA, comboOperand)
	}

	return registerA, registerB, registerC, i, outputs
}

func getComboOperand(literal int, a, b, c int) int {
	switch literal {
	case 0, 1, 2, 3:
		return literal
	case 4:
		return a
	case 5:
		return b
	case 6:
		return c
	// case 7:
	// 	return a
	}

	return 0
}

func adv(regA, operand int) int {
	denominator := int(math.Pow(2, float64(operand)))
	return regA / denominator
}

func bxl(regB, operand int) int {
	return regB ^ operand
}

func bst(operand int) int {
	return (operand % 8 + 8) % 8
}

func jnz(regA, operand, i int) int {
	if regA != 0 {
		return operand - 1
	}

	return i
}

func bxc(regA, regB int) int {
	return regA ^ regB
}