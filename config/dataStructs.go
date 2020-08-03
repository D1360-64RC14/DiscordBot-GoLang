package config

// ConfigStruct :
// Struct do YAML de configuração.
type configStruct struct {
	DiscordBot struct {
		Token string
	} `yaml:"discord_bot"`

	YoutubeAPI struct {
		Token       string
		RequestURLs struct {
			SearchAPI   string `yaml:"search_api"`
			Video       string
			ChannelsAPI string `yaml:"channels_api"`
			Channels    string
		} `yaml:"requestURLs"`
	} `yaml:"youtube_api"`
}
