package main

import "time"

type DatabaseConnection struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	Due       time.Time `json:"due"`
}

type DatabaseConnections []DatabaseConnection
