package controllers

import (
	"net/http"
	"strconv"
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
// @Param date path string true "Date in the format MM-YYYY" Format(MM-YYYY)
// @Success 200 {object} []models.Bill
// @Router /bill/{date} [get]
func GetMonthBills(c *gin.Context) {
	date := c.Param("date")
	parts := strings.Split(date, "-")
	if len(parts) != 2 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format. Use MM-YYYY format."})
		return
	}

	month, err := strconv.Atoi(parts[0])
	if err != nil || month < 1 || month > 12 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid month."})
		return
	}

	year, err := strconv.Atoi(parts[1])
	if err != nil || year < 1000 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid year."})
		return
	}

	var bills []models.Bill
	models.DB.Find(&bills)

	var filteredBills []models.Bill
	for _, bill := range bills {
		billParts := strings.Split(bill.Date, ".")
		billMonth, err := strconv.Atoi(billParts[1])
		if err != nil {
			// Handle parsing error if needed
			continue
		}
		billYear, err := strconv.Atoi(billParts[2])
		if err != nil {
			// Handle parsing error if needed
			continue
		}
		if billMonth == month && billYear == year {
			filteredBills = append(filteredBills, bill)
		}
	}

	c.JSON(http.StatusOK, gin.H{"data": filteredBills})
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
