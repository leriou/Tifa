package manager

type FileInfo struct {
	ModTime string
	IsDir   bool
	Name    string
	Mode    string
	Size    int64
	Md5     string
	Path    string
	NewPath string
	Applied bool
}

func NewFileInfo() *FileInfo {
	return new(FileInfo)
}
