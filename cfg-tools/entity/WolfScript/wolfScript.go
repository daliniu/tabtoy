
package entity

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"encoding/json"
	"errors"
)

type WolfScript struct {
	WolfScript []struct {
		AIID int `json:"AIId"`
		Second int `json:"Second"`
		Condition1 int `json:"Condition1"`
		Action1 struct {
			Action int `json:"action"`
			Arg1 string `json:"arg1"`
		} `json:"Action1"`
		NextID1 int `json:"NextId1"`
		Condition2 int `json:"Condition2"`
		Action2 struct {
			Action int `json:"action"`
			Arg1 string `json:"arg1"`
		} `json:"Action2,omitempty"`
		NextID2 int `json:"NextId2,omitempty"`
	} `json:"wolfScript"`
	WolfScriptAt int `json:"wolfScriptAt"`
	TabName string `json:"tabName"`
}

func NewWolfScriptFromConfig() *WolfScript {
	o, err := InitObjectFromFile(NewWolfScript(0),"/./entity/WolfScript/wolfScript.json")
	if err != nil {
		return nil
	}
	p, ok := o.(*WolfScript)
	if ok {
		return p
	}
	return nil
}

func NewWolfScript(id int64) *WolfScript {
	p := &WolfScript{}
	p.Seq = int(id)
	p.TabName = "wolfScript"
	return p
}

func (p *WolfScript) String() zapcore.Field {
	return zap.Any(p.TabName, p)
}

func (p *WolfScript) Marshal() ([]byte, error) {
	buff, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}
	return buff, nil
}

func (p *WolfScript) UnMarshal(buff []byte) (GameObject, error) {
	err := json.Unmarshal(buff, p)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (p *WolfScript) GetTime() int64 {
	return int64(p.WolfScriptAt)
}

func (p *WolfScript) SetTime(t int64) {
	p.WolfScriptAt = int(t)
}

func (p *WolfScript) ID() int64 {
	return int64(p.Seq)
}

func (p *WolfScript) Name() string {
	return p.TabName
}

func (p *WolfScript) SetID(id int64) {
	p.Seq = int(id)
}

func (p *WolfScript) Instance(o GameObject, err error) (*WolfScript, error) {
	if err != nil {
		return p, err
	}
	p, ok := o.(*WolfScript)
	if ok {
		return p, nil
	}
	return p, errors.New("WolfScript instance fail")
}
