package check

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"regexp"
	"net/url"
	"strings"
	"net"
	"math/rand"
)

func Get_req(url string, servers *string, headers *string, bodys *string) {
	client := &http.Client{}
    req,_ := http.NewRequest("GET",url,nil)
    req.Header.Add("User_Agent", get_random_ua()) //随机UA，消除发包特征
	r,_ := client.Do(req)
	*servers = r.Header.Get("Server")
	defer func() { _ = r.Body.Close() }()
	a, _ := ioutil.ReadAll(r.Body)
	*bodys = string(a)
	dataType, _ := json.Marshal(r.Header)
	*headers = string(dataType)
}

func Get_server(value string, servers string) (result bool) {
	result = false
	if servers == value {
		result = true
	}
	return result
}
func Get_title(value string, bodys string) (result bool) {
	result = false
	compileRegex := regexp.MustCompile("<title>(.*?)</title>")
	matchArr := compileRegex.FindStringSubmatch(bodys)
	if matchArr != nil {
		if matchArr[1] == value {
			result = true
		}
	}
	return result
}

func Get_body(value string, bodys string) (result bool) {
	result = strings.Contains(bodys, value)
	return result
}

func Get_banner(value string, bodys string) (result bool) {
	result = false
	compileRegex := regexp.MustCompile("(?im)<\\s*banner.*>(.*?)<\\s*/\\s*banner>")
	matchArr := compileRegex.FindAllStringSubmatch(bodys, -1)
	var i int
	for i = 0; i < len(matchArr); i++ {
		if value == matchArr[i][1] {
			result = true
		}
	}
	return result
}

func Get_port(value string, urls string) (result bool) {

	u, _ := url.Parse(urls)
	result = false
	host := u.Host
	address := net.ParseIP(host)
	if address == nil {
		result = false
	} else {
		ho := strings.Split(host, ":")
		if len(ho) > 1 {
			port := ho[1]
			if port == value {
				result = true
			}
		}
	}
	return result
}

func Get_head(value string, headers string) (result bool) {
	what_headers, _ := json.Marshal(headers)
	child_string := string(what_headers)
	what_headers = []byte(strings.Replace(child_string, "\"", "", -1))
	what_headers = []byte(strings.Replace(string(what_headers), "\\", "", -1))
	what_headers = []byte(strings.Replace(string(what_headers), "[", "", -1))
	what_headers = []byte(strings.Replace(string(what_headers), "]", "", -1))
	result = strings.Contains(string(what_headers), value)
	return result
}

func Check_model(match string, value string, url string, bodys string, heades string, servers string) (result bool) {
	switch {
	case match == "body_contains":
		result = Get_body(value, bodys)
	case match == "protocol_contains":
		result = false
	case match == "title_contains":
		result = Get_title(value, bodys)
	case match == "banner_contains":
		result = Get_banner(value, bodys)
	case match == "header_contains":
		result = Get_head(value, heades)
	case match == "port_contains":
		result = Get_port(value, url)
	case match == "server":
		result = Get_server(value, servers)
	case match == "title":
		result = Get_title(value, bodys)
	case match == "cert_contains":
		result = false
	case match == "server_contains":
		result = Get_server(value, servers)
	case match == "protocol":
		result = false
	default:
		result = false
	}
	return result
}

func get_random_ua() string {
	USER_AGENTS := []string{"Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; SV1; AcooBrowser; .NET CLR 1.1.4322; .NET CLR 2.0.50727)",
		"Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 6.0; Acoo Browser; SLCC1; .NET CLR 2.0.50727; Media Center PC 5.0; .NET CLR 3.0.04506)",
		"Mozilla/4.0 (compatible; MSIE 7.0; AOL 9.5; AOLBuild 4337.35; Windows NT 5.1; .NET CLR 1.1.4322; .NET CLR 2.0.50727)",
		"Mozilla/5.0 (Windows; U; MSIE 9.0; Windows NT 9.0; en-US)",
		"Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.1; Win64; x64; Trident/5.0; .NET CLR 3.5.30729; .NET CLR 3.0.30729; .NET CLR 2.0.50727; Media Center PC 6.0)",
		"Mozilla/5.0 (compatible; MSIE 8.0; Windows NT 6.0; Trident/4.0; WOW64; Trident/4.0; SLCC2; .NET CLR 2.0.50727; .NET CLR 3.5.30729; .NET CLR 3.0.30729; .NET CLR 1.0.3705; .NET CLR 1.1.4322)",
		"Mozilla/4.0 (compatible; MSIE 7.0b; Windows NT 5.2; .NET CLR 1.1.4322; .NET CLR 2.0.50727; InfoPath.2; .NET CLR 3.0.04506.30)",
		"Mozilla/5.0 (Windows; U; Windows NT 5.1; zh-CN) AppleWebKit/523.15 (KHTML, like Gecko, Safari/419.3) Arora/0.3 (Change: 287 c9dfb30)",
		"Mozilla/5.0 (X11; U; Linux; en-US) AppleWebKit/527+ (KHTML, like Gecko, Safari/419.3) Arora/0.6",
		"Mozilla/5.0 (Windows; U; Windows NT 5.1; en-US; rv:1.8.1.2pre) Gecko/20070215 K-Ninja/2.1.1",
		"Mozilla/5.0 (Windows; U; Windows NT 5.1; zh-CN; rv:1.9) Gecko/20080705 Firefox/3.0 Kapiko/3.0",
		"Mozilla/5.0 (X11; Linux i686; U;) Gecko/20070322 Kazehakase/0.4.5",
		"Mozilla/5.0 (X11; U; Linux i686; en-US; rv:1.9.0.8) Gecko Fedora/1.9.0.8-1.fc10 Kazehakase/0.5.6",
		"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/535.11 (KHTML, like Gecko) Chrome/17.0.963.56 Safari/535.11",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_7_3) AppleWebKit/535.20 (KHTML, like Gecko) Chrome/19.0.1036.7 Safari/535.20",
		"Opera/9.80 (Macintosh; Intel Mac OS X 10.6.8; U; fr) Presto/2.9.168 Version/11.52"}
	length := len(USER_AGENTS)
	index := rand.Intn(length)
	return USER_AGENTS[index]

}
