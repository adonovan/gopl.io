echo "## 예제 코드" > README.md
echo "" >> README.md
echo "\`\`\`" >> README.md
cat main.go >> README.md
echo "\`\`\`" >> README.md
echo "" >> README.md
echo "## 실행 결과" >> README.md
echo "" >> README.md
echo "\`\`\`" >> README.md
go run main.go https://golang.org https://gopl.io https://godoc.org >> README.md
echo "\`\`\`" >> README.md

