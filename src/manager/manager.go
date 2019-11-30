package manager

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"sephiroth/utils"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type FM struct {
	di       *utils.Di
	logger   *utils.Logger
	db       *mgo.Collection
	timeutil *utils.TimeUtil
	total    int
	hidden   bool
}

const (
	DISPLAY_HIDDEN_FILE_DEFAULT = false
	TOTAL_DEFAULT               = 0
	APPLIED_DEFAULT             = false
)

func NewFm() *FM {
	fm := new(FM)
	fm.di = utils.NewDi()
	fm.timeutil = utils.NewTimeUtil()
	fm.logger = utils.NewLogger()
	fm.db = fm.di.GetMongoDB().DB("local").C("files")
	fm.hidden = DISPLAY_HIDDEN_FILE_DEFAULT
	return fm
}

func (t *FM) regex(pattern, str string) []string {
	reg, _ := regexp.Compile(pattern)
	return reg.FindAllString(str, 10)
}

func (fm *FM) SetHidden(flag bool) {
	fm.hidden = flag
}

func (fm *FM) Rename(old, new string) {
	os.Rename(old, new)
	fm.logger.Info("Rename " + old + " to " + new)
}

func (fm *FM) Remove(path string) {
	os.Remove(path)
	fm.logger.Info(" Remove " + path)
}

/**
 * 扫描某文件夹的信息
 */
func (fm *FM) Scan(filepath string) {
	fm.total = TOTAL_DEFAULT
	// 将文件夹读入数据库
	fm.SaveFileInfos(filepath)
	fm.logger.Info(fmt.Sprintf(" Reading files success, total : %d ", fm.total))
}

func (fm *FM) SaveFileInfos(path string) {
	files := fm.ReadFileFromPath(path)
	fm.SaveFileInfo(files)
}

/**
 * 读入某路径下的所有文件
 */
func (fm *FM) ReadFileFromPath(path string) []*FileInfo {
	fsp := make([]*FileInfo, 0)
	paths := make([]string, 0)
	err := filepath.Walk(path,
		func(path string, f os.FileInfo, err error) error {
			if f == nil {
				return err
			}
			if f.IsDir() {
				return nil
			}
			if !fm.hidden && len(fm.regex("/\\.", path)) > 0 {
				return nil
			}
			paths = append(paths, path)
			return nil
		})
	if err != nil {
		fm.logger.Error(fmt.Sprintf("filepath.Walk() returned %v \n", err))
	}
	for _, path := range paths {
		fsp = append(fsp, fm.GetFileMetaInfo(path))
	}
	return fsp
}

/**
 * 获取文件信息
 */
func (fm *FM) GetFileMetaInfo(path string) *FileInfo {
	info, _ := os.Stat(path)
	data, _ := ioutil.ReadFile(path)
	o := NewFileInfo()
	o.Md5 = fmt.Sprintf("%x", md5.Sum(data))                 // md5
	o.ModTime = info.ModTime().Format("2006-01-02 15:04:05") // 修改时间
	o.IsDir = info.IsDir()                                   // 是否目录
	o.Mode = fmt.Sprintf("%s", info.Mode())                  // 文件权限
	o.Name = info.Name()                                     // 文件名
	o.Size = float64(info.Size()) / 1024                     // 文件大小
	o.Applied = APPLIED_DEFAULT                              // 是否被更新过
	o.Path = path                                            // 文件路径
	o.NewPath = path
	o.UpTime = fm.timeutil.GetTime()
	return o
}

/**
 * 保存文件信息
 */
func (fm *FM) SaveFileInfo(files []*FileInfo) {
	for _, f := range files {
		// 检查文件是否重复, 重复则更新,否则插入
		var dbfiles []FileInfo
		condition := bson.M{"md5": f.Md5, "path": f.Path}
		fm.db.Find(condition).All(&dbfiles)
		if len(dbfiles) > 0 {
			for _, item := range dbfiles {
				f.UpTime = fm.timeutil.GetTime()
				fm.db.Update(bson.M{"path": item.Path}, f)
			}
		} else {
			fm.db.Insert(f)
			fm.total++
			fm.logger.Info(" file: " + f.Path + " done")
		}
	}
}

/**
 * 对某文件夹进行更新
 */
func (fm *FM) Apply(filepath string) {
	// 查找文件路径下的旧文件信息
	var files []FileInfo
	condition := bson.M{"path": bson.M{"$regex": filepath}}
	fm.db.Find(condition).All(&files)
	for _, a := range files {
		if a.Path != a.NewPath {
			fm.Rename(a.Path, a.NewPath)
		}
	}
	// 清空旧数据库文件信息
	fm.ClearPath(filepath)
	// 重新导入
	fm.Scan(filepath)
	//fm.ClearAll()
}

func (fm *FM) ClearPath(path string) {
	fm.db.RemoveAll(bson.M{"path": bson.M{"$regex": path}})
}

/**
 * 清理所有文件
 */
func (fm *FM) ClearAll() {
	fm.db.RemoveAll(bson.M{})
}
