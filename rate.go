package goshiki

type Rate struct {
	ID        int64   `json:"id"`
	Score     float32 `json:"score"`
	Status    string  `json:"status"`
	Text      string  `json:"text"`
	TextHTML  string  `json:"text_html"`
	Episodes  int     `json:"episodes"`
	Chapters  int     `json:"chapters"`
	Volumes   int     `json:"volumes"`
	Rewatches int     `json:"rewatches"`
}
