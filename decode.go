package main

import (
	"encoding/json"
	"go-json-parse/typefile"
	"os"
	"fmt"
)

func main() {
	// ファイルから読込みさせる場合
	var result typefile.Result
	data, _ := os.Open("./data/test.json")
	if err := json.NewDecoder(data).Decode(&result); err != nil {
		panic(err)
	}

	// デコードしたデータを表示
	fmt.Printf("took:%v\n", result.Took)
	fmt.Printf("took:%v\n", result.TimedOut)

	fmt.Print("=============\n")

	for _, hits:= range result.Hits.Hits{
		fmt.Printf("_index:%v\n", hits.Index)
		fmt.Printf("_type:%v\n", hits.Type)
		fmt.Printf("_score:%v\n", hits.Score)

		fmt.Printf("user:%v\n",hits.Source.User)
		fmt.Printf("message:%v\n",hits.Source.Message)
		fmt.Printf("date:%v\n",hits.Source.Date)
		fmt.Printf("likes:%v\n",hits.Source.Likes)

		fmt.Print("=============\n")
	}

	/* ElasticSearchのAPIコールしてJSONを取得させる場合
	resp, err := http.Get("https://localhost:9090/_search")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	var result typefile.Result
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		panic(err)
	}

	// デコードしたデータを表示
	fmt.Printf("took:%v\n", result.Took)
	fmt.Printf("took:%v\n", result.TimedOut)

	fmt.Print("=============\n")

	for _, hits:= range result.Hits.Hits{
		fmt.Printf("_index:%v\n", hits.Index)
		fmt.Printf("_type:%v\n", hits.Type)
		fmt.Printf("_score:%v\n", hits.Score)

		fmt.Printf("user:%v\n",hits.Source.User)
		fmt.Printf("message:%v\n",hits.Source.Message)
		fmt.Printf("date:%v\n",hits.Source.Date)
		fmt.Printf("likes:%v\n",hits.Source.Likes)

		fmt.Print("=============\n")
	*/
}