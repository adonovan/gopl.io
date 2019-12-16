package main

import "fmt"

func main() {
	fmt.Println("Hello, 世界")
}

// $ go "run" main.go
// .go 로 끝나는 한 개 이상의 소스 파일을
// 컴파일하고 라이브러리와 링크한 후 결과 실행 파일을 구동

// $ go "build" main.go
// 컴파일한 뒤 바이너리 파일을 생성 (main)
