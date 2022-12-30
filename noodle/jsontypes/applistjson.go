// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    appList, err := UnmarshalAppList(bytes)
//    bytes, err = appList.Marshal()

package jsontypes

import "encoding/json"

func UnmarshalAppList(data []byte) (AppList, error) {
	var r AppList
	err := json.Unmarshal(data, &r)
	return r, err
}

type AppList struct {
	Appcount int64 `json:"appcount"`
	Apps     []App `json:"apps"`
}

type App struct {
	Appid          string `json:"appid"`
	Name           string `json:"name"`
	Website        string `json:"website"`
	License        string `json:"license"`
	Description    string `json:"description"`
	Enhanced       bool   `json:"enhanced"`
	TileBackground string `json:"tile_background"`
	Icon           string `json:"icon"`
	SHA            string `json:"sha"`
}
