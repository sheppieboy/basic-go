package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

const url = "https://jsonplaceholder.typicode.com" //return to us some json

type todo struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {
	resp, err := http.Get(url + "/todos/1")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(-1)
	}

	//defer will wait for response till after
	defer resp.Body.Close() //response could have body, we need to close it as socket wont close

	if resp.StatusCode == http.StatusOK {
		body, err := io.ReadAll(resp.Body)

		if err != nil {
			fmt.Println(os.Stderr, err)
			os.Exit(-1)
		}

		var item todo

		err = json.Unmarshal(body, &item) // use = instead of := here

		if err != nil {
			fmt.Println(os.Stderr, err)
			os.Exit(-1)
		}

		fmt.Printf("%#v\n", item) // use Printf instead of Println
	}
}
