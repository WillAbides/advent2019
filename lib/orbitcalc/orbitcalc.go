package orbitcalc

import (
	"strings"
)

type nodeMap map[string]*node

func (n nodeMap) addNode(name string) {
	if n[name] == nil {
		n[name] = &node{
			name: name,
		}
	}
}

func (n nodeMap) addConnection(src, dest string) {
	n.addNode(src)
	n.addNode(dest)
	d := n[dest]
	s := n[src]
	for _, c := range s.connections {
		if c == d {
			return
		}
	}
	s.connections = append(s.connections, d)
}

type node struct {
	name        string
	connections []*node
	hopCounts map[string]int
}

func (n *node) initHopCounts() {
	if n.hopCounts == nil {
		n.hopCounts = map[string]int{
			n.name: 0,
		}
	}
}

func (n *node) updateHopCount(target string, newCount int) {
	if newCount < 0 {
		return
	}
	n.initHopCounts()
	currentCount, exists := n.hopCounts[target]
	if !exists || newCount < currentCount {
		n.hopCounts[target] = newCount
	}
}

func (n *node) hasCalculatedPath(target string) bool {
	n.initHopCounts()
	_, ok := n.hopCounts[target]
	return ok
}

func (n *node) hopsToTarget(callstack []*node, target string) int {
	for _, n2 := range callstack {
		if n2 == n {
			return -1
		}
	}

	n.initHopCounts()

	if hops, ok := n.hopCounts[target]; ok {
		return hops
	}
	callstack = append(callstack, n)

	for _, c := range n.connections {
		cHops := c.hopsToTarget(callstack, target)
		if cHops < 0 {
			continue
		}
		n.updateHopCount(target, cHops + 1)
	}

	hops, ok := n.hopCounts[target]
	if !ok {
		return -1
	}
	return hops
}

func CalcOrbitTransfers(orbitMap , src, dest string) int {
	orbitMap = strings.TrimSpace(orbitMap)
	nodes := nodeMap{}
	for _, s := range strings.Split(orbitMap, "\n") {
		ss := strings.Split(s, ")")
		nodes.addConnection(ss[0], ss[1])
		nodes.addConnection(ss[1], ss[0])
	}

	s := nodes[src]
	return s.hopsToTarget(nil, dest) - 2
}

type orbits map[string]map[string]struct{}

func (o orbits) addOrbiter(center string, orbiters ...string) {
	if o[center] == nil {
		o[center] = make(map[string]struct{}, len(orbiters))
	}
	for _, orbiter := range orbiters {
		o[center][orbiter] = struct{}{}
	}
}

func (o orbits) getOrbiters(center string) []string {
	mp := o[center]
	res := make([]string, 0, len(mp))
	for s := range mp {
		res = append(res, s)
	}
	return res
}

func (o orbits) addSubOrbiters(center string) int {
	start := len(o[center])
	for _, s := range o.getOrbiters(center) {
		for _, s2 := range o.getOrbiters(s) {
			o.addOrbiter(center, s2)
		}
	}
	return len(o[center]) - start
}

func (o orbits) totalCount() int {
	total := 0
	for _, m := range o {
		total += len(m)
	}
	return total
}

//OrbitCount returns the total number of orbits in an orbit file
//see exinput.txt and Day 6 for input format
func OrbitCount(input string) int {
	o := orbits{}

	for _, s := range strings.Split(input, "\n") {
		ss := strings.Split(s, ")")
		o.addOrbiter(ss[0], ss[1])
	}

	for {
		delta := 0
		for c := range o {
			delta += o.addSubOrbiters(c)
		}
		if delta == 0 {
			break
		}
	}

	return o.totalCount()
}
