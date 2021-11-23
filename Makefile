generate:
	find ./proto -name "*.proto" | xargs protoc \
		--go_out=. \
		--go_opt=module=github.com/WebEngineeringGroupI/genproto-go \
		--go-grpc_out=. \
		--go-grpc_opt=module=github.com/WebEngineeringGroupI/genproto-go

update-submodules:
	git submodule update --init --recursive --remote
