package astnormalization

import (
	"github.com/wundergraph/graphql-go-tools/pkg/ast"
	"github.com/wundergraph/graphql-go-tools/pkg/astvisitor"
)

func removeMergedTypeExtensions(walker *astvisitor.Walker) {
	visitor := removeMergedTypeExtensionsVisitor{
		Walker: walker,
	}
	walker.RegisterLeaveDocumentVisitor(&visitor)
}

type removeMergedTypeExtensionsVisitor struct {
	*astvisitor.Walker
}

func (r *removeMergedTypeExtensionsVisitor) LeaveDocument(operation, definition *ast.Document) {
	operation.RemoveMergedTypeExtensions()
}
