// Golang
// Another server using Gin

package main

import (
	"github.com/gin-gonic/gin"
)

var ninjaWeapons = map[string]string{
	"ninjaStar": "Beginner Ninja Star - Damage 1",
}

// curl localhost:8000/weapon?type=ninjaStar
func GetWeapon(c *gin.Context) {
	weaponType := c.Query("type")
	weaponName, ok := ninjaWeapons[weaponType]
	if !ok {
		c.JSON(404, gin.H{
			weaponType: "",
		})
		return
	}
	c.JSON(200, gin.H{
		weaponType: weaponName,
	})
}

func PostWeapon(c *gin.Context) {
	weaponType := c.Query("type")
	weaponName := c.Query("name")

	// some simple validation
	if len(weaponType) == 0 || len(weaponName) == 0 {
		c.JSON(400, gin.H{
			weaponType: weaponName,
		})
		return
	}

	if _, ok := ninjaWeapons[weaponType]; ok {
		c.JSON(409, gin.H{
			"message": "Weapon alread exists",
		})
		return
	}
	c.JSON(201, gin.H{
		weaponType: weaponName,
	})

}

func notmain() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/weapon", GetWeapon)
	r.POST("/weapon", PostWeapon)
	r.Run("localhost:8000")
}
