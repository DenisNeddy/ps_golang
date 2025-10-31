package weather

import (
	"demo/weather/geo"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

var ErrorWrongFormal = errors.New("WRONG__FORMAT")

func GetWeather(geo geo.GeoData, format int) (string, error) {
	if format < 1 || format > 4 {
		return "", ErrorWrongFormal
	}
	baseUrl, err := url.Parse("https://wttr.in/" + geo.City)
	if err != nil {
		fmt.Println(err.Error())
		return "", errors.New("ERROR__URL")
	}

	params := url.Values{}
	params.Add("format", fmt.Sprint(format))
	baseUrl.RawQuery = params.Encode()
	resp, err := http.Get(baseUrl.String())
	if err != nil {
		fmt.Println(err.Error())
		return "", errors.New("ERROR__HTTP")
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err.Error())
		return "", errors.New("ERROR__READBODY")
	}
	return string(body), nil
}
