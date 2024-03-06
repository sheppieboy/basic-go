package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

type xkcd struct {
    Num        int    `json:"num"`
    Day        string `json:"day"`
    Month      string `json:"month"`
    Year       string `json:"year"`
    Title      string `json:"title"`
    Transcript string `json:"transcript"`
}

func main() {
    if len(os.Args) < 2 {
        fmt.Fprintln(os.Stderr, "Need json file and find file")
        os.Exit(-1)
    }

    fn := os.Args[1]

    if len(os.Args) < 3 {
        fmt.Fprintln(os.Stderr, "No search terms")
        os.Exit(-1)
    }

    var (
        item  xkcd
        terms []string
        input io.ReadCloser
        cnt   int
        err   error
    )

    // Decode file
    input, err = os.Open(fn)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error with file: %s\n", err)
        os.Exit(-1)
    }
    defer input.Close()

    if err = json.NewDecoder(input).Decode(&item); err != nil {
        fmt.Fprintf(os.Stderr, "Messed up json: %s\n", err)
        os.Exit(-1)
    }

    fmt.Fprintf(os.Stderr, "read 1 comic\n")

    // Get search terms
    for _, t := range os.Args[2:] {
        terms = append(terms, strings.ToLower(t))
    }

    // Search
    title := strings.ToLower(item.Title)
    transcript := strings.ToLower(item.Transcript)

    for _, term := range terms {
        if !strings.Contains(title, term) && !strings.Contains(transcript, term) {
            continue
        }
    }

    fmt.Printf("https://xkcd.com/%d/%s/%s/%s %s\n", item.Num, item.Month, item.Day, item.Year, item.Title)
    cnt++

    fmt.Fprintf(os.Stderr, "found %d comics\n", cnt)
}
