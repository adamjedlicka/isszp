package controller

type Config struct {
	Secret string
}

var (
	cfg Config
)

func Configure(config Config) { cfg = config }
