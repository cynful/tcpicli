// 2017-11-06 07:13:23.655517341 +0000 UTC m=+14.032839287
// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots

package cmq

import (
	"encoding/json"
	"github.com/tencentcloudplatform/tcpicli/core"
)

type DeleteQueueResp struct {
	Code      int64  `json:"code"`
	Message   string `json:"message"`
	RequestID string `json:"requestId"`
}

// Implement https://cloud.tencent.com/document/api/406/5836
func DeleteQueue(options ...string) (*DeleteQueueResp, error) {
	resp, err := DoAction("DeleteQueue", options...)
	if err != nil {
		return nil, err
	}
	var s DeleteQueueResp
	err = json.Unmarshal(resp, &s)
	return &s, err
}

func (r *DeleteQueueResp) String(args ...interface{}) (string, error) {
	var b []byte
	var err error
	if len(args) == 0 {
		b, err = json.MarshalIndent(r, "", "  ")
	} else if len(args) == 1 {
		if val, ok := args[0].(string); ok {
			if len(val) == 0 {
				b, err = json.MarshalIndent(r, "", "  ")
			} else {
				b, err = core.FmtOutput(val, r)
			}
		}
	}
	return string(b), err
}
