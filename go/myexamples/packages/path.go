package main

import (
    "fmt"
    "path"
    "os"
)

func main() {
    wd, err := os.Getwd()
    if err != nil {
        panic(err)
    }

    fmt.Println("working directory: ", wd)

    fmt.Println("base: ", path.Base(wd))
    fmt.Println("dir: ", path.Dir(wd))
    fmt.Println("ext: ", path.Ext(wd+"howdy.test"))
    fmt.Println("abs: ", path.IsAbs(wd))
    fmt.Println("abs: ", path.IsAbs("howdy/../../"))

    dirty := "////howdy////lots/of//slashes//yeah/"
    fmt.Println("dirty: ", dirty)
    fmt.Println("clean: ", path.Clean(dirty))

    fmt.Println("joined: ", path.Join("one", "two", "three", "four"))


    dir, file := path.Split(wd+"/lala.txt")
    fmt.Println("dir: ", dir, " file: ", file)

    // Hmmmm, I wonder if path works with URLs.
    var base string
    url := "http://example.com/test/file.txt"
    base, file = path.Split(url)
    fmt.Printf("Got base %v and file %v\n", base, file)

    // Looks like the double slash after http: gets messed up by path.Dir.
    fmt.Printf("Base is %v\nDir is %v\n", path.Base(url), path.Dir(url))
}
