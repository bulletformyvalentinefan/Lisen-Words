package deezer

type Track struct {
	ID			int64	`json:"id"`
	Title		string	`json:"title"`
	Duration	int		`json:"duration"`
	Preview		string	`json:"preview"`
	Artist		Artist	`json:"artist"`
}

type Artist struct {
	Name		string	`json:"name"`
}

type SearchResponse struct {
	Data	[]Track 	`json:"data"`
}