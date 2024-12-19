// Code generated by protoc-gen-vkit. DO NOT EDIT.
// versions:
// - protoc-gen-vkit v1.0.0

package handler

import (
	"github.com/vison888/go-vkit/grpcx"
	
)

func GetList() []interface{} {
	list := make([]interface{}, 0)
	list = append(list, &LogService{})
	
	return list
}

func GetApiEndpoint() []*grpcx.ApiEndpoint {
	return []*grpcx.ApiEndpoint{
		{
			Method:"LogService.One",
			Url:"/api/collector/log.add", 
			ClientStream:false, 
			ServerStream:false,
		},{
			Method:"LogService.Batch",
			Url:"/api/collector/log.addBatch", 
			ClientStream:false, 
			ServerStream:false,
		},
	}
}
