package distance

import "math"

// CalculateDistance - расчёт дистанции между 2 точками
func CalculateDistance(x1, x2, y1, y2 float64) float64 {
	return math.Sqrt(math.Pow(math.Abs(x2)-math.Abs(x1), 2) + math.Pow(math.Abs(y2)-math.Abs(y1), 2))
}
