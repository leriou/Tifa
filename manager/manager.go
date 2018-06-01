package manager

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"lgocommon/utils"
)

type FM struct {
	di   *utils.Di
	tool *utils.Tool
	total int
}

func NewFm() *FM {
	fm := new(FM)
	fm.di = utils.NewDi()
	fm.tool = utils.NewTool()
	return fm
}

func (fm *FM) Read(path string) {
	err := filepath.Walk(path,
		func(path string, f os.FileInfo, err error) error {
			if f == nil {
				return err
			}
			if f.IsDir() {
				return nil
			}
			fInfo := NewFileInfo()
			fm.GetFileMetaInfo(path, fInfo)
			fm.di.GetMongoDB().DB("local").C("files").Insert(fInfo)
			fm.total++
			fm.tool.Logging("INFO", " file :"+path+" done")
			return nil
		})
	if err != nil {
		fmt.Printf("filepath.Walk() returned %v \n", err)
	}
}

func (fm *FM) GetFileMetaInfo(path string, finfo *FileInfo) bool {
	info, _ := os.Stat(path)
	data, _ := ioutil.ReadFile(path)
	finfo.Md5 = fmt.Sprintf("%x", md5.Sum(data))                 // md5
	finfo.ModTime = info.ModTime().Format("2006-01-02 15:04:05") // 修改时间
	finfo.IsDir = info.IsDir()                                   // 是否目录
	finfo.Mode = fmt.Sprintf("%s", info.Mode())                  // 文件权限
	finfo.Name = info.Name()                                     // 文件名
	finfo.Size = info.Size() / 1024                              // 文件大小
	finfo.Applied = false
	finfo.Path = path
	finfo.NewPath = path
	return true
}

func (fm *FM) Rename(old, new string) {
	os.Rename(old, new)
	fm.tool.Logging("INFO"," Rename "+ old +" to "+ new )
}

func (fm *FM) Remove(path string) {
	os.Remove(path)
	fm.tool.Logging("INFO"," Remove "+path)
}

func (fm *FM) Scan(filepath string) {
	fm.total = 0
	// 将文件夹读入数据库 
	fm.Read(filepath)
	fm.tool.Logging("INFO", fmt.Sprintf(" Reading files success, total: %d ",fm.total))
}

func (fm *FM) Apply(filepath string) {
	// 查找文件路径下的文件信息
	// db := fm.di.GetMongoDB().DB("local").C("files")

	// fmt.Println(db.Find())

	// 按照新的文件进行移动和更名处理
}
