package graph

/* ------------------------------ Island Count ------------------------------ */

type Coord struct {
	r int
	c int
}

func IsValidCell(r int, c int, grid [][]byte) bool {
	numberOfRows := len(grid)
	numberOfCols := len(grid[0])
	if r < 0 || r >= numberOfRows {
		return false
	}

	if c < 0 || c >= numberOfCols {
		return false
	}

	return true
}

func ExploreNeighboor(r int, c int, grid [][]byte, stack *[]Coord, visitedNodes map[Coord]bool) {
	if IsValidCell(r, c, grid) {
		_, visited := visitedNodes[Coord{r, c}]
		if grid[r][c] == '1' && !visited {
			*stack = append(*stack, Coord{r, c})
			visitedNodes[Coord{r, c}] = true
		}
	}
}

func NumIslands(grid [][]byte) int {
	visitedNodes := make(map[Coord]bool)
	count := 0

	numberOfRows := len(grid)
	numberOfCols := len(grid[0])

	for r := 0; r < numberOfRows; r++ {
		for c := 0; c < numberOfCols; c++ {
			// Check if the current cell is a non visited land
			_, visited := visitedNodes[Coord{r, c}]
			if grid[r][c] == '1' && !visited {
				// DFS Start
				currentNode := Coord{r, c}
				visitedNodes[currentNode] = true
				stack := []Coord{}
				stack = append(stack, currentNode)

				for len(stack) > 0 {

					// Pop from stack
					poppedNode := stack[len(stack)-1]
					visitedNodes[poppedNode] = true
					stack = stack[:len(stack)-1]

					// Add all neighboor of the poppedNode and check if their coord exist and
					// finaly push them to the stack if they are not already visited
					if IsValidCell(poppedNode.r-1, poppedNode.c, grid) {
						ExploreNeighboor(poppedNode.r-1, poppedNode.c, grid, &stack, visitedNodes)
					}

					if IsValidCell(poppedNode.r+1, poppedNode.c, grid) {
						ExploreNeighboor(poppedNode.r+1, poppedNode.c, grid, &stack, visitedNodes)
					}

					if IsValidCell(poppedNode.r, poppedNode.c-1, grid) {
						ExploreNeighboor(poppedNode.r, poppedNode.c-1, grid, &stack, visitedNodes)
					}

					if IsValidCell(poppedNode.r, poppedNode.c+1, grid) {
						ExploreNeighboor(poppedNode.r, poppedNode.c+1, grid, &stack, visitedNodes)
					}
				}
				// DFS END
				count++
			}
		}
	}
	return count
}
