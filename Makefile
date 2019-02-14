dev:
	@echo "====test leve 1===="
	@echo "> 只是能run <"
	@go run main.go test/test1
	@echo ""
	@echo ""
	@echo "====test leve 2===="
	@echo "> 支持自定义func < "
	@go run main.go test/test2
	@echo ""
	@echo ""
