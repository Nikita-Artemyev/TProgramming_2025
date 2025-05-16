package main

import (
	"fmt"
	"math"
)

// Функция для вычисления y=f(x)
func f(b, x float64) float64 {
	return (1 + math.Pow(math.Sin(math.Pow(b, 3)+math.Pow(x, 3)), 2)) / math.Sqrt(math.Pow(b, 3)+math.Pow(x, 3))
}

// Функция для решения задачи A
func taskA(b, x_start, x_end, deltaX float64) {
	fmt.Println("Результаты для Задачи A (x, y):")
	for x := x_start; x <= x_end; x += deltaX {
		y := f(b, x)
		fmt.Printf("x: %.2f, y: %.4f\n", x, y)
	}
}

// Функция для решения задачи B
func taskB(b float64, xValues []float64) {
	fmt.Println("Результаты для Задачи B (x, y):")
	for _, x := range xValues {
		y := f(b, x)
		fmt.Printf("x: %.2f, y: %.4f\n", x, y)
	}
}

func main() {
	// Задача A
	b := 2.5
	x_n := 1.28
	x_k := 3.28
	delta_x := 0.4

	taskA(b, x_n, x_k, delta_x)

	// Задача B
	xValues := []float64{1.1, 2.4, 3.6, 1.7, 3.9}
	taskB(b, xValues)
}
