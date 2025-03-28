package main

type task struct {
  TaskTitle string `json:"taskTitle"`
}

func (t task) FilterValue() string {
	return t.TaskTitle
}

func (t task) Title() string {
	return t.TaskTitle
}

func (t task) Description() string {
	// TODO: there is an empty space under the title, because of this empty string, somehow remove it...
	// ...maybe create a new delegate and use the render method of it
	return ""
}
