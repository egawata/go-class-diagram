package sample

// sample file
// this is a test

type NameGetter interface {
	GetName() string
}

type Author struct {
	ID   int
	Name string
}

func (a *Author) GetName() string {
	return a.Name
}

type Book struct {
	ID     int
	Name   string
	Author *Author
}

func (b *Book) GetName() string {
	return b.Name
}

func (b *Book) GetID() int {
	return b.ID
}

type Library struct {
	ID    int
	Name  string
	Books []*Book
}

type PublicLibrary struct {
	Library
	City string
}

func (l *Library) GetName() string {
	return l.Name
}

func (l *Library) GetID() int {
	return l.ID
}

func (l *Library) SetName(n string) {
	l.Name = n
}

/*
func main() {
	l := &PublicLibrary{
		Library{
			ID:   101,
			Name: "Tokyo Library",
			Books: []*Book{
				{ID: 201, Name: "History", Author: &Author{ID: 301, Name: "Bob"}},
				{ID: 202, Name: "Science", Author: &Author{ID: 302, Name: "Alice"}},
			},
		},
		"Tokyo",
	}
	fmt.Printf("%#v\n", l)
}
*/
