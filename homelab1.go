package main

import (
	"errors"
	"fmt"
	"math"
)

const pi = 3.1415

func main() {
	printCircleArea(2)
	printTriangleArea(2, 3)
	printTrianglePerimeter(2, 3)
	printRectangleArea(5, 6)
	printRectanglePerimeter(5, 6)

}

func printCircleArea(radius int) {
	circleArea, err := calculateCircleArea(radius)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("Расчет радиуса круга:\n")
	fmt.Printf("_____________________\n")
	fmt.Printf("Радиус круга: %d см.\n", radius)
	fmt.Println("Формула для расчета площади круга: A=πr2\n")

	fmt.Printf("Площадь круга: %f32 см. кв.\n\n", circleArea)

}

func printTriangleArea(side1 int, side2 int) {
	triangleArea, err := calculateTriangleArea(side1, side2)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("Расчет площади треугольника:\n")
	fmt.Printf("____________________________\n")
	fmt.Printf("Сторона 1: %d см.\n", side1)
	fmt.Printf("Сторона 2: %d см.\n", side2)
	fmt.Println("Формула для расчета площади треугольника: S=(A*B)/2\n")
	fmt.Printf("Площадь треугольника: %f32 см. кв.\n\n", triangleArea)
}

func printTrianglePerimeter(side1 int, side2 int) {
	trianglePerimeter, err := calculateTrianglePerimeter(side1, side2)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("Расчет периметра треугольника:\n")
	fmt.Printf("______________________________\n")
	fmt.Printf("Сторона 1: %d см.\n", side1)
	fmt.Printf("Сторона 2: %d см.\n", side2)
	fmt.Println("Формула для расчета периметра треугольника: S=(A+B)*2\n")
	fmt.Printf("Периметр треугольника: %f32 см. кв.\n\n", trianglePerimeter)
}

func printRectanglePerimeter(side3 int, side4 int) {
	rectanglePerimeter, err := calculateRectanglePerimeter(side3, side4)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("Расчет периметра прямоугольника:\n")
	fmt.Printf("________________________________\n")
	fmt.Printf("Сторона 1: %d см.\n", side3)
	fmt.Printf("Сторона 2: %d см.\n", side4)
	fmt.Println("Формула для расчета периметра прямоугольника: S=(A+B)*2\n")
	fmt.Printf("Периметр прямоугольника: %f32 см. кв.\n\n", rectanglePerimeter)
}

func printRectangleArea(side3 int, side4 int) {
	rectangleArea, err := calculateRectangleArea(side3, side4)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("Расчет площади прямоугольника:\n")
	fmt.Printf("______________________________\n")
	fmt.Printf("Сторона 1: %d см.\n", side3)
	fmt.Printf("Сторона 2: %d см.\n", side4)
	fmt.Println("Формула для расчета площади прямоугольника: S=A*B\n")
	fmt.Printf("Площадь прямоугольника: %f32 см. кв.\n\n", rectangleArea)
}

func calculateCircleArea(radius int) (float32, error) {
	if radius <= 0 {
		return float32(0), errors.New("Радиус круга не может быть отрицательным!")
	}

	return float32(radius) * float32(radius) * math.Pi, nil
}

func calculateTriangleArea(side1 int, side2 int) (float32, error) {
	if side1 <= 0 {
		return float32(0), errors.New("Сторона треугольника не может быть отрицательным числом!")
	}
	if side2 <= 0 {
		return float32(0), errors.New("Сторона треугольника не может быть отрицательным числом!")
	}

	return (float32(side1) * float32(side2)) / 2, nil
}

func calculateTrianglePerimeter(side1 int, side2 int) (float32, error) {
	if side1 <= 0 {
		return float32(0), errors.New("Сторона треугольника не может быть отрицательным числом!")
	}
	if side2 <= 0 {
		return float32(0), errors.New("Сторона треугольника не может быть отрицательным числом!")
	}

	return (float32(side1) + float32(side2)) * 2, nil
}

func calculateRectangleArea(side3 int, side4 int) (float32, error) {
	if side3 <= 0 {
		return float32(0), errors.New("Сторона прямоугольника не может быть отрицательным числом!")
	}
	if side4 <= 0 {
		return float32(0), errors.New("Сторона прямоугольника не может быть отрицательным числом!")
	}

	return float32(side3) * float32(side4), nil
}

func calculateRectanglePerimeter(side3 int, side4 int) (float32, error) {
	if side3 <= 0 {
		return float32(0), errors.New("Сторона прямоугольника не может быть отрицательным числом!")
	}
	if side4 <= 0 {
		return float32(0), errors.New("Сторона прямоугольника не может быть отрицательным числом!")
	}
	return (float32(side3) + float32(side4)) * 2, nil
}
