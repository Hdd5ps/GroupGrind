package main

import (
  "github.com/gin-gonic/gin"
)

func main() {
  r := gin.Default()
  r.GET("/api/study-groups", getStudyGroups)
  r.POST("/api/study-groups", createStudyGroup)
  r.Run()
}
package main

import (
  "github.com/gin-gonic/gin"
)

func main() {
  r := gin.Default()

  // Study group routes
  r.GET("/api/study-groups", getStudyGroups)
  r.POST("/api/study-groups", createStudyGroup)
  r.PUT("/api/study-groups/:id", updateStudyGroup)
  r.DELETE("/api/study-groups/:id", deleteStudyGroup)

  // Member routes
  r.POST("/api/study-groups/:id/members", addMember)
  r.DELETE("/api/study-groups/:id/members/:memberId", removeMember)

  r.Run()
}
package main

import (
  "github.com/gin-gonic/gin"
)

func main() {
  r := gin.Default()

  // Study group routes
  r.GET("/api/study-groups", getStudyGroups)
  r.POST("/api/study-groups", createStudyGroup)
  r.PUT("/api/study-groups/:id", updateStudyGroup)
  r.DELETE("/api/study-groups/:id", deleteStudyGroup)

  // Member routes
  r.POST("/api/study-groups/:id/members", addMember)
  r.DELETE("/api/study-groups/:id/members/:memberId", removeMember)

  r.Run()
}
package main

import (
  "github.com/gin-gonic/gin"
)

func main() {
  r := gin.Default()

  // User authentication routes
  r.POST("/api/register", registerUser)
  r.POST("/api/login", loginUser)

  // Protected routes (example)
  r.GET("/api/protected", AuthMiddleware(), protectedEndpoint)

  r.Run()
}
