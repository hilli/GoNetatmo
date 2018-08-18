# GoNetatmo [![CircleCI](https://circleci.com/gh/mariusbreivik/GoNetatmo/tree/develop.svg?style=svg)](https://circleci.com/gh/mariusbreivik/GoNetatmo/tree/develop)

Simple goloang client to retrieve json data from the netatmo weather station via [api.netatmo.com]

More details:

https://dev.netatmo.com/resources/technical/reference

## Get started

1. clone
```
$ git clone git@github.com:mariusbreivik/GoNetatmo.git
```

2. install project dependencies:
```
$ dep ensure
```

3. You need to create an application at https://dev.netatmo.com/ to be able to get your weather station data fra https://api.netatmo.com  
Create the _.env_ file in the project directory containing the following:
```
CLIENT_ID=<clientID from your app at https://dev.netatmo.com/myaccount/>
NETATMO_USERNAME=<your username @ netatmo>
NETATMO_PASSWD=<your password @ netatmo>
CLIENT_SECRET=<your client secret for your app @ netatmo>
```

4. run
```
$ go run main.go
```


