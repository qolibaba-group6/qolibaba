
package routes

import (
    "testing"

    "github.com/gin-gonic/gin"
    "github.com/stretchr/testify/assert"
)

func TestRegisterCompanyRoutes(t *testing.T) {
    router := gin.Default()
    RegisterCompanyRoutes(router, nil)

    routes := router.Routes()
    assert.NotEmpty(t, routes, "routes should not be empty")
}
