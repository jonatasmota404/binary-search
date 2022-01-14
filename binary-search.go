package main

import (
	"errors"
	"fmt"
)

func main() {
	list := generateList(10)
	pos, err := binarySearch(list, 6)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("valor encontrado na posição:", pos)

}

func binarySearch(list []int, value int) (int, error) {
	inicio := 0
	fim := len(list) - 1

	for inicio <= fim {
		meio := (inicio + fim) / 2
		chute := list[meio]
		switch {
		case chute == value:
			return meio, nil
		case chute < value:
			inicio = meio + 1
		case chute > value:
			fim = meio - 1
		}
	}
	return -1, errors.New("valor não encontrado")
}

func generateList(x int) []int {
	nums := []int{}
	for i := 0; i < x; i++ {
		nums = append(nums, i)
	}

	return nums
}
