/*Input: 2 lines of equal length comma separated numbers
 Example:
 	1 4 5
	2 2 6

 This represents alices and bobs score. Each person gets a point if the array[x] is high than the other person
 number is suppose to be at least 1 and no more than 100 but no instructions given for error ouput so not validating
*/
package main

import "bufio"
import "os"
import "fmt"
import "strings"
import "strconv"

func main() {
	s := bufio.NewScanner(os.Stdin)
	s.Scan()
	alice := s.Text()
	s.Scan()
	bob := s.Text()

	alicePoint, bobPoint := getComparisonPoints(alice, bob)

	fmt.Println(alicePoint, bobPoint)
}

func getComparisonPoints(a, b string) (int, int) {
	var aliceTotal, bobTotal int
	aliceArray := getArrayFromString(a)
	bobArray := getArrayFromString(b)

	for i, v := range aliceArray {
		if v > bobArray[i] {
			aliceTotal++
		} else if v < bobArray[i] {
			bobTotal++
		}
	}

	return aliceTotal, bobTotal
}

func getArrayFromString(s string) []int {
	temp := strings.Split(s, " ")
	result := make([]int, len(temp))
	for i, v := range temp {
		d, _ := strconv.Atoi(v)
		result[i] = d
	}
	return result
}
