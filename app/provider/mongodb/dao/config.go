package provider

import "github.com/americanas-go/config"

const (
	appconfig       = "app"
	rootconfig      = appconfig + ".provider.database.mongo"
	LivroCollection = rootconfig + ".livroCollection"
)

func init() {

	config.Add(LivroCollection, "livro", "define the collection of mongodb")

}
