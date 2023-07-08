package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var toMuch = "Вывод ошибки, калькулятор должен принимать на вход числа от 1 до 10 включительно, не более."
var zeroOper = "Вывод ошибки, так как строка " +
	"не является математической операцией."

var moreOper = "Вывод ошибки, так как формат математической операции " +
	"не удовлетворяет заданию — два операнда и один оператор (+, -, /, *)."

var numberSys = "Вывод ошибки, так как используются " +
	"одновременно разные системы счисления."

var negativeRome = "Вывод ошибки, так как в римской системе \n " +
	"нет отрицательных чисел."

var nilROM = "Вывод ошибки, так как в римской системе нет числа 0."
var nilFloat = "Калькулятор умеет работать только с арабскими целыми " +
	"числами или римскими цифрами от 1 до 10 включительно" //

var romanians = map[string]int{
	"C":    100,
	"XC":   90,
	"L":    50,
	"XL":   40,
	"X":    10,
	"IX":   9,
	"VIII": 8,
	"VII":  7,
	"VI":   6,
	"V":    5,
	"IV":   4,
	"III":  3,
	"II":   2,
	"I":    1,
}
var ints = [14]int{
	100, 90, 50, 40, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1,
}
var opers = [4]string{"+", "-", "/", "*"}

func mathOp(a string) []string {
	floatIs := strings.Split(a, "")
	for _, el := range floatIs {
		if el == "." || el == "," {
			panic(nilFloat)
		}
	}

	workStr := strings.Split(a, " ")
	return workStr

}
func sum(a, b *int) int {
	res := *a + *b
	return int(res)
}
func sub(a, b *int) int {
	res := *a - *b
	return int(res)
}
func mult(a, b *int) int {
	res := *a * *b
	return int(res)
}
func div(a, b *int) int {
	res := *a / *b
	return int(res)
}
func intToR(a int) {
	var romanNumFromInt string
	for _, el := range ints {
		for i := el; i <= a; {
			for val, elem := range romanians {
				if elem == i {
					romanNumFromInt += val
					a -= elem
				}
			}
		}

	}
	fmt.Println(romanNumFromInt)
}
func calculator(input string) {
	var s string = strings.TrimSpace(input)
	workS := mathOp(s)

	if len(workS) < 3 {
		panic(zeroOper)
	}
	var operator string
	var sToIntR []string
	var sToIntA []int
	var sRToInt []int
	var out string
	var iRr int

	if strings.ContainsAny(workS[1], "+") || strings.ContainsAny(workS[1], "-") || strings.ContainsAny(workS[1], "*") || strings.ContainsAny(workS[1], "/") {

		for _, el := range workS {
			for _, vl := range opers {
				if el == vl {
					operator += el
				}
			}
		}
	}
	switch {
	case len(operator) < 1:
		{
			panic(zeroOper)
		}
	case (len(operator) > 1):
		panic(moreOper)
	}

	for _, el := range workS {
		for vl := range romanians {
			if vl == el {
				sToIntR = append(sToIntR, vl)
			}
		}

	}

	for _, el := range workS {
		i, _ := strconv.Atoi(el)
		if i > 0 && i <= 10 {
			sToIntA = append(sToIntA, i)

		} else if i > 10 {
			panic(toMuch)
		}

	}

	if len(sToIntR) < 2 && len(sToIntA) < 2 {
		panic(numberSys)

	} else if len(sToIntA) == 2 {

		switch operator {
		case "-":

			a, b := &sToIntA[0], &sToIntA[1]

			out = fmt.Sprintf("%d", sub(a, b))
		case "+":

			a, b := &sToIntA[0], &sToIntA[1]

			out = fmt.Sprintf("%d", sum(a, b))
		case "*":

			a, b := &sToIntA[0], &sToIntA[1]

			out = fmt.Sprintf("%d", mult(a, b))

		case "/":

			a, b := &sToIntA[0], &sToIntA[1]

			out = fmt.Sprintf("%d", div(a, b))

		}

	} else if len(sToIntR) == 2 {

		switch operator {
		case "-":

			for _, vl := range sToIntR {
				for i, _ := range romanians {
					if vl == i {
						sRToInt = append(sRToInt, romanians[i])
					}
				}
			}
			fR, sR := sRToInt[0], sRToInt[1]
			fRi, sRi := &fR, &sR
			iRr = sub(fRi, sRi)

		case "+":
			for _, vl := range sToIntR {
				for i, _ := range romanians {
					if vl == i {
						sRToInt = append(sRToInt, romanians[i])
					}
				}
			}
			fR, sR := sRToInt[0], sRToInt[1]
			fRi, sRi := &fR, &sR
			iRr = sum(fRi, sRi)

		case "*":
			for _, vl := range sToIntR {
				for i, _ := range romanians {
					if vl == i {
						sRToInt = append(sRToInt, romanians[i])
					}
				}
			}
			fR, sR := sRToInt[0], sRToInt[1]
			fRi, sRi := &fR, &sR
			iRr = mult(fRi, sRi)

		case "/":
			for _, vl := range sToIntR {
				for i, _ := range romanians {
					if vl == i {

						sRToInt = append(sRToInt, romanians[i])
					}
				}
			}

			fR, sR := sRToInt[0], sRToInt[1]
			fRi, sRi := &fR, &sR
			iRr = div(fRi, sRi)

		}

	}

	if iRr < 0 && len(sToIntR) == 2 {
		panic(negativeRome)
	} else if iRr == 0 && len(sToIntR) == 2 {
		panic(nilROM)
	} else if iRr > 0 {
		intToR(iRr)

	}

	if len(out) != 0 {
		fmt.Println(out)
	}
}

func main() {
	fmt.Println("Добро пожаловать в ката калькулятор")
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Введите занчение")
		text, _ := reader.ReadString('\n')
		calculator(strings.ToUpper(text))
	}

}

//test
