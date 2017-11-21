package main

import "fmt"

func main() {
	ToLCD(8)
}

func ToLCDMultiple(number ...int) string {
	var (
		result1, result2, result3 string
		digits                    = map[int][]int{
			0: {0, 1, 0, 1, 0, 1, 1, 1, 1},
			8: {0, 1, 0, 1, 1, 1, 1, 1, 1},
		}
	)
	for _, num := range number {
		for pos, onoff := range digits[num] {
			switch pos {
			case 0, 1, 2:
				result1 += Draw(pos, onoff)
			case 3, 4, 5:
				result2 += Draw(pos, onoff)
			case 6, 7, 8:
				result3 += Draw(pos, onoff)
			}
		}
	}
	return fmt.Sprintf("%s\n%s\n%s", result1, result2, result3)
}

func ToLCD(number int) string {
	var (
		result string
		digits = map[int][]int{
			0: {0, 1, 0, 1, 0, 1, 1, 1, 1},
			8: {0, 1, 0, 1, 1, 1, 1, 1, 1},
		}
	)
	for pos, onoff := range digits[number] {
		if pos != 0 && pos%3 == 0 {
			result += "\n"
		}
		result += Draw(pos, onoff)
	}
	return result
}

func Draw(pos int, onoff int) string {
	switch pos {
	case 0, 2:
		return " "
	case 1, 4, 7:
		if onoff == 1 {
			return "_"
		}
	case 3, 5, 6, 8:
		if onoff == 1 {
			return "|"
		}
	default:
		panic("unexpected pos")
	}
	return " "
}
