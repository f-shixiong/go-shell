default:dev1 dev2 dev3 dev4 dev5 dev6 dev7 dev8 dev9
dev1:
	@echo "====test leve 1===="
	@echo "> 支持基本+-*/ <"
	@go run main.go test/test1
dev2:
	@echo "====test leve 2===="
	@echo "> 支持自定义func < "
	@go run main.go test/test2
dev3:
	@echo "====test leve 3===="
	@echo "> 支持复杂数据类型 <"
	@go run main.go test/test3
dev4:
	@echo "====test leve 4===="
	@echo "> 支持for循环 <"
	@go run main.go test/test4
dev5:
	@echo "====test leve 5===="
	@echo "> 支持append <"
	@go run main.go test/test5
dev6:
	@echo "====test leve 6===="
	@echo "> 支持自定义结构体 <"
	@echo "> but 会变成mp     <"
	@echo "> so TODO          <"
	@go run main.go test/test6

dev7:
	@echo "====test leve 7===="
	@echo "> 支持多层结构体   <"
	@go run main.go test/test7
dev8:
	@echo "====test leve 8===="
	@echo "> 支持指针         <"
	@go run main.go test/test8
dev9:
	@echo "====test leve 9===="
	@echo "> 支持  return     <"
	@go run main.go test/test9
dev10:
	@echo "====test leve 10===="
	@echo "> 支持import       <"
	@go run main.go test/test9
