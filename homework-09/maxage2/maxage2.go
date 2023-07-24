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
func MaxAge(users ...any) any {
	var oldest any
	maxAge := 0

	for _, p := range users {
		switch v := p.(type) {
		case employee:
			if v.Age > maxAge {
				maxAge, oldest = v.Age, v
			}
		case customer:
			if v.Age > maxAge {
				maxAge, oldest = v.Age, v
			}
		default:
			continue
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
	fmt.Println("The eldest person among customers", MaxAge(c1, c2))
	fmt.Println("The eldest person among employees", MaxAge(e1, e2))
	fmt.Println("The eldest person", MaxAge(c1, c2, e1, e2))
}
