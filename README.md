# weather-check

Weather-Check is a webserver application that accepts JSON lat and long values to return live weather information from the closest weather station.

GO:
- built on Go (Golang) version 1.15.7

Uses:
- [OpenWeatherMap] for weather information (openweathermap.org)
- [Chi] for router and middleware (github.com/go-chi/chi)

Usage Example:

Sent at Post Request on Port 3000 as JSON, 
{
	"lat": -33.86,
	"long": 151.20
}

Returns,
{
    "name": "Sydney",
    "sys": {
        "country": "AU"
    },
    "coord": {
        "lat": -33.8679,
        "lon": 151.2073
    },
    "weather": [
        {
            "main": "Clear",
            "description": "clear sky"
        }
    ],
    "main": {
        "temp": 294.87,
        "temp_min": 293.3,
        "temp_max": 296.93
    },
    "wind": {
        "speed": 0.89,
        "deg": 335,
        "gust": 4.92
    },
    "cod": 200
}

Performance:
- Calculates and returns weather information as JSON on average 1.23 seconds with 50up/20down connection speeds

Roadmap:
- Add authentication via JWT
- Run as Docker for cloud deploy

ToDo:


Other Ideas:
- look up what inormation google maps gives ? traffic as indicator of density? density - urban density?
- look at map data
    - analyse elevation and undulating terrian
    - find api to indicate urban density, noise floor
    - 
- add boat information, 
- add flight path id
- add disaster alert, bush fires
- low alt aircraft