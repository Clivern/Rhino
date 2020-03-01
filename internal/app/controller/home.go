// Copyright 2020 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package controller

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// Index controller
func Index(c *gin.Context) {

	homeTpl := `<!DOCTYPE html>
<html>
    <head>
        <meta charset="utf-8">
        <meta name="viewport" content="width=device-width, initial-scale=1">
        <title>Rhino</title>
        <link rel="icon" type="image/png" href="https://raw.githubusercontent.com/clivern/Rhino/master/assets/img/gopher.png?v=0.1.1">
        <link href="https://fonts.googleapis.com/css?family=Nunito:200,600" rel="stylesheet" type="text/css">
        <link rel="stylesheet" href="https://fonts.googleapis.com/css?family=Roboto:300,300italic,700,700italic">
        <link rel="stylesheet" href="https://cdn.rawgit.com/necolas/normalize.css/master/normalize.css">
        <link rel="stylesheet" href="https://cdn.rawgit.com/milligram/milligram/master/dist/milligram.min.css">
        <style type="text/css" media="screen">
         .wrapper {
             display: block;
             overflow: hidden;
             position: relative;
             width: 100%
        }
         .wrapper .container {
             max-width: 110rem
        }
         .wrapper>.container {
             padding-bottom: 7.5rem;
             padding-top: 7.5rem
        }
         .header {
             background-color: #f4f5f6;
             padding-top: 1rem
        }
         @media(min-width:40rem) {
             .header {
                 padding-top: 5rem
            }
        }
         .header+section {
             border-top: 0
        }
         .header .container {
             border-top: 0;
             padding-bottom: 7.5rem;
             padding-top: 7.5rem;
             position: relative;
             text-align: center
        }
         .header .title {
             font-family: 'Nunito', sans-serif;
        }
         .header .img {
             height: 15rem;
             margin-bottom: 2rem
        }
         .header .img path {
             animation: 7s a forwards;
             fill: #9b4dca;
             stroke: #9b4dca;
             stroke-dasharray: 38321;
             stroke-miterlimit: 10;
             stroke-width: 15px
        }
         .header .button {
             margin-bottom: 4rem;
             margin-top: 2rem
        }
         @media(min-width:40rem) {
             .header .button {
                 margin-bottom: 4rem;
                 margin-top: 2rem
            }
        }
         @keyframes a {
             0% {
                 fill-opacity: 0;
                 stroke-dashoffset: 38321
            }
             25% {
                 fill-opacity: 0;
                 stroke: #9b4dca
            }
             to {
                 fill-opacity: 1;
                 stroke-dashoffset: 0
            }
        }
         .navigation {
             background: #f4f5f6;
             display: block;
             height: 5.2rem;
             left: 0;
             max-width: 100%;
             position: fixed;
             right: 0;
             top: 0;
             width: 100%;
             z-index: 1
        }
         .navigation .container {
             padding-bottom: 0;
             padding-top: 0
        }
         .navigation .navigation-list {
             list-style: none;
             margin-bottom: 0;
             margin-right: 5rem
        }
         @media(min-width:80rem) {
             .navigation .navigation-list {
                 margin-right: 0
            }
        }
         .navigation .navigation-item {
             float: left;
             margin-bottom: 0;
             margin-left: 2.5rem;
             position: relative
        }
         .navigation .img {
             fill: #9b4dca;
             height: 2rem;
             position: relative;
             top: .3rem
        }
         .navigation .navigation-title, .navigation .title {
             color: #606c76;
             font-family: 'Nunito', sans-serif;
             position: relative
        }
         .navigation .navigation-link, .navigation .navigation-title, .navigation .title {
             display: inline;
             font-size: 1.6rem;
             line-height: 5.2rem;
             padding: 0;
             text-decoration: none
        }
         .navigation .navigation-link.active {
             color: #606c76
        }
         .octocat {
             border: 0;
             color: #f4f5f6;
             fill: #9b4dca;
             height: 5.2rem;
             position: fixed;
             right: 0;
             top: 0;
             width: 5.2rem;
             z-index: 1
        }
         .octocat:hover .octocat-arm {
             animation: b .56s infinite;
             transform-origin: 13rem 10.6rem
        }
         .octocat .octocat-arm, .octocat .octocat-body {
             fill: #f4f5f6
        }
         @keyframes b {
             0%, 50% {
                 transform: rotate(0)
            }
             25%, 75% {
                 transform: rotate(-25deg)
            }
        }
         @media(min-width:40rem) {
             .only-mobile {
                 display: none
            }
        }
         .prettyprint.code {
             border: 0;
             border-left: .3rem solid #9b4dca;
             color: #655d5d
        }
         .prettyprint.code .str {
             color: #4b8b8b
        }
         .prettyprint.code .kwd {
             color: #8464c4
        }
         .prettyprint.code .com {
             color: #adadad
        }
         .prettyprint.code .typ {
             color: #7272ca
        }
         .prettyprint.code .lit {
             color: #9b4dca
        }
         .prettyprint.code .pun {
             color: #5485b6
        }
         .prettyprint.code .clo, .prettyprint.code .opn {
             color: #f4ecec
        }
         .prettyprint.code .atn, .prettyprint.code .tag {
             color: #9b4dca
        }
         .prettyprint.code .atv {
             color: #5485b6
        }
         .prettyprint.code .dec {
             color: #b45a3c
        }
         .prettyprint.code .var {
             color: #ca4949
        }
         .prettyprint.code .fun {
             color: #7272ca
        }
         .prettyprint.code.lang-md * {
             color: #655d5d
        }
        </style>
    </head>
    <body>
        <main class="wrapper" id="app">
            <nav class="navigation">
                <section class="container">
                    <a href="https://github.com/Clivern/Rhino" title="Rhino on Github" target="_blank">
                        <svg class="octocat" viewBox="0 0 250 250">
                            <path d="M0,0 L115,115 L130,115 L142,142 L250,250 L250,0 Z"></path>
                            <path class="octocat-arm" d="M128.3,109.0 C113.8,99.7 119.0,89.6 119.0,89.6 C122.0,82.7 120.5,78.6 120.5,78.6 C119.2,72.0 123.4,76.3 123.4,76.3 C127.3,80.9 125.5,87.3 125.5,87.3 C122.9,97.6 130.6,101.9 134.4,103.2"></path>
                            <path class="octocat-body" d="M115.0,115.0 C114.9,115.1 118.7,116.5 119.8,115.4 L133.7,101.6 C136.9,99.2 139.9,98.4 142.2,98.6 C133.8,88.0 127.5,74.4 143.8,58.0 C148.5,53.4 154.0,51.2 159.7,51.0 C160.3,49.4 163.2,43.6 171.4,40.1 C171.4,40.1 176.1,42.5 178.8,56.2 C183.1,58.6 187.2,61.8 190.9,65.4 C194.5,69.0 197.7,73.2 200.1,77.6 C213.8,80.2 216.3,84.9 216.3,84.9 C212.7,93.1 206.9,96.0 205.4,96.6 C205.1,102.4 203.0,107.8 198.3,112.5 C181.9,128.9 168.3,122.5 157.7,114.1 C157.9,116.9 156.7,120.9 152.7,124.9 L141.0,136.5 C139.8,137.7 141.6,141.9 141.8,141.8 Z"></path>
                        </svg>
                    </a>
                </section>
            </nav>
            <header class="header" id="home">
                <section class="container">
                    <img src="https://raw.githubusercontent.com/clivern/Rhino/master/assets/img/gopher.png?v=0.1.1" width="15%">
                    <h1 class="title">{{ name }}</h1>
                    <p class="description">{{ description }}</p>
                </section>
            </header>
            <template v-for="request in requests">
                <section class="container">
                    <pre class="code prettyprint prettyprinted" style=""><code class="code-content">
Status: {{ request.status }}
Route: {{ request.route }}
URI: {{ request.uri }}
Time: {{ request.time }}
Method: {{ request.method }}
Headers: {{ request.headers }}
StatusCode: {{ request.statusCode }}
Body: {{ request.body }}
                    </code></pre>
                </section>
            </template>
            <section class="container" id="contributing">
                <h3 class="title">Contributing</h3>
                <p>Want to contribute? Follow these<a href="https://github.com/Clivern/Rhino/blob/master/CONTRIBUTING.md" title="Contributing"> Recommendations</a>.</p>
            </section>
            <footer class="footer">
                <section class="container">
                    <p>Crafted with ♥ by<a href="http://clivern.com" title="Clivern" target="_blank"> Clivern</a>. Licensed under the<a href="https://github.com/Clivern/Rhino#license" title="MIT License" target="_blank"> MIT License</a>.</p>
                </section>
            </footer>
        </main>
        <script src="https://unpkg.com/vue@2.6.10/dist/vue.min.js"></script>
        <script src="https://unpkg.com/axios@0.19.0/dist/axios.min.js"></script>
        <script type="text/javascript" charset="utf-8">
            var app = new Vue({
                el: '#app',
                data: {
                    name: "Rhino",
                    description: "HTTP Mocking & Debugging Service.",
                    requests: []
                },
                methods: {

                },
                mounted () {
                    axios
                      .get('APP_PROJECTS_URL')
                      .then(response => (this.requests = response.data))

                }
            })
        </script>
    </body>
</html>
`

	homeTpl = strings.Replace(
		homeTpl,
		"APP_PROJECTS_URL",
		fmt.Sprintf("%s/api/requests", strings.TrimSuffix(viper.GetString("app.domain"), "/")),
        -1,
	)
	c.Writer.WriteHeader(http.StatusOK)
	c.Writer.Write([]byte(homeTpl))
	return
}
