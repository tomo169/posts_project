package controllers

import (
	"back/config"
	"back/models"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// CreatePost creates a new post
func CreatePost(c *gin.Context) {
	var post models.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	email := c.MustGet("email").(string) // Retrieve email from context set by middleware

	// Find the user by email
	var user models.User
	if result := config.DB.Where("email = ?", email).First(&user); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User not found"})
		return
	}

	post.AuthorID = user.ID

	// Create the post with the AuthorID set to the user's ID
	if result := config.DB.Create(&post); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, post)
}

// GetPosts retrieves all posts
func GetPosts(c *gin.Context) {
	limit := 10
	offset := 0

	if l, err := strconv.Atoi(c.Query("limit")); err == nil {
		limit = l
	}
	if o, err := strconv.Atoi(c.Query("offset")); err == nil {
		offset = o
	}

	var posts []models.Post
	// Preload the Author relationship to get author data
	if err := config.DB.Preload("Author").Order("created_at desc").Limit(limit).Offset(offset).Find(&posts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch posts"})
		return
	}

	type PostResponse struct {
		ID         uint      `json:"id"`
		Title      string    `json:"title"`
		Content    string    `json:"content"`
		AuthorName string    `json:"authorName"`
		AuthorID   uint      `json:"authorId"`
		CreatedAt  time.Time `json:"createdAt"`
		UpdatedAt  time.Time `json:"updatedAt"`
	}

	var response []PostResponse
	for _, post := range posts {
		response = append(response, PostResponse{
			ID:         post.ID,
			Title:      post.Title,
			Content:    post.Content,
			AuthorName: post.Author.Name,
			AuthorID:   post.AuthorID,
			CreatedAt:  post.CreatedAt,
			UpdatedAt:  post.UpdatedAt,
		})
	}

	// Respond with posts and included author details
	c.JSON(http.StatusOK, response)
}

// GetPost retrieves a post by id
func GetPost(c *gin.Context) {
	id := c.Param("id")
	var post models.Post
	if result := config.DB.First(&post, id); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}
	c.JSON(http.StatusOK, post)
}

// UpdatePost updates a post by id
func UpdatePost(c *gin.Context) {
	id := c.Param("id")
	var post models.Post
	if result := config.DB.First(&post, id); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Save(&post)
	c.JSON(http.StatusOK, post)
}

// DeletePost deletes a post by id
func DeletePost(c *gin.Context) {
	id := c.Param("id")
	if result := config.DB.Delete(&models.Post{}, id); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Post deleted"})
}

func GetPostsByUser(c *gin.Context) {
	userID := c.Param("userID")

	var posts []models.Post
	// Fetch posts where 'AuthorID' matches 'userID'
	if err := config.DB.Where("author_id = ?", userID).Find(&posts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch posts"})
		return
	}

	if len(posts) == 0 {
		c.JSON(http.StatusOK, []models.Post{})
		return
	}

	c.JSON(http.StatusOK, posts)
}
