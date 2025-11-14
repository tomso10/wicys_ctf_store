package structs

type Item struct {
	Name        string `mapstructure:"name"`
	Description string `mapstructure:"description"`
	Price       int    `mapstructure:"price"`
	Image       string `mapstructure:"image"`
}
