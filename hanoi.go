package hanoi

import "fmt"

/*
Hanoi ...
  Calculate count
*/
func Hanoi(disks int) int {
	hanoi(disks, "1", "2", "3")
	return disks
}

func hanoi(disks int, a, b, c string) {
	if disks == 1 {
		fmt.Println(a, "->", b)
	} else {
		hanoi(disks-1, a, c, b)
		hanoi(disks-1, a, b, c)
		hanoi(disks-1, b, a, c)
	}
}
