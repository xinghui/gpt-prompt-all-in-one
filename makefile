APP = "xh-gpt"

b:
	@echo $(APP)
	@rm -f $(APP)
	@go build -o $(APP)

# build and run
br:	b
	./$(APP)
