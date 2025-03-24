package main

type task struct {
	title string
}

func (t task) FilterValue() string {
	return t.title
}

func (t task) Title() string {
	return t.title
}

func (t task) Description() string {
	// TODO: there is an empty space under the title, because of this empty string, somehow remove it...
	// ...maybe create a new delegate and use the render method of it
	return ""
}
