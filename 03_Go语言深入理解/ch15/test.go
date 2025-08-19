package main
import (
    "fmt"
    "io/ioutil"
    "os"
)

func main() {
    inputFile := "D:\\Learning\\learn_Go\\03_Go语言深入理解\\ch15\\products_copy.txt"
    outputFile := "D:\\Learning\\learn_Go\\products_copy.txt"
    
    buf, err := ioutil.ReadFile(inputFile)
    if err != nil {
        fmt.Fprintf(os.Stderr, "File Error: %s\n", err)
        // panic(err.Error())
        }
    fmt.Printf("%s\n", string(buf))
    err = ioutil.WriteFile(outputFile, buf, 0644) // oct, not hex
    if err != nil {
        panic(err.Error())
    }
}