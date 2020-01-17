# ggl-top - Google's Site's Top URLs

A CLI app to retrieve a top indexed links of a site from the Google Search engine results and to export to a CSV file.

Could be used by SEO experts or any people, who are interested by controlling of site's ranking. Especially useful when the site is recreated under new technology, and the redirects should be provided to keep the SEO ranking.

## Installation

    go get github.com/fedir/ggl-top

## Usage example

    ggl-top -s www.example.com
    ggl-top -s www.example.com -c fr -l fr

## Development

    go build
    ./ggl-top -s www.example.com -c fr -l fr -d

## Credentials

* https://github.com/PuerkitoBio/goquery
* https://gist.github.com/EdmundMartin/eaea4aaa5d231078cb433b89878dbecf
