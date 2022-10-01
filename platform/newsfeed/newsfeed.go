package newsfeed

//Getter interface
type Getter interface {
	GetAll() []Item
}

//Adder interface
type Adder interface {
	Add(item Item)
}

//Item is model
type Item struct {
	Title string `json:"title"`
	Post  string `json:"post"`
}

//Repo is Collection on Item
type Repo struct {
	Items []Item
}

//New returns pointer to repo
func New() *Repo {
	return &Repo{
		Items: []Item{},
	}
}

//Add item to repo
func (r *Repo) Add(item Item) {
	r.Items = append(r.Items, item)
}

//GetAll returns the collection
func (r *Repo) GetAll() []Item {
	return r.Items
}
