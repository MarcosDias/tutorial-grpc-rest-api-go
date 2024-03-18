package server

import (
	"context"

	"github.com/google/uuid"
	v1 "github.com/marcosdias/tutorial-grpc-rest-api-go/gen/product/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	// mutex    sync.Mutex
	products map[string]v1.Product
	v1.UnimplementedProductServiceServer
}

func (server *Server) AddProduct(ctx context.Context, req *v1.AddProductRequest) (*v1.AddProductResponse, error) {
	// server.mutex.Lock()
	// defer server.mutex.Unlock()
	product := v1.Product{
		Id:    uuid.New().String(),
		Name:  req.Name,
		Price: req.Price,
	}

	server.products[product.Id] = product
	// server.mutex.Unlock()

	return &v1.AddProductResponse{ProductId: product.Id}, nil
}

func (server *Server) DeleteProduct(ctx context.Context, req *v1.DeleteProductRequest) (*v1.DeleteProductResponse, error) {
	// server.mutex.Lock()
	// defer server.mutex.Unlock()

	product, ok := server.products[req.ProductId]
	if !ok {
		return nil, status.Errorf(codes.NotFound, "product with ID: %q does not exists", req.ProductId)
	}
	delete(server.products, req.ProductId)

	return &v1.DeleteProductResponse{Product: &product}, nil
}

func (server *Server) ListProducts(ctx context.Context, req *v1.ListProductsRequest) (*v1.ListProductsResponse, error) {
	// server.mutex.Lock()
	// defer server.mutex.Unlock()

	productList := make([]*v1.Product, 0, len(server.products))
	for id := range server.products {
		product := server.products[id]
		productList = append(productList, &product)
	}

	return &v1.ListProductsResponse{Products: productList}, nil
}

func New() *Server {
	return &Server{
		products: make(map[string]v1.Product),
	}
}
