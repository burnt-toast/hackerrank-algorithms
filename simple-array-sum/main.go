package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var size, total int

	fmt.Scanf("%d", &size)
	s := bufio.NewScanner(os.Stdin)
	s.Scan()
	line := s.Text()

	t := strings.Split(line, " ")

	for _, v := range t {
		i, _ := strconv.Atoi(v)
		total += i
	}

	fmt.Println(total)

}
