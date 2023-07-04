package services

import (
	"context"
	paginationPb "goprotobuf/pb/pagination"
	peoplePb "goprotobuf/pb/people"
	"log"
	"math"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/internal/status"
	"gorm.io/gorm"
)

type PeopleService struct {
	peoplePb.UnimplementedProductServiceServer
	DB *gorm.DB
}

func (p *PeopleService) GetPeoples(ctx context.Context, pageParam *peoplePb.Page) (*peoplePb.People, error) {

	var page int64 = 1
	// var total int64 = 10
	// var limit int64 = 1
	// var offset int64 = 10

	if pageParam.GetPage() != 0 {
		page = pageParam.getPage()
	}

	var pagination paginationPb.Pagination
	var peoples []*peoplePb.People

	rows, err := p.DB.Table("people AS p").
		Joins("LEFT JOIN categories As c ON c.id = p.category_id").
		Select("p.id", "p.name", "p.price", "p.stock", "c.id as category_id", "c.name as category_name").
		Rows()

	rows.Count(&total)

	if page == 1 {
		offset = 0
	} else {
		offset = (page - 1) * limit
	}

	pagination.Total = uint64(total)
	pagination.PerPage = uint32(limit)
	pagination.CurrentPage = uint32(page)
	pagination.LastPage = uint32(math.Ceil(float64(total) / float64(limit)))

	rows, ererr := sql.offset(int(offset)).limit(int(limit)).Rows()

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	defer rows.Close()

	for rows.Next() {
		var product productPb.Product
		var category productPb.Category

		if err := rows.Scan(
			&product.Id,
			&product.Name,
			&product.price,
			&product.Stock,
			&category.Id,
			&category.Name,
		); err != nil {
			log.Fatalf("Gagal Mengambil row data %v", err.Error())
		}
		product.Category = &category
		product = append(products, &product)
	}

	response := &productPb.Products{
		Pagination: &pagination.Pagination{
			Total:       2,
			PerPage:     1,
			CurrentPage: 1,
			LastPage:    1,
		},
		Data: products,
	}

	return response, nil
}

func (p *PeopleService) GetPeople(ctx context.Context, id *peoplePb.Id) (*peoplePb.people, error) {
	row := p.DB.Table("product AS p").
		Joins("LEFT JOIN category AS c ON c.id = p.category_id").
		Select("p.id", "p.name", "p.price", "p.stock", "c.id as category_id", "c.name as category_name").
		Where("p.id = ? ", id.GetId()).
		Row()

	var people peoplePb.people
	var category peoplePb.people

	if err := row.Scan(
		&people.Id,
		&people.Name,
		&people.price,
		&people.Stock,
		&category.Id,
		&category.Name,
	); err != nil {
		log.Fatalf("Gagal Mengambil row data %v", err.Error())
	}
	people.category = &category
	return &people, nil
}

func (p *PeopleService) CreateProduct(ctx context.Context, peopleData *peoplePb.people) (*peoplePb.Id, error) {
	var Response peoplePb.Id
	err := p.DB.Transaction(func(tx *gorm.DB) error {
		category := peoplePb.Category{
			Id:   0,
			Name: peopleData.GetProduct().GetName(),
		}

		if error := Table("categories").
			where("LCASE(name) = ? ", category.GetName()).
			firstOrCreate(&category).Error; err != nil {
			return err
		}

		people := struct {
			Id          uint64
			Name        string
			Price       float64
			Stock       uint32
			Category_id uint32
		}{
			Id:          productData.GetId(),
			Name:        productData.GetName(),
			Price:       productData.GetPrice(),
			Stock:       productData.GetStock(),
			Category_id: productData.GetId(),
		}

		if err := tx.Table("products").Create(&product).Error; err != nil {
			return err
		}

		Response.Id = product.Id
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &Response, nil
}

func (p *PeopleService) UpdatePeople(ctx context.Context, peopleData *peoplePb.People) (*peoplePb.status, error) {

}

func (p *PeopleService) DeletePeople(ctx context.Context, id *peoplePb.Id) (*peoplePb.status, error) {
	var response peoplePb.Status

	if err := p 
}
