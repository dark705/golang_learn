package main

import (
	"fmt"
	s "strings"
)

func main() {
	p := fmt.Println
	_, _ = p(s.Count("test", "t"))
	_, _ = p(s.Contains("test", "es"))
	n, _ := p(s.Repeat("test1", 3))
	fmt.Println(n)

	_, _ = p(s.Join([]string{"a", "b", "c", "d"}, "_"))
	_, _ = p(s.Split("a_b_c_d", "_"))
}
