package main

import "fmt"

const english = "English"
const spanish = "Spanish"
const french = "French"
const hindi = "Hindi"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const hindiHelloPrefix = "Namaste, "

func Hello(name, language string) string {

  if name == "" {
    name = "world"
  }
  return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
  switch language {
  case french:
    prefix = frenchHelloPrefix
  case spanish:
    prefix = spanishHelloPrefix
  case hindi:
    prefix = hindiHelloPrefix
  default:
    prefix = englishHelloPrefix
  }
  return
}

func main() {
  fmt.Println(Hello("world", ""))
}
