package generator

type Input struct {
	Key   string
	Value string
}

type Service interface {
	//return path for file and error of generator
	Generate(templateId string, settings []Input) (string, error)
}
