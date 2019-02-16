default:dev1.0 dev2.0 dev3.0 dev4.0 dev5.0 dev6.0 dev6.1 dev6.2 dev7 dev8 dev9 dev10
dev1.0:
	@echo "====test leve 1.0===="
	@echo "> 支持基本+-*/ <"
	@go run main.go test/test1
dev2.0:
	@echo "====test leve 2.0===="
	@echo "> 支持自定义func < "
	@go run main.go test/test2
dev3.0:
	@echo "====test leve 3.0===="
	@echo "> 支持复杂数据类型 <"
	@go run main.go test/test3
dev4.0:
	@echo "====test leve 4.0===="
	@echo "> 支持for循环 <"
	@go run main.go test/test4
dev5.0:
	@echo "====test leve 5.0===="
	@echo "> 支持append <"
	@go run main.go test/test5
dev6.0:
	@echo "====test leve 6.0===="
	@echo "> 支持自定义结构体 <"
	@echo "> but 会变成mp     <"
	@echo "> so TODO          <"
	@go run main.go test/test6
dev6.1:
	@echo "====test leve 6.1===="
	@echo "> 支持结构体func   <"
	@go run main.go test/test6.1
dev6.2:
	@echo "====test leve 6.2===="
	@echo ">支持结构体指针func<"
	@go run main.go test/test6.2

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
	@echo "====test leve 10.0===="
	@echo "> 支持import       <"
	@go run main.go test/test10
dev10.1:
	@echo "====test leve 10.1===="
	@echo "> 支持import缓存   <"
	@go run main.go test/test10.1
dev11:
	@echo "====test leve 11===="
	@echo "> 支持shell       <"
	@go run main.go test/test11
dev12:
	@echo "====test leve 12===="
	@echo "> 支持全部二进制  <"
	@go run main.go test/test12
dev13: 
	@echo "====test leve 13===="
	@echo "> 支持全部数据类型<"
	@go run main.go test/test13
build:
	go build
