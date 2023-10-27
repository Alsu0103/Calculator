package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

func main() {

	fmt.Println("Введите выражение:")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	expression := scanner.Text()

	expression_split := strings.Fields(expression)
	len_expression := len(expression_split)
	operations := []string{"+", "-", "/", "*"}
	rim := []string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X"}
	arab := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}
	var n1, n2 int
	rim_full := []string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X", "XI", "XII", "XIII", "XIV", "XV", "XVI", "XVII", "XVIII", "XIX", "XX", "XXI", "XXII", "XXIII", "XXIV", "XXV", "XXVI", "XXVII", "XXVIII", "XXIX", "XXX", "XXXI", "XXXII", "XXXIII", "XXXIV", "XXXV", "XXXVI", "XXXVII", "XXXVIII", "XXXIX", "XXXX", "XXXXI", "XXXXII", "XXXXIII", "XXXXIV", "XXXXV", "XXXXVI", "XXXXVII", "XXXXVIII", "XXXXIX", "L", "LI", "LII", "LIII", "LIV", "LV", "LVI", "LVII", "LVIII", "LIX", "LX", "LXI", "LXII", "LXIII", "LXIV", "LXV", "LXVI", "LXVII", "LXVIII", "LXIX", "LXX", "LXXI", "LXXII", "LXXIII", "LXXIV", "LXXV", "LXXVI", "LXXVII", "LXXVIII", "LXXIX", "LXXX", "LXXXI", "LXXXII", "LXXXIII", "LXXXIV", "LXXXV", "LXXXVI", "LXXXVII", "LXXXVIII", "LXXXIX", "LXXXX", "LXXXXI", "LXXXXII", "LXXXXIII", "LXXXXIV", "LXXXXV", "LXXXXVI", "LXXXXVII", "LXXXXVIII", "LXXXXIX", "C"}

	found := contains(operations, expression_split[1])
	found1 := contains(rim, expression_split[0]) && contains(rim, expression_split[2])
	found2 := contains(arab, expression_split[0]) && contains(arab, expression_split[2])

	if len_expression == 3 && found && (found1 || found2) {
		if found1 {
			for i := 0; i < 10; i++ {
				if rim[i] == expression_split[0] {
					n1 = i + 1
					break
				}
			}
			for i := 0; i < 10; i++ {
				if rim[i] == expression_split[2] {
					n2 = i + 1
					break
				}
			}
			var ans1 int
			if expression_split[1] == "+" {
				ans1 = n1 + n2 - 1
				if ans1 > 0 {
					fmt.Println(rim_full[ans1])
				} else {
					fmt.Println("Ошибка")
				}
			} else if expression_split[1] == "-" {
				ans1 = n1 - n2 - 1
				if ans1 > 0 {
					fmt.Println(rim_full[ans1])
				} else {
					fmt.Println("Ошибка")
				}
			} else if expression_split[1] == "*" {
				ans1 = n1*n2 - 1
				if ans1 > 0 {
					fmt.Println(rim_full[ans1])
				} else {
					fmt.Println("Ошибка")
				}
			} else {
				ans1 = n1/n2 - 1
				if ans1 > 0 {
					fmt.Println(rim_full[ans1])
				} else {
					fmt.Println("Ошибка")
				}
			}
		} else {
			for i := 0; i < 10; i++ {
				if arab[i] == expression_split[0] {
					n1 = i + 1
					break
				}
			}
			for i := 0; i < 10; i++ {
				if arab[i] == expression_split[2] {
					n2 = i + 1
					break
				}
			}
			if expression_split[1] == "+" {
				fmt.Println(n1 + n2)
			} else if expression_split[1] == "-" {
				fmt.Println(n1 - n2)
			} else if expression_split[1] == "*" {
				fmt.Println(n1 * n2)
			} else {
				fmt.Println(n1 / n2)
			}
		}
	} else {
		fmt.Println("Ошибка")
	}
}
