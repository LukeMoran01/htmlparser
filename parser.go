package htmlparse

import (
	"fmt"
	"io"
	"log"
	"os"

	"golang.org/x/net/html"
)

type Link struct {
	Href string
	Text string
}

type Stack[T any] []T

func (stack *Stack[T]) Push(value T) {
	*stack = append(*stack, value)
}

func (stack *Stack[T]) Pop() T {
	lastIndex := len(*stack) - 1
	value := (*stack)[lastIndex]
	*stack = (*stack)[:lastIndex]
	return value
}

func (stack *Stack[_]) IsEmpty() bool {
	return len(*stack) == 0
}

func HtmlToReader(filename string) io.Reader {
	reader, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	return reader
}

func actOnNode(node *html.Node) Link {
	link := Link{}
	for _, value := range node.Attr {
		if value.Key == "href" {
			link.Href = value.Val
		}
	}
	for n := node.FirstChild; n != nil; n = n.NextSibling {
		if n.Type == 1 {
			link.Text += n.Data
		}
		for m := n.FirstChild; m != nil; m = m.FirstChild {
			if m.Type == 1 {
				link.Text += m.Data
			}
		}
	}
	return link
}

// Current capacity of 100 links
func Dfs(node *html.Node) []Link {
	visited := make(map[*html.Node]struct{})
	nodeStack := Stack[*html.Node]{}
	links := make([]Link, 0, 100)
	nodeStack.Push(node)
	for !nodeStack.IsEmpty() {
		currentNode := nodeStack.Pop()
		fmt.Println(currentNode.Data)
		// NON-DFS Actions
		if currentNode.FirstChild != nil && currentNode.Data == "a" {
			links = append(links, actOnNode(currentNode))
		}
		// BACK TO DFS Actions
		_, found := visited[currentNode]
		if !found {
			visited[currentNode] = struct{}{}
			for adjacentNode := currentNode.FirstChild; adjacentNode != nil; adjacentNode = adjacentNode.NextSibling {
				nodeStack.Push(adjacentNode)
			}
		}
	}
	return links
}
