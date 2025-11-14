package structs

type Token struct {
	Value string `mapstructure:"value"`
	Type  string `mapstructure:"type"`
}
