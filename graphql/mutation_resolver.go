package main

import "context"

type mutationResolver struct {
	server *Server
}

// CreateAccount implements MutationResolver.
func (m *mutationResolver) CreateAccount(ctx context.Context, account AccountInput) (*Account, error) {
	panic("unimplemented")
}

// CreateOrder implements MutationResolver.
func (m *mutationResolver) CreateOrder(ctx context.Context, order OrderInput) (*Order, error) {
	panic("unimplemented")
}

// CreateProduct implements MutationResolver.
func (m *mutationResolver) CreateProduct(ctx context.Context, product ProductInput) (*Product, error) {
	panic("unimplemented")
}

// func (r *mutationResolver) createAccount(ctx context.Context, in AccountInput) (*Account, error) {

// }

// func (r *mutationResolver) createProduct(ctx context.Context, in ProductInput) (*Product, error) {

// }

// func (r *mutationResolver) createOrder(ctx context.Context, in OrderInput) (*Order, error) {

// }
