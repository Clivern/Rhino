<p align="center">
    <img alt="Rhino Logo" src="https://raw.githubusercontent.com/clivern/Rhino/master/assets/img/gopher.png?v=0.0.1" width="150" />
    <h3 align="center">Rhino</h3>
    <p align="center">HTTP Mocking & Debugging Service</p>
    <p align="center">
        <a href="https://travis-ci.com/Clivern/Rhino"><img src="https://travis-ci.com/Clivern/Rhino.svg?branch=master"></a>
        <a href="https://github.com/Clivern/Rhino/releases"><img src="https://img.shields.io/badge/Version-0.0.1-red.svg"></a>
        <a href="https://goreportcard.com/report/github.com/Clivern/Rhino"><img src="https://goreportcard.com/badge/github.com/clivern/Rhino?v=0.0.1"></a>
        <a href="https://github.com/Clivern/Rhino/blob/master/LICENSE"><img src="https://img.shields.io/badge/LICENSE-MIT-orange.svg"></a>
    </p>
</p>

Rhino is an HTTP Mocking & Debugging Service. It enables easy mocking of any HTTP web service for testing and debugging purposes.

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
        "mode": "dev",
        "port": "8080",
        "tls": {
            "status": "off",
            "pemPath": "/cert/server.pem",
            "keyPath": "/cert/server.key"
        }
    },
    "mock": [

    ],
    "debug": [

    ]
}
```

Run Rhino with that config file

```zsh
$ ./Rhino --config=/path/to/config.prod.json
```

Check the release.

```zsh
$ ./Rhino --get=release
```

Test it.

```zsh
$ curl http://127.0.0.1:8080/_health
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
