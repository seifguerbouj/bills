package controllers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/seifguerbouj/bills/models"
)

// @Summary Find all bills
// @Description Get all bills
// @Tags bills
// @Accept json
// @Produce json
// @Success 200 {object} []models.Bill
// @Router /bills [get]
func FindBills(c *gin.Context) {
	var bills []models.Bill
	models.DB.Find(&bills)
	c.JSON(http.StatusOK, gin.H{"data": bills})
}

// @Summary Create a bill
// @Description Create a new bill
// @Tags bills
// @Accept json
// @Produce json
// @Param input body models.CreateBillInput true "Bill input object"
// @Success 200 {object} models.Bill
// @Router /bill [post]
func CreateBill(c *gin.Context) {
	var input models.CreateBillInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	bill := models.Bill{Title: input.Title, Author: input.Author, Amount: input.Amount, Date: input.Date}
	models.DB.Create(&bill)
	c.JSON(http.StatusOK, gin.H{"data": bill})
}

// @Summary Calculate all expenses
// @Description Get total expenses
// @Tags expenses
// @Accept json
// @Produce json
// @Success 200 {object} uint
// @Router /expenses [get]
func AllExpenses(c *gin.Context) {
	var bills []models.Bill
	models.DB.Find(&bills)
	var sum uint
	for _, bill := range bills {
		sum += bill.Amount
	}
	c.JSON(http.StatusOK, gin.H{"data": sum})
}

// @Summary Get bills for a specific month
// @Description Get bills for a specific month
// @Tags bills
// @Accept json
// @Produce json
// @Param date path string true "Date in the format MM" Format(mm)
// @Success 200 {object} []models.Bill
// @Router /bill/{date} [get]
func GetMonthBills(c *gin.Context) {
	var newDate string
	var lastDate string
	date := c.Param("date")
	var bills []models.Bill
	if !strings.Contains(date, "0") && len(date) == 1 {
		newDate = "01.0" + date + ".2023"
		lastDate = "31.0" + date + ".2023"
	} else {
		newDate = "01." + date + ".2023"
		lastDate = "31." + date + ".2023"
	}
	models.DB.Where("date BETWEEN ? AND ?", newDate, lastDate).Find(&bills)
	c.JSON(http.StatusOK, gin.H{"data": bills})
}

// @Summary Update a bill
// @Description Update a bill with the specified ID
// @Tags bills
// @Accept json
// @Produce json
// @Param id path string true "Bill ID"
// @Param input body models.CreateBillInput true "Updated bill object"
// @Success 200 {object} models.Bill
// @Router /bill/{id} [put]
func UpdateBill(c *gin.Context) {
	id := c.Param("id")
	var input models.CreateBillInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var bill models.Bill
	if err := models.DB.First(&bill, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Bill not found"})
		return
	}
	bill.Title = input.Title
	bill.Author = input.Author
	bill.Amount = input.Amount
	bill.Date = input.Date
	models.DB.Save(&bill)
	c.JSON(http.StatusOK, gin.H{"data": bill})
}

// @Summary Delete a bill
// @Description Delete a bill with the specified ID
// @Tags bills
// @Accept json
// @Produce json
// @Param id path string true "Bill ID"
// @Success 204 "No Content"
// @Router /bill/{id} [delete]
func DeleteBill(c *gin.Context) {
	id := c.Param("id")
	var bill models.Bill
	if err := models.DB.First(&bill, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Bill not found"})
		return
	}
	models.DB.Delete(&bill)
	c.JSON(http.StatusOK, gin.H{"data": bill})
}
