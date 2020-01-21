package main

func main() {

}

// // IntList 는 정수의 링크드 리스
// // nil *IntList는 빈 목록을 표시함
// type IntList struct {
// 	Value int
// 	Tail  *IntList
// }

// // Sum Sum
// func (list *IntList) Sum() int {
// 	if list == nil {
// 		return 0
// 	}
// 	return list.Value + list.Tail.Sum()
// }

// var (
// 	mu      sync.Mutex
// 	mapping = make(map[string]string)
// )

// func Lookup(key string) string {
// 	mu.Lock()
// 	defer mu.Unlock()
// 	v := mapping[key]
// 	return v
// }

// var chche = struct {
// 	sync.Mutex
// 	mapping map[string]string
// }{
// 	mapping: make(map[string]string),
// }

// // Lookup Lookup
// func Lookup(key string) string {
// 	chche.Lock()
// 	defer chche.Unlock()
// 	v := chche.mapping[key]
// 	return v
// }
