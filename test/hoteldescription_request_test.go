package test

import (
	"context"
	"testing"

	"github.com/nbio/st"
)

func TestHotelDescriptionRequest_Ok(t *testing.T) {
	testRequest("hoteldescription_response_example.xml", false)
	defer gockOff()

	data, err := acApi.HotelDescriptionRequest(context.Background(), 900376, 0)
	er := getCustomErrorType()
	st.Expect(t, err, er)
	st.Expect(t, data.Code, 900376)
	st.Expect(t, data.Name, "Петрово-Дальнее Пансионат")
	st.Expect(t, data.Zip, "143422")
	st.Expect(t, data.Address, "Московская область, Красногорский район, Петрово-Дальнее село")
	st.Expect(t, data.Story, "Пансионат \"Петрово - Дальнее\" - это загородный дом, городская квартира и номер отеля на отдыхе, расположенный в одном из красивейших уголков Рублевки при слиянии р. Истры с Москвой рекой, на территории усадьбы Петровское. Усадьба была любимым подмосковным имением рода Голициных. В окрестностях расположено пять златоглавых церквей, что позволило назвать в народе эти места \"благостными\". Сегодня, дом отдыха \"Петрово - Дальнее\" представляет собой уникальную зону отдыха с комфортными условиями проживания. Здесь есть все, что нужно для работы и досуга. Расположенные на территории пансионата мини-острова: \"Остров любви\", \"Остров Удачи\", \"Березовая роща\" - прекрасное место для проведения частных праздников: VIP свадеб, летних пикников, детских праздников, а также корпоративных мероприятий и team building.\nПриехав отдохнуть в \"Петрово - Дальнее\" пансионат УДП РФ - гости всегда найдут чем заняться. К услугам гостей: рыбалка, конюшня, футбольное поле, открытый и закрытый теннисные корты, тренажерный зал, детская площадка, зимний сад, рестораны, VIP сауна до 15 гостей (в ней имеется: бильярдная комната, турецкий хамам, финская сауна, купель, комнаты для отдыха), SPA центр (в нем имеется: 2 бассейна, турецкий хамам, финская сауна, ультракрасная сауна, соляная комната, 2 купели в бочках, душ впечатлений, солярий).\nЭлитный пансионат УДП \"Петрово - Дальнее\" имеет в своем распоряжении современный спортивный комплекс, который придется по душе всем любителям активного отдыха, а также гостям, заботящимся о своей фигуре и соблюдающие здоровый образ жизни. Уникальный санаторий \"Петрово - Дальнее\" - это роскошный отдых в пределах Подмосковья.\n")
	st.Expect(t, data.Built, "1977")
	st.Expect(t, data.Reconstructed, "2008")
	st.Expect(t, data.NumberOfBlocks, "")
	st.Expect(t, data.NumberOfRooms, 106)
	st.Expect(t, data.NumberOfFloors, 8)
	st.Expect(t, data.CityCentre, "Центр города - 20 км")
	st.Expect(t, data.Underground, "Тушинская - Таганско-Краснопресненская линия - 9,5 км")
	st.Expect(t, data.Description, "Музей-усадьба \"Архангельское\", Парк Победы на Поклонной Горе")
	st.Expect(t, data.Airport, "Внуково - 28 км, Шереметьево - 44 км, Домодедово - 79 км")
	st.Expect(t, data.RailwayStation, "Белорусский - 35 км")
	st.Expect(t, data.RiverPort, "Северный речной порт - 37 км")
	st.Expect(t, data.SeaPort, "")
	st.Expect(t, data.Map, "http://images.acase.ru/maps_images/900376_map.gif")
	st.Expect(t, data.SpecialRequirements, "1. Гостиница имеет право взимать регистрационный сбор или депозит за дополнительные услуги при размещении. Размер регистрационного сбора гостиница устанавливает самостоятельно..")
	st.Expect(t, data.ObjType.Code, 2)
	st.Expect(t, data.ObjType.Name, "Гостиница")
	st.Expect(t, data.City.Code, 2)
	st.Expect(t, data.City.Name, "Москва")
	st.Expect(t, data.City.Position.Latitude, "55.758959")
	st.Expect(t, data.City.Position.Longitude, "37.616272")
	st.Expect(t, len(data.AlternativeCities.AlternativeCity), 1)
	st.Expect(t, data.AlternativeCities.AlternativeCity[0].Code, 313)
	st.Expect(t, data.AlternativeCities.AlternativeCity[0].Name, "Красногорск")
	st.Expect(t, data.AlternativeCities.AlternativeCity[0].Country.Code, 9)
	st.Expect(t, data.AlternativeCities.AlternativeCity[0].Country.Name, "Россия")
	st.Expect(t, data.AlternativeCities.AlternativeCity[0].AdmUnit1.Code, 1)
	st.Expect(t, data.AlternativeCities.AlternativeCity[0].AdmUnit1.Name, "Не определено")
	st.Expect(t, data.AlternativeCities.AlternativeCity[0].AdmUnit2.Code, 1)
	st.Expect(t, data.AlternativeCities.AlternativeCity[0].AdmUnit2.Name, "Не определено")
	st.Expect(t, data.AlternativeCities.AlternativeCity[0].TypeOfPlace.Code, 2)
	st.Expect(t, data.AlternativeCities.AlternativeCity[0].TypeOfPlace.Name, "Город")
	st.Expect(t, data.AdmUnit1.Code, 1)
	st.Expect(t, data.AdmUnit1.Name, "Не определено")
	st.Expect(t, data.AdmUnit2.Code, 1)
	st.Expect(t, data.AdmUnit2.Name, "Не определено")
	st.Expect(t, data.TypeOfPlace.Code, 2)
	st.Expect(t, data.TypeOfPlace.Name, "Город")
	st.Expect(t, data.Country.Code, 9)
	st.Expect(t, data.Country.Name, "Россия")
	st.Expect(t, data.Position.Latitude, "55.746274")
	st.Expect(t, data.Position.Longitude, "37.147111")
	st.Expect(t, data.Rating.Code, 99)
	st.Expect(t, data.Rating.Name, "-")
	st.Expect(t, data.Stars.Code, 3)
	st.Expect(t, data.Stars.Name, "Уровень 3 из 5")
	st.Expect(t, data.Stars.Value, "3")
	st.Expect(t, data.Stars.Url, "http://images.acase.ru/signs_images/Level 3 grey.svg")
	st.Expect(t, data.Stars.OfficialCertificate.Code, 1)
	st.Expect(t, data.Stars.OfficialCertificate.Name, "Нет")
	st.Expect(t, data.Stars.OfficialCertificate.Certificate.Code, 99)
	st.Expect(t, data.Stars.OfficialCertificate.Certificate.Name, "-")
	st.Expect(t, data.Stars.OfficialCertificate.Certificate.Url, "http://images.acase.ru/signs_images/no_stars_hotel.gif")
	st.Expect(t, data.TimeCheck.In, "18:00")
	st.Expect(t, data.TimeCheck.Out, "16:00")
	st.Expect(t, len(data.Images.Image), 5)
	st.Expect(t, data.Images.Image[0].Url, "http://images.acase.ru/hotels_images/900376_00.jpg")
	st.Expect(t, len(data.Amenities.Amenity), 29)
	st.Expect(t, data.Amenities.Amenity[0].Code, 1)
	st.Expect(t, data.Amenities.Amenity[0].Name, "Отель в зоне отдыха")
	st.Expect(t, data.Amenities.Amenity[0].Url, "http://images.acase.ru/signs_images/hotel_in_recreation_area.gif")
	st.Expect(t, len(data.RoomsList.Room), 4)
	st.Expect(t, data.RoomsList.Room[0].Code, 1313292)
	st.Expect(t, data.RoomsList.Room[0].Name, "Студия")
	st.Expect(t, data.RoomsList.Room[0].Description, "Шикарный номер я балконом, гостинной и кухней")
	st.Expect(t, data.RoomsList.Room[0].MaxNumberOfGuests, 2)
	st.Expect(t, data.RoomsList.Room[0].MaxNumberOfExtraBeds, 1)
	st.Expect(t, data.RoomsList.Room[0].MaxNumberOfExtraBedsAdult, 1)
	st.Expect(t, data.RoomsList.Room[0].MaxNumberOfExtraBedsChild, 1)
	st.Expect(t, data.RoomsList.Room[0].MaxNumberOfExtraBedsInfant, 0)
	st.Expect(t, len(data.SpecialRequirementList.SpecialRequirement), 1)
	st.Expect(t, data.SpecialRequirementList.SpecialRequirement[0].Code, 595)
	st.Expect(t, data.SpecialRequirementList.SpecialRequirement[0].Text, "Гостиница имеет право взимать регистрационный сбор или депозит за дополнительные услуги при размещении. Размер регистрационного сбора гостиница устанавливает самостоятельно.")
	st.Expect(t, len(data.Objects.Object), 9)
	st.Expect(t, data.Objects.Object[0].Code, 800004)
	st.Expect(t, data.Objects.Object[0].Name, "Внуково")
	st.Expect(t, data.Objects.Object[0].Distance, 28.0)
	st.Expect(t, data.Objects.Object[0].ObjectType.Code, 800001)
	st.Expect(t, data.Objects.Object[0].ObjectType.Name, "Аэропорт")
	st.Expect(t, data.Objects.Object[0].ObjectSubType.Code, 9500001)
	st.Expect(t, data.Objects.Object[0].ObjectSubType.Name, "Не имеет значения")
	st.Expect(t, data.Objects.Object[0].City.Code, 2)
	st.Expect(t, data.Objects.Object[0].City.Name, "Москва")
}

func TestHotelDescriptionRequest_Error(t *testing.T) {
	testRequest("hoteldescription_error_example.xml", true)
	defer gockOff()

	_, err := acApi.HotelDescriptionRequest(context.Background(), 900376, 0)

	st.Expect(t, err.Code, "9998")
	st.Expect(t, err.Message, "Доступ запрещен !")
}
