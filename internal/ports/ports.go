package ports

import "github.com/gin-gonic/gin"

// Driver

type RestPort interface {
	CreateRouter() *gin.Engine
	RunRestServer(*gin.Engine) error
}

// API
// Connects Rest to DB
type APIPort interface {
	CallTest()
}

// Driven

type DBPort interface {
	Test()
}
