本程序用于从源代码生成`plugin`中的导出结构体的接口(`interface`)，以便用于类型断言，只支持对只使用基本类型的`method`的包装。
支持对使用基本类型的函数和方法进行包装。

use this program convert export structs to interface, in order to type assert. 

Only wrap funcs and methods use base type.

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
