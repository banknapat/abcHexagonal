package model

type GetRequest struct {
	ID      int    `db:"id"`
	Name    string `db:"name"`
	Type	string `db:"type"`
	Detail  string `db:"detail"`
	URL	 string `db:"url"`
}

type GetRequestResponse struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Type	string `json:"type"`
	Detail  string `json:"detail"`
	URL	 string `json:"url"`
}