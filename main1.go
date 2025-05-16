package main1

import (
	"fmt"
)

func main() {
	// Ввод размерности матрицы
	var n int
	fmt.Print("Введите размерность матрицы n: ")
	fmt.Scan(&n)

	// Создание и заполнение матрицы
	matrix := make([][]float64, n)
	for i := range matrix {
		matrix[i] = make([]float64, n)
	}

	fmt.Println("Введите элементы матрицы построчно:")
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Scan(&matrix[i][j])
		}
	}

	// Вычисление сумм диагоналей
	mainSum, secondarySum := CalculateDiagonals(matrix)

	// Сравнение и вывод результатов
	result := CompareDiagonals(mainSum, secondarySum)

	fmt.Printf("\nСумма главной диагонали: %.2f\n", mainSum)
	fmt.Printf("Сумма побочной диагонали: %.2f\n", secondarySum)
	fmt.Println(result)
}
