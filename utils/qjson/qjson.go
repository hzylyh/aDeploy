/*
 * @Description:
 * @Author: John Holl
 * @Github: https://github.com/hzylyh
 * @Date: 2021-05-10 17:26:45
 * @LastEditors: John Holl
 * @LastEditTime: 2021-05-10 17:27:13
 */
package qjson

type QJson struct {
	ReqInfo map[string]interface{}
}

func (qj *QJson) GetString(key string) string {
	if res, ok := qj.ReqInfo[key].(string); ok {
		return res
	}
	return ""
}

func (qj *QJson) GetNum(key string) float64 {
	if res, ok := qj.ReqInfo[key].(float64); ok {
		return res
	}
	return 0
}

func (qj *QJson) GetJson(key string) *QJson {
	//var (
	//	reqInfo map[string]interface{}
	//	err     error
	//)
	if res, ok := qj.ReqInfo[key].(map[string]interface{}); ok {
		//if err = json.Unmarshal([]byte(res), &reqInfo); err != nil {
		//	return nil
		//}
		return &QJson{ReqInfo: res}
	}
	return nil
}

func (qj *QJson) GetMap(key string) map[string]interface{} {
	if res, ok := qj.ReqInfo[key].(map[string]interface{}); ok {
		//if err = json.Unmarshal([]byte(res), &reqInfo); err != nil {
		//	return nil
		//}
		return res
	}
	return nil
}

func (qj *QJson) GetArray(key string) []interface{} {
	if res, ok := qj.ReqInfo[key].([]interface{}); ok {
		//if err = json.Unmarshal([]byte(res), &reqInfo); err != nil {
		//	return nil
		//}
		return res
	}
	return nil
}
