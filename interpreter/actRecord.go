package interpreter

import "go/constant"

type actRecord struct {
    variables map[string]constant.Value
}

func newActRecord() *actRecord {
    return &actRecord{variables: make(map[string]constant.Value)}
}