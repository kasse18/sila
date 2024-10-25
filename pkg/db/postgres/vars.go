package postgres

const (
	queryContainer  = "SELECT * FROM containers"
	insertContainer = "INSERT INTO containers(id, name, link_small, link_big) values ($1, $2, $3, $4)"
	queryUpload     = "UPDATE containers SET document_id = $1"
)

const (
	queryInitUsers = `CREATE TABLE IF NOT EXISTS containers (
		id bigint NOT NULL,
		name text NOT NULL,
		document_id bigint NULL,
		link_small text NOT NULL,
		link_bit text NOT NULL,
		PRIMARY KEY (id)
	  )`
)
