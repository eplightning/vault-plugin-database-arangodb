package arangodb

import (
	"encoding/json"
	"fmt"
	"github.com/hashicorp/vault/sdk/database/dbplugin/v5"
	"github.com/mitchellh/mapstructure"
)

type PluginConfig struct {
	Endpoints string `mapstructure:"endpoints"`
	Username  string `mapstructure:"username"`
	Password  string `mapstructure:"password"`
	Insecure  bool   `mapstructure:"insecure_tls"`
}

type Permission struct {
	Database   string `json:"database"`
	Collection string `json:"collection"`
	Grant      string `json:"grant"`
}

func configFromRaw(raw map[string]interface{}) (*PluginConfig, error) {
	config := &PluginConfig{}

	if err := mapstructure.Decode(raw, config); err != nil {
		return nil, err
	}

	if config.Endpoints == "" {
		return nil, fmt.Errorf("endpoints are required")
	}

	if config.Username == "" || config.Password == "" {
		return nil, fmt.Errorf("username and password are required")
	}

	return config, nil
}

func permissionsFromStatements(statements dbplugin.Statements) ([]Permission, error) {
	permissions := make([]Permission, 0, len(statements.Commands))

	for _, statement := range statements.Commands {
		permission := Permission{}

		if err := json.Unmarshal([]byte(statement), &permission); err != nil {
			return nil, err
		}

		if permission.Database == "" {
			return nil, fmt.Errorf("database field in permission statements cannot be empty")
		}

		if permission.Grant != "rw" && permission.Grant != "ro" && permission.Grant != "none" {
			return nil, fmt.Errorf("grant field in permission statements must be one of: rw, ro, none")
		}

		permissions = append(permissions, permission)
	}

	return permissions, nil
}
