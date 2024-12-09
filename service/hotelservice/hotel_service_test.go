package hotelservice

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"hotel_with_test/entity"
	"hotel_with_test/params"
	"testing"
)

type MockHotelRepository struct {
	mock.Mock
}

func (m *MockHotelRepository) Insert(ctx context.Context, hotel entity.Hotel) (entity.Hotel, error) {
	args := m.Called(ctx, hotel)
	return args.Get(0).(entity.Hotel), args.Error(1)
}

func (m *MockHotelRepository) GetHotelByID(ctx context.Context, id string) (entity.Hotel, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(entity.Hotel), args.Error(1)
}

func (m *MockHotelRepository) GetAllHotels(ctx context.Context) ([]entity.Hotel, error) {
	args := m.Called(ctx)
	return args.Get(0).([]entity.Hotel), args.Error(1)
}

func TestHotelService_Create(t *testing.T) {
	// Arrange
	repo := new(MockHotelRepository)
	hotelService := HotelService{
		repo: repo,
	}

	req := params.HotelCreateRequest{
		Name:     "Test Hotel",
		Location: "Test Location",
		Rating:   4.5,
	}

	expectedHotel := entity.Hotel{
		ID:       primitive.NewObjectID(),
		Name:     req.Name,
		Location: req.Location,
		Rating:   req.Rating,
	}

	repo.On("Insert", mock.Anything, mock.MatchedBy(func(hotel entity.Hotel) bool {
		return hotel.Name == req.Name && hotel.Location == req.Location && hotel.Rating == req.Rating
	})).Return(expectedHotel, nil)

	ctx := context.Background()

	// Act
	resp, err := hotelService.Create(ctx, req)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, expectedHotel.ID, resp.ID)
	assert.Equal(t, expectedHotel.Name, resp.Name)
	assert.Equal(t, expectedHotel.Location, resp.Location)
	assert.Equal(t, expectedHotel.Rating, resp.Rating)
	repo.AssertExpectations(t)
}
