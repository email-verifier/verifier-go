# Official Go Library for verifier.meetchopra.com

# Installation
```github.com/email-verifier/verifier-go```

# Usage
Below is the example of how to use the library

```go
package main

import (
   "fmt"
   "github.com/email-verifier/verifier-go"
)

func main() {
   response := emailVerifier.verify("meet@gmail.com", "newasd")
   fmt.Println(response)
}
```
