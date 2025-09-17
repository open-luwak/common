package metadata

type Preloader struct {
	Var    string `toml:"var"`
	Method string `toml:"method"`
	Param  any    `toml:"param"`
}
