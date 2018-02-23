package typefile

/*
type ResultShards struct {
	Total      int `json:"total"`
	Successful int `json:"successful"`
	Failed     int `json:"failed"`
}

type ResultHitSource struct {
	User    string `json:"user"`
	Message string `json:"message"`
	Date    string `json:"date"`
	Likes   int    `json:"likes"`
}

type ResultHit struct {
	Index  string          `json:"_index"`
	Type   string          `json:"_type"`
	Id     string          `json:"_id"`
	Score  float32         `json:"_score"`
	Source ResultHitSource `json:"_source"`
}

type ResultHits struct {
	Total    int         `json:"total"`
	MaxScore float32     `json:"max_score"`
	Hits     []ResultHit `json:"hits"`
}

type Result struct {
	Took     int          `json:"took"`
	TimedOut bool         `json:"timed_out"`
	Shards   ResultShards `json:"_shards"`
	Hits     ResultHits   `json:"hits"`
}
*/

// 構造体をネスト定義することが可能なので、上記のコードに適用したもの
// ネスト定義する構造体には型名が必要ないのでそれも削除できる
type Result struct {
	Took     int  `json:"took"`
	TimedOut bool `json:"timed_out"`
	Shards   struct {
		Total      int `json:"total"`
		Successful int `json:"successful"`
		Failed     int `json:"failed"`
	} `json:"_shards"`
	Hits struct {
		Total    int     `json:"total"`
		MaxScore float32 `json:"max_score"`
		Hits     []struct {
			Index  string  `json:"_index"`
			Type   string  `json:"_type"`
			Id     string  `json:"_id"`
			Score  float32 `json:"_score"`
			Source struct {
				User    string `json:"user"`
				Message string `json:"message"`
				Date    string `json:"date"`
				Likes   int    `json:"likes"`
			} `json:"_source"`
		} `json:"hits"`
	} `json:"hits"`
}