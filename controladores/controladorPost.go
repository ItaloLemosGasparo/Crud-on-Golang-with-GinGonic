package controladores

import (
	"projeto/inicializadores"
	"projeto/modelos"

	"github.com/gin-gonic/gin"
)

func CriarPost(c *gin.Context) {
	// receber dados
	var body struct {
		Nome string
		Cod  int
	}

	c.Bind(&body)

	//Criar o post
	post := modelos.Post{Nome: body.Nome, Cod: body.Cod}

	result := inicializadores.BD.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	//retornar ele
	c.JSON(200, gin.H{
		"post": post,
	})
}

// Busca todos os posts realizados
func BuscarPosts(c *gin.Context) {
	//Receber os posts
	var posts []modelos.Post
	inicializadores.BD.Find(&posts)

	//apresentar eles
	c.JSON(200, gin.H{
		"posts": posts,
	})
}

// Busca 1 post especificado pelo id na url
func BuscarPost(c *gin.Context) {
	// receber o id pela URL
	id := c.Param("id")

	//Localizar os posts
	var post modelos.Post
	inicializadores.BD.First(&post, id)

	//apresentar ele
	c.JSON(200, gin.H{
		"post": post,
	})
}

func AtualizarPost(c *gin.Context) {
	// receber o id pela URL
	id := c.Param("id")

	// receber os dados
	var body struct {
		Nome string
		Cod  int
	}

	c.Bind(&body)

	// Localizar o post
	var post modelos.Post
	inicializadores.BD.First(&post, id)

	//atualizar o post
	inicializadores.BD.Model(&post).Updates(modelos.Post{
		Nome: body.Nome,
		Cod:  body.Cod,
	})

	//apresentar ele
	c.JSON(200, gin.H{
		"post": post,
	})
}

func DeletarPost(c *gin.Context) {
	// receber o id pela URL
	id := c.Param("id")

	//Deletar o post
	inicializadores.BD.Delete(&modelos.Post{}, id)

	//Apresentar ele
	c.Status(200)
}
