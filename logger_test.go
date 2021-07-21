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
	loggerProduction.Info("info")
	loggerProduction.Debug("debug")

	loggerDevelop := NewZLogger(true, false)
	loggerDevelop.Info("info")
	loggerDevelop.Debug("debug")
}
