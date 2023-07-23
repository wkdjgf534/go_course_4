package main

import (
	"fmt"
	"sort"
)

type Ager interface {
	age() int
}

// Customer -
type Customer struct {
	Age int
}

// Employee -
type Employee struct {
	Age int
}

func (e *Employee) age() int { return e.Age }

func (c *Customer) age() int { return c.Age }

//type ager interface{ age() int }

// MaxAge - return the eldest person from a collection
func MaxAge(p ...Ager) int {
	sort.Slice(p, func(i, j int) bool {
		return p[i].age() > p[j].age()
	})
	return p[0].age()
}

func main() {
	c1 := Customer{Age: 55}
	c2 := Customer{Age: 33}

	e1 := Employee{Age: 25}
	e2 := Employee{Age: 45}

	fmt.Println("The eldest person between customers", MaxAge(&c1, &c2))
	fmt.Println("The eldest person between employers", MaxAge(&e1, &e2))
	fmt.Println("The eldest person", MaxAge(&c1, &c2, &e1, &e2))
}
