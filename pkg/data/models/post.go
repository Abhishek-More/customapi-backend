package models

type Post struct {
	ID    int `json:"_id,omitempty" bson:"_id,omitempty"`
	Title string `json:"title,omitempty" bson:"title,omitempty"`
	Body  string `json:"body,omitempty" bson:"body,omitempty"`
	URL   string `json:"url,omitempty" bson:"url,omitempty"`
}
