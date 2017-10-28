package json


type TreeNode struct {
	Id int64 `json:"id"`
	Name string`json:"text"`
	Children bool `json:"children"`
    Icon string `json:"icon"`
	Type string `json:type`
	State map[string]interface{} `json:"state"`
}