package main1

// CalculateDiagonals вычисляет суммы главной и побочной диагоналей
func CalculateDiagonals(matrix [][]float64) (float64, float64) {
	size := len(matrix)
	mainSum := 0.0      // Сумма главной диагонали
	secondarySum := 0.0 // Сумма побочной диагонали

	for i := 0; i < size; i++ {
		mainSum += matrix[i][i]             // Элементы главной диагонали
		secondarySum += matrix[i][size-1-i] // Элементы побочной диагонали
	}

	return mainSum, secondarySum
}

// CompareDiagonals сравнивает суммы диагоналей и возвращает результат
func CompareDiagonals(mainSum, secondarySum float64) string {
	switch {
	case mainSum > secondarySum:
		return "Сумма главной диагонали больше"
	case mainSum < secondarySum:
		return "Сумма побочной диагонали больше"
	default:
		return "Суммы диагоналей равны"
	}
}
