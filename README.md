### pluginloader
用于简化调用 plugin 内的函数。 Help developer call func in go plugin easy.

### ChangeLog
20190412 add new method: CallValue, 
被调用的函数可以返回任意数量的返回值，返回值形式为: ([]reflect.Value,error)。
CallValue Allow any number of return values,return type: ([]reflect.Value,error)

20190413 add new method: MakeFunc，
用函数名从`plugin`中构建`func`。
make func from plugin by name.


### 内容

```
type PluginLoader struct {
	...
}

///Call return type must be: (res,error)
func (p *PluginLoader) Call(funcName string, p0 ...interface{}) (interface{}, error)

//CallValue Allow any number of return values,return type: []reflect.Value,error
func (p *PluginLoader) CallValue(funcName string, p0 ...interface{}) ([]reflect.Value, error)

//MakeFunc point a func ptr to plugin
func (s *PluginLoader) MakeFunc(fptr interface{}, name string) error 
```

### 用法（usage）:

```
import "github.com/rocket049/pluginloader"

p, err := pluginloader.NewPluginLoader( "path_to_plugin" )
if err != nil {
	panic(err)
}

res, err := p.Call("NameOfFunc", p0,p1,p3,...)
//...

ret := p.CallValue("NameOfFunc", p0,p1,p3,...)
//...

var Foo func(arg string)(string,error)
p.MakeFunc(&Foo,"Foo")
//call Foo(something)

```

### 注意(attention)：

#### Call
被调用的函数可以无返回值，也可以返回1个返回值，或者1个返回值跟1个`error`。

被 `Call` 的函数返回值格式只能是： `ResType` 或 `(ResType, error)` -- `ResType`可以是任何类型。

当只有0个或1个返回值时，`Call` 返回的 `error` 始终是 `nil`。

The func called can have no return value, or return 1 value or 1 value and 1 error。

The format MUST be: `ResType` OR `(ResType, error)` -- `ResType` can be any type.

If there is 0 or 1 return value,the `error` return from `Call` is always `nil`.

#### CallValue
被调用的函数可以返回任意数量的返回值，返回值形式为: ([]reflect.Value,error)

The func called can return any number values, return type: ([]reflect.Value,error)

### 命令行工具 pluginwrap
本程序用于从源代码生成`plugin`中的可导出结构体的接口(`interface`)，以便用于类型断言。

use this program convert export structs to interface, in order to type assert.

安装(install)：

`go get github.com/rocket049/pluginloader/cmd/pluginwrap`

用法（usage）：

`pluginwrap path/to/plugin/dir`

生成的文件(generate)：

`dirWrap.go`