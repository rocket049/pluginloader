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

被 `Call` 的函数返回值格式只能是： `(ResType, error)`

The func called must only have the return format: `(ResType, error)`
