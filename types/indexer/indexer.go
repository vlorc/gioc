package indexer

import "strconv"

type Int int

type String string

func (i Int) Value() int {
	return int(i)
}

func (i Int) String() string {
	return strconv.FormatInt(int64(i), 10)
}

func (i String) Value() int {
	v, _ := strconv.ParseInt(string(i), 10, 63)
	return int(v)
}

func (i String) String() string {
	return string(i)
}
