# sync

## [Mutex](http://en.wikipedia.org/wiki/Mutualexclusion)

여러 고루틴에서 공유하는 데이터를 보호해야 할 때 사용

다음과 같은 함수를 제공

```go
func (m *Mutex) Lock()
```

```go
func (m *Mutex) Unlock()
```

임계 영역(critical section)의 코드를 실행하기 전에는 뮤텍스의 `Lock()` 메서드로 잠금을 하고, 처리 완료 후에는 `Unlock()` 메서드로 잠금을 해제한다.

Lock, Unlock 함수는 반드시 짝을 맞춰야 하며 짝이 맞지 않으면 데드락(deadlock, 교착 상태)이 발생하므로 주의해야 한다.
