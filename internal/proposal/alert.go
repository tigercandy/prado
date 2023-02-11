package proposal

import (
	"encoding/json"
	"time"
)

type AlertMessage struct {
	ProjectName string      `json:"project_name"`
	Env         string      `json:"env"`
	TraceID     string      `json:"trace_id"`
	Host        string      `json:"host"`
	Uri         string      `json:"uri"`
	Method      string      `json:"method"`
	ErrMsg      interface{} `json:"err_msg"`
	ErrStack    string      `json:"err_stack"`
	Timestamp   time.Time   `json:"timestamp"`
}

func (a *AlertMessage) Marshal() (jsonRaw []byte) {
	jsonRaw, _ = json.Marshal(a)
	return
}

type NotifyHandler func(msg *AlertMessage)
