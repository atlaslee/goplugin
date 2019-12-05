# GoPlugin
GoPlugin is an interface based plugin framework written in golang.

# Usage for plugin framework developer：

	package helloworld

	import ("github.com/atlaslee/goplugin")

	// Define the extension interface yourself
	type MyExtension interface {
		GetWord() string
	}

	func init() {
		manager := goplugin.NewManager("hello_world")
		plubinManager.AddExtension(goplugin.Extension{Id:"get_word", MyExtension})
	}

	// Use plugins for your system
	func HelloWorld(manager *goplugin.Manager) {
		implements := manager.GetImplements("get_word")
		getWord, _ := implements[0].(MyExtension)
		print(getWord.GetWord())
	}

# Usage for plugin developer：

	package helloplugin

	import ("github.com/atlaslee/goplugin")

	// Implement the extension interface
	type MyPlugin struct {}

	func (this *Myplugin) GetWord() string {
		return "Hello world"
	}

	// Regist your plugin
	func init() {
		// Extension format：ManagerId/ExtensionId
		goplugin.Register(&goplugin.Plugin{Extension: "hello_world/get_word", Implement: &MyImplement{}})
	}

# Install a plugin by goplugin script：

	~/helloworld$ goplugin get -p helloworld github.com/atlaslee/helloplugin

# Install a plugin manually：

	~/helloworld$ cp -rf ~/helloplugin src/
	~/helloworld$ echo 'package helloworld' > plugin.go
	~/helloworld$ echo 'import ("github.com/atlaslee/helloplugin")' >> plugin.go
