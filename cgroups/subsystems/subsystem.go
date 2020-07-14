package subsystems

// ResourceConfig 用于传递资源限制配置的结构体，包含内存限制，CPU时间片权重，CPU核心数
type ResourceConfig struct {
	MemoryLimit string
	CpuShare    string
	CpuSet      string
}

// Subsystem 接口， 每个SU币system可以实现以下四个接口
type Subsystem interface {
	// 返回subsystem的名字，比如cpu memory。
	Name() string
	// 设置某粉cgroup在这个Subsystem中的资源限制
	Set(path string, res *ResourceConfig) error
	// 将进程添加到某个cgroup中
	Apply(path string, pid int) error
	// 移除某个cgroup
	Remove(path string) error
}

// 通过不同的subsystem初始化实例创建资源限制处理链数组
var (
	SubsystemsIns = []Subsystem{
		&CpusetSubSystem{},
		&MemorySubSystem{},
		&CpuSubSystem{},
	}
)
