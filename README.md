# Mac文件管理

通过将某目录中的文件路径和文件元属性信息进行汇总

使用程序批量修改文件系统的文件名字等

目前打算支持`文件删除/移动/重命名`

# 接口设计


|方法|功能||
|:-:|:-:|:-:|
|scan|扫描一个新的文件夹,将所有文件清单都存入数据库||
|apply|应用一批文件修改数据||

# 待办事项
1. 是否添加隐藏文件 (.开头)
2.  
