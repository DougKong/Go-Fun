package main

import "github.com/go-martini/martini"

func main() {
  m := martini.Classic()
  m.Get("/:name", func(params martini.Params) string {
    return "Hello" + params["name"]
  })

  m.Post("/new", func () string {
    return "hey"
  })
  m.Run()
}