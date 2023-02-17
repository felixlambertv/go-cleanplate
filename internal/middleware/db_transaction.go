package middleware

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
)

func DbTransactionMiddleware(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		txHandle := db.Begin()
		log.Print("begin db transaction")

		defer func() {
			if r := recover(); r != nil {
				txHandle.Rollback()
			}
		}()

		c.Set("db_trx", txHandle)
		c.Next()
	}
}
