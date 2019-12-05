# goplugin
goplugin是一个GO语言的接口式插件框架

# Manager使用方法：
`
package helloworld

import ("github.com/atlaslee/goplugin")

// 设计插件接口
type MyExtension interface {
	GetWord() string
}

// 初始化插件管理器
func init() {
	// 创建插件管理器
	manager := goplugin.NewManager("hello_world")

	// 创建扩展点
	plubinManager.AddExtension(goplugin.Extension{Id:"get_word", MyExtension})
}

// 通过插件实现功能
func HelloWorld(manager *goplugin.Manager) {
	implements := manager.GetImplements("get_word")
	getWord, _ := implements[0].(MyExtension)
	print(getWord.GetWord())
}
`

# Plugin使用方法：
`
package helloplugin

import ("github.com/atlaslee/goplugin")

// 实现插件接口
type MyPlugin struct {}

func (this *Myplugin) GetWord() string {
	return "Hello world"
}

// 注册插件
func init() {
	// Extension格式：管理器Id/扩展点Id
	goplugin.Register(&goplugin.Plugin{Extension: "hello_world/get_word", Implement: &MyImplement{}})
}

# 插件安装方法：
`
~/helloworld$ goplugin get -p helloworld github.com/atlaslee/helloplugin
`

# 手动安装插件：

`
~/helloworld$ cp -rf ~/helloplugin src/
~/helloworld$ echo -e 'package helloworld\nimport("helloplugin")\n' > plugin.go
`
