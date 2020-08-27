package main

import (
  "fmt"
  "strings"
)

func main() {
 firstn := "ashish"
 lastn := "tyagi"
 fmt.Println("This is print line")
 fmt.Println(convertor(firstn,lastn))
}
func convertor(fn, ln string) (s1, s2 string) {
 fn = strings.Title(fn)
 ln = strings.ToUpper(ln)
 return fn, ln
}
