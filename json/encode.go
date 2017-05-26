package main

import (
	"encoding/json"
	"fmt"
)

type cartoon struct {
	name        string `json:"name"`        //名称
	author      string `json:"author"`      //作者
	protagonist string `json:"protagonist"` //主角
}
