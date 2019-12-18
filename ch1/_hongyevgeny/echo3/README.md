## 예제 코드

```
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(strings.Join(os.Args[1:]," "))
}
```

## 실행 결과

```
hello nice to meet you
```
