package main

import (
	"container/list"
	"fmt"
	"sort"
)

type Node struct {
	symbol    byte
	wordIndex int
	parent    *Node
	suffixRef *Node
	finalRef  *Node
	nextNodes map[byte]*Node
}

type AhoCorasick struct {
	root  *Node
	words []string
}

func NewNode(symbol byte, parent *Node) *Node {
	return &Node{
		symbol:    symbol,
		wordIndex: -1,
		parent:    parent,
		nextNodes: make(map[byte]*Node),
	}
}

func NewAhoCorasick() *AhoCorasick {
	root := NewNode(0, nil)
	root.suffixRef = root
	root.parent = root
	return &AhoCorasick{
		root:  root,
		words: []string{},
	}
}

func (n *Node) FindNextNode(symbol byte) *Node {
	return n.nextNodes[symbol]
}

func (n *Node) IsFinal() bool {
	return n.wordIndex >= 0
}

func (ac *AhoCorasick) AddString(patt string) {
	curNode := ac.root

	for i := 0; i < len(patt); i++ {
		symbol := patt[i]
		nextNode := curNode.FindNextNode(symbol)

		if nextNode == nil {
			nextNode = NewNode(symbol, curNode)
			curNode.nextNodes[symbol] = nextNode
		}

		curNode = nextNode
	}

	curNode.wordIndex = len(ac.words)
	ac.words = append(ac.words, patt)
}

func (ac *AhoCorasick) Init() {
	ac.createRefs()
}

func (ac *AhoCorasick) createRefs() {
	queue := list.New()

	for _, nextNode := range ac.root.nextNodes {
		nextNode.suffixRef = ac.root
		queue.PushBack(nextNode)
	}

	for queue.Len() > 0 {
		tempNode := queue.Remove(queue.Front()).(*Node)

		for _, nextNode := range tempNode.nextNodes {
			suffixRefSet := false
			nodeForSuffixRef := tempNode

			for nodeForSuffixRef != ac.root {
				nodeForSuffixRef = nodeForSuffixRef.suffixRef
				foundNode := nodeForSuffixRef.FindNextNode(nextNode.symbol)

				if foundNode != nil {
					if !suffixRefSet {
						suffixRefSet = true
						nextNode.suffixRef = foundNode
					}

					if foundNode.IsFinal() {
						nextNode.finalRef = foundNode
						break
					}
				}
			}

			if !suffixRefSet {
				nextNode.suffixRef = ac.root
			}

			queue.PushBack(nextNode)
		}
	}
}

func (ac *AhoCorasick) Search(text string) [][2]int {
	var ans [][2]int

	curNode := ac.root

	for i := 0; i < len(text); i++ {
		textSymbol := text[i]
		nextNode := curNode.FindNextNode(textSymbol)

		for nextNode == nil {
			if curNode == ac.root {
				nextNode = ac.root
				break
			}
			curNode = curNode.suffixRef
			nextNode = curNode.FindNextNode(textSymbol)
		}

		curNode = nextNode

		for nextNode != ac.root {
			if nextNode.IsFinal() {
				ans = append(ans, [2]int{i - len(ac.words[nextNode.wordIndex]) + 2, nextNode.wordIndex + 1})
			}
			nextNode = nextNode.suffixRef
		}
	}

	sort.Slice(ans, func(i, j int) bool {
		if ans[i][0] == ans[j][0] {
			return ans[i][1] < ans[j][1]
		}
		return ans[i][0] < ans[j][0]
	})

	return ans
}

func main() {
	ahoCorasick := NewAhoCorasick()

	var temp string
	var pattAmount int
	var patt string

	fmt.Scan(&temp)
	fmt.Scan(&pattAmount)

	for i := 0; i < pattAmount; i++ {
		fmt.Scan(&patt)
		ahoCorasick.AddString(patt)
	}

	ahoCorasick.Init()

	ans := ahoCorasick.Search(temp)

	for _, pair := range ans {
		fmt.Println(pair[0], pair[1])
	}
}
