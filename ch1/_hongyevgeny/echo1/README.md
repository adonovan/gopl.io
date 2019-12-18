## 예제 코드

```
package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	for i :=1;i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
```

## 실행 결과

```
hello nice to meet you
```
