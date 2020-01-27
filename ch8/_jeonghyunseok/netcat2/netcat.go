// 이번에는 읽고 쓰기 가능

package main

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	go mustCopy(os.Stdout, conn)
	mustCopy(conn, os.Stdin)
}
func mustCopy(dst is.Writer, src. io.Reader) {
	if _, err := io.Copy(dst, src); err !=nil {
		log.Fatal(err)
	}
}