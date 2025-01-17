/*
Copyright 2021 The terraform-docs Authors.

Licensed under the MIT license (the "License"); you may not
use this file except in compliance with the License.

You may obtain a copy of the License at the LICENSE file in
the root directory of this source tree.
*/

package format

import (
	"bytes"
	"strings"

	tomlsdk "github.com/BurntSushi/toml"

	"github.com/terraform-docs/terraform-docs/print"
	"github.com/terraform-docs/terraform-docs/terraform"
)

// toml represents TOML format.
type toml struct {
	*print.Generator

	config *print.Config
}

// NewTOML returns new instance of TOML.
func NewTOML(config *print.Config) Type {
	return &toml{
		Generator: print.NewGenerator("toml", config.ModuleRoot),
		config:    config,
	}
}

// Generate a Terraform module as toml.
func (t *toml) Generate(module *terraform.Module) error {
	copy := copySections(t.config, module)

	buffer := new(bytes.Buffer)
	encoder := tomlsdk.NewEncoder(buffer)

	if err := encoder.Encode(copy); err != nil {
		return err
	}

	t.Generator.Funcs(print.WithContent(strings.TrimSuffix(buffer.String(), "\n")))

	return nil

}

func init() {
	register(map[string]initializerFn{
		"toml": NewTOML,
	})
}
