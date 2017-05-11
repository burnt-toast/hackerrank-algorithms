/*Input: 3 lines
 Example:
 	3 - number of rows/columns in square
 	1 4 5 - square row 1
	2 2 6 - square row 2
	5 7 8 - square row 3

 Calculate the absolute difference between the two sums of the matrix's diagonals as a single integer
*/
package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	var squareSize int
	fmt.Scanf("%d", &squareSize)

	square := readSquare(squareSize)

	pds, sds := getDiagonalSums(square)

	fmt.Println(math.Abs(float64(pds) - float64(sds)))

}

func readSquare(size int) [][]int {
	square := make([][]int, size)
	s := bufio.NewScanner(os.Stdin)

	for i := 0; i < size; i++ {
		temp, err := readSquareLine(s, size)
		if err != nil {
			fmt.Println("SH*& is broken")
		} else {
			square[i] = temp
		}

	}
	return square
}

func readSquareLine(s *bufio.Scanner, size int) ([]int, error) {
	result := make([]int, size)
	s.Scan()
	line := s.Text()
	stringSlice := strings.Split(line, " ")
	for i, v := range stringSlice {
		d, err := strconv.Atoi(v)
		if err != nil {
			return nil, err
		}
		result[i] = d
	}

	return result, nil
}

func getDiagonalSums(square [][]int) (int, int) {
	var pds, sds int
	size := len(square[0])

	for i := 0; i < size; i++ {
		pdsPoint := (size - 1) - i
		pds += square[i][pdsPoint]
		sds += square[i][i]
	}

	return pds, sds
}
