// 2017-11-17 13:23:41.625311 -0800 PST m=+34.038171982
// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots

package vpc

import (
	"encoding/json"
	"github.com/tencentcloudplatform/tcpicli/core"
)

type DescribeNetworkAclResp struct {
	Code     int64  `json:"code"`
	CodeDesc string `json:"codeDesc"`
	Data     []struct {
		CreateTime         string `json:"createTime"`
		NetworkAclEntrySet struct {
			Egress []struct {
				Action     int64  `json:"action"`
				CidrIP     string `json:"cidrIp"`
				Desc       string `json:"desc"`
				IPProtocol string `json:"ipProtocol"`
				PortRange  string `json:"portRange"`
			} `json:"egress"`
			Ingress []struct {
				Action     int64  `json:"action"`
				CidrIP     string `json:"cidrIp"`
				Desc       string `json:"desc"`
				IPProtocol string `json:"ipProtocol"`
				PortRange  string `json:"portRange"`
			} `json:"ingress"`
		} `json:"networkAclEntrySet"`
		NetworkAclID   string        `json:"networkAclId"`
		NetworkAclName string        `json:"networkAclName"`
		SubnetNum      int64         `json:"subnetNum"`
		SubnetSet      []interface{} `json:"subnetSet"`
		UnVpcID        string        `json:"unVpcId"`
		VpcCidrBlock   string        `json:"vpcCidrBlock"`
		VpcID          string        `json:"vpcId"`
		VpcName        string        `json:"vpcName"`
	} `json:"data"`
	Message    string `json:"message"`
	TotalCount int64  `json:"totalCount"`
}

// Implement https://cloud.tencent.com/document/api/215/1441
func DescribeNetworkAcl(options ...string) (*DescribeNetworkAclResp, error) {
	resp, err := DoAction("DescribeNetworkAcl", options...)
	if err != nil {
		return nil, err
	}
	var s DescribeNetworkAclResp
	err = json.Unmarshal(resp, &s)
	return &s, err
}

func (r *DescribeNetworkAclResp) String(args ...interface{}) (string, error) {
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
