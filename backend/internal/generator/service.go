package generator

type Input struct {
	Key   string
	Value string
}

type Service interface {
	//return path for file and error of generator
	Generate(templateId string, input []Input) (string, error)
}
