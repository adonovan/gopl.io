echo "## 예제 코드" > README.md
echo "" >> README.md
echo "\`\`\`" >> README.md
cat main.go >> README.md
echo "\`\`\`" >> README.md
echo "" >> README.md
echo "## 실행 결과" >> README.md
echo "" >> README.md
echo "\`\`\`" >> README.md
go run main.go hello nice to meet you >> README.md
echo "\`\`\`" >> README.md

