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
	peerUin := "peerUin=100000655192"
	region := "Region=bj"
	zoneId := "subnetSet.0.zoneId=800001"
	vpcNameInit := "vpcName=tcpiclivpcinit"
	vpcName0 := "vpcName=tcpiclivpc"
	vpcName1 := "vpcName=pcxonly"
	cidrBlock0 := "cidrBlock=10.0.0.0/16"
	cidrBlock1 := "cidrBlock=192.168.0.0/16"
	subnetCidr0 := "subnetSet.0.cidrBlock=10.0.0.0/24"
	subnetCidr1 := "subnetSet.0.cidrBlock=192.168.0.0/24"
	subnetName0 := "subnetSet.0.subnetName=tcpclisub0"
	subnetName1 := "subnetSet.0.subnetName=tcpclisub1"
	routeTableName := "routeTableName=tcpiclirt"
	destinationCidrBlock0 := subnetCidr1
	nextHub0 := ""
	nextType0 := "4"
	peeringConnectionName := "peeringConnectionName=tcpiclipcx"
	// -- for modifying subnet attributes --
	subnetNameMod := "subnetName=tcpclisubmod0"
	// -- for adding and deleting routes in Subnet API actions --
	destinationCidrBlockAdd := "routeSet.0.destinationCidrBlock=172.217.6.46/32"
	nextTypeAdd := "routeSet.0.nextType=" + nextType0
	// -- for adding and deleting NACL stuff --
	networkAclName := "networkAclName=tcpiclinacl"
	networkAclNameMod := networkAclName + "new"
	ruleDirectionIn := "ruleDirection=1"
	entrySetDesc := "networkAclEntrySet.0.desc=apitest"
	entrySetIpProtocol := "networkAclEntrySet.0.ipProtocol=all"
	entrySetCidrIp := "networkAclEntrySet.0.cidrIp=0.0.0.0/0"
	entrySetPortRange := "networkAclEntrySet.0.portRange=ALL"
	entrySetAction := "networkAclEntrySet.0.action=1"
	// -- vpn gw stuff --
	period := "period=1"
	createPeriod := "period=0"
	vpnBandwidth := "bandwidth=5"
	vpnGwName := "vpnGwNamename=apitest"
	vpnGwNameMod := "vpnGwName=apitestmod"
	userGwName := "userGwName=apitest"
	userGwAddr := "userGwAddr=183.30.0.1"
	userGwNameMod := "userGwName=apitestmod"
	// -- direct connect gw stuff --
	directConnectGatewayName := "directConnectGatewayName=apitest"
	directConnectGatewayNameMod := "directConnectGatewayName=apitestmod"
	directConnectGatewayType := "type=1" // 0 = non-NAT, 1 = NAT
	localIPTranslationOriginalIP := "localIPTranslation.0.originalIP=10.0.0.8"
	localIPTranslationTranslationIP := "localIPTranslation.0.translationIP=10.10.0.8"
	localIPTranslationOriginalIPOld := "oldOriginalIP=10.0.0.8"
	localIPTranslationTranslationIPOld := "oldTranslationIP=10.10.0.8"
	localIPTranslationOriginalIPNew := "originalIP=10.0.0.11"
	localIPTranslationTranslationIPNew := "translationIP=10.11.0.8"
	// -- NAT gw stuff --
	maxConcurrent := "maxConcurrent=1000"
	natName := "natName=apitest"
	natNameNew := "natName=apitestnew"
	autoAllocEipNum := "autoAllocEipNum=1"
	// -- ENI stuff --
	eniName := "eniName=apitest"
	eniDescription := "eniDescription=createdbyapi" // just do this instead of changing the name like for everything else
	eniPrimaryIp := "privateIpAddressSet.0.primary=true"
	privateIpAddress := "privateIpAddressSet.0.privateIpAddress=10.0.0.10"

	gen := &autogen.Gen{
		DocRoot: "https://cloud.tencent.com/document/api/",
		Seq: []string{
			// -- don't work or need to engage vpc team to figure out api --
			// "AssociateVip",
			// "AttachClassicLinkVpc",
			// "DescribeVpcLimit",

			"CreateVpc",
			"DO tcpicli vpc CreateVpc " + region + " " + vpcName1 + " " + cidrBlock1,
			`SET vpcId0=tcpicli -f "{{range .Data}}{{.UnVpcID}}{{end}}" vpc DescribeVpcEx ` + region + " " + vpcName0,
			"DescribeVpcEx",
			`SET vpcId1=tcpicli -f "{{range .Data}}{{.UnVpcID}}{{end}}" vpc DescribeVpcEx ` + region + " " + vpcName1,
			"DescribeVpcClassicLink",
			"ModifyVpcAttribute",
			"CreateSubnet",
			"DO tcpicli vpc CreateSubnet vpcId=$vpcId1 " + region + " " + subnetCidr1 + " " + subnetName1 + " " + zoneId,
			"DescribeSubnetEx",
			`SET subnetId0=tcpicli -f "{{range .Data}}{{.UnSubnetID}}{{end}}" vpc DescribeSubnetEx vpcId=$vpcId0 ` + region,
			"DescribeSubnet",
			`SET subnetId1=tcpicli -f "{{range .Data}}{{.UnSubnetID}}{{end}}" vpc DescribeSubnetEx vpcId=$vpcId1 ` + region,
			"CreateVpcPeeringConnection",
			"DescribeVpcPeeringConnections",
			`SET peeringConnectionId=tcpicli -f "{{range .Data}}{{.peeringConnectionId}}{{end}}" vpc DescribeVpcPeeringConnections ` + region + " " + peeringConnectionName,
			"CreateRouteTable",
			"DescribeRouteTable",
			`SET routeTableId=tcpicli -f "{{range .Data}}{{.UnRouteTableID}}{{end}}" vpc DescribeRouteTable vpcId=$vpcId0 ` + region + " " + routeTableName,
			"ModifyRouteTableAttribute",
			"CreateRoute",
			"AssociateRouteTable",
			"CreateNetworkAcl",
			"DescribeNetworkAcl",
			`SET networkAclId=tcpicli -f "{{range .Data}}{{.NetworkAclID}}{{end}}" vpc DescribeNetworkAcl vpcId=$vpcId0 ` + region + " " + networkAclName,
			"ModifyNetworkAcl",
			"ModifyNetworkAclEntry",
			"CreateSubnetAclRule",
			"DeteleSubnetAclRule",
			"ModifySubnetAttribute",
			"InquiryVpnPrice",
			// "CreateVpn", -- API action only supports prepaid 1 month minimum purchase. --
			// "DescribeVpnGw",
			// `SET vpnGwId=tcpicli -f "{{range .Data}}{{.unVpnGwId}}{{end}}" vpc DescribeVpnGw ` + region + " " + vpnGwName,
			// "ModifyVpnGw",
			"DescribeUserGwVendor",
			"AddUserGw",
			"DescribeUserGw",
			`SET userGwId=tcpicli -f "{{range .Data}}{{.UserGwID}}{{end}}" vpc DescribeUserGw vpcId=$vpcId0 ` + region + " " + userGwName,
			"ModifyUserGw",
			"CreateDirectConnectGateway",
			"DescribeDirectConnectGateway",
			`SET directConnectGatewayId=tcpicli -f "{{range .Data}}{{.DirectConnectGatewayID}}{{end}}" vpc DescribeDirectConnectGateway vpcId=$vpcId0 ` + region + " " + directConnectGatewayName,
			"ModifyDirectConnectGateway",
			"CreateLocalIPTranslationNatRule",
			"DescribeLocalIPTranslationNatRule",
			"ModifyLocalIPTranslationNatRule",
			"CreateLocalIPTranslationNatRule",
			"InquiryNatPrice",
			"CreateNatGateway",
			"DescribeNatGateway",
			`SET natId=tcpicli -f "{{range .Data}}{{.NatID}}{{end}}" vpc DescribeNatGateway ` + region + " " + natName,
			`SET originalEipSet=tcpicli -f "{{range .Data}}{{.EipSet}}{{end}}" vpc DescribeNatGateway ` + region + " " + natName + " | tr -d '[]'",
			"ModifyNatGateway",
			"EipBindNatGateway",
			// "EipUnBindNatGateway",
			"CreateNetworkInterface",
			"DescribeNetworkInterfaces",
			`SET networkInterfaceId=tcpcli -f "{{range .Data.Data}}{{.networkInterfaceId}}{{end}}" vpc DescribeNetworkInterfaces vpcId=$vpcId0 ` + region + " " + eniName,
			"ModifyNetworkInterface",
			// -- clean everything up --
			"DeleteNetworkInterface",
			"DeleteNatGateway",
			"DeleteLocalIPTranslationNatRule",
			"DeleteDirectConnectGateway",
			"DeleteUserGw",
			"DeleteNetworkAcl",
			"DeleteRoute",
			// -- RT doesn't delete if bound to subnet. No API action to dissociate RT from subnet. RT gets deleted when subnet it is bound to gets deleted. --
			"DeleteRouteTable",
			"DeleteVpcPeeringConnection",
			// -- Bound RT gets deleted here. Failure on 'DeleteRouteTable' expected. If regen to 'DeleteRouteTable.go' needed, comment out 'AssociateRouteTable' above  --
			"DeleteSubnet",
			"DeleteVpc",
			// -- clean up manual stuff --
			"DO tcpicli vpc DeleteSubnet vpcId=$vpcId1 subnetId=$subnetId1 " + region,
			"DO tcpicli vpc DeleteVpc vpcId=$vpcId1 " + region,
		},
		FuncMap: map[string][]string{
			"CreateVpc": []string{"215/1309",
				region,
				vpcNameInit,
				cidrBlock0,
			},
			"DescribeVpcEx": []string{"215/1372",
				region,
			},
			"DescribeVpcClassicLink": []string{"215/2112",
				region,
			},
			"DescribeVpcLimit": []string{"215/1844",
				region,
			},
			"ModifyVpcAttribute": []string{"215/1310",
				region,
				"vpcId=$vpcId0",
				vpcName0,
			},
			"AssociateVip": []string{"215/1361",
				region,
				"vpcId=$vpcId0",
			},
			"AttachClassicLinkVpc": []string{"215/2098",
				region,
				"vpcId=$vpcId0",
			},
			"DeleteVpc": []string{"215/1307",
				region,
				"vpcId=$vpcId0",
			},
			"CreateSubnet": []string{"215/1314",
				region,
				"vpcId=$vpcId0",
				subnetCidr0,
				subnetName0,
				zoneId,
			},
			"DescribeSubnetEx": []string{"215/1371",
				region,
				"vpcId=$vpcId0",
			},
			"DescribeSubnet": []string{"215/1311",
				region,
				"vpcId=$vpcId0",
				"subnetId=$subnetId0",
			},
			"CreateVpcPeeringConnection": []string{"215/2107",
				region,
				peeringConnectionName,
				"vpcId=$vpcId0",
				"peerVpcId=$vpcId1",
				peerUin,
			},
			"DescribeVpcPeeringConnections": []string{"215/2101"},
			"DeleteVpcPeeringConnection": []string{"215/2104",
				region,
				"peeringConnectionId=$peeringConnectionId",
			},
			"CreateRouteTable": []string{"215/1419",
				region,
				"vpcId=$vpcId0",
				"subnetId=$subnetId0",
				routeTableName,
				destinationCidrBlock0,
				nextType0,
				nextHub0,
			},
			"DescribeRouteTable": []string{"215/1420",
				region,
				"vpcId=$vpcId0",
				routeTableName,
			},
			"ModifyRouteTableAttribute": []string{"215/1417",
				region,
				"vpcId=$vpcId0",
				"routeTableId=$routeTableId",
			},
			"CreateRoute": []string{"215/5688",
				region,
				destinationCidrBlockAdd,
				nextTypeAdd,
				"routeSet.0.nextHub=$peeringConnectionId",
				"vpcId=$vpcId0",
				"routeTableId=$routeTableId",
			},
			"AssociateRouteTable": []string{"215/1416",
				region,
				"vpcId=$vpcId0",
				"routeTableId=$routeTableId",
				"subnetId=$subnetId0",
			},
			"InquiryVpnPrice": []string{"215/5104",
				region,
				period,
				vpnBandwidth,
			},
			"CreateVpn": []string{"215/5106",
				region,
				createPeriod,
				vpnBandwidth,
				vpnGwName,
				"vpcId=$vpcId0",
			},
			"DescribeVpnGw": []string{"215/5108",
				region,
			},
			"ModifyVpnGw": []string{"215/5107",
				region,
				vpnGwNameMod,
				"vpcId=$vpcId0",
				"vpnGwId=$vpnGwId",
			},
			"DescribeUserGwVendor": []string{"215/5120",
				region,
			},
			"DescribeUserGw": []string{"215/5119",
				region,
			},
			"ModifyUserGw": []string{"215/5118",
				region,
				userGwNameMod,
				"userGwId=$userGwId",
			},
			"CreateDirectConnectGateway": []string{"215/4824",
				region,
				directConnectGatewayName,
				directConnectGatewayType,
				"vpcId=$vpcId0",
			},
			"DescribeDirectConnectGateway": []string{"215/4827",
				region,
				directConnectGatewayName,
				"vpcId=$vpcId0",
			},
			"ModifyDirectConnectGateway": []string{"215/4826",
				region,
				directConnectGatewayNameMod,
				"vpcId=$vpcId0",
				"directConnectGatewayId=$directConnectGatewayId",
			},
			"CreateLocalIPTranslationNatRule": []string{"215/5185",
				region,
				localIPTranslationOriginalIP,
				localIPTranslationTranslationIP,
				"vpcId=$vpcId0",
				"directConnectGatewayId=$directConnectGatewayId",
			},
			"DescribeLocalIPTranslationNatRule": []string{"215/5188",
				region,
				"vpcId=$vpcId0",
				"directConnectGatewayId=$directConnectGatewayId",
			},
			"ModifyLocalIPTranslationNatRule": []string{"215/5187",
				region,
				"vpcId=$vpcId0",
				"directConnectGatewayId=$directConnectGatewayId",
				localIPTranslationOriginalIPOld,
				localIPTranslationTranslationIPOld,
				localIPTranslationOriginalIPNew,
				localIPTranslationTranslationIPNew,
			},
			"DeleteLocalIPTranslationNatRule": []string{"215/5186",
				region,
				localIPTranslationOriginalIP,
				localIPTranslationTranslationIP,
				"vpcId=$vpcId0",
				"directConnectGatewayId=$directConnectGatewayId",
			},
			"InquiryNatPrice": []string{"215/4091",
				region,
				maxConcurrent,
			},
			"CreateNatGateway": []string{"215/4094",
				region,
				maxConcurrent,
				natName,
				autoAllocEipNum,
				"vpcId=$vpcId0",
			},
			"DescribeNatGateway": []string{"215/4088",
				region,
				natName,
				"vpcId=$vpcId0",
			},
			"ModifyNatGateway": []string{"215/4086",
				region,
				natNameNew,
				"vpcId=$vpcId0",
				"natId=$natId",
			},
			"EipBindNatGateway": []string{"215/4093",
				region,
				autoAllocEipNum,
				"vpcId=$vpcId0",
				"natId=$natId",
			},
			"EipUnBindNatGateway": []string{"215/4093",
				region,
				"vpcId=$vpcId0",
				"natId=$natId",
				"assignedEipSet=$originalEipSet",
			},
			"CreateNetworkInterface": []string{"215/4811",
				region,
				eniName,
				eniPrimaryIp,
				privateIpAddress,
				"vpcId=$vpcId0",
				"subnetId=$subnetId0",
			},
			"DescribeNetworkInterfaces": []string{"215/4814",
				region,
				eniName,
			},
			"ModifyNetworkInterface": []string{"215/5372",
				region,
				eniDescription,
				"networkInterfaceId=$networkInterfaceId",
				"vpcId=$vpcId0",
			},
			"DeleteNetworkInterface": []string{"215/4813",
				region,
				"networkInterfaceId=$networkInterfaceId",
				"vpcId=$vpcId0",
			},
			"DeleteNatGateway": []string{"215/4087",
				region,
				"vpcId=$vpcId0",
				"natId=$natId",
			},
			"DeleteDirectConnectGateway": []string{"215/4825",
				region,
				"vpcId=$vpcId0",
				"directConnectGatewayId=$directConnectGatewayId",
			},
			"DeleteUserGw": []string{"215/5117",
				region,
				"userGwId=$userGwId",
			},
			"AddUserGw": []string{"215/5116",
				region,
				userGwName,
				userGwAddr,
			},
			"CreateNetworkAcl": []string{"215/1437",
				region,
				networkAclName,
				"vpcId=$vpcId0",
			},
			"DescribeNetworkAcl": []string{"215/1441",
				region,
				"vpcId=$vpcId0",
			},
			"ModifyNetworkAcl": []string{"215/1443",
				region,
				networkAclNameMod,
				"vpcId=$vpcId0",
				"networkAclId=$networkAclId",
			},
			"ModifyNetworkAclEntry": []string{"215/1444",
				region,
				ruleDirectionIn,
				"vpcId=$vpcId0",
				"networkAclId=$networkAclId",
				entrySetDesc,
				entrySetIpProtocol,
				entrySetCidrIp,
				entrySetPortRange,
				entrySetAction,
			},
			"CreateSubnetAclRule": []string{"215/1438",
				region,
				"vpcId=$vpcId0",
				"subnetIds.0=$subnetId0",
				"networkAclId=$networkAclId",
			},
			"DeteleSubnetAclRule": []string{"215/1442",
				region,
				"vpcId=$vpcId0",
				"subnetIds.0=$subnetId0",
				"networkAclId=$networkAclId",
			},
			"ModifySubnetAttribute": []string{"215/1313",
				region,
				subnetNameMod,
				"vpcId=$vpcId0",
				"subnetId=$subnetId0",
			},
			"DeleteNetworkAcl": []string{"215/1439",
				region,
				"vpcId=$vpcId0",
				"networkAclId=$networkAclId",
			},
			"DeleteRoute": []string{"215/5689",
				region,
				destinationCidrBlockAdd,
				nextTypeAdd,
				"routeSet.0.nextHub=$peeringConnectionId",
				"vpcId=$vpcId0",
				"routeTableId=$routeTableId",
			},
			"DeleteRouteTable": []string{"215/1418",
				region,
				"vpcId=$vpcId0",
				"routeTableId=$routeTableId",
			},
			"DeleteSubnet": []string{"215/1312",
				region,
				"vpcId=$vpcId0",
				"subnetId=$subnetId0",
			},
		},
		PkgName: "vpc",
		Pkg:     new(Pkg),
	}

	gen.Run()
}
