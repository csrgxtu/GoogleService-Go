package main

import "fmt"
import "github.com/parnurzeal/gorequest"

func main() {
  request := gorequest.New()
  resp, body, errs := request.Get("https://www.google.com").End()
  if errs != nil {
    fmt.Printf("errs")
  }
  resp.Body.Close()
  fmt.Printf(body)
}
