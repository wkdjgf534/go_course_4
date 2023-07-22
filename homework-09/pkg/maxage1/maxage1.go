package maxage1

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

func MaxAge(people) int {
	max = products[0]
	for _, product := range products {
		if product.Price > max.Price {
			max = product
		}
	}
	return max
}
