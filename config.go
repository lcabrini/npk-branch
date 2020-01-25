package main

import (
    "github.com/lcabrini/npk-common"
)

var config = npk.Configuration {
    Database: npk.DatabaseConfig {
        Host: "localhost",
        Port: "5432",
        SSLMode: "disabled",
        Username: "npk",
        Password: "%npK-s3Kr3T%",
    },
}
