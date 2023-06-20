ifeq ($(OS),Windows_NT)
	EXE_EXT := .exe
else
	EXE_EXT := 
endif

build-admin:
	cd cmd/admin; go build -ldflags "-s -w" -o ../../bin/cprotect-admin$(EXE_EXT)

build-activator:	
	cd cmd/activator; go build -ldflags "-s -w -X 'main.Product=$(CPROTECT_PRODUCT)' -X 'main.Password=$(CPROTECT_PASSWORD)'" -o ../../bin/$(CPROTECT_PRODUCT)-activator-cli$(EXE_EXT)

build-activator-gui:
	cd cmd/activator-gui; go build -ldflags "-s -w -X 'main.Product=$(CPROTECT_PRODUCT)' -X 'main.Password=$(CPROTECT_PASSWORD)'" -o ../../bin/$(CPROTECT_PRODUCT)-activator$(EXE_EXT)

build-activator-gui-windows:
	cd cmd/activator-gui; GOOS=windows go build -ldflags "-s -w -X 'main.Product=$(CPROTECT_PRODUCT)' -X 'main.Password=$(CPROTECT_PASSWORD)' -H=windowsgui" -o ../../bin/$(CPROTECT_PRODUCT)-activator.exe

release:
	git tag v0.2.0
	git push origin HEAD