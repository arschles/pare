package config

// Command represents a command to execute
type Command struct {
	Exec      string `toml:"exec"`
	Crash     bool   `toml:"crash"`
	Directory string `toml:"directory"`
}
