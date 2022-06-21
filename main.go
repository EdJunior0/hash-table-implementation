package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type node struct {
	key  int
	prox *node
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func hash(key int, tableSize int) int {
	constant := 0.618
	value := float64(key) * constant
	_, fractionalFloat := math.Modf(value)
	return int(fractionalFloat * float64(tableSize))
}

func searchEmptyNode(node *node, colisions *int) *node {
	*colisions++
	if node == nil {
		return node
	}
	return searchEmptyNode(node.prox, colisions)
}

func hash2(key int) int64 {
	multipliedKey := key * key
	keyInString := strconv.FormatInt(int64(multipliedKey), 2)
	// fmt.Println("key in binary = ", keyInString)
	// fmt.Println("size of string = ", len(keyInString))
	halfOfString := int(len(keyInString) / 2)
	// fmt.Println("half of string = ", halfOfString)
	differenceOfHalf := int((len(keyInString) - halfOfString) / 2)
	// fmt.Println("difference of half = ", differenceOfHalf)

	mask := ""

	for i := 0; i < len(keyInString); i++ {
		if i < differenceOfHalf || i > halfOfString+differenceOfHalf {
			mask += "0"
		}
		if i >= differenceOfHalf && i < halfOfString+differenceOfHalf {
			mask += "1"
		}
	}

	// fmt.Println("mask = ", mask)
	firstNumber, _ := strconv.ParseInt(keyInString, 2, 64)
	secondNumber, _ := strconv.ParseInt(mask, 2, 64)
	composed := firstNumber & secondNumber
	myKey := composed >> int64(differenceOfHalf)

	// fmt.Println("Key =", myKey)

	return myKey
}

func main() {

	input, err := os.ReadFile("input.txt")
	check(err)
	inputSplited := strings.Split(string(input), " ")
	colisions := 0

	hashTableSize, _ := strconv.Atoi(inputSplited[0])
	slice := make([]node, hashTableSize)

	for index, value := range inputSplited {
		valueInt, _ := strconv.Atoi(value)
		hash := hash2(valueInt)
		nodeWithKey := node{key: valueInt, prox: nil}
		if index != 0 {
			if slice[hash].key == 0 {
				slice[hash] = nodeWithKey
			} else {
				emptyNode := searchEmptyNode(slice[hash].prox, &colisions)
				emptyNode = &nodeWithKey
				_ = emptyNode
			}
		}
	}

	fmt.Println("Função hash utilizada: Método da multiplicação (em código, a função de nome hash)")
	fmt.Println("Método de tratamento de colisões utilizado: Encadeamento externo")
	fmt.Println("Total de colisões: ", colisions)

}
