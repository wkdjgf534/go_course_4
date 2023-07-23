package main

import (
	"fmt"
)

type customer struct {
	Age int
}

type employee struct {
	Age int
}

// MaxAge - return the eldest person from a collection
func MaxAge(p ...any) any {
	var oldest any
	maxAge := 0
	for _, v := range p {
		switch person := v.(type) {
		case employee:
			if person.Age > maxAge {
				maxAge = person.Age
				oldest = person
			}

		case customer:
			if person.Age > maxAge {
				maxAge = person.Age
				oldest = person
			}
		default:
			return 0
		}
	}
	return oldest
}

func main() {
	c1 := customer{Age: 66}
	c2 := customer{Age: 33}

	e1 := employee{Age: 77}
	e2 := employee{Age: 45}

	fmt.Println(MaxAge())
	fmt.Println("The eldest person among customer", MaxAge(c1, c2))
	fmt.Println("The eldest person among employee", MaxAge(e1, e2))
	fmt.Println("The eldest person", MaxAge(c1, c2, e1, e2))
}
