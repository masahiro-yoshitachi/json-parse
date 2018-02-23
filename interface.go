package main

import (
	"encoding/json"
	"os"
	"fmt"
)

func main() {
	var result interface{}

	data, _ := os.Open("./data/test.json")
	if err := json.NewDecoder(data).Decode(&result); err != nil {
		panic(err)
	}

	took, _ := result.(map[string]interface{})["took"].(float64)
	timed_out, _ := result.(map[string]interface{})["timed_out"].(bool)
	shards_total, _ := result.(map[string]interface{})["_shards"].(map[string]interface{})["total"].(float64)
	shards_successful, _ := result.(map[string]interface{})["_shards"].(map[string]interface{})["successful"].(float64)

	fmt.Printf("took:%v\n", int(took))
	fmt.Printf("timed_out:%v\n", timed_out)
	fmt.Printf("shards_total:%v\n", shards_total)
	fmt.Printf("shards_successful:%v\n", shards_successful)

	hits_total, _ := result.(map[string]interface{})["hits"].(map[string]interface{})["total"].(float64)
	hits_max_score, _ := result.(map[string]interface{})["hits"].(map[string]interface{})["max_score"].(float64)

	fmt.Printf("hits_total:%v\n", hits_total)
	fmt.Printf("hits_max_score:%v\n", hits_max_score)

	hits, _ := result.(map[string]interface{})["hits"].(map[string]interface{})["hits"].([]interface{})
	for _, hits:= range hits{
		_index, _:= hits.(map[string]interface{})["_index"].(string)
		_type, _:= hits.(map[string]interface{})["_type"].(string)
		_id, _:= hits.(map[string]interface{})["_id"].(float64)
		_score, _:= hits.(map[string]interface{})["_score"].(float64)

		fmt.Print("=============\n")
		fmt.Printf("_index:%v\n", _index)
		fmt.Printf("_type:%v\n", _type)
		fmt.Printf("_id:%v\n", int(_id))
		fmt.Printf("_score:%v\n", _score)

		source, _:= hits.(map[string]interface{})["_source"]

		user, _:= source.(map[string]interface{})["user"]
		message, _:= source.(map[string]interface{})["message"]
		date, _:= source.(map[string]interface{})["date"]
		likes, _:= source.(map[string]interface{})["likes"]

		fmt.Printf("user:%v\n", user)
		fmt.Printf("message:%v\n", message)
		fmt.Printf("date:%v\n", date)
		fmt.Printf("likes:%v\n", likes)
	}
}