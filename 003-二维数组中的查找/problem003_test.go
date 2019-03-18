package problem003

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

type para struct {
	board [][]int
	target int
}

type ans struct {
	find bool
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
				board: [][]int{
					[]int{ 1, 2, 8, 9, },
    				[]int{ 2, 4, 9, 12, },
    				[]int{ 4, 7, 10, 13, },
    				[]int{ 6, 8, 11, 15, },	
				},
				target: 7,
			},
			a: ans{
				find: true,
			},
		},
		question{
			p: para{
				board: [][]int{
					[]int{ 1, 2, 8, 9, },
    				[]int{ 2, 4, 9, 12, },
    				[]int{ 4, 6, 10, 13, },
    				[]int{ 6, 8, 11, 15, },		
				},
				target: 7,
			},
			a: ans{
				find: false,
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.find, Find(p.board, p.target), "输入:%v", p)
	}
}