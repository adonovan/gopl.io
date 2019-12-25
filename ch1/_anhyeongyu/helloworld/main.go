package main

import "fmt"

func main() {
	fmt.Println("Hello, 世界") // UTF-8 지원
}

// $ go "run" main.go
// .go 로 끝나는 한 개 이상의 소스 파일을
// 컴파일하고 라이브러리와 링크한 후 결과 실행 파일을 구동

// $ go "build" main.go
// 컴파일한 뒤 바이너리 파일을 생성 (main)

// go install
// GOPATH의 bin 폴더에 실행파일 저장
// GOPATH의 bin 폴더가 PATH로 잡혀있다면 어디서든 실행 가능

// 세미콜론은 컴파일 시에 자동 삽입
