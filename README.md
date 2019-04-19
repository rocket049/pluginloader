### pluginloader
用于简化调用 plugin 内的函数。 Help developer call func in go plugin easy.

### ChangeLog

20190419-2 update pluginwrap: delete InitxxxFuncs , add `func GetxxxFuncs(p *pluginloader.PluginLoader) *xxxFuncs` ,avoid namespace conflict.
修改 pluginwrap ，删除生成的文件中的 `InitxxxFuncs`，改为：`func GetxxxFuncs(p *pluginloader.PluginLoader) *xxxFuncs`，
函数都被包含在 `xxxFuncs`结构体中避免命名冲突。

20190419 新增用于调用未定义结构体的对象：`UnknownObject`。 can use undefined object.

20190418 pluginwrap is almost perfect now! pluginwrap几乎已经完美了！
现在除了用户自定义类型，已经可以使用所有被导入库的类型。
It can deal with all types import from packages, except User Defined types in main.

20190413 add new method: MakeFunc，
用函数名从`plugin`中构建`func`。
make func from plugin by name.

20190412 add new method: CallValue, 
被调用的函数可以返回任意数量的返回值，返回值形式为: ([]reflect.Value,error)。
CallValue Allow any number of return values,return type: ([]reflect.Value,error)


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


//20190419 new
//UnknownObject field 'V' MUST be valueof object type *struct{...}
//成员'V' 必须是结构体指针的 Value: *struct{...}
type UnknownObject struct {
	V reflect.Value
}
//NewUnknownObject parameter 'v' MUST be valueof object type *struct{...}
//参数'v' 必须是结构体指针的 Value: *struct{...}
func NewUnknownObject(v reflect.Value) *UnknownObject 

//Get 得到结构体成员的 Value
//get the value of a field
func (s *UnknownObject) Get(name string) reflect.Value

//Call 运行结构体的 method
//call the method of the struct
func (s *UnknownObject) Call(fn string, args ...interface{}) []reflect.Value
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

// Use UnknownObject. NewFoo return 'foo *Foo'
v, err := p.CallValue("NewFoo")
if err != nil {
	t.Fatal(err)
}
obj := NewUnknownObject(v[0])

id: = obj.Get("Id").Int()

err = obj.Call("Set", nil)

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
本程序用于从`plugin`源代码生成:

1. 可导出结构体的接口(`interface`)，以便用于类型断言。 convert export structs to interface, in order to type assert.
2. 包装可调用的函数。 Wrap funcs.

#### 安装(install)：

`go get github.com/rocket049/pluginloader/cmd/pluginwrap`

#### 用法（usage）：

`pluginwrap path/to/plugin/dir`

生成的文件(generate)：

`dirWrap.go`