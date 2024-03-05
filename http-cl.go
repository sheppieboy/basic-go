package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main(){
	resp, err := http.Get("http://localhost:8080/" + os.Args[1])
	if err != nil{
		fmt.Fprintln(os.Stderr, err)
		os.Exit(-1)
	}

	//defer will wait for response till after
	defer resp.Body.Close() //response could have body, we need to close it as socket wont close


	if resp.StatusCode == http.StatusOK{
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(os.Stderr, err)
			os.Exit(-1)
		}

		fmt.Println(string(body))
	}
}