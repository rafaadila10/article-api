package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// validator instance global
var validate = validator.New()

// validation request body
func ValidateJSON(model interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		// bind JSON to struct
		if err := c.ShouldBindJSON(model); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
			c.Abort()
			return
		}

		// run validation
		if err := validate.Struct(model); err != nil {
			errors := make(map[string]string)
			for _, e := range err.(validator.ValidationErrors) {
				switch e.Field() {
				case "Title":
					errors["title"] = "Title is required and must be at least 20 characters long"
				case "Content":
					errors["content"] = "Content is required and must be at least 200 characters long"
				case "Category":
					errors["category"] = "Category is required and must be at least 3 characters long"
				case "Status":
					errors["status"] = "Status is required and must be either 'publish', 'draft', or 'thrash'"
				}
			}
			c.JSON(http.StatusBadRequest, gin.H{"validation_error": errors})
			c.Abort()
			return
		}

		// to next handler if passed
		c.Set("validated_body", model)
		c.Next()
	}
}
