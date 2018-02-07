// +build ignore

package main

import (
	. "."
	"github.com/tencentcloudplatform/tcpicli/autogen"
)

type Pkg struct{}

func (p *Pkg) DoAction(action string, query ...string) ([]byte, error) {
	return DoAction(action, query...)
}

func main() {
	region := "Region=gz"
	queueName := "queueName=tcpicligen"
	topicName := "topicName=tcpicligen"

	gen := &autogen.Gen{
		DocRoot: "https://cloud.tencent.com/document/api/",
		Seq: []string{
			"CreateQueue",
			"ListQueue",
			"GetQueueAttributes",
			"SetQueueAttributes",
			"SendMessage",
			`DO sleep 3`,
			"ReceiveMessage",
			`DO tcpicli cmq SendMessage msgBody="tcpicligen2" ` + region + " " + queueName,
			`DO sleep 3`,
			`SET receiptHandle=tcpicli -f "{{.ReceiptHandle}}" cmq ReceiveMessage ` + region + " " + queueName,
			`DO echo $receiptHandle`,
			"DeleteMessage",
			"BatchSendMessage",
			`DO sleep 3`,
			"BatchReceiveMessage",
			`DO sleep 3`,
			`SET batchMessage1=tcpicli -f '{{range .MsgInfoList}}{{if eq .MsgBody "test 3"}}{{.ReceiptHandle}}{{end}}{{end}}' cmq BatchReceiveMessage numOfMsg=2 ` + region + " " + queueName,
			`DO echo $batchMessage1`,
			`DO sleep 3`,
			`SET batchMessage2=tcpicli -f '{{range .MsgInfoList}}{{if eq .MsgBody "test 5"}}{{.ReceiptHandle}}{{end}}{{end}}' cmq BatchReceiveMessage numOfMsg=2 ` + region + " " + queueName,
			`DO echo $batchMessage2`,
			"BatchDeleteMessage",
			"CreateTopic",
			"ListTopic",
			"GetTopicAttributes",
			"SetTopicAttributes",
			"Subscribe",
			"PublishMessage",
			"ListSubscriptionByTopic",
			"GetSubscriptionAttributes",
			"SetSubscriptionAttributes",
			"Unsubscribe",
			"DeleteTopic",
			"DeleteQueue",
		},
		FuncMap: map[string][]string{
			"CreateQueue": []string{"406/5832",
				region,
				queueName,
			},
			"ListQueue": []string{"406/5833",
				region,
			},
			"GetQueueAttributes": []string{"406/5834",
				region,
				queueName,
			},
			"SetQueueAttributes": []string{"406/5835",
				region,
				queueName,
				"pollingWaitSeconds=5", // used to arbitrarily change queue attributes for function gen
			},
			"SendMessage": []string{"406/5837",
				region,
				queueName,
				"msgBody=tcpicli gen test",
			},
			"ReceiveMessage": []string{"406/5839",
				region,
				queueName,
			},
			"DeleteMessage": []string{"406/5840",
				region,
				queueName,
				"receiptHandle=$receiptHandle",
			},
			"BatchSendMessage": []string{"406/5838",
				region,
				queueName,
				"msgBody.1=test 1",
				"msgBody.2=test 2",
				"msgBody.3=test 3",
				"msgBody.4=test 4",
				"msgBody.5=test 5",
				"msgBody.6=test 6",
			},
			"BatchReceiveMessage": []string{"406/5924",
				region,
				queueName,
				"numOfMsg=2",
			},
			"BatchDeleteMessage": []string{"406/5841",
				region,
				queueName,
				"receiptHandle.1=$batchMessage1",
				"receiptHandle.2=$batchMessage2",
			},
			"DeleteQueue": []string{"406/5836",
				region,
				queueName,
			},
			"CreateTopic": []string{"406/7405",
				region,
				topicName,
			},
			"ListTopic": []string{"406/7407",
				region,
			},
			"GetTopicAttributes": []string{"406/7408",
				region,
				topicName,
			},
			"SetTopicAttributes": []string{"406/7406",
				region,
				topicName,
				"maxMsgSize=65535",
			},
			"PublishMessage": []string{"406/7411",
				region,
				topicName,
				"msgBody=tcpicli test message",
			},
			"BatchPublishMessage": []string{"406/7411",
				region,
				topicName,
				"msgBody.1=tcpicli batch test message 1",
				"msgBody.2=tcpicli batch test message 2",
			},
			"Subscribe": []string{"406/7414",
				region,
				topicName,
				"subscriptionName=tcpiclitestsub",
				"protocol=queue",
				"endpoint=tcpicligen",
			},
			"ListSubscriptionByTopic": []string{"406/7415",
				region,
				topicName,
			},
			"GetSubscriptionAttributes": []string{"406/7418",
				region,
				topicName,
				"subscriptionName=tcpiclitestsub",
			},
			"SetSubscriptionAttributes": []string{"406/7416",
				region,
				topicName,
				"subscriptionName=tcpiclitestsub",
				"filterTag.0=test",
			},
			"Unsubscribe": []string{"406/7417",
				region,
				topicName,
				"subscriptionName=tcpiclitestsub",
			},
			"DeleteTopic": []string{"406/7409",
				region,
				topicName,
			},
		},
		PkgName: "cmq",
		Pkg:     new(Pkg),
	}
	gen.Run()
}
