package mymath

import "math"

func Ceil(x float64) float64 {
	return math.Ceil(x)
}

func Floor(x float64) float64 {
	return math.Floor(x)
}
func Pow(x, y float64) float64 {
	return math.Pow(x, y)
}
func Min(x, y float64) float64 {
	return math.Min(x, y)
}

func Sqrt(x float64) float64 {
	return math.Sqrt(x)
}

func Sin(x float64) float64 {
	return math.Sin(x)
}

func Cos(x float64) float64 {
	return math.Cos(x)
}

func Add(x, y float64) float64 {
	return x + y
}

func Sub(x, y float64) float64 {
	return x - y
}

func Mul(x, y float64) float64 {
	return x * y
}

func Div(x, y float64) float64 {
	return x / y
}

func Abs(x float64) float64 {
	return math.Abs(x)
}

func Max(x, y float64) float64 {
	return math.Max(x, y)
}

func Yn(n int, x float64) float64 {
	return math.Yn(n, x)
}
