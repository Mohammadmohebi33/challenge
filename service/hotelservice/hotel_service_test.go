package hotelservice

import (
	"context"
	"errors"
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
	hotelService := NewHotelService(repo)

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

	resp, err := hotelService.Create(ctx, req)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, expectedHotel.ID, resp.ID)
	assert.Equal(t, expectedHotel.Name, resp.Name)
	assert.Equal(t, expectedHotel.Location, resp.Location)
	assert.Equal(t, expectedHotel.Rating, resp.Rating)
	repo.AssertExpectations(t)
}

func TestHotelService_Create_Error(t *testing.T) {
	// Arrange
	repo := new(MockHotelRepository)
	hotelService := NewHotelService(repo)

	req := params.HotelCreateRequest{
		Name:     "Test Hotel",
		Location: "Test Location",
		Rating:   4.5,
	}

	repo.On("Insert", mock.Anything, mock.Anything).Return(entity.Hotel{}, errors.New("insert error"))

	ctx := context.Background()

	resp, err := hotelService.Create(ctx, req)

	// Assert
	assert.Error(t, err)
	assert.Equal(t, params.HotelCreateResponse{}, resp)
	repo.AssertExpectations(t)
}

func TestHotelService_GetHotelById(t *testing.T) {

	repo := new(MockHotelRepository)
	hotelService := NewHotelService(repo)

	id := "test-id"
	expectedHotel := entity.Hotel{
		ID:       primitive.NewObjectID(),
		Name:     "Test Hotel",
		Location: "Test Location",
		Rating:   4.5,
	}

	repo.On("GetHotelByID", mock.Anything, id).Return(expectedHotel, nil)

	ctx := context.Background()

	// Act
	hotel, err := hotelService.GetHotelById(ctx, id)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, expectedHotel, hotel)
	repo.AssertExpectations(t)
}

func TestHotelService_GetHotelById_Error(t *testing.T) {
	// Arrange
	repo := new(MockHotelRepository)
	hotelService := NewHotelService(repo)

	id := "test-id"

	repo.On("GetHotelByID", mock.Anything, id).Return(entity.Hotel{}, errors.New("not found"))

	ctx := context.Background()

	// Act
	hotel, err := hotelService.GetHotelById(ctx, id)

	// Assert
	assert.Error(t, err)
	assert.Equal(t, entity.Hotel{}, hotel)
	repo.AssertExpectations(t)
}

func TestHotelService_GetAllHotel(t *testing.T) {
	// Arrange
	repo := new(MockHotelRepository)
	hotelService := NewHotelService(repo)

	expectedHotels := []entity.Hotel{
		{
			ID:       primitive.NewObjectID(),
			Name:     "Hotel 1",
			Location: "Location 1",
			Rating:   4.5,
		},
		{
			ID:       primitive.NewObjectID(),
			Name:     "Hotel 2",
			Location: "Location 2",
			Rating:   4.0,
		},
	}

	repo.On("GetAllHotels", mock.Anything).Return(expectedHotels, nil)

	ctx := context.Background()

	// Act
	hotels, err := hotelService.GetAllHotel(ctx)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, expectedHotels, hotels)
	repo.AssertExpectations(t)
}

func TestHotelService_GetAllHotel_Error(t *testing.T) {
	// Arrange
	repo := new(MockHotelRepository)
	hotelService := NewHotelService(repo)

	repo.On("GetAllHotels", mock.Anything).Return([]entity.Hotel{}, errors.New("database error"))

	ctx := context.Background()

	hotels, err := hotelService.GetAllHotel(ctx)

	assert.Error(t, err)
	assert.Empty(t, hotels)
	repo.AssertExpectations(t)
}
