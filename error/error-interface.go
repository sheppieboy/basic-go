package main

import "fmt"

type errFoo struct{
	err error
	path string
}

func (e errFoo) Error() string{
	return fmt.Sprintf("%s:", e.path, e.err)
}

func XYZ(a int) error {
	return nil
}

func main(){
	//err := XYZ(1)

	var err error = XYZ(1) //BAD: interface gets a nil concrete ptr

	if err != nil{
		fmt.Println("oops")
	}else{
		fmt.Println("OK!")
	}
}