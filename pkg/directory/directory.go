package directory

type Directory interface {
	GetProjectBasePath() string
}

type directory struct {
	BasePath string
}

func NewDirectory(path string) Directory {
	return &directory{
		BasePath: path,
	}
}

func (d *directory) GetProjectBasePath() string {
	return d.BasePath
}
