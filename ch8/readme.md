# Chapter 8

## Example typing submit

|Submit|Submit|Submit|
| ----- | ----- | ----- |
|✅ 강익준     |⬜️ 노영일 |⬜️ 원치훈     |
|⬜️ 김태호     |⬜️ 배지원 |⬜️ 이민철     |
|⬜️ 김현성     |⬜️ 백경준 |✅ 이범용     |
|⬜️ 김형진     |⬜️ 서병선 |⬜️ 이재성     |
|⬜️ 김혜곤     |⬜️ 손민성 |⬜️ 임지애     |
|⬜️ 권용민(on) |⬜️ 신진환 |⬜️ 정현석     |
|⬜️ 김혜림     |⬜️ 안현규 |⬜️ 한상윤     |
|             |         |⬜️ 홍예브게니  |


✅: submit
⬜️: submit not

## Tips

### cake

- cake_test.go 에는 default 에서 testing.Verbose() 가 쓰이는데 이건 testing.Init() 이전에 불리는 셈이라 문제가 있다. 각각의 Benchmark 함수 안에 default 를 정의하거나 혹은 다른 방법이 필요하다. 
