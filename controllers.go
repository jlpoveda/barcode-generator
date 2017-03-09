package main

import (
    "net/http"

    "github.com/boombuler/barcode"
    "github.com/boombuler/barcode/qr"
    "log"
    "image/png"
    "bytes"
    "strconv"
    "github.com/gorilla/mux"
    "github.com/boombuler/barcode/ean"
    "github.com/boombuler/barcode/code128"
    "github.com/boombuler/barcode/code39"
)

func GetQr(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)

    code, err := qr.Encode(vars["code"], qr.L, qr.Unicode)
    if err != nil {
        log.Fatal(err)
    }
    code = scaleCode(code, 300, 300)
    sendImage(w, code)
}

func GetEan(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)

    code, err := ean.Encode(vars["code"])
    if err != nil {
        log.Fatal(err)
    }
    var c barcode.Barcode
    c = scaleCode(code, 300, 300)
    sendImage(w, c)
}

func GetCode128(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)

    code, err := code128.Encode(vars["code"])
    if err != nil {
        log.Fatal(err)
    }
    var c barcode.Barcode
    c = scaleCode(code, 300, 300)
    sendImage(w, c)
}

func GetCode39(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)

    code, err := code39.Encode(vars["code"], false, false)
    if err != nil {
        log.Fatal(err)
    }
    var c barcode.Barcode
    c = scaleCode(code, 300, 300)
    sendImage(w, c)
}

func scaleCode(code barcode.Barcode, w int, h int) barcode.Barcode {
    code, err := barcode.Scale(code, w, h)
    if err != nil {
        log.Fatal(err)
    }

    return code
}

func sendImage(w http.ResponseWriter, qrcode barcode.Barcode) {
    buffer := new(bytes.Buffer)
    png.Encode(buffer, qrcode)

    w.Header().Set("Content-Type", "image/png")
    w.Header().Set("Content-Length", strconv.Itoa(len(buffer.Bytes())))
    if _, err := w.Write(buffer.Bytes()); err != nil {
        log.Println("unable to write image.")
    }
}