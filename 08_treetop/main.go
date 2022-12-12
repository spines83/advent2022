package main

import (
	"advent2022/utils"
	"fmt"
	"strconv"
)

func main() {

	fileName := "input.txt"
	lineCount := utils.CountLines(fileName)
	grid := make([][]int, lineCount)

	lineCounter := 0
	lines, _ := utils.IterLines(fileName)
	for currentLine := range lines {
		grid[lineCounter] = make([]int, len(currentLine))
		for i, character := range currentLine {
			tree, _ := strconv.Atoi(string(character))
			grid[lineCounter][i] = tree
		}
		lineCounter += 1
	}

	testGrid(grid)
}

func testGrid(grid [][]int) {
	p1count := 0
	p2maxScenicScore := 0
	for i, row := range grid {
		for j := range row {
			northTest, northCount := testNorth(grid, i, j)
			southTest, southCount := testSouth(grid, i, j)
			eastTest, eastCount := testEast(grid, i, j)
			westTest, westCount := testWest(grid, i, j)
			if northTest || southTest || eastTest || westTest {
				p1count++
			}

			scenicScore := northCount * southCount * eastCount * westCount
			if scenicScore > p2maxScenicScore {
				p2maxScenicScore = scenicScore
			}
		}
	}

	fmt.Printf("%d\n", p1count)
	fmt.Printf("%d\n", p2maxScenicScore)
}

// Returns (isEdgeVisible, numVisibleTrees)
func testNorth(grid [][]int, i, j int) (bool, int) {
	tree := grid[i][j]
	for x := i - 1; x >= 0; x-- {
		if grid[x][j] >= tree {
			return false, i - x
		}
	}
	return true, i
}

// Returns (isEdgeVisible, numVisibleTrees)
func testSouth(grid [][]int, i, j int) (bool, int) {
	tree := grid[i][j]
	for x := i + 1; x < len(grid); x++ {
		if grid[x][j] >= tree {
			return false, x - i
		}
	}
	return true, len(grid) - i - 1
}

// Returns (isEdgeVisible, numVisibleTrees)
func testWest(grid [][]int, i, j int) (bool, int) {
	tree := grid[i][j]
	for y := j - 1; y >= 0; y-- {
		if grid[i][y] >= tree {
			return false, j - y
		}
	}
	return true, j
}

// Returns (isEdgeVisible, numVisibleTrees)
func testEast(grid [][]int, i, j int) (bool, int) {
	tree := grid[i][j]
	for y := j + 1; y < len(grid[0]); y++ {
		if grid[i][y] >= tree {
			return false, y - j
		}
	}
	return true, len(grid[0]) - j - 1
}
