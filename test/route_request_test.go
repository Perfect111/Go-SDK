package test

import (
	"testing"
	"gopkg.in/h2non/gock.v1"
	"github.com/nbio/st"
)

func TestRoute1Request_Ok(t *testing.T) {
	testRequest("route1_response_example.xml")
	defer gock.Off()

	data, err := acApi.RouteRequest("Моск","",0,0,0,0)
	er := getCustomErrorType()
	st.Expect(t, err, er)
	st.Expect(t, data.Action.Name, "LIST")
	st.Expect(t, data.Action.Parameters.StartPoint.Name, "Моск")
	st.Expect(t, len(data.TransferPointList.TransferPoint) , 4)
	st.Expect(t, data.TransferPointList.TransferPoint[0].Code, 2)
	st.Expect(t, data.TransferPointList.TransferPoint[0].TransferPointType.Code, 1)
	st.Expect(t, data.TransferPointList.TransferPoint[0].TransferPointType.Name, "Город")
	st.Expect(t, data.TransferPointList.TransferPoint[0].Country.Code, 9)
	st.Expect(t, data.TransferPointList.TransferPoint[0].Country.Name, "Россия")
	st.Expect(t, data.TransferPointList.TransferPoint[0].AdmUnit1.Code, 1)
	st.Expect(t, data.TransferPointList.TransferPoint[0].AdmUnit1.Name, "Не определено")
	st.Expect(t, data.TransferPointList.TransferPoint[0].AdmUnit2.Code, 1)
	st.Expect(t, data.TransferPointList.TransferPoint[0].AdmUnit2.Name, "Не определено")
	st.Expect(t, data.TransferPointList.TransferPoint[0].TypeOfPlace.Code, 2)
	st.Expect(t, data.TransferPointList.TransferPoint[0].TypeOfPlace.Name, "Город")
	st.Expect(t, data.TransferPointList.TransferPoint[0].City.Code, 2)
	st.Expect(t, data.TransferPointList.TransferPoint[0].City.Name, "Москва")
	st.Expect(t, data.TransferPointList.TransferPoint[1].Code, 100115)
	st.Expect(t, data.TransferPointList.TransferPoint[1].TransferPointType.Code, 2)
	st.Expect(t, data.TransferPointList.TransferPoint[1].TransferPointType.Name, "Гостиница")
	st.Expect(t, data.TransferPointList.TransferPoint[1].Country.Code, 9)
	st.Expect(t, data.TransferPointList.TransferPoint[1].Country.Name, "Россия")
	st.Expect(t, data.TransferPointList.TransferPoint[1].AdmUnit1.Code, 23)
	st.Expect(t, data.TransferPointList.TransferPoint[1].AdmUnit1.Name, "Калининградская область")
	st.Expect(t, data.TransferPointList.TransferPoint[1].AdmUnit2.Code, 1)
	st.Expect(t, data.TransferPointList.TransferPoint[1].AdmUnit2.Name, "Не определено")
	st.Expect(t, data.TransferPointList.TransferPoint[1].TypeOfPlace.Code, 2)
	st.Expect(t, data.TransferPointList.TransferPoint[1].TypeOfPlace.Name, "Город")
	st.Expect(t, data.TransferPointList.TransferPoint[1].City.Code, 19)
	st.Expect(t, data.TransferPointList.TransferPoint[1].City.Name, "Калининград")
	st.Expect(t, data.TransferPointList.TransferPoint[1].Hotel.Code, 100115)
	st.Expect(t, data.TransferPointList.TransferPoint[1].Hotel.Name, "Москва")
	st.Expect(t, data.TransferPointList.TransferPoint[2].Code, 1000411)
	st.Expect(t, data.TransferPointList.TransferPoint[2].TransferPointType.Code, 4)
	st.Expect(t, data.TransferPointList.TransferPoint[2].TransferPointType.Name, "Ж/д вокзалы")
	st.Expect(t, data.TransferPointList.TransferPoint[2].Country.Code, 9)
	st.Expect(t, data.TransferPointList.TransferPoint[2].Country.Name, "Россия")
	st.Expect(t, data.TransferPointList.TransferPoint[2].AdmUnit1.Code, 1)
	st.Expect(t, data.TransferPointList.TransferPoint[2].AdmUnit1.Name, "Не определено")
	st.Expect(t, data.TransferPointList.TransferPoint[2].AdmUnit2.Code, 1)
	st.Expect(t, data.TransferPointList.TransferPoint[2].AdmUnit2.Name, "Не определено")
	st.Expect(t, data.TransferPointList.TransferPoint[2].TypeOfPlace.Code, 2)
	st.Expect(t, data.TransferPointList.TransferPoint[2].TypeOfPlace.Name, "Город")
	st.Expect(t, data.TransferPointList.TransferPoint[2].City.Code, 3)
	st.Expect(t, data.TransferPointList.TransferPoint[2].City.Name, "Санкт-Петербург")
	st.Expect(t, data.TransferPointList.TransferPoint[2].Object.Code, 1000411)
	st.Expect(t, data.TransferPointList.TransferPoint[2].Object.Name, "Московский")
	st.Expect(t, data.TransferPointList.TransferPoint[2].ObjectSubType.Code, 9500001)
	st.Expect(t, data.TransferPointList.TransferPoint[2].ObjectSubType.Name, "Не имеет значения")
	st.Expect(t, data.TransferPointList.TransferPoint[2].ObjectType.Code, 1000039)
	st.Expect(t, data.TransferPointList.TransferPoint[2].ObjectType.Name, "Ж/д вокзалы")
	st.Expect(t, data.TransferPointList.TransferPoint[3].Code, 1000391)
	st.Expect(t, data.TransferPointList.TransferPoint[3].TransferPointType.Code, 4)
	st.Expect(t, data.TransferPointList.TransferPoint[3].TransferPointType.Name, "Ж/д вокзалы")
	st.Expect(t, data.TransferPointList.TransferPoint[3].Country.Code, 9)
	st.Expect(t, data.TransferPointList.TransferPoint[3].Country.Name, "Россия")
	st.Expect(t, data.TransferPointList.TransferPoint[3].AdmUnit1.Code, 26)
	st.Expect(t, data.TransferPointList.TransferPoint[3].AdmUnit1.Name, "Нижегородская область")
	st.Expect(t, data.TransferPointList.TransferPoint[3].AdmUnit2.Code, 1)
	st.Expect(t, data.TransferPointList.TransferPoint[3].AdmUnit2.Name, "Не определено")
	st.Expect(t, data.TransferPointList.TransferPoint[3].TypeOfPlace.Code, 2)
	st.Expect(t, data.TransferPointList.TransferPoint[3].TypeOfPlace.Name, "Город")
	st.Expect(t, data.TransferPointList.TransferPoint[3].City.Code, 20)
	st.Expect(t, data.TransferPointList.TransferPoint[3].City.Name, "Нижний Новгород")
	st.Expect(t, data.TransferPointList.TransferPoint[3].Object.Code, 1000391)
	st.Expect(t, data.TransferPointList.TransferPoint[3].Object.Name, "Нижний Новгород-Московский")
	st.Expect(t, data.TransferPointList.TransferPoint[3].ObjectSubType.Code, 9500001)
	st.Expect(t, data.TransferPointList.TransferPoint[3].ObjectSubType.Name, "Не имеет значения")
	st.Expect(t, data.TransferPointList.TransferPoint[3].ObjectType.Code, 1000039)
	st.Expect(t, data.TransferPointList.TransferPoint[3].ObjectType.Name, "Ж/д вокзалы")
}

