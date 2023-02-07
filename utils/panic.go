// Package utils 日志相关组件包
package utils

import (
	"fmt"
	"github.com/qingconglaixueit/wechatbot/pkg/logger"
	"runtime/debug"
	"strings"
)

// DumpPanicStack 进程崩溃重启，并捕获 panic 输出
func DumpPanicStack(err error) {
	if e := recover(); e != nil {
		logger.Info("[Panic] err:%s, stack:%s", err.Error(), GetStackInfo())
	}
}

// GetStackInfo 获取调用栈信息
func GetStackInfo() string {
	return strings.Replace(fmt.Sprintf("%s", debug.Stack()), "\t", "    ", -1)
}
