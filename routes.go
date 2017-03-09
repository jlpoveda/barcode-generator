package main

import "net/http"

type Route struct {
    Name        string
    Method      string
    Pattern     string
    HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
    Route{
        "Qr",
        "GET",
        "/qr/{code}",
        GetQr,
    },
    Route{
        "Ean",
        "GET",
        "/ean/{code}",
        GetEan,
    },
    Route{
        "Code128",
        "GET",
        "/code128/{code}",
        GetCode128,
    },
    Route{
        "Code39",
        "GET",
        "/code39/{code}",
        GetCode39,
    },
}