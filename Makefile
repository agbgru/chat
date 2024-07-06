docker_image_for_generating = docker_image_for_generating
pwd = $(PWD)

lint_docker: ## Run linter in the docker.
	docker run --rm -v $(CURDIR):/app -w /app --env GOTAGS=testdocker golangci/golangci-lint:v1.58.0 golangci-lint run -v --timeout 4m30s