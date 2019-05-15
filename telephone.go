package main

import (
	"fmt"
	"regexp"
)

func main() {
	phone1 := "+1 (613)-995-0253"
	phone2 := "613-995-0253"
	phone3 := "1 613 995 0253"
	phone4 := "613.995.0253"

	matched := regexp.MustCompile(`^\+?([0-9]{1})?[\( ]?[\(]?([0-9]{3})\)?[-. ]?([0-9]{3})[-. ]?([0-9]{4})$`)

	fmt.Printf("\nPhone: %v\t:%v\n", phone1, matched.MatchString(phone1))
	fmt.Printf("Phone: %v\t:%v\n", phone2, matched.MatchString(phone2))
	fmt.Printf("Phone: %v\t\t:%v\n", phone3, matched.MatchString(phone3))
	fmt.Printf("Phone: %v\t\t\t:%v\n", phone4, matched.MatchString(phone4))

	replace_string := regexp.MustCompile(`[\+\(\)\-\. ]`)
	fmt.Printf("%q\n", replace_string.ReplaceAllString(phone1, ""))
	fmt.Printf("%q\n", replace_string.ReplaceAllString(phone2, ""))
	fmt.Printf("%q\n", replace_string.ReplaceAllString(phone3, ""))
	fmt.Printf("%q\n", replace_string.ReplaceAllString(phone4, ""))
	}