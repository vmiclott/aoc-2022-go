package day11

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type binaryOperation struct {
	operation string
	values    [2]int
}

func (op binaryOperation) evaluate(item *int) error {
	switch op.operation {
	case "+":
		*item = (op.values[0]%mod + op.values[1]%mod) % mod
		return nil
	case "*":
		*item = (op.values[0] % mod * op.values[1] % mod) % mod
		return nil
	}
	return fmt.Errorf("unsupported operation %s", op.operation)
}

func parseValue(formulaElement string, fallback int) int {
	res, err := strconv.Atoi(formulaElement)
	if err != nil {
		return fallback
	}
	return res
}

func parseBinaryOperation(formula string, old int) (binaryOperation, error) {
	formulaElements := strings.Split(formula, " ")
	return binaryOperation{operation: formulaElements[1], values: [2]int{parseValue(formulaElements[0], old), parseValue(formulaElements[2], old)}}, nil
}

func monkeyItems(line string) (queue, error) {
	items := queue{}
	for _, s := range strings.Split(line[len("  Starting items: "):], ",") {
		item, err := strconv.Atoi(strings.TrimSpace(s))
		if err != nil {
			return nil, err
		}
		items.push(&item)
	}
	return items, nil
}

func monkeyOperation(line string) (operation, error) {
	formula := line[len("  Operation: new = "):]
	return func(item *int) error {
		binaryOperation, err := parseBinaryOperation(formula, *item)
		if err != nil {
			return err
		}
		return binaryOperation.evaluate(item)
	}, nil
}

func monkeyTest(line1 string, line2 string, line3 string) (test, error) {
	divisor, err := strconv.Atoi(line1[len("  Test: divisible by "):])
	mod *= divisor // lcm would be sufficient, but all divisors in input are prime
	if err != nil {
		return nil, err
	}
	caseTrue, err := strconv.Atoi(line2[len("    If true: throw to monkey "):])
	if err != nil {
		return nil, err
	}
	caseFalse, err := strconv.Atoi(line3[len("    If false: throw to monkey "):])
	if err != nil {
		return nil, err
	}
	return func(item int) int {
		if item%divisor == 0 {
			return caseTrue
		}
		return caseFalse
	}, nil
}

func parse(inputFileName string) (monkeys, error) {
	readFile, err := os.Open(inputFileName)
	if err != nil {
		return nil, err
	}
	defer readFile.Close()

	scanner := bufio.NewScanner(readFile)
	scanner.Split(bufio.ScanLines)

	monkeys := monkeys{}
	for scanner.Scan() {
		scanner.Scan()
		monkeyItems, err := monkeyItems(scanner.Text())
		if err != nil {
			return nil, err
		}
		scanner.Scan()
		monkeyOperation, err := monkeyOperation(scanner.Text())
		if err != nil {
			return nil, err
		}
		scanner.Scan()
		testLine1 := scanner.Text()
		scanner.Scan()
		testLine2 := scanner.Text()
		scanner.Scan()
		testLine3 := scanner.Text()
		monkeyTest, err := monkeyTest(testLine1, testLine2, testLine3)
		if err != nil {
			return nil, err
		}
		scanner.Scan()

		monkeys = append(monkeys, &monkey{
			items:     monkeyItems,
			operation: monkeyOperation,
			test:      monkeyTest,
		})
	}

	return monkeys, nil
}
