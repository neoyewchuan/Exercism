package twofer

import "fmt"

func ShareWith(para1 string) string {
	var name = para1
	if para1 == "" {
		name = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}
