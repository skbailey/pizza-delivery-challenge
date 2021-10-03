.PHONY: docker.build docker.build.run docker.run run test test.coverage

input_file = directions/pizza_delivery_input.txt
ifdef input
	input_file = $(input)
endif

deliverer_count = 1
ifdef deliverers
	deliverer_count = $(deliverers)
endif

run:
	@go run main.go --file-input $(input_file) --deliverer-count $(deliverer_count)

test:
	@go test -v ./...

test.coverage:
	@go test -race -covermode=atomic -coverprofile=coverage.out ./...

image_name = delivery/pizza
ifdef image
	image_name = $(image)
endif

docker.build:
	@docker build -t $(image_name):latest .

docker.build.run: docker.build
	@docker run --rm -it --name pizza-bae $(image_name):latest go run main.go --file-input $(input_file) --deliverer-count $(deliverer_count)

docker.run:
	@docker run --rm -it --name pizza-bae $(image_name):latest go run main.go --file-input $(input_file) --deliverer-count $(deliverer_count)
