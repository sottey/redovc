package models

type Task struct {
	ID      int      `json:"id"`
	Project []string `json:"projects"`
	Tags    []string `json:"tags"`
	//Due       time.Time `json:"due,omitempty"`
	Completed bool   `json:"completed"`
	Archived  bool   `json:"archived"`
	Subject   string `json:"subject"`
}

/*
type Task struct {
	"id":
	"uuid": "0b777f23-abc5-4978-90f3-3d3beb581e13",
	"subject": "Migrate stuff from Apple Music to Amazon Music",
	"projects": [
		"home"
	],
	"tags": [],
	"due": "",
	"completed": true,
	"completed_date": "2023-11-17T17:42:37-08:00",
	"status": "completed",
	"archived": true,
	"is_priority": false,
	"notes": null,
	"recur": "",
	"recur_until": "",
	"prev_recur_todo_uuid": ""
}
*/
