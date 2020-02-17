package main

import (
    // "bufio"
    // "fmt"
    "io/ioutil"
    // "os"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {

    d1 := []byte("hello\ngo\n")
    err := ioutil.WriteFile("fs.md", d1, 0644)
    check(err)

    tm := "---\n" +
        "date: \n" +
        "title: \n" +
        "description: \n" +
        "images: \n" +
        "categories: \n" +
        "tags: \n"

    content := []byte(tm)
    errFile := ioutil.WriteFile("test.md", content, 0644)
    check(errFile)

}
