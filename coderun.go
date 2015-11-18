package coderun

type CodeStorageEnum int

const (
	CodeStorageInline CodeStorageEnum = 1
	CodeStorageFile   CodeStorageEnum = 2
)

type CodeType struct {
	Id          string
	CompileTool string
	BuildTool   string
	ExecTool    string
}

type CodeModule struct {
	Id         string
	CodeTypeId string
	Path       string
}

type CodeBase struct {
	Id          string
	CodeStorage CodeStorageEnum
	Txt         string
	Title       string
}

type CodeOp struct {
	Op      string
	Command string
	Params  []interface{}
}

type Code struct {
	Id         string
	Title      string
	CodeTypeId string
	Modules    []*CodeModule
	CodeBases  []*CodeBase

	// ops will be used for compile, build and exec
	// if nil for each of those, it will use standard tool from CodeType
	Ops map[string]*CodeOp
}

func (c *Code) CodeType() *CodeType {
	return nil
}

func (c *Code) Compile() error {
	return nil
}

func (c *Code) Build() error {
	return nil
}

func (c *Code) Exec() (interface{}, error) {
	return nil
}
