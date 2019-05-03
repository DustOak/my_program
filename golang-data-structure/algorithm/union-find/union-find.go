package union_find

type Api interface {
	//在p和q之间添加一条链接
	Union(p, q int)
	//p(0到n-1)所在的分量的标识符
	Find(p int) int
	//如果p和q存在于同一个分量中则返回true
	Connected(p, q int) bool
	//连通分量的数量
	Count() int
}

type UF struct {
	id    []int
	count int
}

func (u *UF) Union(p, q int) {
	pid := u.Find(p)
	qid := u.Find(q)
	if pid == qid {
		return
	}
	for i := 0; i < u.count; i++ {
		if u.id[i] == pid {
			u.id[i] = qid
		}
	}
	u.count--
}

func (u *UF) Find(p int) int {
	return u.id[p]
}

func (u *UF) Connected(p, q int) bool {
	return u.Find(p) == u.Find(q)
}

func (u *UF) Count() int {
	return u.count
}

func NewUF(count int) *UF {
	temp := &UF{id: make([]int, count), count: count}
	for k := range temp.id {
		temp.id[k] = k
	}
	return temp
}
