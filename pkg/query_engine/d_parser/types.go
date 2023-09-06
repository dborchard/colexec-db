package parser

// Statement generic interface
type Statement interface {
	stmt()
	travel()
}

// SelectStatement Select statement
type SelectStatement struct {
	Fields         Fields
	Sources        Sources
	WhereCondition *ExprNode
}

func (*SelectStatement) stmt() {}

func (stmt *SelectStatement) travel() {

}

type Field struct {
	Name string
}
type Fields []*Field

type Source struct {
	Name string
}
type Sources []*Source

type ExprNodeType int

const (
	Unknown     ExprNodeType = 0
	BinaryNode  ExprNodeType = 1
	FieldNode   ExprNodeType = 2
	IntegerNode ExprNodeType = 3
	FloatNode   ExprNodeType = 4
	StringNode  ExprNodeType = 5
)

type ExprNode struct {
	Type  ExprNodeType
	Left  *ExprNode
	Right *ExprNode

	FloVal float64
	IntVal int64
	StrVal string
	Name   string
	Op     int
}