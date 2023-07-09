package https

import (
	contentStore "artvortex-api/contentStore"
	"artvortex-api/graph"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/labstack/echo/v4"
)

func GraphQLQuery(c echo.Context) error {
	graphqlHandler := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))
	graphqlHandler.ServeHTTP(c.Response(), c.Request())
	return nil
}

func GraphQLPlayground(c echo.Context) error {
	playgroundHandler := playground.Handler("GraphQL", "/query")
	playgroundHandler.ServeHTTP(c.Response(), c.Request())
	return nil
}

/**
 * just for testing
 */
func IPFSGetHandler(c echo.Context) error {
	cid := c.Param("cid")
	ipfsContent, err := contentStore.IPFSGetByCID(cid,"")

	contentSlice := ipfsContent[0:512]
	contentType := http.DetectContentType(contentSlice)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}
	format := contentType

	return c.Blob(http.StatusOK, format, ipfsContent)
}

/**
 * just for testing
 */
func IPFSAddFromURLHandler(c echo.Context) error {
	url := c.Param("url")
	content, err := contentStore.GetObjectFromURL(url)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}
	cid, err := contentStore.IPFSAddContent(content, "")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK,map[string]string{
		"cid": cid,
	})
}
