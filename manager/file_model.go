package manager

type FileInfo struct {
	ModTime string
	IsDir   bool
	Name    string
	Mode    string
	Size    float64
	Md5     string
	Path    string
	NewPath string
	Applied bool
	UpTime  string
}

func NewFileInfo() *FileInfo {
	return new(FileInfo)
}
