package models

import (
	
)


type Bill struct {

 ID  uint `json:"id" gorm:"primary_key"`
 Title  string `json:"title"`
 Author  string `json:"author"`
 Amount  uint `json:"amount"`
 Date   string `json:"date"`
}

type CreateBillInput struct {
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
	Amount  uint `json:"amount"`
    Date   string `json:"date"`
  }