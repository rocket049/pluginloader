本程序用于从源代码生成`plugin`中的导出结构体的接口(`interface`)，以便用于类型断言。

use this program convert export structs to interface, in order to type assert.

安装(install)：

go get github.com/rocket049/pluginloader/cmd/pluginwrap

用法（usage）：

pluginwrap path/to/plugin/dir

生成的文件(generate)：

dirWrap.go
