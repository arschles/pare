package config

// File represents the top level config file
type File struct {
	Targets map[string]*Target `toml:"targets"`
}
