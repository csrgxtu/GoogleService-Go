package main

import "fmt"
import "time"
import "github.com/franela/goreq"

func main() {
  type Item struct {
    q string
    hl  string
    lr  string
    start int
    num int
  }

  item := Item {
    q: "ubuntu",
    hl: "en",
    lr: "en",
    start: 1,
    num: 10,
  }

  goreq.SetConnectTimeout(100 * time.Millisecond)
  
  res, err := goreq.Request{
    Uri: "https://www.google.com/search",
    QueryString: item,
  }.Do()

  fmt.Println(res, err)
}
