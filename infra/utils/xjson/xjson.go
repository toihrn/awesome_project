package xjson

import (
	"github.com/bytedance/sonic"
	"github.com/sirupsen/logrus"
)

const (
	MarshalErr = "MarshalErr"
)

func ToString(i interface{}) string {
	str, err := sonic.MarshalString(i)
	if err != nil {
		logrus.Errorf("[ToString] marshal err: %v", err)
		return MarshalErr
	}
	return str
}
