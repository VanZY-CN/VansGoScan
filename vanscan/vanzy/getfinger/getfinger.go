package getfinger

import (
	"go_scan/vanzy/check"
	jsoniter "github.com/json-iterator/go"
	"fmt"
	"os"
	"strconv"
	"time"
	"bufio"
	_ "embed"
 )
 
type Rules [][]struct {
	Match   string `json:"match"`
	Content string `json:"content"`
}

type Fofa_dic struct {
	RuleID         string `json:"rule_id"`
	Level          string `json:"level"`
	Softhard       string `json:"softhard"`
	Product        string `json:"product"`
	Company        string `json:"company"`
	Category       string `json:"category"`
	ParentCategory string `json:"parent_category"`
	Rules          Rules  `json:"rules"`
}

//go:embed fofa.json
var FoFaFingerData []byte

type FoFaFinger []Fofa_dic

var FoFa FoFaFinger


func Run(url string, Fi1e bool) {
	var bodys string
	var headers string
	var servers string
	check.Get_req(url, &servers, &headers, &bodys)
	var jsons = jsoniter.ConfigCompatibleWithStandardLibrary
	err := jsons.Unmarshal(FoFaFingerData, &FoFa)
	if err != nil {
		panic(err.Error())
	}
	unmarshelledConfigs := FoFa
	if Fi1e == false {
		fmt.Println(url + ":")
	}
	for _, configObj := range unmarshelledConfigs {
		i := 0
		for i = 0; i < len(configObj.Rules); i++ {
			if len(configObj.Rules[i]) >= 2 {
				arr := make([]bool, len(configObj.Rules))
				j := 0
				for j = 0; j < len(configObj.Rules[i]); j++ {
					arr = append(arr, check.Check_model(configObj.Rules[i][j].Match, configObj.Rules[i][j].Content, url, bodys, headers, servers))
				}
				if !(ContainsInSlice(arr, false)) {
					if ContainsInSlice(arr, true) && Fi1e == true{
						result_writefile(url, configObj.Product)
					}
					if ContainsInSlice(arr, true) && Fi1e == false {
						result_print(url, configObj.Product)
					}
				}

			} else {
				if check.Check_model(configObj.Rules[i][0].Match, configObj.Rules[i][0].Content, url, bodys, headers, servers) && Fi1e == true{
					result_writefile(url, configObj.Product)
				}
				if check.Check_model(configObj.Rules[i][0].Match, configObj.Rules[i][0].Content, url, bodys, headers, servers) && Fi1e == false{
					result_print(url, configObj.Product)
				}
			}
		}
	}
}

func ContainsInSlice(items []bool, item bool) bool {
	for _, eachItem := range items {
		if eachItem == item {
			return true
		}
	}
	return false
}

func result_print(url string, value string){
	fmt.Println("	" + value)
}

func result_writefile(url string, value string) {
	year := time.Now().Year()
	years := strconv.Itoa(int(year))
	month := time.Now().Month()
	months := strconv.Itoa(int(month))
	day := time.Now().Day()
	days := strconv.Itoa(int(day))
	var filepath string
	filepath = years + "." + months + "." + days + ".txt"
	file, err := os.OpenFile(filepath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("文件打开失败", err)
	}
	defer file.Close()
	write := bufio.NewWriter(file)
	write.WriteString(url + " :" + value + "\n")
	write.Flush()
}
