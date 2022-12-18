package main

import (
	"sync"
	"sync/atomic"
)

const blockSize = 1000

func seqBfs(g Graph, v int) {
	var cur, next []*Node

	g.Reset(v)
	cur = append(cur, g.Get(v))

	for d := 1; len(cur) > 0; d++ {
		for _, n := range cur {
			for _, i := range g.Adjacent(n.idx) {
				a := g.Get(i)
				if a.distance == -1 {
					a.distance = int32(d)
					next = append(next, a)
				}
			}
		}

		cur = next
		next = []*Node{}
	}
}

func parBfs(g Graph, v int) {
	l := g.Size()
	var cur, next = make(chan []*Node, l), make(chan []*Node, l)

	g.Reset(v)
	cur <- []*Node{g.Get(v)}

	for d := 1; len(cur) > 0; d++ {
		goroutinesCnt := maxGoroutines
		if len(cur) < goroutinesCnt {
			goroutinesCnt = len(cur)
		}

		var wg sync.WaitGroup
		for i := 0; i < goroutinesCnt; i++ {
			wg.Add(1)

			go func() {
				defer wg.Done()

				for {
					select {
					case nodes := <-cur:
						var adjs []*Node
						for _, n := range nodes {
							for _, a := range g.Adjacent(n.idx) {
								adj := g.Get(a)
								if atomic.CompareAndSwapInt32(&adj.distance, -1, int32(d)) {
									adjs = append(adjs, adj)
								}
							}
							if len(adjs) > blockSize {
								next <- adjs
								adjs = []*Node{}
							}
						}
						if len(adjs) > 0 {
							next <- adjs
						}
					default:
						return
					}
				}
			}()
		}
		wg.Wait()

		cur, next = next, cur
	}
}
