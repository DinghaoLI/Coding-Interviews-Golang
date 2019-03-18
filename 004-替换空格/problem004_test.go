package problem004

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

type para struct {
	str []byte
	length int
}

type ans struct {
	res []byte
}

type question struct {
	p para
	a ans
}

func Test_OK(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		question{
			p: para{
				str: []byte{'a',' ','b',' ','c','x','x','x','x'},
				length: 5,
			},
			a: ans{
				res: []byte{'a','%','2','0','b','%','2','0','c'},
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		replaceSpace(p.str, p.length)
		ast.Equal(string(a.res), string(p.str), "输入:%v", p)
	}
}