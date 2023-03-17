APP = "xx-gpt"

b:
	@echo $(APP)
	@rm $(APP)
	@go build -o $(APP)

# build and run
br:	b
	./$(APP)
