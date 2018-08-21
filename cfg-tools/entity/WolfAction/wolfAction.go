
package entity

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"encoding/json"
	"errors"
)

type WolfAction struct {
	WolfAction []struct {
		ActionID int `json:"ActionId"`
		Function string `json:"Function"`
		ArgName string `json:"ArgName"`
	} `json:"wolfAction"`
	WolfActionAt int `json:"wolfActionAt"`
	TabName string `json:"tabName"`
}

func NewWolfActionFromConfig() *WolfAction {
	o, err := InitObjectFromFile(NewWolfAction(0),"/./entity/WolfAction/wolfAction.json")
	if err != nil {
		return nil
	}
	p, ok := o.(*WolfAction)
	if ok {
		return p
	}
	return nil
}

func NewWolfAction(id int64) *WolfAction {
	p := &WolfAction{}
	p.Seq = int(id)
	p.TabName = "wolfAction"
	return p
}

func (p *WolfAction) String() zapcore.Field {
	return zap.Any(p.TabName, p)
}

func (p *WolfAction) Marshal() ([]byte, error) {
	buff, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}
	return buff, nil
}

func (p *WolfAction) UnMarshal(buff []byte) (GameObject, error) {
	err := json.Unmarshal(buff, p)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (p *WolfAction) GetTime() int64 {
	return int64(p.WolfActionAt)
}

func (p *WolfAction) SetTime(t int64) {
	p.WolfActionAt = int(t)
}

func (p *WolfAction) ID() int64 {
	return int64(p.Seq)
}

func (p *WolfAction) Name() string {
	return p.TabName
}

func (p *WolfAction) SetID(id int64) {
	p.Seq = int(id)
}

func (p *WolfAction) Instance(o GameObject, err error) (*WolfAction, error) {
	if err != nil {
		return p, err
	}
	p, ok := o.(*WolfAction)
	if ok {
		return p, nil
	}
	return p, errors.New("WolfAction instance fail")
}
