package main

import (
	"encoding/json"
    "go-json-parse/typefile"
	"log"
	"fmt"
)

func main() {
	// JSONテストデータ
	data := []byte(`{
					"took": 1,
					"timed_out": false,
					"_shards":{
						"total" : 1,
						"successful" : 1,
						"failed" : 0
					},
					"hits":{
						"total" : 1,
						"max_score": 1.3862944,
						"hits" : [
							{
								"_index" : "twitter",
								"_type" : "tweet",
								"_id" : "0",
								"_score": 1.3862944,
								"_source" : {
									"user" : "kimchy",
									"message": "trying out Elasticsearch",
									"date" : "2009-11-15T14:12:12",
									"likes" : 0
								}
							},
							{
								"_index" : "facebook",
								"_type" : "comment",
								"_id" : "1",
								"_score": 2.3213183,
								"_source" : {
									"user" : "mikky",
									"message": "trying out Elasticsearch",
									"date" : "2012-12-27T19:16:24",
									"likes" : 100
								}
							}
						]
					}
				}`)

	var result typefile.Result
	// UnmarshalでのJSONデコード
	if err := json.Unmarshal(data,&result); err != nil {
		log.Fatal(err)
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

}