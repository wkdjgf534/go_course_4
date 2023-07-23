package main

import (
	"fmt"
	"sort"
)

type ager interface {
	age() int
}

type customer struct {
	Age int
}

type employee struct {
	Age int
}

func (e *employee) age() int { return e.Age }

func (c *customer) age() int { return c.Age }

//type ager interface{ age() int }

// MaxAge - return the eldest person from a collection
func MaxAge(p ...ager) int {
	if len(p) == 0 {
		return 0
	}

	sort.Slice(p, func(i, j int) bool {
		return p[i].age() > p[j].age()
	})
	return p[0].age()
}

func main() {
	c1 := customer{Age: 55}
	c2 := customer{Age: 33}

	e1 := employee{Age: 25}
	e2 := employee{Age: 45}

	fmt.Println(MaxAge())
	fmt.Println("The eldest person among customers", MaxAge(&c1, &c2))
	fmt.Println("The eldest person among employees", MaxAge(&e1, &e2))
	fmt.Println("The eldest person", MaxAge(&c1, &c2, &e1, &e2))
}
