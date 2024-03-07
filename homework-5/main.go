package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

type sku struct{
	item, price string
}

var items = []sku{
	{"shoes", "46"},
	{"socks", "23"},
	{"sandles", "25"},
	{"clogs", "34"},
	{"pants", "30"},
	{"shorts", "20"},
}

func doQuery(cmd, parms string) error{
	resp, err := http.Get("http://localhost:8080/"+cmd+"?"+parms)

	if err != nil{
		fmt.Fprintf(os.Stderr, "err %s = %v\n", parms, err)
	}

	defer resp.Body.Close()
	fmt.Fprintf(os.Stderr, "got %s = %d (no err)\n",parms, resp.StatusCode)
	return nil
}

func runAdds(){
	for {
		for _, s := range items{
			if err := doQuery("create", "item="+s.item+"&price="+s.price); err !=nil{
				return
			}
		}
	}
}

func runUpdates(){
	for {
		for _, s := range items{
			if err := doQuery("update", "item="+s.item+"&price="+s.price); err !=nil{
				return
			}
		}
	}
}


func runDeletes(){
	for {
		for _, s := range items{
			if err := doQuery("delete", "item="+s.item); err !=nil{
				return
			}
		}
	}
}

func main(){
	go runAdds()
	go runDeletes()
	go runUpdates()
	time.Sleep(5*time.Second)
}