# GoNetatmo

Simple goloang client to retrieve data from the netatmo weather station via [api.netatmo.com](https://dev.netatmo.com/resources/technical/reference)

## Get started

1. clone
```
$git clone git@github.com:mariusbreivik/GoNetatmo.git
```

2. install project dependencies:
```
$dep ensure
```

3. create .env file in the project directory containing the following:
```
CLIENT_ID=<The client ID from your app at https://dev.netatmo.com/myaccount/>
NETATMO_USERNAME=<your username @ netatmo>
NETATMO_PASSWD=<your password @ netatmo>
CLIENT_SECRET=<your client secret for your app @ netatmo>
```

4. run
```
$go run main.go
```


