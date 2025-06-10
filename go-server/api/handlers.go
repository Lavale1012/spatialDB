package api

import (
	"fmt"
	"net/http"
	"sort"

	// "strings"
	api "github.com/Lavale1012/vector-db/go-server/clients"
	"github.com/Lavale1012/vector-db/go-server/config"

	math "github.com/Lavale1012/vector-db/go-server/core"
	"github.com/Lavale1012/vector-db/go-server/models"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

type EmbedRequest struct {
	ID   string `json:"id"`
	Text string `json:"text"`
}
type SearchRequest struct {
	Text string `json:"text"`
	K    int    `json:"k"`
}
type SearchResponse struct {
	ID string `json:"id"`
	// Optional, if you want to include the vector in the response
	Text  string  `json:"text"`
	Score float64 `json:"score"` // Optional, if you want to include the score in the response
}

func EmbedReq(c *gin.Context) {

	var req []EmbedRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Println("Error binding JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}
	for _, r := range req {
		vector, err := api.EmbedClient(r.Text)
		if err != nil {
			fmt.Println("Error in embed_client:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get embedding vector"})
			return
		}

		model := models.VectorModel{
			ID:     r.ID,
			Text:   r.Text,
			Vector: pq.Float64Array(vector),
		}

		if err := config.DB.Create(&model).Error; err != nil {
			fmt.Println("Error saving vector to database:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save vector to database"})
			return
		}

	}

	c.JSON(http.StatusOK, gin.H{
		"message": "All vectors saved to database successfully",
	})
}

func SearchReq(c *gin.Context) {
	var req SearchRequest
	var vectors []models.VectorModel
	var results []SearchResponse

	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Println("Error binding JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}
	queryVector, err := api.EmbedClient(req.Text)
	if err != nil {
		fmt.Println("Error in embed_client:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get embedding vector"})
		return
	}
	if err := config.DB.Find(&vectors).Error; err != nil {
		fmt.Println("Error retrieving vectors from database:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve vectors from database"})
		return
	}
	for _, vector := range vectors {
		score := math.CosineSimilarity([]float64(vector.Vector), queryVector)
		results = append(results, SearchResponse{
			ID:   vector.ID,
			Text: vector.Text,
			// Optional, if you want to include the vector in the response
			Score: score, // Uncomment if you add Score to SearchResponse struct

		})
	}
	sort.Slice(results, func(i, j int) bool {
		return results[i].Score > results[j].Score // Sort by score descending
	})
	if req.K > 0 && req.K < len(results) {
		results = results[:req.K]
	}
	c.JSON(http.StatusOK, gin.H{"results": results})

}
