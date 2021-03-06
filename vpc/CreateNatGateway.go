// 2018-01-09 09:23:20.675573 -0800 PST m=+82.001390284
// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots

package vpc

import (
	"encoding/json"
	"github.com/tencentcloudplatform/tcpicli/core"
)

type CreateNatGatewayResp struct {
	BillID   string `json:"billId"`
	Code     int64  `json:"code"`
	CodeDesc string `json:"codeDesc"`
	Message  string `json:"message"`
}

// Implement https://cloud.tencent.com/document/api/215/4094
func (c *VpcClient) CreateNatGateway(options ...string) (*CreateNatGatewayResp, error) {
	resp, err := c.DoAction("CreateNatGateway", options...)
	if err != nil {
		return nil, err
	}
	var s CreateNatGatewayResp
	err = json.Unmarshal(resp, &s)
	return &s, err
}

func CreateNatGateway(options ...string) (*CreateNatGatewayResp, error) {
	return DefaultClient.CreateNatGateway(options...)
}

func (r *CreateNatGatewayResp) String(args ...interface{}) (string, error) {
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
