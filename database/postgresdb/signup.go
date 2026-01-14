package postgresdb

import (
    "fmt"
    "net/http"

    "github.com/AliMumtazDev/Go_Chat_App/models"
    "github.com/gin-gonic/gin"
)

func (u *StorageImpl) SignUpdb(c *gin.Context, req *models.User) *models.User {
    // Check duplicates (email)
    var emailExists bool
    if err := u.db.QueryRow("SELECT EXISTS (SELECT 1 FROM employeedata WHERE email = $1)", req.Email).Scan(&emailExists); err != nil {
        fmt.Printf("Error checking existing email: %v\n", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error01"})
        return nil
    }
    if emailExists {
        c.JSON(http.StatusConflict, gin.H{"error": "User already exists with this email"})
        return nil
    }

    // Check duplicates (username)
    var usernameExists bool
    if err := u.db.QueryRow("SELECT EXISTS (SELECT 1 FROM employeedata WHERE username = $1)", req.Username).Scan(&usernameExists); err != nil {
        fmt.Printf("Error checking existing username: %v\n", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error02"})
        return nil
    }
    if usernameExists {
        c.JSON(http.StatusConflict, gin.H{"error": "Username not available, choose another"})
        return nil
    }

    if _, err := u.db.Exec("INSERT INTO employeedata (username, email, password) VALUES ($1, $2, $3)", req.Username, req.Email, req.Password); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user", "details": err.Error()})
        return nil
    }

    return req
}
