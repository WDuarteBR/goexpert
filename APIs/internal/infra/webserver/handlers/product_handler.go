package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"github.com/wduartebr/goexpert/apis/internal/dto"
	"github.com/wduartebr/goexpert/apis/internal/entity"
	"github.com/wduartebr/goexpert/apis/internal/infra/database"
	pkg_entity "github.com/wduartebr/goexpert/apis/pkg/entity"
)

type ProductHandler struct {
	ProductDB database.ProductInterface
}

func NewProductHandler(db database.ProductInterface) *ProductHandler {
	return &ProductHandler{
		ProductDB: db,
	}
}

// CreateProduct godoc
// @Summary 		Create Product
// @Description 	Create Product
// @Tags 			products
// @Accept 			json
// @Produce 		json
// @Param 			request 	body	dto.CreateProductInput	true	"product request"
// @Success			201
// @Failure			500			{object}	Error
// @Router			/product	[post]
// @Security 		ApiKeyAuth
func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product dto.CreateProductInput
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	p, err := entity.NewProduct(product.Name, product.Price)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.ProductDB.CreateProduct(p)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// AllProducts godoc
// @Summary 		List Products
// @Description 	List Product
// @Tags 			products
// @Accept 			json
// @Produce 		json
// @Param 			page	query	string	false	"page number"
// @Param 			limit	query	string 	false	"limit"
// @Success			200			{array}		entity.Product
// @Failure			404			{object}	Error
// @Failure			500			{object}	Error
// @Router			/product/all	[get]
// @Security 		ApiKeyAuth
func (h *ProductHandler) AllProducts(w http.ResponseWriter, r *http.Request) {
	page := r.URL.Query().Get("page")
	intPage, err := strconv.Atoi(page)
	if err != nil {
		intPage = 0
	}

	limit := r.URL.Query().Get("limit")
	intLimit, err := strconv.Atoi(limit)
	if err != nil {
		intLimit = 0
	}

	sort := r.URL.Query().Get("sort")

	products, err := h.ProductDB.FindAll(intPage, intLimit, sort)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&products)

}

// GetProducts godoc
// @Summary 		Get a products
// @Description 	Get a products
// @Tags 			products
// @Accept 			json
// @Produce 		json
// @Param 			id		path 	string	true	"product ID" Format(uuid)
// @Success			200			{object}	entity.Product
// @Failure			404
// @Failure			500			{object}	Error
// @Router			/product/{id}	[get]
// @Security 		ApiKeyAuth
func (h *ProductHandler) GetProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	product, err := h.ProductDB.FindById(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(product)
}

// UpdateProduct godoc
// @Summary 		Update a  product
// @Description 	Update a product
// @Tags 			products
// @Accept 			json
// @Produce 		json
// @Param 			id		path 	string	true	"product ID" Format(uuid)
// @Param 			request 	body	dto.CreateProductInput	true	"product request"
// @Success			200
// @Failure			404
// @Failure			500			{object}	Error
// @Router			/product/{id}	[put]
// @Security 		ApiKeyAuth
func (h *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var product entity.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	product.ID, err = pkg_entity.ParseID(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	_, err = h.ProductDB.FindById(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}

	err = h.ProductDB.UpdateProduct(&product)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

}

// DeleteProducts godoc
// @Summary 		Delete a product
// @Description 	Delete a products
// @Tags 			products
// @Accept 			json
// @Produce 		json
// @Param 			id		path 	string	true	"product ID" Format(uuid)
// @Success			200
// @Failure			404
// @Failure			500			{object}	Error
// @Router			/product/{id}	[delete]
// @Security 		ApiKeyAuth
func (h *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err := pkg_entity.ParseID(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	err = h.ProductDB.DeleteProduct(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

}
