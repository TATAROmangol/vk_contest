package main

import (
	"bufio"
	"os"
    . "vk_contest/application"
)

func main(){
    in := bufio.NewReader(os.Stdin)
    out := bufio.NewWriter(os.Stdout)
    outErr := bufio.NewWriter(os.Stderr)
    defer out.Flush()
    defer outErr.Flush()

    app := NewApplication(in, out, outErr)
    app.Run()
}