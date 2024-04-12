package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"
)
var TrueArr = []string{"1", "t", "T", "TRUE", "true","True"}
var FalseArr = []string{"0", "f", "F", "FALSE", "false","False"}
 
var jsonInput1 = `{
	"null_1": {
		"NULL": "T"
	  },
	"bool_1": {
		"BOOL": "T"
	  },
	"number_1": {
	  "N": "1.50"
	},
	"string_1": {
	  "S": "784498 "
	},
	"string_2": {
	  "S": "2014-07-16T20:55:46Z"
	},
	"map_1": {
	  "M": {
		"bool_1": {
		  "BOOL": "truthy"
		},
		"null_1": {
		  "NULL ": "true"
		},
		"list_1": {
		  "L": [
			{
			  "S": ""
			},
			{
			  "N": "011"
			},
			{
			  "N": "5215s"
			},
			{
			  "BOOL": "f"
			},
			{
			  "NULL": "0"
			}
		  ]
		}
	  }
	},
	"list_2": {
	  "L": "noop"
	},
	"list_3": {
	  "L": [
		"noop"
	  ]
	},
	"": {
	  "S": "noop"
	}
  }`
func main() {
	var input map[string]interface{}
	
	json.Unmarshal([]byte(jsonInput1), &input)

	ans, ok := ConvertMap(input)
	if !ok {
		return
	}
		jsonAns , err := json.Marshal(ans)
		if err != nil {
			return
		}
		fmt.Println(string(jsonAns))
}

func CleanString(str string) (string , bool) {

	ans := strings.TrimSpace(str)
	if ans == "" {
		return ans,false
	}
	return ans,true
}

func ConvertString(inter interface{}) (string,bool) {
	if str,ok := inter.(string); ok {
		str, ok = CleanString(str)
		if !ok {
			return "",false
		}
		time ,err := time.Parse(time.RFC3339, str)
		if err != nil {
			return str,true
		}
		return strconv.FormatInt(time.Unix(),10),true
	}
	return "",false
}

func ConvertNumber(inter interface{}) (interface{},bool) {
	if str,ok := inter.(string); ok {
		str, ok = CleanString(str)
		if !ok {
			return "",false
		}
		i, err := strconv.ParseInt(str, 10, 64)
		if err != nil {
			f, err := strconv.ParseFloat(str, 64)
				if err != nil {
					return "", false
				}
				return f,true
		}
		return i,true
	}
	return "",false
}

func ConvertBoolean(inter interface{}) (bool,bool) {
	if str,ok := inter.(string); ok {
		str, ok = CleanString(str)
		if !ok {
			return false,false
		}
		for _,v := range TrueArr {
			if v == str {
				return true,true;
			}
		}
		for _,v := range FalseArr {
			if v == str {
				return false,true;
			}
		}
	}
	return false,false
}
func ConvertNull(inter interface{}) (interface{},bool) {
	if str,ok := inter.(string); ok {
		str, ok = CleanString(str)
		if !ok {
			return false,false
		}
		for _,v := range TrueArr {
			if v == str {
				return nil,true;
			}
		}
	}
	return false,false
}

func ConvertList(inter interface{}) (interface{},bool) {
	ans := []interface{}{}
	if Arr,ok := inter.([]interface{}); ok {
		
		for _,v := range Arr {
			if obj, ok := v.(map[string]interface{}); ok {
				

			for k1,v1 := range obj {
				k1, ok := CleanString(k1)
			if !ok {
				continue
			}
			if k1 == "S" {
				if value,ok := ConvertString(v1); ok {
					ans = append(ans, value)
				}
			} else if k1 == "N" {

				if value,ok := ConvertNumber(v1); ok {
					ans = append(ans, value)
				}
				
			} else if k1 == "L" {

				if value,ok := ConvertList(v1); ok {
					ans = append(ans, value)
				}
			
			} else if k1 == "M" {
				if value,ok := ConvertMap(v1); ok {
					ans = append(ans, value)
				}
				
			} else if k1 == "BOOL" {

				if value,ok := ConvertBoolean(v1); ok {
					ans = append(ans, value)	
				}
			} else if k1 == "NULL" {

				if value,ok := ConvertNull(v1); ok {
					ans = append(ans, value)
				}
			}
		}
		}

		}
	}
	if len(ans) != 0 {
		return ans,true
	}
	return ans,false
}

func ConvertMap(inter interface{}) (interface{},bool) {
	
	
	output := map[string]interface{}{}
	if Map,ok := inter.(map[string]interface{}); ok {
		
		for k,v := range Map {
			k, ok := CleanString(k)
			if !ok {
				continue
			}
			if obj, ok := v.(map[string]interface{}); ok {

			for k1,v1 := range obj {

			if k1 == "S" {
				if value,ok := ConvertString(v1); ok {
					output[k] = value;
				}
			} else if k1 == "N" {

				if value,ok := ConvertNumber(v1); ok {
					output[k] = value;
				}
				
			} else if k1 == "L" {

				if value,ok := ConvertList(v1); ok {
					output[k] = value;
				}
			
			} else if k1 == "M" {
				
				if value,ok := ConvertMap(v1); ok {
					output[k] = value;
				}
				
			} else if k1 == "BOOL" {

				if value,ok := ConvertBoolean(v1); ok {
					output[k] = value;	
				}
			} else if k1 == "NULL" {

				if value,ok := ConvertNull(v1); ok {
					output[k] = value;
				}
			}
		}
		}

		}
	}
	if len(output) != 0 {
		return output,true
	}
	return output,false
}