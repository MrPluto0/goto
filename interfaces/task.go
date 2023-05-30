package interfaces

type Task struct {
	Name    string   `json:"name"`
	Alias   []string `json:"alias"`
	Command string   `json:"command"`
}
