package testmod

import (
    "fmt"
)

func Hi(name string) string {
    fmt.Print(111)
    return fmt.Sprintf("Hi, %s", name)
}
