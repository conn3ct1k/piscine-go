package main

import "github.com/01-edu/z01"

func PrintStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

type Door struct {
	state bool
}

func OpenDoor(ptrDoor *Door) {
	PrintStr("Door Opening...\n")
	ptrDoor.state = true
}

func CloseDoor(ptrDoor *Door) bool {
	PrintStr("Door Closing...\n")
	ptrDoor.state = false
	return true
}

func IsDoorOpen(Door *Door) bool {
	PrintStr("is the Door opened ?\n")
	return Door.state
}

func IsDoorClose(ptrDoor *Door) bool {
	PrintStr("is the Door closed ?\n")
	return ptrDoor.state
}

func main() {
	door := Door{}

	OpenDoor(&door)
	IsDoorClose(&door)
	if IsDoorOpen(&door) {
		CloseDoor(&door)
	}
}
