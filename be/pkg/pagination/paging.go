package pagination

type Paging struct {
	Page   int   `json:"page"`
	Limit  int   `json:"limit"`
	Offset int   `json:"offset"`
	Total  int64 `json:"total"`
}

func (p *Paging) Process() {
	if p.Page < 1 {
		p.Page = 1
	}
	if p.Limit < 1 || p.Limit > 100 {
		p.Limit = 10
	}
}
