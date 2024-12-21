package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	registerA, registerB, registerC, program := readInput("input.txt")
	outputs := executeProgram(registerA, registerB, registerC, program)

	stringsOutputs := []string{}
	for _, output := range outputs {
		stringsOutputs = append(stringsOutputs, strconv.Itoa(output))
	}

	joinedOutput := strings.Join(stringsOutputs, ",")
	fmt.Println("Output: ", joinedOutput)

	//* I'm not ashamed to admit that I'm not sure why this works, but it does
	//* The << 3 trick supposedly is the key here, and why I don't really understand
	lowestRegisterA := getLowestRegisterA(0, registerB, registerC, 1, program)
	fmt.Println("Lowest Register A:", lowestRegisterA)
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

func executeProgram(registerA, registerB, registerC int, program []int) []int {
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

	// fmt.Println("Outputs: ", outputs)

	return outputs
}

func executeInstruction(instruction int, operand int, registerA, registerB, registerC int, i int) (int, int, int, int, []int) {
	outputs := []int{}
	switch instruction {
	case 0:
		comboOperand := getComboOperand(operand, registerA, registerB, registerC)
		registerA = adv(registerA, comboOperand)
	case 1:
		registerB = bxl(registerB, operand)
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

func bxc(regB, regC int) int {
	return regB ^ regC
}

func getLowestRegisterA(regA, regB, regC int, compareIndex int, program []int) int {
    result := make(map[int]bool)
    var smallest int = -1

    for n := 0; n < 8; n++ {
        A2 := (regA << 3) | n
        output := executeProgram(A2, regB, regC, program)
        if compareOutputs(output, program[len(program)-compareIndex:]) {
            if compareOutputs(output, program) {
                result[A2] = true
            } else {
                possible := getLowestRegisterA(A2, regB, regC, compareIndex+1, program)
                if possible != 0 {
                    result[possible] = true
                }
            }
        }
    }

    for key := range result {
        if smallest == -1 || key < smallest {
            smallest = key
        }
    }

    if smallest == -1 {
        return 0
    }
    return smallest
}

func compareOutputs(output, target []int) bool {
    if len(output) != len(target) {
        return false
    }
    for i := range output {
        if output[i] != target[i] {
            return false
        }
    }
    return true
}