# bahttext
convert float number to thai text baht currency

inspiration:: https://github.com/joofjang/numgothai

## Example
```go
package main

import (
	"fmt"
	"github.com/smalldog124/bahttext"
)

func main()  {
	amount := 2000.0
	thaiText := bahttext.THBText(amount)
	fmt.Println(thaiText)
	// สองพันบาทถ้วน
}
```
