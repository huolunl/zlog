/*
   @Author:huolun
   @Date:2021/7/21
   @Description
*/
package zlog

import (
	"testing"
)

func TestNewZLoggerProduction(t *testing.T) {
	loggerProduction := NewZLogger(false, false)
	loggerProduction.Debug("debug")
	loggerProduction.Info("info")
	loggerProduction.Warn("warn")
	loggerProduction.Error("error")

	loggerDevelop := NewZLogger(true, false)
	loggerDevelop.Debug("debug")
	loggerDevelop.Info("info")
	loggerDevelop.Warn("warn")
	loggerDevelop.Error("error")

}
