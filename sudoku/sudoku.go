package main

import (
	"fmt"
	"os"
)

// Global variable table, to access it from any function with no pointers and any variables
var (
	table1 [9][9]int
	table2 [9][9]int
)

func main() {
	args := []string(os.Args[1:])
	if len(args) == 1 {
		args = SplitWhiteSpaces(args[0])
	}
	if !input_corrected(args) {
		fmt.Println("Error")
		return
	}
	fill_table(args, &table1)

	fill_table(args, &table2)
	if !one_solution_base(table1) {
		fmt.Println("Error")
		return
	}
	if sudoku_solver1(0, 0) {
		sudoku_solver2(8, 8)
		if table1 == table2 {
			print_table(table1)
		} else {
			fmt.Println("Error")
			print_table(table1)
			fmt.Println()
			print_table(table2)
		}
	} else {
		fmt.Println("Error")
		return
	}
	fmt.Println()
}

func input_corrected(args []string) bool {
	// To check the correctness of entering arguments into the program
	if len(args) != 9 {
		return false
	}
	for i := 0; i < len(args); i++ {
		if len(args[i]) != 9 {
			return false
		}
	}
	for i := 0; i < len(args); i++ {
		for j := 0; j < len(args[i]); j++ {
			if (args[i][j] <= 48 || args[i][j] > 57) && !(args[i][j] == '.') {
				return false
			}
		}
	}
	return true
}

func fill_table(args []string, table *[9][9]int) {
	// To fill a table with data from Args
	for i := 0; i < len(args); i++ {
		for j := 0; j < len(args[i]); j++ {
			if args[i][j] == '.' {
				continue
			} else {
				table[i][j] = int(args[i][j] - '0')
			}
		}
	}
}

func print_table(table [9][9]int) {
	// To print a table
	for _, l := range table {
		for i, c := range l {
			fmt.Print(c)
			if i != len(l)-1 {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func one_solution_base(table [9][9]int) bool {
	// If the input values ​​are less than 16, then the Sudoku has multiple solutions
	counter := 0
	for i := range table {
		for j := range table[i] {
			if table[i][j] > 0 {
				counter++
			}
		}
	}
	return counter > 16
}

func is_valid(x int, y int, value int, table [9][9]int) bool {
	// Checks a number vertically, horizontally and inside a 3x3 square for no matches
	for i := range [9]int{} {
		if table[i][x] == value && i != y {
			return false
		}
	}

	for i := range [9]int{} {
		if table[y][i] == value && i != x {
			return false
		}
	}

	for i := 0; i < 9; i++ {
		counter := [10]int{}
		for j := 0; j < 9; j++ {
			counter[table1[i][j]]++
		}
		if check_copies(counter) {
			return false
		}

	}

	for i := 0; i < 9; i++ {
		counter := [10]int{}
		for j := 0; j < 9; j++ {
			counter[table1[j][i]]++
		}
		if check_copies(counter) {
			return false
		}

	}

	sx, sy := int(x/3)*3, int(y/3)*3
	for dy := range [3]int{} {
		for dx := range [3]int{} {
			if table[sy+dy][sx+dx] == value && sy+dy != y && sx+dx != x {
				return false
			}
		}
	}
	return true
}

func check_copies(counter [10]int) bool {
	for i, count := range counter {
		if i == 0 {
			continue
		}
		if count > 1 {
			return true
		}
	}
	return false
}

func next(x int, y int) (int, int) {
	// To move to the next x and y coordinate value
	nextX, nextY := (x+1)%9, y
	if nextX == 0 {
		nextY = y + 1
	}
	return nextX, nextY
}

func prev(x int, y int) (int, int) {
	// To move to the prev x and y coordinate value
	nextX, nextY := x-1, y
	if nextX < 0 {
		nextX = 8
	}
	if nextX == 8 {
		nextY = y - 1
	}
	return nextX, nextY
}

func sudoku_solver1(x int, y int) bool {
	// Sudoku solver, if the function returns false, then this sudoku has no solution
	if y == 9 {
		return true
	}
	if table1[y][x] != 0 {
		return sudoku_solver1(next(x, y))
	} else {
		for i := range [9]int{} {
			v := i + 1
			if is_valid(x, y, v, table1) {
				table1[y][x] = v
				if sudoku_solver1(next(x, y)) {
					return true
				}
				table1[y][x] = 0
			}
		}
		return false
	}
}

func sudoku_solver2(x int, y int) bool {
	// Sudoku solver, if the function returns false, then this sudoku has no solution
	if y == -1 {
		return true
	}
	if table2[y][x] != 0 {
		return sudoku_solver2(prev(x, y))
	} else {
		for i := range [9]int{} {
			v := i + 1
			if is_valid(x, y, v, table2) {
				table2[y][x] = v
				if sudoku_solver2(prev(x, y)) {
					return true
				}
				table2[y][x] = 0
			}
		}
		return false
	}
}

func SplitWhiteSpaces(s string) []string {
	word := ""
	var result []string
	for _, w := range s {
		if w == ' ' && word != "" {
			result = append(result, word)
			word = ""
		} else if w == ' ' {
			continue
		} else {
			word += string(w)
		}
	}
	if word != "" {
		result = append(result, word)
	}

	return result
}