func TestRoute4Request_Ok(t *testing.T) {
	testRequest("route4_response_example.xml")
	defer gock.Off()

	data, err := acApi.RouteRequest("Шере","",0,2,0,1)
	er := getCustomErrorType()
	st.Expect(t, err, er)
	st.Expect(t, data.Action.Name, "LIST")
	st.Expect(t, data.Action.Parameters.StartPoint.Name, "Шере")
	st.Expect(t, data.Action.Parameters.EndPoint.Code, 2)
	st.Expect(t, data.Action.Parameters.EndPoint.Type, 1)
	st.Expect(t, len(data.TransferPointList.TransferPoint) , 5)
	st.Expect(t, data.TransferPointList.TransferPoint[0].Code, 1200089)
	st.Expect(t, data.TransferPointList.TransferPoint[0].TransferPointType.Code, 3)
	st.Expect(t, data.TransferPointList.TransferPoint[0].TransferPointType.Name, "Аэропорт")
	st.Expect(t, data.TransferPointList.TransferPoint[0].Country.Code, 9)
	st.Expect(t, data.TransferPointList.TransferPoint[0].Country.Name, "Россия")
	st.Expect(t, data.TransferPointList.TransferPoint[0].AdmUnit1.Code, 1)
	st.Expect(t, data.TransferPointList.TransferPoint[0].AdmUnit1.Name, "Не определено")
	st.Expect(t, data.TransferPointList.TransferPoint[0].AdmUnit2.Code, 1)
	st.Expect(t, data.TransferPointList.TransferPoint[0].AdmUnit2.Name, "Не определено")
	st.Expect(t, data.TransferPointList.TransferPoint[0].TypeOfPlace.Code, 2)
	st.Expect(t, data.TransferPointList.TransferPoint[0].TypeOfPlace.Name, "Город")
	st.Expect(t, data.TransferPointList.TransferPoint[0].City.Code, 2)
	st.Expect(t, data.TransferPointList.TransferPoint[0].City.Name, "Москва")
	st.Expect(t, data.TransferPointList.TransferPoint[0].Object.Code, 1200089)
	st.Expect(t, data.TransferPointList.TransferPoint[0].Object.Name, "Шереметьево B")
	st.Expect(t, data.TransferPointList.TransferPoint[0].ObjectSubType.Code, 9500001)
	st.Expect(t, data.TransferPointList.TransferPoint[0].ObjectSubType.Name, "Не имеет значения")
	st.Expect(t, data.TransferPointList.TransferPoint[0].ObjectType.Code, 800001)
	st.Expect(t, data.TransferPointList.TransferPoint[0].ObjectType.Name, "Аэропорт")
	st.Expect(t, data.TransferPointList.TransferPoint[1].Code, 1200090)
	st.Expect(t, data.TransferPointList.TransferPoint[1].TransferPointType.Code, 3)
	st.Expect(t, data.TransferPointList.TransferPoint[1].TransferPointType.Name, "Аэропорт")
	st.Expect(t, data.TransferPointList.TransferPoint[1].Country.Code, 9)
	st.Expect(t, data.TransferPointList.TransferPoint[1].Country.Name, "Россия")
	st.Expect(t, data.TransferPointList.TransferPoint[1].AdmUnit1.Code, 1)
	st.Expect(t, data.TransferPointList.TransferPoint[1].AdmUnit1.Name, "Не определено")
	st.Expect(t, data.TransferPointList.TransferPoint[1].AdmUnit2.Code, 1)
	st.Expect(t, data.TransferPointList.TransferPoint[1].AdmUnit2.Name, "Не определено")
	st.Expect(t, data.TransferPointList.TransferPoint[1].TypeOfPlace.Code, 2)
	st.Expect(t, data.TransferPointList.TransferPoint[1].TypeOfPlace.Name, "Город")
	st.Expect(t, data.TransferPointList.TransferPoint[1].City.Code, 2)
	st.Expect(t, data.TransferPointList.TransferPoint[1].City.Name, "Москва")
	st.Expect(t, data.TransferPointList.TransferPoint[1].Object.Code, 1200090)
	st.Expect(t, data.TransferPointList.TransferPoint[1].Object.Name, "Шереметьево C")
	st.Expect(t, data.TransferPointList.TransferPoint[1].ObjectSubType.Code, 9500001)
	st.Expect(t, data.TransferPointList.TransferPoint[1].ObjectSubType.Name, "Не имеет значения")
	st.Expect(t, data.TransferPointList.TransferPoint[1].ObjectType.Code, 800001)
	st.Expect(t, data.TransferPointList.TransferPoint[1].ObjectType.Name, "Аэропорт")
	st.Expect(t, data.TransferPointList.TransferPoint[2].Code, 1200091)
	st.Expect(t, data.TransferPointList.TransferPoint[2].TransferPointType.Code, 3)
	st.Expect(t, data.TransferPointList.TransferPoint[2].TransferPointType.Name, "Аэропорт")
	st.Expect(t, data.TransferPointList.TransferPoint[2].Country.Code, 9)
	st.Expect(t, data.TransferPointList.TransferPoint[2].Country.Name, "Россия")
	st.Expect(t, data.TransferPointList.TransferPoint[2].AdmUnit1.Code, 1)
	st.Expect(t, data.TransferPointList.TransferPoint[2].AdmUnit1.Name, "Не определено")
	st.Expect(t, data.TransferPointList.TransferPoint[2].AdmUnit2.Code, 1)
	st.Expect(t, data.TransferPointList.TransferPoint[2].AdmUnit2.Name, "Не определено")
	st.Expect(t, data.TransferPointList.TransferPoint[2].TypeOfPlace.Code, 2)
	st.Expect(t, data.TransferPointList.TransferPoint[2].TypeOfPlace.Name, "Город")
	st.Expect(t, data.TransferPointList.TransferPoint[2].City.Code, 2)
	st.Expect(t, data.TransferPointList.TransferPoint[2].City.Name, "Москва")
	st.Expect(t, data.TransferPointList.TransferPoint[2].Object.Code, 1200091)
	st.Expect(t, data.TransferPointList.TransferPoint[2].Object.Name, "Шереметьево D")
	st.Expect(t, data.TransferPointList.TransferPoint[2].ObjectSubType.Code, 9500001)
	st.Expect(t, data.TransferPointList.TransferPoint[2].ObjectSubType.Name, "Не имеет значения")
	st.Expect(t, data.TransferPointList.TransferPoint[2].ObjectType.Code, 800001)
	st.Expect(t, data.TransferPointList.TransferPoint[2].ObjectType.Name, "Аэропорт")
	st.Expect(t, data.TransferPointList.TransferPoint[3].Code, 1200092)
	st.Expect(t, data.TransferPointList.TransferPoint[3].TransferPointType.Code, 3)
	st.Expect(t, data.TransferPointList.TransferPoint[3].TransferPointType.Name, "Аэропорт")
	st.Expect(t, data.TransferPointList.TransferPoint[3].Country.Code, 9)
	st.Expect(t, data.TransferPointList.TransferPoint[3].Country.Name, "Россия")
	st.Expect(t, data.TransferPointList.TransferPoint[3].AdmUnit1.Code, 1)
	st.Expect(t, data.TransferPointList.TransferPoint[3].AdmUnit1.Name, "Не определено")
	st.Expect(t, data.TransferPointList.TransferPoint[3].AdmUnit2.Code, 1)
	st.Expect(t, data.TransferPointList.TransferPoint[3].AdmUnit2.Name, "Не определено")
	st.Expect(t, data.TransferPointList.TransferPoint[3].TypeOfPlace.Code, 2)
	st.Expect(t, data.TransferPointList.TransferPoint[3].TypeOfPlace.Name, "Город")
	st.Expect(t, data.TransferPointList.TransferPoint[3].City.Code, 2)
	st.Expect(t, data.TransferPointList.TransferPoint[3].City.Name, "Москва")
	st.Expect(t, data.TransferPointList.TransferPoint[3].Object.Code, 1200092)
	st.Expect(t, data.TransferPointList.TransferPoint[3].Object.Name, "Шереметьево E")
	st.Expect(t, data.TransferPointList.TransferPoint[3].ObjectSubType.Code, 9500001)
	st.Expect(t, data.TransferPointList.TransferPoint[3].ObjectSubType.Name, "Не имеет значения")
	st.Expect(t, data.TransferPointList.TransferPoint[3].ObjectType.Code, 800001)
	st.Expect(t, data.TransferPointList.TransferPoint[3].ObjectType.Name, "Аэропорт")
	st.Expect(t, data.TransferPointList.TransferPoint[4].Code, 1200093)
	st.Expect(t, data.TransferPointList.TransferPoint[4].TransferPointType.Code, 3)
	st.Expect(t, data.TransferPointList.TransferPoint[4].TransferPointType.Name, "Аэропорт")
	st.Expect(t, data.TransferPointList.TransferPoint[4].Country.Code, 9)
	st.Expect(t, data.TransferPointList.TransferPoint[4].Country.Name, "Россия")
	st.Expect(t, data.TransferPointList.TransferPoint[4].AdmUnit1.Code, 1)
	st.Expect(t, data.TransferPointList.TransferPoint[4].AdmUnit1.Name, "Не определено")
	st.Expect(t, data.TransferPointList.TransferPoint[4].AdmUnit2.Code, 1)
	st.Expect(t, data.TransferPointList.TransferPoint[4].AdmUnit2.Name, "Не определено")
	st.Expect(t, data.TransferPointList.TransferPoint[4].TypeOfPlace.Code, 2)
	st.Expect(t, data.TransferPointList.TransferPoint[4].TypeOfPlace.Name, "Город")
	st.Expect(t, data.TransferPointList.TransferPoint[4].City.Code, 2)
	st.Expect(t, data.TransferPointList.TransferPoint[4].City.Name, "Москва")
	st.Expect(t, data.TransferPointList.TransferPoint[4].Object.Code, 1200093)
	st.Expect(t, data.TransferPointList.TransferPoint[4].Object.Name, "Шереметьево F")
	st.Expect(t, data.TransferPointList.TransferPoint[4].ObjectSubType.Code, 9500001)
	st.Expect(t, data.TransferPointList.TransferPoint[4].ObjectSubType.Name, "Не имеет значения")
	st.Expect(t, data.TransferPointList.TransferPoint[4].ObjectType.Code, 800001)
	st.Expect(t, data.TransferPointList.TransferPoint[4].ObjectType.Name, "Аэропорт")
}

func TestRouteRequest_Error(t *testing.T) {
	testRequest("route_error_example.xml")
	defer gock.Off()

	_, err := acApi.RouteRequest("","",0,0,0,0)

	st.Expect(t, err.Code, "9998")
	st.Expect(t, err.Message, "Доступ запрещен !")
}

