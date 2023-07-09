package https

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func StartServer(port string) error {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/query", GraphQLQuery)
	e.GET("/playground", GraphQLPlayground)


	//ipfs apis just for demo
	//eg. http://localhost:8000/get-ipfs/https://placekitten.com/200/300
	e.GET("/get-ipfs/:cid", IPFSGetHandler)
	//eg. http://localhost:8000/add-ipfs/QmXW4a4qoySdr7FnqjGanjTsrALpD856dGnLjEq7QconsS
	e.GET("/add-ipfs/:url", IPFSAddFromURLHandler)

	return e.Start(":" + port)

}
