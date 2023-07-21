package maxage

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
