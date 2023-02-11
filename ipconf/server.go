package ipconf

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/xiongshou/im/common/config"
	"github.com/xiongshou/im/ipconf/domain"
	"github.com/xiongshou/im/ipconf/source"
)

// RunMain 启动web容器
func RunMain(path string) {
	config.Init(path)
	source.Init() //数据源要优先启动， etcd，都会去监听eventchan
	domain.Init() // 初始化调度层

	s := server.Default(server.WithHostPorts(":6789"))
	s.GET("/ip/list", GetIpInfoList)
	s.Spin()
}
