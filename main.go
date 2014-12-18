package main

import (
    "flag"
    "fmt"
    "time"
)

func main() {
    port := flag.Int("p", 3000, "port to run on default 3000")
    delay := flag.Duration("d", time.Second*2, "delay to show this thing")

    flag.Parse()

    flag.VisitAll(func(f *flag.Flag) {
        fmt.Printf("%s=%v\n", f.Name, f.Value)
    })
    fmt.Println("Args: ", flag.Args())

    time.Sleep(*delay)
    fmt.Printf("Fake server running on port %d after waiting %s", *port, *delay)
}
