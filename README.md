# ArangoDB database plugin for Vault

Database plugin for HashiCorp Vault providing support for ArangoDB secret management.

Due to usage of v5 database plugin interface, Vault >= 1.6.0 is required.

Heavily inspired by https://github.com/hashicorp/vault-plugin-database-elasticsearch

## Features

- Static / dynamic roles
- Root credential rotation support
- Adjustable database / collection permissions per role
- Using official ArangoDB Go client library - both single server & cluster setups are fully supported

## How to build

```make```

## Usage

NOTE: Vault plugins need to be registered before use. This happens automatically when Vault is running in dev mode.

```
vault secrets enable database

vault write database/config/arango \                         
    plugin_name="vault-plugin-database-arangodb" \
    allowed_roles="*" \                                      
    endpoints="http://localhost:8529" \
    username="root" \
    password="password"

vault write database/roles/my \
    db_name=arango \
    creation_statements="{\"database\": \"test\", \"grant\": \"rw\"}"

vault write database/roles/my2 \
    db_name=arango \
    creation_statements="{\"database\": \"test\", \"collection\": \"col\", \"grant\": \"ro\"}"

vault read database/creds/my
vault read database/creds/my2
```

## TODO

- Tests
- Proper build system with support for more platforms
