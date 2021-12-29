package entity

type Activity struct {
	ID        int64       `json:"id"`
	Email     string      `json:"email,omitempty"`
	Title     string      `json:"title,omitempty"`
	CreatedAt string      `json:"created_at,omitempty"`
	UpdatedAt string      `json:"updated_at,omitempty"`
	DeletedAt interface{} `json:"deleted_at,omitempty"`
}

type ActivityWNull struct {
	ID        int64       `json:"id"`
	Email     string      `json:"email,omitempty"`
	Title     string      `json:"title,omitempty"`
	CreatedAt string      `json:"created_at,omitempty"`
	UpdatedAt string      `json:"updated_at,omitempty"`
	DeletedAt interface{} `json:"deleted_at"`
}

type EmptyObject struct {
}

type ActivityCreateRequest struct {
	Email string `json:"email"`
	Title string `json:"title"`
}

type TodoCreateRequest struct {
	ActivityGroupId int64  `json:"activity_group_id"`
	Title           string `json:"title"`
	IsActive        bool   `json:"is_active"`
}
