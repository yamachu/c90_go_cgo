package main

import "C"
import "aobayama"

//export myCLibrary
func myCLibrary(user *C.char) *C.char {
	go_str := aobayama.GoAobaUser(C.GoString(user))
	return C.CString(go_str)
}

func init () {}

func main () {}