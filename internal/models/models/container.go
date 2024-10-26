package models

type Container struct {
	ID         int64  `json:"id"`
	Name       string `json:"name"`
	DocumentID int64  `json:"document_id"`
	LinkSmall  string `json:"link_small"`
	LinkBig    string `json:"link_big"`
}

type CreateContainer struct {
	Name      string `json:"name"`
	LinkSmall string `json:"link_small"`
	LinkBig   string `json:"link_big"`
}
