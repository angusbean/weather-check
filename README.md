# weather-check

Weather-Check is a webserver application that accepts JSON lat and long values to return live weather information from the closest weather station.

GO:
- built on Go (Golang) version 1.15.7

Uses:
- [OpenWeatherMap] for weather information (openweathermap.org)
- [Chi] for router and middleware (github.com/go-chi/chi)
- [Go-Redis] for key value store (github.com/go-redis/redis/v7)
- [twinj] for uuid creation (github.com/twinj/uuid)

Requires:
- Redis for as Key Value store

Directions:
1) Create an acount and generate a valid OpenWeatherMap API token
2) Clone the repo
3) Update the values in .env-example and update name to .env
4) Ensure Redis is running with '$ redis-server'
5) Run with './run.sh'

Usage Example:

Send a 'Post' Request on 'localhost:3000/' as JSON, 
{
	"lat": -33.86,
	"long": 151.20
}

Application Returns:
{
    "name": "Ostrovnoy",
    "sys": {
        "country": "RU",
        "sunrise": 1629854503,
        "sunset": 1629912408
    },
    "coord": {
        "lat": 68.0531,
        "long": 0
    },
    "weather": [
        {
            "main": "Clouds",
            "description": "overcast clouds"
        }
    ],
    "main": {
        "temp": 282.16,
        "temp_min": 282.16,
        "temp_max": 282.16,
        "pressure": 1022,
        "humidity": 58
    },
    "visiblity": 0,
    "clouds": {
        "all": 87
    },
    "wind": {
        "speed": 4.79,
        "deg": 75,
        "gust": 3.79
    },
    "cod": 200
}

Performance:
- Calculates and returns weather information as JSON on average 1.2 seconds with 50up/20down connection speeds

Roadmap:
- Add authentication via JWT
- Use db for authentication 
- Run as Docker for cloud deploy
- Add stream via gRPC

Other Ideas:
- add density indicators (traffic? 4G coverage? housing data?)
- ArcGIS urban density calculators
- add low altitude airspace information (flight paths? ADS-B information)