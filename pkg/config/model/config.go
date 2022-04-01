package configmodel

type Fields struct {
	Server    Server    `yaml:"server"`
	RPCServer RPCServer `yaml:"rpc-server"`
}

type Server struct {
	Port int `yaml:"port"`
}

type RPCServer struct {
	Port int `yaml:"port"`
}
