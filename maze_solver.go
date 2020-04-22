/**
* @author mariogmarq
* @version 1.0
* @description basic maze solver made in go
 */

package main

import (
	"errors"
	"fmt"
)

const ENTRANCE int = 6
const EXIT int = 7
const PATH int = 2

type maze struct {
	cell [][]int
}

func main() {
	lab := maze{[][]int{
		{1, 1, 7, 1, 1, 1, 1, 1},
		{1, 1, 0, 1, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 1, 0, 1},
		{1, 1, 1, 0, 1, 1, 1, 1},
		{1, 1, 1, 0, 0, 0, 0, 1},
		{1, 1, 1, 1, 1, 6, 1, 1},
	}}
	y, x, err := lab.FindEntrance()
	if err != nil {
		fmt.Println(err)
	} else {
		lab.solve(y, x)
	}
}

/**
* @brief checks if you can move to a certain position
* @param the position to be checked
* @retval true if the move is posible or false if not
 */
func (m maze) valid(y int, x int) bool {
	if m.cell[y][x] == 0 {
		return true
	}
	return false
}

/**
* @brief checks if a certain position is the exit
* @param the position to be checked
* @return true if the position is the exit or false if not
 */
func (m maze) isExit(y int, x int) bool {
	if m.cell[y][x] == EXIT {
		return true
	}
	return false
}

/**
* @brief finds the entrance of the maze
* @return the position fo the entrance, in case that there is not entrance it throws an error
 */

func (m maze) FindEntrance() (int, int, error) {
	for y := range m.cell {
		for x := range m.cell[0] {
			if m.cell[y][x] == ENTRANCE {
				return y, x, nil
			}
		}

	}
	return -1, -1, errors.New("There is no entrance to the maze")
}

/**
* @brief Solves the maze showing the path to the screen
 */

//It uses recursivity in order to undo the incorrect moves, but due to this once the maze is solved it begins to be unsolved
//Thats the reason of the Println statement in this function
func (m maze) solve(posy int, posx int) {
	if posy < (len(m.cell) - 1) {
		if m.valid(posy+1, posx) {
			m.cell[posy+1][posx] = PATH
			m.solve(posy+1, posx)
			m.cell[posy+1][posx] = 0
		} else if m.isExit(posy+1, posx) {
			for i := 0; i < len(m.cell); i++ {
				fmt.Println(m.cell[i])
			}
		}
	}

	if posy > 0 {
		if m.valid(posy-1, posx) {
			m.cell[posy-1][posx] = PATH
			m.solve(posy-1, posx)
			m.cell[posy-1][posx] = 0
		} else if m.isExit(posy-1, posx) {
			for i := 0; i < len(m.cell); i++ {
				fmt.Println(m.cell[i])
			}
		}
	}
	if posx < len(m.cell[0])-1 {
		if m.valid(posy, posx+1) {
			m.cell[posy][posx+1] = PATH
			m.solve(posy, posx+1)
			m.cell[posy][posx+1] = 0
		} else if m.isExit(posy, posx+1) {
			for i := 0; i < len(m.cell); i++ {
				fmt.Println(m.cell[i])
			}
		}
	}
	if posx > 0 {
		if m.valid(posy, posx-1) {
			m.cell[posy][posx-1] = PATH
			m.solve(posy, posx-1)
			m.cell[posy][posx-1] = 0
		} else if m.isExit(posy, posx-1) {
			for i := 0; i < len(m.cell); i++ {
				fmt.Println(m.cell[i])
			}
		}
	}
}
