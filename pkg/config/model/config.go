package configmodel

type Fields struct {
	Server Server `yaml:"server"`
}

type Server struct {
	Port string `yaml:"port"`
}
