package types

type Row []Cell

func (r *Row) String() (ret []string) {
	for _, cell := range *r {
		ret = append(ret, cell.String())
	}
	return
}
