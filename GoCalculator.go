package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func checkRoman(input string) bool {
	romanSymbols := "IVXLCDM"
	for _, c := range input {
		if !contains(romanSymbols, string(c)) {
			return false
		}
	}
	return true
}

func checkArabic(input string) bool {
	_, err := strconv.Atoi(input)
	return err == nil
}

func contains(s string, char string) bool {
	for _, c := range s {
		if string(c) == char {
			return true
		}
	}
	return false
}

func romanToInt(roman string) int {
	romanNumerals := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	result := 0
	prev := 0
	for i := len(roman) - 1; i >= 0; i-- {
		current := romanNumerals[rune(roman[i])]
		if prev > current {
			result -= current
		} else {
			result += current
		}
		prev = current
	}

	return result
}


func intToRoman(num int) string {
	arabicNumbers := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	romanNumerals := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	roman := ""

	for i := 0; i < len(arabicNumbers); i++ {
		for num >= arabicNumbers[i] {
			roman += romanNumerals[i]
			num -= arabicNumbers[i]
		}
	}

	return roman
}

func main() {

	var row, numbOne, operator, numbSec string

	fmt.Println("Ввод операции:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	row = scanner.Text()
	fmt.Println()
	fmt.Println(row)

	number := strings.Split(row, " ")

	if len(number) != 3 {
		panic("Формат данных не соответствует")
	}

	numbOne = number[0]
	operator = number[1]
	numbSec = number[2]

	isRoman := checkRoman(numbSec)
	isArabic := checkArabic(numbOne)

	if isRoman && !isArabic {
		fmt.Println("Введено римское число")
	} else if !isRoman && isArabic {
		fmt.Println("Введено арабское число")
	} else {
		panic("Некорректное число")
	}

	if isArabic == true {

		a1, _ := strconv.Atoi(numbOne)
		b1, _ := strconv.Atoi(numbSec)

		if !(1 <= a1 && a1 <= 10) || !(1 <= b1 && b1 <= 10) {
			panic("Значения чисел не соответствуют")
		}

		switch operator {
		case "+":
			fmt.Println(a1 + b1)
		case "-":
			fmt.Println(a1 - b1)
		case "*":
			fmt.Println(a1 * b1)
		case "/":
			fmt.Println(a1 / b1)

		}
	} else if isRoman == true {

		a1 := romanToInt(numbOne)
		b1 := romanToInt(numbSec)

		if !(1 <= a1 && a1 <= 10) || !(1 <= b1 && b1 <= 10) {
			panic("Значения чисел не соответствуют")
		}

		switch operator {
		case "+":
			fmt.Println(intToRoman(a1 + b1))
			//fmt.Println(a1 + b1)
		case "-":
			if a1 < b1 {
				panic("Первое число меньше второго")
			} else {
				fmt.Println(intToRoman(a1 - b1))
				//fmt.Println(a1 - b1)
			}
		case "*":
			fmt.Println(intToRoman(a1 * b1))
			//fmt.Println(a1 * b1)
		case "/":
			if a1 < b1 || a1%b1 != 0 {
				panic("Первое число меньше второго или деление с остатком")
			} else {
				fmt.Println(intToRoman(a1 / b1))
				//fmt.Println(a1 - b1)
			}

		}
	}

}
