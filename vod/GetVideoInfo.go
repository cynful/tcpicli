// 2017-12-01 13:20:52.728205 -0800 PST m=+3.676851475
// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots

package vod

import (
	"encoding/json"
	"github.com/tencentcloudplatform/tcpicli/core"
)

type GetVideoInfoResp struct {
	BasicInfo struct {
		ClassificationID   int64         `json:"classificationId"`
		ClassificationName string        `json:"classificationName"`
		CoverURL           string        `json:"coverUrl"`
		CreateTime         int64         `json:"createTime"`
		Description        string        `json:"description"`
		Duration           int64         `json:"duration"`
		ExpireTime         int64         `json:"expireTime"`
		Name               string        `json:"name"`
		PlayerID           int64         `json:"playerId"`
		Size               int64         `json:"size"`
		SourceVideoURL     string        `json:"sourceVideoUrl"`
		Status             string        `json:"status"`
		Tags               []interface{} `json:"tags"`
		Type               string        `json:"type"`
		UpdateTime         int64         `json:"updateTime"`
	} `json:"basicInfo"`
	Code     int64  `json:"code"`
	CodeDesc string `json:"codeDesc"`
	Message  string `json:"message"`
	MetaData struct {
		AudioStreamList []struct {
			Bitrate      int64  `json:"bitrate"`
			Codec        string `json:"codec"`
			SamplingRate int64  `json:"samplingRate"`
		} `json:"audioStreamList"`
		Bitrate         int64  `json:"bitrate"`
		Container       string `json:"container"`
		Duration        int64  `json:"duration"`
		Height          int64  `json:"height"`
		Md5             string `json:"md5"`
		Size            int64  `json:"size"`
		VideoStreamList []struct {
			Bitrate int64  `json:"bitrate"`
			Codec   string `json:"codec"`
			Fps     int64  `json:"fps"`
			Height  int64  `json:"height"`
			Width   int64  `json:"width"`
		} `json:"videoStreamList"`
		Width int64 `json:"width"`
	} `json:"metaData"`
	SampleSnapshotInfo struct {
		SampleSnapshotList []struct {
			Definition int64    `json:"definition"`
			ImageUrls  []string `json:"imageUrls"`
			Interval   int64    `json:"interval"`
			SampleType string   `json:"sampleType"`
		} `json:"sampleSnapshotList"`
	} `json:"sampleSnapshotInfo"`
	TranscodeInfo struct {
		TranscodeList []struct {
			AudioStreamList []struct {
				Bitrate      int64  `json:"bitrate"`
				Codec        string `json:"codec"`
				SamplingRate int64  `json:"samplingRate"`
			} `json:"audioStreamList"`
			Bitrate         int64  `json:"bitrate"`
			Container       string `json:"container"`
			Definition      int64  `json:"definition"`
			Duration        int64  `json:"duration"`
			Height          int64  `json:"height"`
			Md5             string `json:"md5"`
			Size            int64  `json:"size"`
			URL             string `json:"url"`
			VideoStreamList []struct {
				Bitrate int64  `json:"bitrate"`
				Codec   string `json:"codec"`
				Fps     int64  `json:"fps"`
				Height  int64  `json:"height"`
				Width   int64  `json:"width"`
			} `json:"videoStreamList"`
			Width int64 `json:"width"`
		} `json:"transcodeList"`
	} `json:"transcodeInfo"`
}

// Implement https://cloud.tencent.com/document/api/266/8586
func GetVideoInfo(options ...string) (*GetVideoInfoResp, error) {
	resp, err := DoAction("GetVideoInfo", options...)
	if err != nil {
		return nil, err
	}
	var s GetVideoInfoResp
	err = json.Unmarshal(resp, &s)
	return &s, err
}

func (r *GetVideoInfoResp) String(args ...interface{}) (string, error) {
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
