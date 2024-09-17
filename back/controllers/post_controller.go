package controllers

import (
	"back/config"
	"back/models"
	"net/http"

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

	// First, find the user by email
	var user models.User
	if result := config.DB.Where("email = ?", email).First(&user); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User not found"})
		return
	}

	// Set AuthorID from the found user
	post.AuthorID = user.ID

	// Now create the post with the AuthorID set to the user's ID
	if result := config.DB.Create(&post); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, post)
}

// GetPosts retrieves all posts
func GetPosts(c *gin.Context) {
	var posts []models.Post
	if result := config.DB.Find(&posts); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, posts)
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
