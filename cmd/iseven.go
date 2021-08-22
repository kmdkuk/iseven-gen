package cmd

import (
	"context"
	"fmt"
	"io"
	"strings"
)

func WriteHeader(w io.Writer) {
	fmt.Fprintf(w, `package iseven

type Number interface{}

func is(lhs, rhs interface{}) bool {
	return fmt.Sprint(lhs) == fmt.Sprint(rhs)
}

func IsEven(number interface{}) bool {
    if (is(number, "even") || is(number, "Even") || is(number, "eVen") || is(number, "evEn") || is(number, "eveN") || is(number, "EVen") || is(number, "EvEn") || is(number, "EveN") || is(number, "eVEn") || is(number, "eVeN") || number == "evEN" || is(number, "eVEN") || is(number, "EvEN") || is(number, "EVeN") || is(number, "EVEn") || is(number, "EVEN")) {
		return true
	} 
	`)
}

func WriteContent(ctx context.Context, w io.Writer) {
	num := 1
	for {
		select {
		case <-ctx.Done():
			return
		default:
			WriteNumber(w, num)
			num++
		}
	}
}

func WriteNumber(w io.Writer, num int) {
	str := numToStr(num)
	fmt.Fprintf(w, `else if(is(number, %d) || is(number, "%d") || is(number, "%s") || is(number, "%s") || is(number, "%s")) {
		return %v
	}
	`, num, num, strings.ToLower(str), str, strings.ToUpper(str), isEven(num))
}

func numToStr(num int) string {
	upTo19 := []string{
		"",
		"One",
		"Two",
		"Three",
		"Four",
		"Five",
		"Six",
		"Seven",
		"Eight",
		"Nine",
		"Ten",
		"Eleven",
		"Twelve",
		"Thirteen",
		"Fourteen",
		"Fifteen",
		"Sixteen",
		"Seventeen",
		"Eighteen",
		"Nineteen",
	}
	tensPlace := []string{
		"",
		"",
		"Twenty",
		"Thirty",
		"Fourty",
		"Fifty",
		"Sixty",
		"Seventy",
		"Eighty",
		"Ninety",
	}
	hundred := "Hundred"
	unit := []string{
		"",
		"Thousand",
		"Million",
		"Billion",
		"Trillion",
	}
	str := []string{}
	d := digit(num)
	for i, n := range d {
		u := i % 3
		uu := i / 3
		s := ""
		switch u {
		case 0:
			if i+1 < len(d) && u == 0 && d[i+1] == 1 {
				continue
			}
			s = s + upTo19[n]
			if uu > 0 {
				s = s + " " + unit[uu] + " "
			}
		case 1:
			switch n {
			case 0:
				continue
			case 1:
				s = upTo19[n+10]
			default:
				s = tensPlace[n]
				if d[i-1] != 0 {
					s = s + "-"
				}
			}
		case 2:
			if n == 0 {
				continue
			}
			s = upTo19[n] + " " + hundred + " "
		default:
			continue

		}
		str = append(str, s)
	}
	r := reverseStr(str)
	return strings.Join(r, "")
}

func digit(num int) []int {
	var list []int
	i := num
	for i > 0 {
		list = append(list, i%10)
		i = i / 10
	}
	return list
}

func reverseStr(list []string) []string {
	for i := 0; i < len(list)/2; i++ {
		j := len(list) - i - 1
		list[i], list[j] = list[j], list[i]
	}
	return list
}

func isEven(num int) bool {
	return num%2 == 0
}

func WriteFooter(w io.Writer) {
	fmt.Fprintf(w, `return false
}
	`)
}
