package Utilitys

import (
	"encoding/json"
	"io/ioutil"
)

type Duration struct {
	StartTimeStamp string `json:"StartTimeStamp"`
}
type Message struct {
	GoMessage  interface{} `json:"GoMessage"`
	AppMessage string      `json:"AppMessage"`
}
type Object struct {
	Name string      `json:"Name"`
	Obj  interface{} `json:"Obj"`
}
type LogInstance struct {
	Duration `json:"Meta"`
	Message  `json:"Message"`
	Object   `json:"Object"`
}
type Log struct {
	Key             string        `json:"Key"`
	value           []LogInstance `json:"value"`
	FinishTimeStamp string        `json:"StartTimeStamp"`
}
type ResultSet struct {
	Code        string
	Description string
}

func Logger(name, appMessage string, obj, last interface{}) *LogInstance {
	return &LogInstance{
		Duration: Duration{
			StartTimeStamp: GetDate(""),
		},
		Message: Message{
			GoMessage:  last,
			AppMessage: appMessage,
		},
		Object: Object{
			Name: name,
			Obj:  obj,
		},
	}
}

func WriteLog(log *LogInstance) {
	b, _ := json.Marshal(&log)
	ioutil.WriteFile("d:\\filename.txt", b, 0644)
}
