package controllers

import (
	"github.com/alfathmuqoddas/go-crud-alf/initializers"
	"github.com/alfathmuqoddas/go-crud-alf/models"
	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {
	//get data off request body
	var body struct {
		Body string
		Title string
	}

	c.Bind(&body)

	//create a post
	post := models.Post{Title: body.Title, Body: body.Body}
    result := initializers.DB.Create(&post)
	
	if result.Error != nil {
		c.Status(400)
		return
	}

	//return it
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsIndex(c *gin.Context) {
	//get all posts
	var posts []models.Post
	initializers.DB.Find(&posts)

	//responds
	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func PostById(c *gin.Context) {
    //get id from url
	id := c.Param("id")

	//get all posts
	var post models.Post
	initializers.DB.First(&post, id)

	//responds
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostUpdate(c *gin.Context) {
	//get id off url
	id := c.Param("id")

	//get the data
	var body struct {
		Body string
		Title string
	}

	c.Bind(&body)

	//find the post where updating
	var post models.Post
	initializers.DB.First(&post, id)

	//update it
	initializers.DB.Model(&post).Updates(models.Post{Title: body.Title, Body: body.Body,})

	//respond with it
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostDelete(c *gin.Context) {
	//get the post id
	id := c.Param("id")

	//delete it
	initializers.DB.Delete(&models.Post{}, id)

	//respond
	c.Status(200)
}