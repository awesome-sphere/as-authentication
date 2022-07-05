package serializer

type BookingSerializer struct {
	MovieName   string `json:"movie_name" binding:"required"`
	Username    string `json:"username" binding:"required"`
	TheaterName string `json:"theater_name" binding:"required"`
	Date        string `json:"date" binding:"required"`
	Time        string `json:"time" binding:"required"`
	SeatNumber  string `json:"seat_number" binding:"required"`
}
