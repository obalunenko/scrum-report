// Package assets handles generated static.
// To update views - add changes to template htmls.
package assets

import (
	"embed"
	"path/filepath"
)

const (
	dir = "templates"
)

// content holds our puzzles inputs content.
//
//go:embed templates/*
var content embed.FS

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	return content.ReadFile(filepath.Clean(
		filepath.Join(dir, name)))
}

// MustAsset loads and returns the asset for the given name.
// It panics if the asset could not be found or
// could not be loaded.
func MustAsset(name string) []byte {
	res, err := Asset(name)
	if err != nil {
		panic(err)
	}

	return res
}
