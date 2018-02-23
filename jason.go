package main

import (
	"github.com/antonholmquist/jason"
	"os"
	"fmt"
)

func main() {
	// ファイルから読込みさせる場合
	data, _ := os.Open("./data/test.json")

	v, err := jason.NewObjectFromReader(data)
	if err != nil {
		panic(err)
	}

	// Read values
	took, err := v.GetInt64("took")
	timed_out, err := v.GetBoolean("timed_out")

	// Read nested values
	_shards_total, err := v.GetInt64("_shards","total")
	_shards_successful, err := v.GetInt64("_shards","successful")
	_shards_failed, err := v.GetInt64("_shards","failed")

	hits_total, err := v.GetInt64("hits","total")
	hits_max_score, err := v.GetInt64("hits","max_score")

	fmt.Printf("took:%v\n", took)
	fmt.Printf("timed_out:%v\n", timed_out)
	fmt.Printf("_shards_total:%v\n", _shards_total)
	fmt.Printf("_shards_successful:%v\n", _shards_successful)
	fmt.Printf("_shards_failed:%v\n", _shards_failed)
	fmt.Printf("hits_total:%v\n", hits_total)
	fmt.Printf("hits_max_score:%v\n", hits_max_score)

	// Loop through array
	hits_hits, err := v.GetObjectArray("hits", "hits")
	for _, hits := range hits_hits{
		_index, _ := hits.GetString("_index")
		_type, _ := hits.GetString("_type")
		_id, _ := hits.GetString("_id")
		_score, _ := hits.GetString("_score")

		_source_user, _ := hits.GetString("_source","user")
		_source_message, _ := hits.GetString("_source","message")
		_source_date, _ := hits.GetString("_source","date")
		_source_likes, _ := hits.GetString("_source","likes")

		fmt.Printf("_index:%v\n", _index)
		fmt.Printf("_type:%v\n", _type)
		fmt.Printf("_id:%v\n", _id)
		fmt.Printf("_score:%v\n", _score)

		fmt.Printf("_source_user:%v\n", _source_user)
		fmt.Printf("_source_message:%v\n", _source_message)
		fmt.Printf("_source_date:%v\n", _source_date)
		fmt.Printf("_source_likes:%v\n", _source_likes)
	}

}