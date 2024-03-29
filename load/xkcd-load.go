package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
)


func getAComic(i int) []byte {
	url := fmt.Sprintf("http://xkcd.com/%d/info.0.json", i)

	resp, err := http.Get(url)

	if err != nil { //did not get to website
		fmt.Fprintf(os.Stderr, "can't read: %s\n", err)
		os.Exit(-1)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK{
		fmt.Fprintf(os.Stderr, "skipping %d: go %d\n", i, resp.StatusCode)
		return nil
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil { //did not get to website
		fmt.Fprintf(os.Stderr, "invalid body: %s\n", err)
		os.Exit(-1)
	}

	return body
}

func main(){
	var (
		output io.WriteCloser = os.Stdout
		err	error
		cnt int
		fails int
		data []byte
	)

	if len(os.Args) > 0 {
		output, err = os.Create(os.Args[1])

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(-1)
		}

		defer output.Close()
	}
	fmt.Println("[")
	defer fmt.Println("]")

	for i:= 1; fails < 2; i++{
		if data = getAComic(i); data == nil{
			fails++
			continue
		}

		if cnt>0{
			fmt.Fprint(output, ",")
		}
	
		_, err = io.Copy(output, bytes.NewBuffer((data)))
	
		if err != nil{
			fmt.Fprintf(os.Stderr, "stopped: %s]n", err)
		}

		fails = 0;
		cnt++
	}

	

	fmt.Fprintf(os.Stderr, "read %d comics\n", cnt)
}