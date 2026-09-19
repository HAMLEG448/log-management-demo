package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"log-management/backend/config"
	"log-management/backend/entity"
	sevices "log-management/backend/services"

	"github.com/gin-gonic/gin"
)

func IngestAWSFile(c *gin.Context) {
	fmt.Println("1. AWS request received")

	fileHeader, err := c.FormFile("file")
	if err != nil {
		fmt.Println("ERROR FormFile:", err)

		c.JSON(http.StatusBadRequest, gin.H{
			"error": "file is required",
		})
		return
	}

	fmt.Println("2. File received:", fileHeader.Filename)

	file, err := fileHeader.Open()
	if err != nil {
		fmt.Println("ERROR Open:", err)

		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to open file",
		})
		return
	}
	defer file.Close()

	fmt.Println("3. File opened")

	data, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("ERROR ReadAll:", err)

		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to read file",
		})
		return
	}

	fmt.Println("4. File read, bytes:", len(data))

	var inputs []entity.AWSCloudTrailInput

	if err := json.Unmarshal(data, &inputs); err != nil {
		var single entity.AWSCloudTrailInput

		if err := json.Unmarshal(data, &single); err != nil {
			fmt.Println("ERROR JSON:", err)

			c.JSON(http.StatusBadRequest, gin.H{
				"error": "invalid AWS JSON file",
			})
			return
		}

		inputs = []entity.AWSCloudTrailInput{single}
	}

	fmt.Println("5. JSON parsed, records:", len(inputs))

	var logs []entity.Log

	for _, input := range inputs {
		normalizedLog, err := sevices.NormalizeAWSLog(input)

		if err != nil {
			fmt.Println("ERROR Normalize:", err)

			c.JSON(http.StatusBadRequest, gin.H{
				"error": "failed to normalize AWS log",
			})
			return
		}

		logs = append(logs, normalizedLog)
	}

	fmt.Println("6. Normalized:", len(logs))

	if len(logs) > 0 {
		fmt.Println("7. Saving to database...")

		if err := config.DB.Create(&logs).Error; err != nil {
			fmt.Println("ERROR Database:", err)

			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "failed to save AWS logs",
			})
			return
		}
	}

	fmt.Println("8. Database save completed")

	c.JSON(http.StatusCreated, gin.H{
		"message": "AWS logs ingested successfully",
		"count":   len(logs),
		"logs":    logs,
	})
}
