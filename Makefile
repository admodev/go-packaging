BINARY_NAME=packaging

.PHONY: build install clean

build:
	go build -o bin/$(BINARY_NAME) ./cmd/$(BINARY_NAME)

install: build
	@echo "Installing $(BINARY_NAME)..."
	@if [ "$(OS)" = "Windows_NT" ]; then \
		mkdir -p $(USERPROFILE)\\bin && \
		cp bin\\$(BINARY_NAME).exe $(USERPROFILE)\\bin\\$(BINARY_NAME).exe && \
		echo "Installed to $(USERPROFILE)\\bin\\$(BINARY_NAME).exe"; \
	else \
		sudo cp bin/$(BINARY_NAME) /usr/local/bin/$(BINARY_NAME) && \
		echo "Installed to /usr/local/bin/$(BINARY_NAME)"; \
	fi

clean:
	rm -rf bin
