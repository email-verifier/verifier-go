# Official Go Library for verifier.meetchopra.com
A Go Library to verify invalid domains, disposable, non-existent emails. [know more](https://verifier.meetchopra.com)

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

`true` if email exsits
`false` if email does not exsits
