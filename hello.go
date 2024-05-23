package main

import "fmt"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

func Hello(name string, lang string) string {
    if name == "" {
        name = "World"
    }
    if lang == "Spanish" {
        return spanishHelloPrefix + name
    }
    if lang == "French" {
        return frenchHelloPrefix + name
    }
    return englishHelloPrefix + name
}

func main() {
    fmt.Println(Hello("world", ""))
}
