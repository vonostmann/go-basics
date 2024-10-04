package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	// Создаем хэш-таблицу для хранения чисел и их индексов
	numMap := make(map[int]int)

	// Проходимся по массиву
	for i, num := range nums {
		// Вычисляем комплемент
		complement := target - num

		// Проверяем, есть ли комплемент в хэш-таблице
		if idx, found := numMap[complement]; found {
			// Если найден, возвращаем индексы
			return []int{idx, i}
		}

		// Если комплемента нет, добавляем текущее число и его индекс в хэш-таблицу
		numMap[num] = i
	}

	// В случае, если решения нет (по условию задачи такого быть не должно)
	return []int{}
}

func main() {
	// Примеры использования функции twoSum

	// Пример 1
	nums1 := []int{2, 7, 11, 15}
	target1 := 9
	result1 := twoSum(nums1, target1)
	fmt.Printf("Индексы: %v (Элементы: %d и %d)\n", result1, nums1[result1[0]], nums1[result1[1]])

	// Пример 2
	/*nums2 := []int{3, 2, 4}
	target2 := 5
	result2 := twoSum(nums2, target2)
	fmt.Printf("Индексы: %v (Элементы: %d и %d)\n", result2, nums2[result2[0]], nums2[result2[1]])

	// Пример 3
	nums3 := []int{3, 3}
	target3 := 6
	result3 := twoSum(nums3, target3)
	fmt.Printf("Индексы: %v (Элементы: %d и %d)\n", result3, nums3[result3[0]], nums3[result3[1]])
	*/
}
