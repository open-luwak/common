package metadata

type ConditionalRole struct {
	Software string `toml:"software"`
	Role     string `toml:"role"`
	Check    string `toml:"check"`
}
