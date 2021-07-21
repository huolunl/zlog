/*
   @Author:huolun
   @Date:2021/7/20
   @Description
*/
package zlog

import "github.com/gofrs/uuid"

func GetUUID() string {
	uid, err := uuid.NewV4()
	if err != nil {
		return "make uuid error"
	}
	return uid.String()
}
