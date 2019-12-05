package goplugin

import(
	"errors"
	"reflect"
	"strings"
)

var IncorrectExtensionPath = errors.New("Extension path incorrect.")
var IncorrectInterface = errors.New("Incorrect interface.")
var UndefinedManager = errors.New("Undefined plugin manager.")
var UndefinedExtension = errors.New("Undefined plugin extension.")

// 插件的描述和实现，由插件方实现
type Plugin struct {
	// 必要属性，指定实现的扩展点
	// 格式：管理器Id/扩展点Id
	// 例子：helloworld/get_word
	Extension string

	// 必要属性。插件实现扩展点的接口
	Implement interface{}

	// 可选属性。插件名称。可用于动态管理
	Name string

	// 可选属性。插件描述。可用于动态管理
	Description string

	// 可选属性。插件版本。可用于动态管理
	Version string

	// 可选属性。插件开发者。可用于动态管理
	Author string

	// 可选属性。插件版权。可用于动态管理
	Copyright string
}

func NewPlugin(
	extension string,
	implement interface{},
	name string,
	description string,
	version string,
	author string,
	copyright string) *Plugin {

	return &Plugin{
		Name: name,
		Description: description,
		Version: version,
		Author: author,
		Copyright: copyright,
		Extension: extension}
}

// 扩展点。由系统方设计的可以扩展的环节
type Extension struct {
	// 必要属性。扩展点的Id。管理器内唯一
	Id string

	// 可选属性。扩展点接口的类型，可以通过reflect.TypeOf((*Foo)(nil)).Elem()获取。用于插件绑定时做类型检查
	// 如果没提供NilImplement，则插件在绑定期不做类型检查，只能在运行时检查
	InterfaceType reflect.Type

	// 可选属性。扩展点的名称。可用于动态管理
	Name string

	// 可选属性。扩展点的用途描述。可用于动态管理
	Description string

	// 插件。一个扩展点可以支持多个插件接入
	plugins []*Plugin
}

func NewExtension(
	id string,
	interfaceType reflect.Type,
	name string,
	description string) *Extension {

	return &Extension{
		Id: id,
		InterfaceType: interfaceType,
		Name: name,
		Description: description}
}

// 插件管理器。
type Manager struct {
	// 必要属性。插件管理器的Id，插件注册的时候需要指定
	// 如果管理器的Id出现重复，则后面注册的管理器会取代前面注册的管理器
	Id string
	implements map[string][]interface{} // 插件实现，方便系统调用
	extensions map[string]*Extension // 扩展点
}

// 管理全局的插件管理器
var managers = make(map[string]*Manager)

// 新建一个插件管理器
func NewManager(Id string) (manager *Manager) {
	manager = &Manager{
		Id: Id,
		implements: make(map[string][]interface{}),
		extensions: make(map[string]*Extension)}

	managers[Id] = manager

	return
}

// 增加扩展点
func (this *Manager) AddExtension(extension *Extension) {
	this.extensions[extension.Id] = extension
}

// 获得插件实现
func (this *Manager) GetImplements(extensionId string) []interface{} {
	// 获得实现并返回
	implements, _ := this.implements[extensionId]
	return implements
}

// 注册插件
func Register(plugin *Plugin) (err error) {
	// 分割Extension字符串，得到managerId和extensionId
	pieces := strings.Split(plugin.Extension, "/")
	if len(pieces) != 2 {
		return IncorrectExtensionPath
	}

	managerId := pieces[0]
	extensionId := pieces[1]

	// 获取插件管理器
	manager, ok := managers[managerId]
	if !ok {
		return UndefinedManager
	}

	// 获取扩展点
	extension, ok := manager.extensions[extensionId]
	if !ok {
		return UndefinedExtension
	}

	// 如果扩展点设置了接口类型，则对插件提供的实现做类型检查
	if extension.InterfaceType != nil {
		pluginType := reflect.TypeOf(plugin.Implement)
		if !pluginType.Implements(extension.InterfaceType) {
			return IncorrectInterface
		}
	}

	// 初始化扩展点的插件数组
	if extension.plugins == nil {
		extension.plugins = make([]*Plugin, 1)
		extension.plugins[0] = plugin

		manager.implements[extensionId] = make([]interface{}, 1)
		manager.implements[extensionId][0] = plugin.Implement
	} else {
		// 把插件和实现放到插件数组
		extension.plugins = append(extension.plugins, plugin)
		manager.implements[extensionId] = append(manager.implements[extensionId], plugin.Implement)
	}

	return nil
}
