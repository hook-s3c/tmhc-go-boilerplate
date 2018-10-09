package main

import (
    "fmt"
    "io"
    "os"
)

func main() {
    hello(os.Stdout)
}

func hello(out io.Writer) {
    fmt.Fprintln(out, "Hello TMHC!")
}
