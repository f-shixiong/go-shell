default:dev1 dev2 dev3 dev4 dev5
dev1:
	@echo "====test leve 1===="
	@echo "> 只是能run <"
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
