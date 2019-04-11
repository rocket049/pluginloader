### pluginloader
用于简化调用 plugin 内的函数。 Help developer call func in go plugin easy.

### 内容

```
type PluginLoader struct {
	...
}
func (p *PluginLoader) Call(funcName string, p0 ...interface{}) (interface{}, error)
```

### 用法（usage）:

```
import "github.com/rocket049/pluginloader"

p, err := pluginloader.NewPluginLoader( "path_to_plugin" )
if err != nil {
	panic(err)
}
res, err := p.Call("NameOfFunc", p0,p1,p3,...)

```

### 注意(attention)：

被调用的函数可以无返回值，也可以返回1个返回值，或者1个返回值跟1个`error`。

被 `Call` 的函数返回值格式只能是： `ResType` 或 `(ResType, error)` -- `ResType`可以是任何类型。

当只有0个或1个返回值时，`Call` 返回的 `error` 始终是 `nil`。

The func called can have no return value, or return 1 value or 1 value and 1 error。

The format MUST be: `ResType` OR `(ResType, error)` -- `ResType` can be any type.

If there is 0 or 1 return value,the `error` return from `Call` is always `nil`.

### 命令行工具 pluginwrap
本程序用于从源代码生成`plugin`中的可导出结构体的接口(`interface`)，以便用于类型断言。

use this program convert export structs to interface, in order to type assert.

安装(install)：

`go get github.com/rocket049/pluginloader/cmd/pluginwrap`

用法（usage）：

`pluginwrap path/to/plugin/dir`

生成的文件(generate)：

`dirWrap.go`