package main

import (
    "bufio"
    "fmt"
    "flag"
    "os"
    "time"
)

func check(e error) {
  if e != nil {
    panic(e)
  }
}

func add(filename *string, args []string) {
  var item string

  t := time.Now()

  f, err := os.OpenFile(*filename, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
  check(err)

  item = args[2]

  fmt.Println(item)
  n, err := f.WriteString(t.Format("2006-01-02") + " " + item + "\n")

  fmt.Println(n)
  check(err)

  f.Close()
}

func ls(f *os.File) {
  scanner := bufio.NewScanner(f)

  for scanner.Scan() {
    buffer := scanner.Text()

    fmt.Println(buffer)
  }
}

func main() {
  var cmd string
  filename := flag.String("filename", "todo.txt", "a filename to parse")

  flag.Parse()

  args := os.Args

  if len(args) > 1 {
    cmd = args[1]
  } else {
    cmd = "ls"
  }

  f, err := os.Open(*filename)
  check(err)

  if cmd == "ls" {
    ls(f)
  } else if (cmd == "add") {
    add(filename, args)
  }
}
