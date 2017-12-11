package config

// Target represents a single build target
type Target struct {
	Commands []string `toml:"commands"`
}
