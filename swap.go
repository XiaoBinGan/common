package common

import "encoding/json"

//通过json tag 进行结构体赋值
func SwapTo(request,res interface{})error{
	bytes, err := json.Marshal(request)
	if err!=nil{
		return err
	}
	return json.Unmarshal(bytes,res)
}