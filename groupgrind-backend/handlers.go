package main

import (
  "net/http"
  "github.com/gin-gonic/gin"
)

func getStudyGroups(c *gin.Context) {
  // Logic to fetch study groups
}

func createStudyGroup(c *gin.Context) {
  // Logic to create a new study group
}
func getPredictions(data []float32) ([]float32, error) {
	// Logic to call TensorFlow Serving API and get predictions
  }
  package main

  import (
    "net/http"
    "github.com/gin-gonic/gin"
    "strconv"
  )
  
  // Handler to get all study groups
  func getStudyGroups(c *gin.Context) {
    // Logic to fetch study groups from the database
    c.JSON(http.StatusOK, gin.H{"message": "Get all study groups"})
  }
  
  // Handler to create a new study group
  func createStudyGroup(c *gin.Context) {
    var newGroup StudyGroup
    if err := c.ShouldBindJSON(&newGroup); err != nil {
      c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
      return
    }
    // Logic to save the new study group to the database
    c.JSON(http.StatusCreated, gin.H{"message": "Study group created"})
  }
  
  // Handler to update an existing study group
  func updateStudyGroup(c *gin.Context) {
    id := c.Param("id")
    var updatedGroup StudyGroup
    if err := c.ShouldBindJSON(&updatedGroup); err != nil {
      c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
      return
    }
    // Logic to update the study group in the database
    c.JSON(http.StatusOK, gin.H{"message": "Study group updated", "id": id})
  }
  
  // Handler to delete a study group
  func deleteStudyGroup(c *gin.Context) {
    id := c.Param("id")
    // Logic to delete the study group from the database
    c.JSON(http.StatusOK, gin.H{"message": "Study group deleted", "id": id})
  }
  
  // Handler to add a member to a study group
  func addMember(c *gin.Context) {
    id := c.Param("id")
    var newMember Member
    if err := c.ShouldBindJSON(&newMember); err != nil {
      c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
      return
    }
    // Logic to add the member to the study group in the database
    c.JSON(http.StatusOK, gin.H{"message": "Member added", "groupId": id})
  }
  
  // Handler to remove a member from a study group
  func removeMember(c *gin.Context) {
    id := c.Param("id")
    memberId := c.Param("memberId")
    // Logic to remove the member from the study group in the database
    c.JSON(http.StatusOK, gin.H{"message": "Member removed", "groupId": id, "memberId": memberId})
  }
  package main

  import (
    "net/http"
    "github.com/gin-gonic/gin"
    "golang.org/x/crypto/bcrypt"
    "github.com/dgrijalva/jwt-go"
    "time"
  )
  
  // Secret key for JWT
  var jwtKey = []byte("your_secret_key")
  
  // User struct for registration and login
  type User struct {
    ID       uint   `json:"id" gorm:"primary_key"`
    Username string `json:"username" gorm:"unique"`
    Password string `json:"password"`
  }
  
  // JWT claims struct
  type Claims struct {
    Username string `json:"username"`
    jwt.StandardClaims
  }
  
  // Handler to register a new user
  func registerUser(c *gin.Context) {
    var user User
    if err := c.ShouldBindJSON(&user); err != nil {
      c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
      return
    }
  
    // Hash the password
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
      c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
      return
    }
    user.Password = string(hashedPassword)
  
    // Save the user to the database
    if err := db.Create(&user).Error; err != nil {
      c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
      return
    }
  
    c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
  }
  
  // Handler to login a user
  func loginUser(c *gin.Context) {
    var user User
    var loginDetails User
    if err := c.ShouldBindJSON(&loginDetails); err != nil {
      c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
      return
    }
  
    // Find the user by username
    if err := db.Where("username = ?", loginDetails.Username).First(&user).Error; err != nil {
      c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
      return
    }
  
    // Check the password
    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginDetails.Password)); err != nil {
      c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
      return
    }
  
    // Create JWT token
    expirationTime := time.Now().Add(24 * time.Hour)
    claims := &Claims{
      Username: user.Username,
      StandardClaims: jwt.StandardClaims{
        ExpiresAt: expirationTime.Unix(),
      },
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    tokenString, err := token.SignedString(jwtKey)
    if err != nil {
      c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
      return
    }
  
    c.JSON(http.StatusOK, gin.H{"token": tokenString})
  }
  
  // Middleware to protect routes
  func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
      tokenString := c.GetHeader("Authorization")
      if tokenString == "" {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
        c.Abort()
        return
      }
  
      claims := &Claims{}
      token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
        return jwtKey, nil
      })
      if err != nil || !token.Valid {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
        c.Abort()
        return
      }
  
      c.Set("username", claims.Username)
      c.Next()
    }
  }
  
  // Example protected endpoint
  func protectedEndpoint(c *gin.Context) {
    username := c.MustGet("username").(string)
    c.JSON(http.StatusOK, gin.H{"message": "Hello " + username})
  }
  
  