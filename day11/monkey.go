package day11

import (
	"fmt"
)

var mod int

type operation func(item *int) error

type test func(item int) int

type monkey struct {
	items            queue
	operation        operation
	test             test
	inspectionsCount int
}

type monkeys []*monkey

func min(xs [2]int) (int, int) {
	if xs[0] < xs[1] {
		return 0, xs[0]
	}
	return 1, xs[1]
}

func (monkeys monkeys) monkeyBusiness() int {
	twoHighestInspectionCounts := [2]int{monkeys[0].inspectionsCount, monkeys[1].inspectionsCount}
	for i := 2; i < len(monkeys); i++ {
		minIdx, min := min(twoHighestInspectionCounts)
		if monkeys[i].inspectionsCount > min {
			twoHighestInspectionCounts[minIdx] = monkeys[i].inspectionsCount
		}
	}

	return twoHighestInspectionCounts[0] * twoHighestInspectionCounts[1]
}

func (monkey *monkey) print() {
	if debug {
		fmt.Println("========")
		if len(monkey.items) > 0 {
			fmt.Printf("[%d", *monkey.items[0])
			for i := 1; i < len(monkey.items); i++ {
				fmt.Printf(", %d", *monkey.items[i])
			}
			fmt.Println("]")
		}
		fmt.Printf("inspectionsCount: %d\n", monkey.inspectionsCount)
	}
}

func (monkey *monkey) inspect(item *int) error {
	monkey.inspectionsCount++
	err := monkey.operation(item)
	if err != nil {
		return err
	}
	return nil
}

func (monkey *monkey) getBored(item *int) {
	if monkeysGetBored {
		*item /= 3
	}
}

func (monkey *monkey) processItem(monkeys monkeys) error {
	item := monkey.items.pop()
	err := monkey.inspect(item)
	if err != nil {
		return err
	}
	monkey.getBored(item)
	targetMonkey := monkeys[monkey.test(*item)]
	targetMonkey.items.push(item)
	return nil
}

func (monkey *monkey) processAllItems(monkeys monkeys) error {
	for len(monkey.items) > 0 {
		err := monkey.processItem(monkeys)
		if err != nil {
			return err
		}
	}
	return nil
}
