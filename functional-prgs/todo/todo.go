package todo

type Todo struct {
	Text string
	Db   *Db // reference to the database
}

func NewTodo() Todo {
	return Todo{
		Text: "",
		Db:   NewDb(),
	}
}

func (t *Todo) Write(text string) {
	if t.Db.IsAuthorized() {
		t.Text = text
	} else {
		panic("not authorized to write todo")
	}
}

func (t *Todo) Append(text string) {
	if t.Db.IsAuthorized() {
		t.Text += text
	} else {
		panic("not authorized to append todo")
	}
}
