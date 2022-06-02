build-admin:
	cd cmd/admin; go build -ldflags "-s -w" -o ../../bin/cprotect-admin.exe

build-activator:
	cd cmd/activator; go build -ldflags "-s -w -X main.Product=Geezr-Go -X 'main.Password=$(CProtectPass)'" -o ../../bin/cprotect-activator.exe