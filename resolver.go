package main

import (
	"context"
	"strings"
)

// Resolver is struct which has fields defined in the schema file.
type VegetableResolver struct {
	v *Vegetable
}

func (r *VegetableResolver) Name() string {
	return r.v.name
}

func (r *VegetableResolver) Price() int32 {
	return int32(r.v.price)
}

func (r *VegetableResolver) Image() *string {
	return r.v.image
}

func (q *query) Vegetable(ctx context.Context, args struct{ Name string }) *VegetableResolver {
	v, ok := vegetables[strings.ToLower(args.Name)]
	if ok {
		return &VegetableResolver{v: &v}
	}
	return nil
}
