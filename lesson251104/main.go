package main

import "fmt"

func sum(arr [10]int) int {
	total := 0
	for i := 0; i < len(arr); i++ {
		total += arr[i]
	}
	return total
}

func multiply(s []int, factor int) []int {
	for i := 0; i < len(s); i++ {
		s[i] *= factor
	}
	return s
}

func mainDiagonal(matrix [][]int) []int {
	diag := make([]int, len(matrix))
	for i := 0; i < len(matrix); i++ {
		diag[i] = matrix[i][i]
	}
	return diag
}

func secondaryDiagonal(matrix [][]int) []int {
	diag := make([]int, len(matrix))
	n := len(matrix)
	for i := 0; i < n; i++ {
		diag[i] = matrix[i][n-1-i]
	}
	return diag
}

func unique(arr []int) []int {
	seen := make(map[int]bool)
	result := []int{}
	for _, v := range arr {
		if !seen[v] {
			seen[v] = true
			result = append(result, v)
		}
	}
	return result
}

func main() {
	fmt.Println("----------- Task 1: ")
	num_arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sum := sum(num_arr)
	fmt.Println("Sum: ", sum)

	fmt.Println("\n----------- Task 2: ")
	slice := []int{2, 4, 6, 8, 10}
	slice2 := make([]int, len(slice))
	copy(slice2, slice)
	multiplied := multiply(slice, 3)
	fmt.Println("Result (slice): ", slice)
	println(slice)
	fmt.Println("Result (multiply): ", multiplied)
	println(multiplied)
	multiplied2 := multiply(slice2, 3)
	fmt.Println("Result (slice2): ", slice2)
	println(slice2)
	fmt.Println("Result (multiply2): ", multiplied2)
	println(multiplied2)

	fmt.Println("\n----------- Task 3: ")
	numbers := make([]int, 0, 20)
	for i := 1; i <= 10; i++ {
		numbers = append(numbers, i)
	}
	fmt.Println("Len:", len(numbers), "Cap:", cap(numbers))
	part := numbers[3:8]
	copy_part := make([]int, len(part))
	copy(copy_part, part)
	fmt.Println("Original slice: ", numbers)
	println(numbers)
	fmt.Println("Part:", part)
	println(part)
	fmt.Println("Copy of part:", copy_part)
	println(copy_part)
	part[0] = part[0] * 100
	part[4] = part[4] * 100
	fmt.Println("Changed part:", part)
	println(part)
	fmt.Println("Copy of part:", copy_part)
	println(copy_part)

	fmt.Println("\n----------- Task 4: ")
	matrix := make([][]int, 3)
	for i := 0; i < 3; i++ {
		matrix[i] = make([]int, 3)
		for j := 0; j < 3; j++ {
			matrix[i][j] = i + j
		}
	}
	fmt.Println("Matrix:")
	for i := 0; i < 3; i++ {
		fmt.Println(matrix[i])
	}
	fmt.Println("Main:", mainDiagonal(matrix))
	fmt.Println("Secondary:", secondaryDiagonal(matrix))

	fmt.Println("\n----------- Task 5: ")
	arr_with_dups := []int{1, 2, 3, 2, 4, 5, 1, 2, 3, 7}
	fmt.Println("Original array:", arr_with_dups)
	unique_arr := unique(arr_with_dups)
	fmt.Println("Unique array:", unique_arr)
}
