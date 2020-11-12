PROJECT		= "github.com/eplightning/vault-plugin-database-arangodb"
GOFILES		= $(shell find . -name "*.go")

default: vault-plugin-database-arangodb

vault-plugin-database-arangodb: $(GOFILES)
	CGO_ENABLED=0 go build -ldflags "-s -w" ./cmd/vault-plugin-database-arangodb

clean:
	rm -f vault-plugin-database-arangodb

.PHONY: default clean