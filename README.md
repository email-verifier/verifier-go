# Official Go Library for verifier.meetchopra.com

# Installation
```go get github.com/email-verifier/verifier-go```

# Usage
Below is the example of how to use the library

```go
package main

import (
	"fmt"

	emailVerifier "github.com/email-verifier/verifier-go"
)

func main() {
	response := emailVerifier.Verify("meet@gmail.com", "newasd")
	fmt.Println(response)
}

```
