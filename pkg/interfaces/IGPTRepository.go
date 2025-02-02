package interfaces

type IGPTRepository interface {
	ChatWithImage(string, string) (*string, error)
	ChatRequest(string) (*string, error)
}
