package configmodel

type Server struct {
	HTTP HTTPServer `yaml:"http"`
	RPC  RPCServer  `yaml:"rpc"`
}

type HTTPServer struct {
	Port int `yaml:"port"`
}

type RPCServer struct {
	Port int `yaml:"port"`
}
