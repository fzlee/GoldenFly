package common

type Pagination struct {
	Page int
	Size int
}

func (self *Pagination) GetOffset() int {
	if self.Page > 0 && self.Size > 0 {
		return (self.Page - 1) * self.Size
	}
	return 0
}

func (self *Pagination) GetLimit() int {
	if self.Size > 0 {
		return self.Size
	} else {
		return 12
	}
}
