package structs

type TokenInfo struct {
	KothUnredeemed    int `mapstructure:"koth_unredeemed"`
	KothRedeemed      int `mapstructure:"koth_redeemed"`
	PurpleUnredeemed  int `mapstructure:"purple_unredeemed"`
	PurpleRedeemed    int `mapstructure:"purple_redeemed"`
	SponsorUnredeemed int `mapstructure:"sponsor_unredeemed"`
	SponsorRedeemed   int `mapstructure:"sponsor_redeemed"`
	CTFUnredeemed     int `mapstructure:"ctf_unredeemed"`
	CTFRedeemed       int `mapstructure:"ctf_redeemed"`
}
