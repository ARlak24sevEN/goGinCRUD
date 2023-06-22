package controllers

import (
	"goGin/initializers"
	"goGin/models"

	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {

	//Get data of req body

	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)
	//Create post
	post := models.Post{Title: body.Title, Body: body.Body}

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}
	//Return it
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsIndex(c *gin.Context) {
	// Get the post
	var posts []models.Post
	initializers.DB.Find(&posts)

	//Response
	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func PostsShow(c *gin.Context) {
	//Get id from url
	id := c.Param("id")
	// Get the post
	var post models.Post

	initializers.DB.First(&post, id)

	//Response
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostUpdate(c *gin.Context) {
	//Get the if off the url
	id := c.Param("id")

	//Get the data off req body
	var body struct {
		Body  string
		Title string
	}
	c.Bind(&body)

	//find the post were updating
	var post models.Post
	initializers.DB.First(&post, id)

	//Update it
	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body:  body.Body,
	})

	//Respond with it
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostDelete(c *gin.Context) {
	//get the id off the url
	id := c.Param("id")

	//Delete the post
	initializers.DB.Delete(&models.Post{}, id)

	//Response
	c.JSON(200, "delete id : "+id)
}
