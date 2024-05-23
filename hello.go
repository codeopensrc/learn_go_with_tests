package main

import "fmt"

const (
    englishHelloPrefix = "Hello, "
    spanishHelloPrefix = "Hola, "
    frenchHelloPrefix = "Bonjour, "
)

func Hello(name string, lang string) string {
    if name == "" {
        name = "World"
    }

    return greetingPrefix(lang) + name
}

func greetingPrefix(lang string) (prefix string) {

    switch lang {
        case "Spanish":
            prefix = spanishHelloPrefix
        case "French":
            prefix = frenchHelloPrefix
        default:
            prefix = englishHelloPrefix
    }
    return
}

func main() {
    fmt.Println(Hello("world", ""))
}
