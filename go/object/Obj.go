package object

// Obj
type Obj struct {
	Name string
	Id   int
}

// NewObj
func NewObj() *Obj {
	return &Obj{
		Name: "",
		Id:   -1,
	}
}

// SetId
func (o *Obj) SetId(id int) {
	o.Id = id
}
