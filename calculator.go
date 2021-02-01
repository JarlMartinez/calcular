package calcular

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Calculator struct{}

// 1st part beteween () makes this func a 'Reciever function'
// which means that the 'struct' now has this function inside.
func (calculator) operate(input string, operator string) int {
	cleanInput := strings.Split(input, operator)
	value1 := parsear(cleanInput[0])
	value2 := parsear(cleanInput[1])

	switch operator {
	case "+":
		return value1 + value2
	case "-":
		return value1 - value2
	case "/":
		return value1 / value2
	case "*":
		return value1 * value2
	default:
		fmt.Println(operator, "Operator not supported D:")
		return 0
	}
}

func parsear(input string) int {
	valor, _ := strconv.Atoi(input)
	return valor
}

func ReadInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}
