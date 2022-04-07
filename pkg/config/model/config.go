package configmodel

// Fields Provide all config fields
type Fields struct {
	Server    Server    `yaml:"server"`
	RPCServer RPCServer `yaml:"rpc-server"`
}

// Server HTTPServer Config
type Server struct {
	Port int `yaml:"port"`
}

// RPCServer gRPCServer Config
type RPCServer struct {
	Port int `yaml:"port"`
}
