/*First number is length of array, we dont use it. Second input is space separated long integers.
Task is to add the integers up. The total may exceed the 32-bit range so we used a 64-bit int
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var size, total int64

	fmt.Scanf("%d", &size)
	s := bufio.NewScanner(os.Stdin)
	s.Scan()
	line := s.Text()

	t := strings.Split(line, " ")

	for _, v := range t {
		i, _ := strconv.ParseInt(v, 10, 64)
		total += i
	}

	fmt.Println(total)

}
