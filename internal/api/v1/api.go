package v1

import "github.com/gin-gonic/gin"

type API struct {
	super *gin.RouterGroup
}

func NewAPI(superGroup *gin.RouterGroup) *API {
	return &API{
		super: superGroup,
	}
}

func (h *API) Router() *gin.RouterGroup {
	//api := h.super.Group("/api/v1")
	h.super.Use(CORSMiddleware())
	return h.super
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		c.Next()
	}
}