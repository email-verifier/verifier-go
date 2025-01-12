# Official Go Library for VerifyRight
A Go Library to verify invalid domains, disposable, non-existent emails. [know more](https://verifyright.co)

# Installation
Install by running the below command

```bash
go get github.com/email-verifier/verifier-go
```

# Usage
Below is the example of how to use the library

```go
package main

import (
	"fmt"

	emailVerifier "github.com/email-verifier/verifier-go"
)

func main() {
	response := emailVerifier.Verify("meet@gmail.com", "ACCESS_TOKEN")
	fmt.Println(response)
}

```
`emailVerifier.Verify` will return a boolean in response.

`true` if email exists
`false` if email does not exist

# Error Handling
The library will return `false` in case of any errors (invalid token, network issues, etc.).
