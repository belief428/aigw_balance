package persist

type IGateway interface {
	// GetCode 获取编号
	GetCode() string
	// GetName 获取名称
	GetName() string
	// GetBuildCount 获取水平平衡调控数
	GetBuildCount() int
	// GetHouseCount 获取垂直平衡调控数
	GetHouseCount() int
}

type IArchive interface {
	// GetName 获取名称
	GetName() string

	// GetBuild 获取建筑附加信息
	GetBuild() IArchiveBuild

	// GetRetTemp 获取回温
	GetRetTemp() float32
}

type IArchiveBuild interface {
	// GetArea 获取面积
	GetArea() float32
}
