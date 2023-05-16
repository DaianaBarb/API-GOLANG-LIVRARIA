package response

type LivroResponse struct {
	ID        string `json:"id"`
	Nome      string `json:"name"`
	Autor     string `json:"autor"`
	Editora   string `json:"editora"`
	AnoFabric string `json:"anoFabric"`
}
