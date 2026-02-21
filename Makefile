.DEFAULT_GOAL := help

.PHONY: help app-proto app-go-start app-purge app-deploy-tag app-delete-tag

help:
	@echo "📋 Available commands (LOCAL DEVELOPMENT):"
	@echo ""
	@echo "  🚀 DESARROLLO:"
	@echo "  make app-start               	- Start app with go run (inicia PostgreSQL)"
	@echo "  make app-stop               	- Stop app with go run (para PostgreSQL)"
	@echo ""
	@echo "  🛠️  UTILIDADES:"
	@echo "  make app-proto            		- Generate protobuf code"
	@echo "  make app-purge            		- Remove artifacts and stop PostgreSQL"
	@echo ""
	@echo "  ☸️  KUBERNETES (deployment/):"
	@echo "  make app-deploy-tag <version>  - Create and push git tag (e.g., 1.2.3)"
	@echo "  make app-delete-tag <version>  - Delete git tag locally and remotely"
	@echo ""

app-proto:
	bash bin/app/proto.sh

app-start:
	bash bin/app/start.sh

app-stop:
	bash bin/app/stop.sh

app-purge:
	bash bin/app/purge.sh

app-deploy-tag:
	bash bin/app/deploy-tag.sh $(filter-out $@,$(MAKECMDGOALS))

app-delete-tag:
	bash bin/app/delete-tag.sh $(filter-out $@,$(MAKECMDGOALS))

%:
	@: