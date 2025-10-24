package main

import (
	"fmt"
	"strings"
)

func numberToText(num int) string {
	units := map[int]string{
		1: "один", 2: "два", 3: "три", 4: "четыре",
		5: "пять", 6: "шесть", 7: "семь", 8: "восемь", 9: "девять",
	}
	teens := map[int]string{
		10: "десять", 11: "одиннадцать", 12: "двенадцать",
		13: "тринадцать", 14: "четырнадцать", 15: "пятнадцать",
		16: "шестнадцать", 17: "семнадцать", 18: "восемнадцать", 19: "девятнадцать",
	}
	tens := map[int]string{
		2: "двадцать", 3: "тридцать", 4: "сорок",
		5: "пятьдесят", 6: "шестьдесят", 7: "семьдесят",
		8: "восемьдесят", 9: "девяносто",
	}
	hundreds := map[int]string{
		1: "сто", 2: "двести", 3: "триста",
		4: "четыреста", 5: "пятьсот", 6: "шестьсот",
		7: "семьсот", 8: "восемьсот", 9: "девятьсот",
	}
	var thousandsWord string

	wrap := func(base int) {
		hundred := base / 100
		if word, exists := hundreds[hundred]; exists {
			thousandsWord += word + " "
		}
		ten := base % 100
		if word, exists := teens[ten]; exists {
			thousandsWord += word + " "
		} else {
			ten = ten / 10
			thousandsWord += tens[ten] + " "
			unit := base % 10
			if unit > 0 {
				thousandsWord += units[unit] + " "
			}
		}
	}

	thousand := num / 1000
	wrap(thousand)

	switch {
	case thousand%10 == 1:
		thousandsWord += "тысяча" + " "
	case thousand%100 >= 2 && thousand%100 <= 4:
		thousandsWord += "тысячи" + " "
	default:
		thousandsWord += "тысяч" + " "
	}
	hundred := num % 1000
	wrap(hundred)

	return strings.TrimSpace(thousandsWord)
}

func main() {
	var num int
	fmt.Scanln(&num)
	if num >= 12307 {
		fmt.Printf("Число %d слишком большое. Введите число меньше 12307.\n", num)
		return
	}
	serverError := false
	for num < 12307 {
		switch {
		case num < 0:
			num *= -1
		case num%7 == 0:
			num *= 39
		case num%9 == 0:
			num = num*13 + 1
			continue
		default:
			num = (num + 2) * 3
		}

		if num%9 == 0 && num%13 == 0 {
			serverError = true
			fmt.Println("service error")
			break
		} else {
			num += 1
		}

	}
	if !serverError {
		textNum := numberToText(num)
		fmt.Printf("Результат: %d(%s)\n", num, textNum)
	}

}
