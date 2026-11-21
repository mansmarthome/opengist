package render

import (
	"bytes"

	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/text"
)

type headingIDTransformer struct{}

func (t *headingIDTransformer) Transform(node *ast.Document, _ text.Reader, _ parser.Context) {
	_ = ast.Walk(node, func(n ast.Node, entering bool) (ast.WalkStatus, error) {
		if !entering {
			return ast.WalkContinue, nil
		}
		heading, ok := n.(*ast.Heading)
		if !ok || heading.Attributes() == nil {
			return ast.WalkContinue, nil
		}
		value, hasID := heading.AttributeString("id")
		heading.RemoveAttributes()
		id, ok := headingIDBytes(value)
		if hasID && ok && safeHeadingID(id) {
			heading.SetAttribute([]byte("id"), id)
		}
		return ast.WalkContinue, nil
	})
}

func headingIDBytes(value any) ([]byte, bool) {
	switch typed := value.(type) {
	case []byte:
		return typed, len(typed) > 0
	case string:
		return []byte(typed), typed != ""
	default:
		return nil, false
	}
}

func safeHeadingID(id []byte) bool {
	if len(id) == 0 || len(id) > 128 || bytes.IndexByte(id, ' ') >= 0 {
		return false
	}
	for _, c := range id {
		switch {
		case c >= 'a' && c <= 'z', c >= 'A' && c <= 'Z', c >= '0' && c <= '9':
		case c == '-' || c == '_' || c == ':' || c == '.':
		default:
			return false
		}
	}
	return true
}
