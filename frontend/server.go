package main

import(
  "net/http"
  "github.com/gin-gonic/gin"
)

const ROOT_URL = "https://easi-pro.org"
const STATIC_CONTENT_URL = "https://s3-us-west-2.amazonaws.com/qolty/"

// func redirToHTTPS(w http.ResponseWriter, req *http.Request) {
//   http.Redirect(w, req, ROOT_URL + req.RequestURI, http.StatusMovedPermanently)
// }

// RequireAuth is a middleware to see if the user has a session.
func RequireAuth() gin.HandlerFunc {
  return func(c *gin.Context) {
    s, _ := c.Request.Cookie("session")
    if s == nil {
      c.Redirect(http.StatusMovedPermanently, ROOT_URL)
    }
  }
}

func main() {
  // Initiate the proxy server.
  // go func(){
  //   if err := http.ListenAndServe(":80", http.HandlerFunc(redirToHTTPS)); err != nil {
  //     panic("error initiating the proxy")
  //   } 
  // }()

  // Initiate the HTTPS server.
  router := gin.Default()
  router.LoadHTMLGlob("templates/*")

  authenticated := router.Group("/")
  authenticated.Use()
  {
    authenticated.GET("/", func(c *gin.Context) {
      c.HTML(http.StatusOK, "home", gin.H{
        "HOME_ACTIVE": "active",
      })
    })
    authenticated.GET("/waitingArea", func(c *gin.Context) {
      c.HTML(http.StatusOK, "waiting_area", gin.H{
        "WORKFLOWS_ACTIVE": "active",
        "WAITING_ACTIVE": "active",
      })
    })
    authenticated.GET("/waitingArea/survey", func(c *gin.Context) {
      c.HTML(http.StatusOK, "waiting_area_survey", gin.H{
        "WORKFLOWS_ACTIVE": "active",
        "WAITING_ACTIVE": "active",
      })
    })
    authenticated.GET("/patientPortal", func(c *gin.Context) {
      c.HTML(http.StatusOK, "patient_portal", gin.H{
        "WORKFLOWS_ACTIVE": "active",
        "PORTAL_ACTIVE": "active",
      })
    })
    authenticated.GET("/orderProcess", func(c *gin.Context) {
      c.HTML(http.StatusOK, "order_process", gin.H{
        "WORKFLOWS_ACTIVE": "active",
        "ORDER_ACTIVE": "active",
      })
    })
    authenticated.GET("/patientChart", func(c *gin.Context) {
      c.HTML(http.StatusOK, "patient_chart", gin.H{
        "WORKFLOWS_ACTIVE": "active",
        "PT_CHART_ACTIVE": "active",
      })
    })
    authenticated.GET("/patientChart/view", func(c *gin.Context) {
      c.HTML(http.StatusOK, "patient_chart_view", gin.H{
        "WORKFLOWS_ACTIVE": "active",
        "PT_CHART_ACTIVE": "active",
      })
    })
    authenticated.GET("/populationChart", func(c *gin.Context) {
      c.HTML(http.StatusOK, "population_chart", gin.H{
        "WORKFLOWS_ACTIVE": "active",
        "POP_CHART_ACTIVE": "active",
      })
    })
  }

  router.Run(":80")
  // router.RunTLS(":443", "/etc/letsencrypt/live/easi-pro.org/fullchain.pem", "/etc/letsencrypt/live/easi-pro.org/privkey.pem")
}