package config

// Target represents a single build target
type Target struct {
	Commands map[string]*Command `toml:"commands"`
}
