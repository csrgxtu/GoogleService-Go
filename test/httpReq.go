package main

import "net/http"
// import "io/ioutil"
import "fmt"

func main() {
  resp, err := http.Get("https://www.google.com")
  if err != nil {
    fmt.Printf("req err")
  } else {
    fmt.Printf(resp)
  }

  // defer resp.Body.Close()
  // body, err := ioutil.ReadAll(resp.Body)
  // fmt.Printf(body)
}
