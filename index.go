package main

import "github.com/go-martini/martini"

func main() {
  m := martini.Classic()

  m.Get("/index.html", func() string{
    return "hello, world"
  })

  m.Post("/index.html", func() string{
    return "hello, world"
  })

  m.Get("/result.html", func() string{
    return "hello, world"
  })

  m.Post("/result.html", func() string{
    return "hello, world"
  })

  m.Get("/search.json", func() string{
    return "hello, world"
  })

  m.Post("/search.json", func() string{
    return "hello, world"
  })
  
  m.Run()
}
