package structs

type Account struct {
	Name     string `mapstructure:"name"`
	Team     string `mapstructure:"team"`
	Password string `mapstructure:"password"`
	Balance  int    `mapstructure:"balance"`
}
