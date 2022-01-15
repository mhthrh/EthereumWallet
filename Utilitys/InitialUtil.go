package Utilitys

import (
	"fmt"
)

type ConfigInterface interface {
	Print()
	Update()
}

type Config struct {
	AppName    string `json:"AppName"`
	AppIsLive  string `json:"AppIsLive"`
	Version    string `json:"Version"`
	ExpireDate string `json:"ExpireDate"`

	Map map[string]interface{}
}

func NewConfig() ConfigInterface {
	c := new(Config)
	//cr := NewKey()
	//bFile, err := ioutil.ReadFile("AppFiles//config.json")
	//
	//if err != nil {
	//	fmt.Println("Can't open json file!\n", err)
	//	return nil
	//}
	//str, err := cr.Decrypt(string(bFile))
	//if err != nil {
	//	fmt.Println("Can't Decrypt file!\n", err)
	//	return nil
	//}
	//
	//jsonMap := make(map[string]interface{})
	//err = json.Unmarshal([]byte(str), &jsonMap)
	//if err != nil {
	//	fmt.Println(err)
	//	return nil
	//}
	//m, b := dumpMap("", jsonMap)
	//if !b {
	//	return nil
	//}
	//c.Map = m
	return c
}

func dumpMap(space string, m map[string]interface{}) (map[string]interface{}, bool) {
	for _, v := range m {
		if mv, ok := v.(map[string]interface{}); ok {
			dumpMap(space+"\t", mv)
		}
	}
	if m == nil {
		return nil, false
	}
	return m, true
}

func (c *Config) Print() {

	for k, v := range c.Map {
		fmt.Printf(" %s ==> %s \n", k, v)
	}
}

func (c *Config) Update() {
	//var answer string
	//var key string
	//var value string
	////var cr1 myCrypto.CryptoInterface
	//cr1 := NewKey()
	//
	//fmt.Print("You want change configuration, are you sure ? (y/n)")
	//fmt.Scan(&answer)
	//answer = strings.ToLower(answer)
	//if answer == "y" {
	//	fmt.Print("Key, please:")
	//	fmt.Scan(&key)
	//	fmt.Print("value, please: ")
	//	fmt.Scan(&value)
	//
	//	if val, ok := c.Map[key]; ok {
	//		if val == value {
	//			fmt.Print("The key was found, but values are same. by")
	//		} else {
	//			c.Map[key] = value
	//			fmt.Print("The key was found, and value set successfully, save it?(y/n)")
	//			answer = "1"
	//			fmt.Scan(&answer)
	//			if answer == "y" {
	//				jsonString, _ := json.Marshal(c.Map)
	//				s, _ := GetPath()
	//				if b, _ := ExistDir(s + "\\AppFiles"); !b {
	//					if err := os.Truncate(s+"\\AppFiles\\config.json", 0); err != nil {
	//						fmt.Printf("Failed to truncate: %v", err)
	//					}
	//
	//				}
	//				file, err := os.Create("\\AppFiles\\config.json")
	//				if err != nil {
	//					fmt.Println("Can't create file!")
	//					return
	//				}
	//				str, _ := cr1.Encrypt(string(jsonString))
	//				_, err = file.Write([]byte(str))
	//				if err != nil {
	//					fmt.Println("Can't write to file!")
	//					return
	//				}
	//			}
	//		}
	//	} else {
	//		fmt.Println("Key not found. by")
	//
	//	}
	//} else {
	//	fmt.Println("Ok, by")
	//}
}
