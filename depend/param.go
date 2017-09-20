package depend

import "strconv"

type ParamString string
type ParamNumber int64

func NewParamString(v string) Param{
	return ParamString(v[1 : len(v) - 1])
}

func NewParamNumber(v string) Param{
	return ParamNumber(ParamString(v).Number())
}

func(ps ParamString)String() string {
	return string(ps)
}

func(ps ParamString)Number() int64 {
	v, err := strconv.ParseInt(ps.String(), 10, 0)
	if nil != err {
		panic(err)
	}
	return v
}

func(ps ParamString)Boolean() bool {
	return "" == ps
}

func(pn ParamNumber)String() string {
	return strconv.FormatInt(pn.Number(), 10)
}

func(pn ParamNumber)Number() int64 {
	return int64(pn)
}

func(pn ParamNumber)Boolean() bool {
	return 0 != pn
}