package url

// Values 는 문자열 키를 값 목록과 매핑한다
type Values map[string][]string

// Get 은 주어진 키와 연관된 첫 번째 값을 반환하거나
// 값이 없으면 ""을 반환한다
func (v Values) Get(key string) string {
	if vs := v[key]; len(vs) > 0 {
		return vs[0]
	}
	return ""
}

// Add 는 값을 키에 추가한다.
// 이 메소드는 키와 연관된 기존 값에 추가한다
func (v Values) Add(key, value string) {
	v[key] = append(v[key], value)
}