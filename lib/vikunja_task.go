package lib

type VikunjaTask struct {
	Assignees              []interface{} `json:"assignees"`
	Attachments            []interface{} `json:"attachments"`
	BucketID               int           `json:"bucket_id"`
	CoverImageAttachmentID int           `json:"cover_image_attachment_id"`
	Created                string        `json:"created"`
	CreatedBy              struct{}      `json:"created_by"`
	Description            string        `json:"description"`
	Done                   bool          `json:"done"`
	DoneAt                 string        `json:"done_at"`
	DueDate                string        `json:"due_date"`
	EndDate                string        `json:"end_date"`
	HexColor               string        `json:"hex_color"`
	ID                     int           `json:"id"`
	Identifier             string        `json:"identifier"`
	Index                  int           `json:"index"`
	IsFavorite             bool          `json:"is_favorite"`
	KanbanPosition         int           `json:"kanban_position"`
	Labels                 []interface{} `json:"labels"`
	PercentDone            int           `json:"percent_done"`
	Position               int           `json:"position"`
	Priority               int           `json:"priority"`
	ProjectID              int           `json:"project_id"`
	RelatedTasks           struct{}      `json:"related_tasks"`
	Reminders              []interface{} `json:"reminders"`
	RepeatAfter            int           `json:"repeat_after"`
	RepeatMode             int           `json:"repeat_mode"`
	StartDate              string        `json:"start_date"`
	Subscription           struct{}      `json:"subscription"`
	Title                  string        `json:"title"`
	Updated                string        `json:"updated"`
}
