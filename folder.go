package main

import (
  "fmt"
  "os"
  "path/filepath"
)

func main() {
  filepath.Walk("/users/", func(path string, info os.FileInfo, err error) error{
    fmt.Println(path)
    return nil
  })
}
