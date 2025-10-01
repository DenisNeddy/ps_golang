package bins

import "time"

type Bin struct {
	Id        string    `json:"id"`
	Private   bool      `json:"private"`
	CreatedAt time.Time `json:"createdAt"`
	Name      string    `json:"name"`
}

func NewBin(id string, private bool, name string) *Bin {
	return &Bin{
		Id:        id,
		Private:   private,
		CreatedAt: time.Now(),
		Name:      name,
	}
}

type BinList []Bin

func (l *BinList) Add(bin *Bin) {
	*l = append(*l, *bin)
}
