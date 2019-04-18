本程序用于从源代码生成`plugin`中的导出结构体的接口(`interface`)，以便用于类型断言，以及对函数和方法进行包装。
函数、方法的返回值、参数类型不能是用户自定义类型。
pluginwrap几乎已经完美了！
现在除了用户自定义类型，已经可以使用所有被导入库的类型。

*如果必须使用自定义类型，有两种办法：*

1. 请使用 `pluginloader.Call` 或 `pluginloader.CallValue` 调用。
2. 把需要导出的复杂类型做成第三方库（`import "your/package"`），不要在`plugin`源代码中定义。

Use this program convert export structs to interface, in order to type assert. 
Wrap funcs and methods.
pluginwrap is almost perfect now! 
It can deal with all types import from packages, except User Defined types in main.

*If you must export self defined type, please put the code in a package can be imported by others.*



安装(install)：

go get github.com/rocket049/pluginloader/cmd/pluginwrap

用法（usage）：

pluginwrap path/to/plugin/dir

生成的文件(generate)：

dirWrap.go

### 调用
#### 使用对象(`struct`)

```
	p, err := pluginloader.NewPluginLoader("foo.so")
	if err != nil {
		panic(err)
	}
	iface, err := p.Call("NewFoo")
	if err != nil {
		panic(err)
	}
	foo := iface.(IFoo)
	// use foo.Method
```

#### 使用`func`

```
	p, err := pluginloader.NewPluginLoader("foo.so")
	if err != nil {
		panic(err)
	}
	
	// MUST call InitxxxFuncs(p) before call funcs, xxx = plugin名字
	InitfooFuncs(p)
	// use funcs in plugin foo
```
