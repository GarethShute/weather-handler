# Weather Handler
## About
A simple microservice to gather a summary of the weather forecast for the week.
This hooks into the accuweather.com API, an account will be required to generate an API key to make use of this app.
Create account here: https://developer.accuweather.com/

## API endpoints
/weather/dayforecast    -   Returns a 5 day forecast
/citylookup/<cityname>  -   Returns location keys as json array of all cities that match the search name including country and region

## Json config format
config.json  
```
{
    "port" : <port_number>      //Optional - Will default to 8080
    "apikey" : <apikey>         //Required - Generate on AccuWeather website
    "areacode" : <area_code>    //Required - Use citylookup endpoint to get code, London can be used for example with the following... "areacode" : "328328"
}
```

## Response format
Example reponse for London location code
```yaml
{"DailyForecasts":[{"MaximumTemp":17.2,"MinimumTemp":7.3,"Unit":"C","IconRef":6,"IconUrl":"https://developer.accuweather.com/sites/default/files/06-s.png","IconPhrase":"Mostly cloudy","RainProbability":2},{"MaximumTemp":18.6,"MinimumTemp":9.2,"Unit":"C","IconRef":4,"IconUrl":"https://developer.accuweather.com/sites/default/files/04-s.png","IconPhrase":"Intermittent clouds","RainProbability":2},{"MaximumTemp":19.7,"MinimumTemp":13,"Unit":"C","IconRef":4,"IconUrl":"https://developer.accuweather.com/sites/default/files/04-s.png","IconPhrase":"Intermittent clouds","RainProbability":25},{"MaximumTemp":20.5,"MinimumTemp":12,"Unit":"C","IconRef":3,"IconUrl":"https://developer.accuweather.com/sites/default/files/03-s.png","IconPhrase":"Partly sunny","RainProbability":4},{"MaximumTemp":21.2,"MinimumTemp":13.2,"Unit":"C","IconRef":6,"IconUrl":"https://developer.accuweather.com/sites/default/files/06-s.png","IconPhrase":"Mostly cloudy","RainProbability":1}]}
```