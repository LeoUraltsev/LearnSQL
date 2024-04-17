package main

import "time"

type User struct {
	Id       int64
	Login    string
	Email    string
	Password string
}

type Spending struct {
	Id         int64     `db:"id"`
	Price      uint      `db:"price"`
	CreatedAt  time.Time `db:"created_at"`
	UserID     int64     `db:"user_id"`
	CategoryID int64     `db:"categores_id"`
}

type Category struct {
	Id   int64
	Name string
}
