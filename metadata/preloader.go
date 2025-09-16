package metadata

type Preloader struct {
	Var    string `toml:"var"`
	Method string `toml:"method"`
	Params any    `toml:"params"`
}
