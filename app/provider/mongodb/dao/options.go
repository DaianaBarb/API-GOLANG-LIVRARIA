package provider

import "github.com/americanas-go/config"

type Options struct {
	LivroCollection string
}

func NewOptions() (*Options, error) {
	o := &Options{}

	err := config.UnmarshalWithPath(rootconfig, &o)
	if err != nil {
		return nil, err
	}

	return o, nil
}
