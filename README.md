<p align="center">
    <img alt="Rhino Logo" src="https://raw.githubusercontent.com/clivern/Rhino/master/assets/img/gopher.png?v=1.4.0" width="150" />
    <h3 align="center">Rhino</h3>
    <p align="center">HTTP Mocking & Debugging Service</p>
    <p align="center">
        <a href="https://github.com/Clivern/Rhino/actions"><img src="https://github.com/Clivern/Rhino/workflows/Build/badge.svg"></a>
        <a href="https://github.com/Clivern/Rhino/actions"><img src="https://github.com/Clivern/Rhino/workflows/Release/badge.svg"></a>
        <a href="https://github.com/Clivern/Rhino/releases"><img src="https://img.shields.io/badge/Version-1.4.0-red.svg"></a>
        <a href="https://goreportcard.com/report/github.com/Clivern/Rhino"><img src="https://goreportcard.com/badge/github.com/clivern/Rhino?v=1.4.0"></a>
        <a href="https://hub.docker.com/r/clivern/rhino"><img src="https://img.shields.io/badge/Docker-Latest-green"></a>
        <a href="https://github.com/Clivern/Rhino/blob/master/LICENSE"><img src="https://img.shields.io/badge/LICENSE-MIT-orange.svg"></a>
    </p>
</p>

Rhino is an HTTP Mocking & Debugging Service. It enables easy mocking of any HTTP web service for testing and debugging purposes. Also it can simulate high latencies and failures to make sure your services have the capability to withstand and recover from failures. It supports cross-origin resource sharing (CORS) so it can be used as a backend for single page applications.


## Documentation

### Usage

Get [the latest binary.](https://github.com/Clivern/Rhino/releases)

```zsh
$ curl -sL https://github.com/Clivern/Rhino/releases/download/x.x.x/Rhino_x.x.x_OS_x86_64.tar.gz | tar xz
```

Create the config file `config.prod.json`

```json
{
    "app": {
        "mode": "prod or dev",
        "port": "8080",
        "domain": "http://127.0.0.1:8080",
        "tls": {
            "status": "off",
            "pemPath": "/cert/server.pem",
            "keyPath": "/cert/server.key"
        }
    },
    "mock": [
        {
            "path": "/api/v2/service1/mock/:id",
            "request": {
                "method": "get"
            },
            "response": {
                "statusCode": 200,
                "headers": [
                    {"key": "Content-Type", "value": "application/json"}
                ],
                "body": "{\"id\": \":id\"}"
            },
            "chaos": {
                "latency": "0s",
                "failRate": "0%"
            }
        },
        {
            "path": "/api/v2/service2/mock/:id",
            "request": {
                "method": "get",
                "parameters": {
                    "var_param": ":var_param",
                    "fixed_param": 10
                }
            },
            "response": {
                "statusCode": 200,
                "headers": [
                    {"key": "Content-Type", "value": "application/json"}
                ],
                "body": "@json:@config_dir/route.response.json"
            },
            "chaos": {
                "latency": "0s",
                "failRate": "0%"
            }
        }
    ],
    "debug": [
        {
            "path": "/api/v2/service/debug",
            "chaos": {
                "latency": "0s",
                "failRate": "0%"
            }
        }
    ],
    "log": {
        "level": "info",
        "output": "stdout or /var/log/rhino.log",
        "format": "text or json"
    }
}
```

Run Rhino with that config file

```zsh
$ ./rhino serve -c /custom/path/config.prod.json
```

Check the release.

```zsh
$ ./rhino version
```

Test it.

```zsh
$ curl http://127.0.0.1:8080/_health
```

You can use fake data flags inside response body and rhino will auto generate them. Here is the full list of supported types:

```bash
Latitude: @fake(:lat)
Longitude: @fake(:long)
CreditCardNumber: @fake(:cc_number)
CreditCardType: @fake(:cc_type)
Email: @fake(:email)
DomainName: @fake(:domain_name)
IPV4: @fake(:ipv4)
IPV6: @fake(:ipv6)
Password: @fake(:password)
PhoneNumber: @fake(:phone_number)
MacAddress: @fake(:mac_address)
URL: @fake(:url)
UserName: @fake(:username)
TollFreeNumber: @fake(:toll_free_number)
E164PhoneNumber: @fake(:e_164_phone_number)
TitleMale: @fake(:title_male)
TitleFemale: @fake(:title_female)
FirstName: @fake(:first_name)
FirstNameMale: @fake(:first_name_male)
FirstNameFemale: @fake(:first_name_female)
LastName: @fake(:last_name)
Name: @fake(:name)
UnixTime: @fake(:unix_time)
Date: @fake(:date)
Time: @fake(:time)
MonthName: @fake(:month_name)
Year: @fake(:year)
DayOfWeek: @fake(:day_of_week)
DayOfMonth: @fake(:day_of_month)
Timestamp: @fake(:timestamp)
Century: @fake(:century)
TimeZone: @fake(:timezone)
TimePeriod: @fake(:time_period)
Word: @fake(:word)
Sentence: @fake(:sentence)
Paragraph: @fake(:paragraph)
Currency: @fake(:currency)
Amount: @fake(:amount)
AmountWithCurrency: @fake(:amount_with_currency)
UUIDHypenated: @fake(:uuid_hyphenated)
UUID: @fake(:uuid_digit)
```


### Docker

Clone and then run docker containers.

```zsh
$ git clone https://github.com/Clivern/Rhino.git
$ cd Rhino/deployment/docker-compose
$ docker-compose up -d
```


## Versioning

For transparency into our release cycle and in striving to maintain backward compatibility, Rhino is maintained under the [Semantic Versioning guidelines](https://semver.org/) and release process is predictable and business-friendly.

See the [Releases section of our GitHub project](https://github.com/clivern/rhino/releases) for changelogs for each release version of Rhino. It contains summaries of the most noteworthy changes made in each release.


## Bug tracker

If you have any suggestions, bug reports, or annoyances please report them to our issue tracker at https://github.com/clivern/rhino/issues


## Security Issues

If you discover a security vulnerability within Rhino, please send an email to [hello@clivern.com](mailto:hello@clivern.com)


## Contributing

We are an open source, community-driven project so please feel free to join us. see the [contributing guidelines](CONTRIBUTING.md) for more details.


## License

© 2020, clivern. Released under [MIT License](https://opensource.org/licenses/mit-license.php).

**Rhino** is authored and maintained by [@clivern](http://github.com/clivern).
