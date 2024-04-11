package algo

import (
	"strconv"
	"strings"
)

func IsNumberPresent(list []int, numToSearch int) bool {
	for _, val := range list {
		if val == numToSearch {
			return true
		}
	}
	return false
}

func IsNumberOdd(num int) bool {
	return num % 2 != 0
}

func IsStringPresentIgnoreCase(list []string, stringToSearch string) bool {
	for _, val := range list{
		if EqualsIgnoreCase(val, stringToSearch){
			return true
		}
	}
	return false
}

func EqualsIgnoreCase(string1 string, string2 string) bool {
	return strings.EqualFold(string1, string2)
}

func SumOfIntArray(list []int) (sum int) {
	for _, val := range list{
		sum = sum + val
	}
	return
}

func ReverseString(s string) (reversedString string){
	for _, v := range s {
		reversedString = string(v) + reversedString
	}
	return		
}

func FizzBuzzTillN(n int) []string{
	var array []string
	for i := 1; i<=n; i++{
		var s string = ""
		if(i % 3 == 0){
			s = "Fizz"
		}
		if(i % 5 == 0){
			s += "Buzz"
		}
		if(s == ""){
			s = strconv.Itoa(i)
		}
		array = append(array, s)
	}
	return array
}

