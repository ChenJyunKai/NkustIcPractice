package main

import (
	"encoding/json"
	"fmt"
)

/*
	TODO:
	熟悉Marshal、Unmarshal用法，並知道如何在不同的結構中應對自如
*/

func main() {
	//簡單轉換
	AstructToBstruct()
}

/*
	1. 練習結構Astruct轉結構Bstruct
*/
type Astruct struct {
	ValueA int64  `json:"value_a,omitempty"`
	ValueB string `json:"value_b,omitempty"`
}

type Bstruct struct {
	ValueA int64  `json:"value_a,omitempty"`
	ValueB string `json:"value_b,omitempty"`
}

/*
	2. 轉換的差異性
*/
type AstructWithDiff struct {
	ValueA int64  `json:"value_a,omitempty"`
	ValueB string `json:"value_b,omitempty"`
	ValueC string `json:"value_c,omitempty"`
}

type BstructWithDiff struct {
	ValueA int64  `json:"value_a,omitempty"`
	ValueB string `json:"value_b,omitempty"`
}

/*
	3. 當結構相似但型態不一樣的轉換
*/
type AstructWithDiffType struct {
	ValueA int64 `json:"value_a,omitempty"`
	ValueB int   `json:"value_b,omitempty"`
}

type BstructWithDiffType struct {
	ValueA int    `json:"value_a,omitempty"`
	ValueB string `json:"value_b,omitempty"`
}

/*
	4. 當其中一個為interface時
*/
type AstructWithInterface struct {
	ValueA int64       `json:"value_a,omitempty"`
	ValueB interface{} `json:"value_b,omitempty"`
}

type BstructWithInterface struct {
	ValueA int64  `json:"value_a,omitempty"`
	ValueB string `json:"value_b,omitempty"`
}

//練習結構Astruct轉結構Bstruct
func AstructToBstruct() (err error) {
	aStruct := Astruct{
		ValueA: 100,
		ValueB: "100",
	}

	aStructByte, err := json.Marshal(aStruct)
	if err != nil {
		fmt.Println(err)
	}

	output := &Bstruct{}
	err = json.Unmarshal(aStructByte, output)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(output)

	return err
}
