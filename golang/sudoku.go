package main

import (
	"fmt"
)

// Sudoku table
type Sudoku [9][9]uint8

// SudokuMeta is metadata
type SudokuMeta struct {
	correct   bool
	full      bool
	rowMap    [9]uint16
	cellMap   [9]uint16
	columnMap [9]uint16
}

func findBest(sudokudata SudokuMeta, sudoku Sudoku) (i uint16, j uint16, numbers []uint8, solved bool) {
	var spotmap uint16
	var counter, x, y, mincounter uint16
	var ncount uint8
	var bestnumbers, spotnumbers []uint8

	mincounter = 10 // Minimal count of possibilities per location
	cellsWithNumbers := 0

	for i = 0; i < 9; i++ {
		for j = 0; j < 9; j++ {
			spotnumbers = nil
			if sudoku[i][j] > 0 {
				cellsWithNumbers++
				continue
			}
			spotmap = sudokudata.rowMap[i] | sudokudata.columnMap[j] | sudokudata.cellMap[3*int(j/3)+int(i/3)]
			counter = 0 // how many possibilities are at this location
			for ncount = 1; ncount < 10; ncount++ {
				// numbers go from 1 to 9 not 0 to 8
				if ((spotmap >> ncount) & 1) == 0 {
					counter++
					spotnumbers = append(spotnumbers, ncount)
				}
			}
			printBitMap("spotmap", spotmap)
			printBitMap("rowmap   ", sudokudata.rowMap[j])
			printBitMap("columnmap", sudokudata.columnMap[i])
			printBitMap("cellmap  ", sudokudata.cellMap[3*int(j/3)+int(i/3)])
			fmt.Printf("Koht: %d,%d cell: %d counter: %d\n", i, j, 3*int(j/3)+int(i/3), counter)
			if counter == 1 {
				return i, j, spotnumbers, false
			}
			if counter < mincounter && counter > 0 {
				mincounter = counter
				x = i
				y = j
				bestnumbers = spotnumbers
			}
		}
	}
	fmt.Printf("Numbritega on %d\n", cellsWithNumbers)
	if mincounter == 10 {
		fmt.Println("Ei ole valikuid")
		if isSudokuSolved(sudoku) {
			return 10, 10, nil, true
		}
		return 10, 10, nil, false
	}
	fmt.Printf("Parim koht on %d %d\n", x, y)
	return x, y, bestnumbers, false
}

func printSudoku(sudoku Sudoku) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			fmt.Print(sudoku[i][j])
			fmt.Print(" ")
		}
		fmt.Println()
	}
}

func printSudokuWithMark(sudoku Sudoku, x uint16, y uint16) {
	var i, j uint16
	for i = 0; i < 9; i++ {
		for j = 0; j < 9; j++ {
			if i == x && j == y {
				fmt.Printf("[%d]", sudoku[i][j])
			} else {
				fmt.Printf(" %d ", sudoku[i][j])
			}
		}
		fmt.Println()
	}
}

func isThatSudoku(sudoku Sudoku) (sudokudata SudokuMeta) {
	var rowCorrect, cellCorrect, columnCorrect bool

	sudokudata.correct = false
	sudokudata.full = false

	for i := 0; i < 9; i++ {
		rowCorrect, sudokudata.rowMap[i] = isSudokuRowCorrect(i, sudoku)
		if !rowCorrect {
			fmt.Printf("Row %d incorrect\n", i)
			return
		}

		columnCorrect, sudokudata.columnMap[i] = isSudokuColumnCorrect(i, sudoku)
		if !columnCorrect {
			fmt.Printf("Column %d incorrect\n", i)
			return
		}

		cellCorrect, sudokudata.cellMap[i] = isSudokuCellCorrect(i, sudoku)
		if !cellCorrect {
			fmt.Printf("Cell %d incorrect\n", i)
			return
		}
	}
	sudokudata.correct = true
	return
}

func isSudokuRowCorrect(row int, sudoku Sudoku) (correct bool, rowMap uint16) {
	rowMap = 0
	for i := 0; i < 9; i++ {
		if sudoku[row][i] == 0 {
			continue
		}
		if rowMap&(1<<sudoku[row][i]) > 0 {

			return false, 0
		}
		rowMap = rowMap | (1 << sudoku[row][i])
	}
	return true, rowMap
}

func isSudokuColumnCorrect(column int, sudoku Sudoku) (correct bool, columnMap uint16) {
	columnMap = 0
	for i := 0; i < 9; i++ {
		if sudoku[i][column] == 0 {
			continue
		}
		if columnMap&(1<<sudoku[i][column]) > 0 {
			return false, 0
		}
		columnMap += 1 << sudoku[i][column]
	}
	return true, columnMap
}

func isSudokuCellCorrect(cell int, sudoku Sudoku) (correct bool, cellMap uint16) {
	var startx, starty int
	cellMap = 0
	startx = 3 * (cell % 3)
	starty = (cell / 3) * 3

	fmt.Printf("Cell: %d, Pos %d,%d\n", cell, startx, starty)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if sudoku[i+startx][j+starty] != 0 {
				if cellMap&(1<<sudoku[i+startx][j+starty]) > 0 {
					return false, 0
				}
				cellMap = cellMap | (1 << sudoku[i+startx][j+starty])
			}
		}
	}
	return true, cellMap
}

func printBitMap(nimi string, bitmap uint16) {
	fmt.Printf("Bitmap %s: %16.16b\n", nimi, bitmap)
}

func isSudokuSolved(sudoku Sudoku) (solved bool) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if sudoku[i][j] == 0 {
				return false
			}
		}
	}
	return true
}

func solveSudoku(sudoku Sudoku, sudokudata SudokuMeta) (allsolved bool) {
	var i, j uint16
	var numbers []uint8
	var solved bool

	i, j, numbers, solved = findBest(sudokudata, sudoku)

	if solved {
		fmt.Println("Sudoku lahendatud")
		printSudoku(sudoku)
		return solved
	}

	if numbers == nil {
		return false
	}

	for _, number := range numbers {
		fmt.Printf("Sama koha numbrid, parim koht %d,%d number: %d\n", i, j, number)
		sudoku[i][j] = number
		printSudokuWithMark(sudoku, i, j)
		solved = solveSudoku(sudoku, isThatSudoku(sudoku))
		if solved {
			return solved
		}
	}

	//time.Sleep(3 * time.Second)
	fmt.Printf("Bad solution\n")
	return false
}

func main() {

	sudoku := Sudoku{
		{8, 8, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 3, 6, 0, 0, 0, 0, 0},
		{0, 7, 0, 0, 9, 0, 2, 0, 0},

		{0, 5, 0, 0, 0, 7, 0, 0, 0},
		{0, 0, 0, 0, 4, 5, 7, 0, 0},
		{0, 0, 0, 1, 0, 0, 0, 3, 0},

		{0, 0, 1, 0, 0, 0, 0, 6, 8},
		{0, 0, 8, 5, 0, 0, 0, 1, 0},
		{0, 9, 0, 0, 0, 0, 4, 0, 0},
	}
	printSudoku(sudoku)
	sudokudata := isThatSudoku(sudoku)

	solveSudoku(sudoku, sudokudata)

}
