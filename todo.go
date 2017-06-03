package main

import (
    "bufio"
    "fmt"
    "flag"
    "os"
)

func check(e error) { 
    if e != nil {
        panic(e)
    }
}

func main() {
  filename := flag.String("filename", "todo.txt", "a filename to parse")

  flag.Parse()

  f, err := os.Open(*filename)
  check(err)

  scanner := bufio.NewScanner(f)

  for scanner.Scan() {
    buffer := scanner.Text()

    fmt.Println(buffer)
  }
}
