package main

import (
  "github.com/01-edu/z01"
  "os"
  "fmt"
)

func main() {
  args := os.Args[1:]
  //Create empty board
  var SolutionTable [9][9]rune
  //Fill the board with arg values
  for i:=range args {
    temp_rune := []rune(args[i])
    for j:=range temp_rune {
      SolutionTable[i][j] = temp_rune[j]
    }
  }

  //********ADD CONDITION FOR VALIDITY OF GIVEN NUMBERS*************
  if checkBoard(SolutionTable) {
    //Solve sudoku
    solve(&SolutionTable, 0, 0)

    for i:=0; i<9; i++ {
      for j:=0; j<9; j++ {
        z01.PrintRune(SolutionTable[i][j])
      }
      z01.PrintRune(10)
    }
  } else {
    fmt.Println("Error")
  }

}

func solve(SolutionTable *[9][9]rune, row, col int) bool {
  temp_board := *SolutionTable
  //Last column, Go to next row
  if col == 9 {
    row++
    col = 0
    //Last row and last column, Finish everything
    if row == 9 {
      return true
    }
  }

  //If cell is already filled, skip it, go to next value in row
  if temp_board[row][col] != '.' {
    *SolutionTable = temp_board
    return solve(SolutionTable, row, col+1)
  }

  for i:=1; i<=9;i++ {
    r := getRune(i)

    *SolutionTable = temp_board
    if isValid(SolutionTable, row, col, r) {
      //Place this rune to board after checking it validity
      temp_board[row][col] = r
      //Check if it works for next values
      *SolutionTable = temp_board
      if solve(SolutionTable, row, col + 1) {
        return true
      }
      //Condition when it does not work, Return empty entry
      temp_board[row][col] = '.'
      *SolutionTable = temp_board
    }
  }
  return false
}

//Check for validity of given rune accodring to row, col and subgrid
func isValid(SolutionTable *[9][9]rune, row, col int, r rune) bool {
  temp_board := *SolutionTable
  //Check column
  for i:=0; i<9; i++ {
    if temp_board[row][i] == r {
      return false
    }
  }

  //Check row
  for i:=0; i<9; i++ {
    if temp_board[i][col] == r {
      return false
    }
  }

  //Check subgrid
  horiz_grid := col / 3
  vert_grid := row / 3
  left_grid := horiz_grid * 3
  top_grid := vert_grid * 3

  for i:=0; i<3; i++ {
    for j:=0; j<3; j++ {
      if temp_board[top_grid + i][left_grid + j] == r {
        return false
      }
    }
  }
  return true
}

//Get rune from int
func getRune(i int) rune {
  num := 1
  r := '1'
  for num < i {
    num++
    r++
  }
  return r
}

//Check board before solving
func checkBoard(SolutionTable [9][9]rune) bool {
  for i:=0; i<9; i++ {
    for j:=0; j<9; j++ {
      if SolutionTable[i][j] == 46 {
        continue
      }
      if !boardValid(SolutionTable, i, j) {
        return false
      }
    }
  }
  return true
}

func boardValid(SolutionTable [9][9]rune, i, j int) bool {
  for m := 0; m < 9; m++ {
		if m != j && SolutionTable[i][m] == SolutionTable[i][j] {
      return false
    }
		if (m != i && SolutionTable[m][j] == SolutionTable[i][j]) {
      return false
    }
	}
	for m := i / 3 * 3; m < i / 3 * 3 + 3; m++ {
		for n := j / 3 * 3; n < j / 3 * 3 + 3; n++ {
			if m != i && n != j && SolutionTable[m][n] == SolutionTable[i][j] {
        return false
      }
		}
	}
	return true
}
