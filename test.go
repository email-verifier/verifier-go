package main

import (
   "fmt"
   "github.com/email-verifier/verifier-go"
)

func main() {
   response := emailVerifier.verify('meet@gmail.com', 'newasd')
   fmt.Println(response)
}
