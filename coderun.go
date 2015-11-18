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

type Code struct {
	Id          string
	Title       string
	CodeTypeId  string
	CodeStorage CodeStorageEnum
	Txt         string
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
