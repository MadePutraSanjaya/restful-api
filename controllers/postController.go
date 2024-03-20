package controllers

import (
	"crud/initializers"
	"crud/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func PostCreate(c *gin.Context) {
	// Generate new UUID
	id := uuid.New()

	// get data
	var body struct {
		Title string
		Body  string
	}

	c.Bind(&body)

	// create post
	post := models.Post{ID: id, Title: body.Title, Body: body.Body}
	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return

	}

	// Response
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsIndex(c *gin.Context) {
	// get data
	var posts []models.Post
	result := initializers.DB.Find(&posts)

	// Response
	c.JSON(200, gin.H{
		"posts": posts,
	})

	if result.Error != nil {
		c.Status(400)
		return
	}

}

func PostsShow(c *gin.Context) {
	// get id
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
        c.JSON(400, gin.H{"error": "invalid ID format"})
        return
    }

	// get data
	var posts models.Post
	result := initializers.DB.Find(&posts, id)

	// Response
	c.JSON(200, gin.H{
		"posts": posts,
	})

	if result.Error != nil {
		c.Status(400)
		return
	}
}

func PostsUpdate(c *gin.Context) {
	// Generate new UUID
	uuid := uuid.New()

	// get id
	id := c.Param("id")

	// get data
	var body struct {
		Title string
		Body  string
	}

	c.Bind(&body)

	// find the post where will updating
	var posts models.Post
	result := initializers.DB.Find(&posts, id)

	// update data
	initializers.DB.Model(&posts).Updates(models.Post{
		ID:    uuid,
		Title: body.Title,
		Body:  body.Body,
	})

	// response
	c.JSON(200, gin.H{
		"posts": posts,
	})

	if result.Error != nil {
		c.Status(400)
		return
	}

}

func PostsDelete(c *gin.Context) {
	// get id
	id := c.Param("id")

	// find the post where will deleting
	var posts models.Post
	result := initializers.DB.Delete(&posts, id)

	// response
	c.JSON(200, gin.H{
		"message": "Delete Post Successfully",
	})

	if result.Error != nil {
		c.Status(400)
		return
	}
}
