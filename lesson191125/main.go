package main

import "fmt"

func SumAll(nums []int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func filterEvenNumbers(nums []int) []int {
	var evens []int
	for _, num := range nums {
		if num%2 == 0 {
			evens = append(evens, num)
		}
	}
	return evens
}

func ReverseSlice(nums []int) {
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}

func DeleteElementAtIndex(nums []int, index int) []int {
	if index < 0 || index >= len(nums) {
		return nums
	}
	return append(nums[:index], nums[index+1:]...)
}

func Insert(nums []int, index int, value int) []int {
	if index < 0 || index > len(nums) {
		return nums
	}
	nums = append(nums, 0)             // Увеличиваем срез на 1 элемент
	copy(nums[index+1:], nums[index:]) // Сдвигаем элементы вправо
	nums[index] = value                // Вставляем новое значение
	return nums
}

func Chunk(nums []int, size int) [][]int {
	if size <= 0 {
		return nil
	}
	var chunks [][]int                     // Результирующий срез чанков
	for i := 0; i < len(nums); i += size { // Шаг с размером "size"
		end := i + size      // Вычисляем конец текущего чанка
		if end > len(nums) { // Проверяем, не вышли ли за пределы среза
			end = len(nums) // Корректируем конец, если нужно
		}
		chunks = append(chunks, nums[i:end]) // Добавляем текущий чанк в результат
	}
	return chunks
}

func main() {
	numbers := []int{1, 2, 3, 4, 5, 12, 17, 23, 34, 45}
	result := SumAll(numbers)
	println("The sum is:", result)

	println("----------------")
	evenNumbers := filterEvenNumbers(numbers)
	println("Even numbers:")
	for _, num := range evenNumbers {
		println(num)
	}

	println("----------------")
	ReverseSlice(numbers)
	println("Reversed numbers:")
	fmt.Println(numbers)

	println("----------------")
	updatedNumbers := DeleteElementAtIndex(numbers, 3)
	println("After deleting element at index 3:")
	fmt.Println(updatedNumbers)

	println("----------------")
	newNumbers := Insert(updatedNumbers, 4, 99)
	println("After inserting 99 at index 4:")
	fmt.Println(newNumbers)
	newNumbers2 := Insert(newNumbers, 4, 101)
	println("After inserting 101 at index 4:")
	fmt.Println(newNumbers2)

	println("----------------")
	chunked := Chunk(newNumbers2, 3)
	println("Chunked numbers (size 3):")
	fmt.Println(chunked)
}
