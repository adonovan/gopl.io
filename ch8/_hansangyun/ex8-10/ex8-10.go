package main

/*
연습문제 8-10

HTTP 요청은 http.Request 구조체의 부가적인 Cancel 채널을 닫아서 취소할 수 있다.
8.6절의 웹 크롤러가 취소할 수 있게 수정하라.
힌트: 편의 함수 http.Get은 Request를 사용자가 직접 수정할 수 없다.
대신 http.NewRequest로 요청을 생성하고 Cancel 필드를 지정한 후 http.DefaultClient.Do(req)를 호출해 요청을 수행하라.
*/
