echo "## 예제 코드" > README.md
echo "" >> README.md
echo "\`\`\`" >> README.md
cat main.go >> README.md
echo "\`\`\`" >> README.md
echo "" >> README.md
echo "## 실행 결과" >> README.md
echo "" >> README.md
echo "\`\`\`" >> README.md
go run main.go &
sleep 5
curl http://localhost:8000 >> README.md
curl http://localhost:8000/count >> README.md
curl http://localhost:8000 >> README.md
curl http://localhost:8000/count >> README.md
echo "\`\`\`" >> README.md

