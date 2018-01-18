package test

import (
	"testing"
	"github.com/nbio/st"
)

func TestHotelListRequest_Ok(t *testing.T)  {
	testRequest("hotellist_response_example.xml", false)
	defer gockOff()

	data, err := acApi.HotelListRequest(0,68,0,0,"","")
	er := getCustomErrorType()
	st.Expect(t, err, er)
	st.Expect(t, len(data.Country), 1)
	st.Expect(t, data.Country[0].Code, 9)
	st.Expect(t, data.Country[0].Name, "Россия")
	st.Expect(t, data.Country[0].Position.Latitude, "57.5158")
	st.Expect(t, data.Country[0].Position.Longitude, "82.6172")
	st.Expect(t, len(*data.Country[0].Cities), 1)
	st.Expect(t, (*data.Country[0].Cities)[0].Code, 2)
	st.Expect(t, (*data.Country[0].Cities)[0].Name, "Москва")
	st.Expect(t, (*data.Country[0].Cities)[0].AdmUnit1.Code, 1)
	st.Expect(t, (*data.Country[0].Cities)[0].AdmUnit1.Name, "")
	st.Expect(t, (*data.Country[0].Cities)[0].AdmUnit2.Code, 1)
	st.Expect(t, (*data.Country[0].Cities)[0].AdmUnit2.Name, "Не определено")
	st.Expect(t, (*data.Country[0].Cities)[0].TypeOfPlace.Code, 2)
	st.Expect(t, (*data.Country[0].Cities)[0].TypeOfPlace.Name, "Город")
	st.Expect(t, (*data.Country[0].Cities)[0].Position.Latitude, "55.758959")
	st.Expect(t, (*data.Country[0].Cities)[0].Position.Longitude, "37.616272")
	st.Expect(t, len(*(*data.Country[0].Cities)[0].Hotels), 1)
	st.Expect(t, (*(*data.Country[0].Cities)[0].Hotels)[0].Code, 900376)
	st.Expect(t, (*(*data.Country[0].Cities)[0].Hotels)[0].Name, "Петрово-Дальнее Пансионат")
	st.Expect(t, (*(*data.Country[0].Cities)[0].Hotels)[0].Zip, "143422")
	st.Expect(t, (*(*data.Country[0].Cities)[0].Hotels)[0].Address, "Московская область, Красногорский район, Петрово-Дальнее село")
	st.Expect(t, (*(*data.Country[0].Cities)[0].Hotels)[0].Underground, "Тушинская - Таганско-Краснопресненская линия - 9,5 км")
	st.Expect(t, (*(*data.Country[0].Cities)[0].Hotels)[0].CityCentre, "Центр города - 20 км")
	st.Expect(t, (*(*data.Country[0].Cities)[0].Hotels)[0].Description, "Музей-усадьба \"Архангельское\", Парк Победы на Поклонной Горе")
	st.Expect(t, (*(*data.Country[0].Cities)[0].Hotels)[0].Url, "http://images.acase.ru/hotels_images/900376_00.jpg")
	st.Expect(t, (*(*data.Country[0].Cities)[0].Hotels)[0].Position.Latitude, "55.746274")
	st.Expect(t, (*(*data.Country[0].Cities)[0].Hotels)[0].Position.Longitude, "37.147111")
	st.Expect(t, (*(*data.Country[0].Cities)[0].Hotels)[0].Rating.Code, 99)
	st.Expect(t, (*(*data.Country[0].Cities)[0].Hotels)[0].Rating.Name, "-")
	st.Expect(t, (*(*data.Country[0].Cities)[0].Hotels)[0].Type.Code, 2)
	st.Expect(t, (*(*data.Country[0].Cities)[0].Hotels)[0].Type.Name, "Гостиница")
	st.Expect(t, (*(*data.Country[0].Cities)[0].Hotels)[0].Stars.Code, 3)
	st.Expect(t, (*(*data.Country[0].Cities)[0].Hotels)[0].Stars.Name, "Уровень 3 из 5")
	st.Expect(t, (*(*data.Country[0].Cities)[0].Hotels)[0].Stars.Value, "3")
	st.Expect(t, (*(*data.Country[0].Cities)[0].Hotels)[0].Stars.Url, "http://images.acase.ru/signs_images/Level 3 grey.svg")
	st.Expect(t, (*(*data.Country[0].Cities)[0].Hotels)[0].Stars.OfficialCertificate.Code, 2)
	st.Expect(t, (*(*data.Country[0].Cities)[0].Hotels)[0].Stars.OfficialCertificate.Name, "Государственный")
	st.Expect(t, (*(*data.Country[0].Cities)[0].Hotels)[0].Stars.OfficialCertificate.Certificate.Code, 2)
	st.Expect(t, (*(*data.Country[0].Cities)[0].Hotels)[0].Stars.OfficialCertificate.Certificate.Name, "4 звезда")
	st.Expect(t, (*(*data.Country[0].Cities)[0].Hotels)[0].Stars.OfficialCertificate.Certificate.Url, "http://images.acase.ru/signs_images/four_stars_hotel.gif")
	st.Expect(t, (*(*data.Country[0].Cities)[0].Hotels)[0].Stars.OfficialCertificate.Certificate.ValidBefore, "27.12.2018")
	st.Expect(t, (*(*data.Country[0].Cities)[0].Hotels)[0].FreeSale.Code, 2)
	st.Expect(t, (*(*data.Country[0].Cities)[0].Hotels)[0].FreeSale.Name, "Нет")
	st.Expect(t, len((*(*data.Country[0].Cities)[0].Hotels)[0].Amenities.Amenity), 29)
	st.Expect(t, (*(*data.Country[0].Cities)[0].Hotels)[0].Amenities.Amenity[0].Code, 1)
	st.Expect(t, (*(*data.Country[0].Cities)[0].Hotels)[0].Amenities.Amenity[1].Code, 2)
	st.Expect(t, (*(*data.Country[0].Cities)[0].Hotels)[0].Amenities.Amenity[2].Code, 69)

	st.Expect(t, len(data.AlternativeCountry), 1)
	st.Expect(t, data.AlternativeCountry[0].Code, 9)
	st.Expect(t, data.AlternativeCountry[0].Name, "Россия")
	st.Expect(t, data.AlternativeCountry[0].Position.Latitude, "57.5158")
	st.Expect(t, data.AlternativeCountry[0].Position.Longitude, "82.6172")
	st.Expect(t, len(*data.AlternativeCountry[0].Cities), 1)
	st.Expect(t, (*data.AlternativeCountry[0].Cities)[0].Code, 313)
	st.Expect(t, (*data.AlternativeCountry[0].Cities)[0].Name, "Krasnogorsk")
	st.Expect(t, (*data.AlternativeCountry[0].Cities)[0].AdmUnit1.Code, 1)
	st.Expect(t, (*data.AlternativeCountry[0].Cities)[0].AdmUnit1.Name, "")
	st.Expect(t, (*data.AlternativeCountry[0].Cities)[0].AdmUnit2.Code, 1)
	st.Expect(t, (*data.AlternativeCountry[0].Cities)[0].AdmUnit2.Name, "Не определено")
	st.Expect(t, (*data.AlternativeCountry[0].Cities)[0].TypeOfPlace.Code, 2)
	st.Expect(t, (*data.AlternativeCountry[0].Cities)[0].TypeOfPlace.Name, "город")
	st.Expect(t, (*data.AlternativeCountry[0].Cities)[0].Position.Latitude, "55.758959")
	st.Expect(t, (*data.AlternativeCountry[0].Cities)[0].Position.Longitude, "37.616272")
	st.Expect(t, len(*(*data.AlternativeCountry[0].Cities)[0].Hotels), 1)
	st.Expect(t, (*(*data.AlternativeCountry[0].Cities)[0].Hotels)[0].Code, 900376)
	st.Expect(t, (*(*data.AlternativeCountry[0].Cities)[0].Hotels)[0].Name, "Петрово-Дальнее Пансионат")
	st.Expect(t, (*(*data.AlternativeCountry[0].Cities)[0].Hotels)[0].Zip, "143422")
	st.Expect(t, (*(*data.AlternativeCountry[0].Cities)[0].Hotels)[0].Address, "Московская область, Красногорский район, Петрово-Дальнее село")
	st.Expect(t, (*(*data.AlternativeCountry[0].Cities)[0].Hotels)[0].Underground, "Тушинская - Таганско-Краснопресненская линия - 9,5 км")
	st.Expect(t, (*(*data.AlternativeCountry[0].Cities)[0].Hotels)[0].CityCentre, "Центр города - 20 км")
	st.Expect(t, (*(*data.AlternativeCountry[0].Cities)[0].Hotels)[0].Description, "Музей-усадьба \"Архангельское\", Парк Победы на Поклонной Горе")
	st.Expect(t, (*(*data.AlternativeCountry[0].Cities)[0].Hotels)[0].Url, "http://images.acase.ru/hotels_images/900376_00.jpg")
	st.Expect(t, (*(*data.AlternativeCountry[0].Cities)[0].Hotels)[0].Position.Latitude, "55.746274")
	st.Expect(t, (*(*data.AlternativeCountry[0].Cities)[0].Hotels)[0].Position.Longitude, "37.147111")
	st.Expect(t, (*(*data.AlternativeCountry[0].Cities)[0].Hotels)[0].Rating.Code, 99)
	st.Expect(t, (*(*data.AlternativeCountry[0].Cities)[0].Hotels)[0].Rating.Name, "-")
	st.Expect(t, (*(*data.AlternativeCountry[0].Cities)[0].Hotels)[0].Type.Code, 2)
	st.Expect(t, (*(*data.AlternativeCountry[0].Cities)[0].Hotels)[0].Type.Name, "Гостиница")
	st.Expect(t, (*(*data.AlternativeCountry[0].Cities)[0].Hotels)[0].Stars.Code, 3)
	st.Expect(t, (*(*data.AlternativeCountry[0].Cities)[0].Hotels)[0].Stars.Name, "Уровень 3 из 5")
	st.Expect(t, (*(*data.AlternativeCountry[0].Cities)[0].Hotels)[0].Stars.Value, "3")
	st.Expect(t, (*(*data.AlternativeCountry[0].Cities)[0].Hotels)[0].Stars.Url, "http://images.acase.ru/signs_images/Level 3 grey.svg")
	st.Expect(t, (*(*data.AlternativeCountry[0].Cities)[0].Hotels)[0].Stars.OfficialCertificate.Code, 1)
	st.Expect(t, (*(*data.AlternativeCountry[0].Cities)[0].Hotels)[0].Stars.OfficialCertificate.Name, "Нет")
	st.Expect(t, (*(*data.AlternativeCountry[0].Cities)[0].Hotels)[0].Stars.OfficialCertificate.Certificate.Code, 5)
	st.Expect(t, (*(*data.AlternativeCountry[0].Cities)[0].Hotels)[0].Stars.OfficialCertificate.Certificate.Name, "1 звезда")
	st.Expect(t, (*(*data.AlternativeCountry[0].Cities)[0].Hotels)[0].Stars.OfficialCertificate.Certificate.Url, "http://images.acase.ru/signs_images/one_star_hotel.gif")
	st.Expect(t, (*(*data.AlternativeCountry[0].Cities)[0].Hotels)[0].FreeSale.Code, 2)
	st.Expect(t, (*(*data.AlternativeCountry[0].Cities)[0].Hotels)[0].FreeSale.Name, "Нет")
	st.Expect(t, len((*(*data.AlternativeCountry[0].Cities)[0].Hotels)[0].Amenities.Amenity), 29)
	st.Expect(t, (*(*data.AlternativeCountry[0].Cities)[0].Hotels)[0].Amenities.Amenity[0].Code, 1)
	st.Expect(t, (*(*data.AlternativeCountry[0].Cities)[0].Hotels)[0].Amenities.Amenity[1].Code, 2)
	st.Expect(t, (*(*data.AlternativeCountry[0].Cities)[0].Hotels)[0].Amenities.Amenity[2].Code, 69)

}

func TestHotelListRequest_Error(t *testing.T) {
	testRequest("hotellist_error_example.xml", true)
	defer gockOff()

	_, err := acApi.HotelListRequest(0,0,0,0,"","")

	st.Expect(t, err.Code, "9998")
	st.Expect(t, err.Message, "Доступ запрещен !")
}

