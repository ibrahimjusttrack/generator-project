package generator

type Input struct {
	Key   string `json:"accessor"`
	Value string `json:"value"`
}

type Service interface {
	//return path for file and error of generator
	Generate(templateId string, settings []Input) (string, error)
}
