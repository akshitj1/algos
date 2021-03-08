package ds

const (
	BIN_DIGITS = 2
	INT_BITS   = 30
)

type TrieNode struct {
	c [BIN_DIGITS]*TrieNode
}

type Trie struct {
	root *TrieNode
}

func NewTrie() *Trie {
	return &Trie{}
}

func (t *Trie) insert(nd *TrieNode, x, i int) *TrieNode {
	if nd == nil {
		nd = &TrieNode{}
	}
	if i < 0 {
		return nd
	}
	ib := (x >> uint(i)) & 1
	nd.c[ib] = t.insert(nd.c[ib], x, i-1)
	return nd
}

func (t *Trie) Insert(x int) {
	t.root = t.insert(t.root, x, INT_BITS-1)
}

func (t *Trie) query(nd *TrieNode, x, i int) int {
	if i < 0 {
		return 0
	}
	bi := (x >> uint(i)) & 1
	if nd.c[bi] == nil {
		bi = 1 - bi
	}
	return bi<<uint(i) + t.query(nd.c[bi], x, i-1)
}

func (t *Trie) Query(x int) int {
	return t.query(t.root, x, INT_BITS-1)
}
