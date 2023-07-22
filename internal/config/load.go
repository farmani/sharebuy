package config

import (
	"fmt"
	"log"
	"strings"

	"github.com/davecgh/go-spew/spew"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/structs"
	"github.com/knadh/koanf/v2"
)

const (
	delimiter = "."
	separator = ","

	envPrefix = "SHAREBUY"
	tagName   = "koanf"

	upTemplate     = "====================== Loaded Configuration ======================"
	bottomTemplate = "=================================================================="
)

func Load(print bool) *Config {
	k := koanf.New(delimiter)

	if err := k.Load(structs.Provider(Default(), "koanf"), nil); err != nil {
		log.Fatalf("error loading default: %s", err)
	}

	if err := LoadEnv(k); err != nil {
		log.Fatalf("error loading environment variables: %v", err)
	}

	config := Config{}
	var tag = koanf.UnmarshalConf{Tag: tagName}
	if err := k.UnmarshalWithConf("", &config, tag); err != nil {
		log.Fatalf("error unmarshalling config: %v", err)
	}

	if print {
		log.Printf("%s\n%v\n%s\n", upTemplate, spew.Sdump(config), bottomTemplate)
	}

	return &config
}

func LoadEnv(k *koanf.Koanf) error {
	callback := func(source string) string {
		base := strings.ToLower(strings.TrimPrefix(source, envPrefix))
		return strings.ReplaceAll(base, separator, delimiter)
	}

	if err := k.Load(env.Provider(envPrefix, delimiter, callback), nil); err != nil {
		return fmt.Errorf("error loading environment variables: %s", err)
	}

	return nil
}
