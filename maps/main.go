package main

import "fmt"

func CountUniqInts(nums []int) map[int]int {
	counts := make(map[int]int)
	for _, num := range nums {
		counts[num]++
	}
	return counts
}

func WordCount(s string) map[string]int {
	counts := make(map[string]int)
	word := ""
	for _, char := range s {
		if char == ' ' || char == '\n' || char == '\t' || char == ',' || char == '.' {
			if word != "" {
				counts[word]++
				word = ""
			}
		} else {
			word += string(char)
		}
	}
	if word != "" {
		counts[word]++
	}
	return counts
}

func IsAnagram(a, b string) bool {
	if len(a) != len(b) {
		return false
	}
	counts := make(map[rune]int)
	for _, char := range a {
		counts[char]++
	}
	for _, char := range b {
		counts[char]--
		if counts[char] < 0 {
			return false
		}
	}
	return true
}

func MapFlip(m map[string]int) map[int]string {
	flipped := make(map[int]string)
	for key, value := range m {
		flipped[value] = key
	}
	return flipped
}

func MergeMaps(m1, m2 map[string]int) map[string]int {
	merged := make(map[string]int)
	for key, value := range m1 {
		merged[key] = value
	}
	for key, value := range m2 {
		if _, exists := merged[key]; !exists {
			merged[key] = 0
		}
		merged[key] = merged[key] + value
	}
	return merged
}

type User struct {
	Name string
	City string
}

func GroupByCity(users []User) map[string][]User {
	cityGroups := make(map[string][]User)
	for _, user := range users {
		cityGroups[user.City] = append(cityGroups[user.City], user)
	}
	return cityGroups
}

func main() {
	fmt.Println("---- 1 ----")
	nums := []int{1, 2, 2, 3, 4, 4, 4, 5}
	counts := CountUniqInts(nums)
	fmt.Println(counts)

	fmt.Println("\n---- 2 ----")
	text := "Эх, если бы да кабы, во рту грибы росли, то был бы не рот, а целый огород"
	wordCounts := WordCount(text)
	fmt.Println(wordCounts)

	fmt.Println("\n---- 3 ----")
	a := "кабан"
	b := "банка"
	fmt.Printf("\"%s\" и \"%s\" анаграмма? %v\n", a, b, IsAnagram(a, b))

	fmt.Println("\n---- 4 ----")
	originalMap := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}
	flippedMap := MapFlip(originalMap)
	fmt.Println("оригинал", originalMap)
	fmt.Println("перевернутое", flippedMap)

	fmt.Println("\n---- 5 ----")
	map1 := map[string]int{"a": 3, "c": 5, "f": 3}
	map2 := map[string]int{"f": 2, "d": 4, "a": 5}
	mergedMap := MergeMaps(map1, map2)
	fmt.Println("map1:", map1)
	fmt.Println("map2:", map2)
	fmt.Println("merged:", mergedMap)

	fmt.Println("\n---- 6 ----")
	users := []User{
		{Name: "Вася", City: "Москва"},
		{Name: "Рус", City: "Алматы"},
		{Name: "Дима", City: "Москва"},
		{Name: "Вова", City: "Астана"},
		{Name: "Саша", City: "Алматы"},
		{Name: "Федя", City: "Алматы"},
	}
	groupedUsers := GroupByCity(users)
	fmt.Println(groupedUsers)
	fmt.Println("\nПользователи по городам:")
	for city, usersInCity := range groupedUsers {
		fmt.Printf("Город: %s\n", city)
		for _, user := range usersInCity {
			fmt.Printf(" - %s\n", user.Name)
		}
	}
}
