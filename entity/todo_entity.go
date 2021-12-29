package entity

type Todo struct {
	ID              int64       `json:"id"`
	ActivityGroupId int64       `json:"activity_group_id"`
	Title           string      `json:"title"`
	IsActive        bool        `json:"is_active"`
	Priority        string      `json:"priority"`
	CreatedAt       string      `json:"created_at"`
	UpdatedAt       string      `json:"updated_at"`
	DeletedAt       interface{} `json:"deleted_at,omitempty"`
}
