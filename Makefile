.PHONY: run test

input_file = directions/pizza_delivery_input.txt
ifdef input
	input_file = input
endif

deliverer_count = 1
ifdef deliverers
	deliverer_count = $(deliverers)
endif

run:
	@go run main.go -file-input $(input_file) -deliverer-count $(deliverer_count)

test:
	@go test -v ./...
