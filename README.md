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

被调用的函数只能返回1个返回值或者1个返回值跟1个`error`。

被 `Call` 的函数返回值格式只能是： `ResType` 或 `(ResType, error)` -- `ResType`可以是任何类型。

当只有一个返回值时，`Call` 返回的 `error` 始终是 `nil`。

The func called must only have 1 or 2 return values, the format MUST be: `ResType` OR `(ResType, error)` -- `ResType` can be any type.

If there is only 1 return value,the `error` return from `Call` is always `nil`.
