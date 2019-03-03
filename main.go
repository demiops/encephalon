package main

import "time"
import "strconv"
// import "log"
import "github.com/gin-gonic/gin"
import "github.com/olivere/elastic"
// import "github.com/teris-io/shortid"
import "net/http"
import "./collectors"
//import "fmt"
//import "context"


const (
    elasticIndexName    = "assets"
    elasticTypeName     = "document"
)

var (
    elasticClient *elastic.Client
)

func main() {
    var err error
    //ctx := context.Background()
    for {
        elasticClient, err = elastic.NewClient(
            elastic.SetURL("http://localhost:9200"),
            elastic.SetSniff(false),
        )
        if err != nil {
            time.Sleep(3 * time.Second)
        } else {
            break
        }
    }

    go collectors.GoogleCloud("./collectors/google-key.json","api-project-xxxxxxxx")
    go collectors.DigitalOcean()
    go collectors.Azure()
    go collectors.Aws()

    router := gin.Default()

    router.GET("/networks", getnetworks)
    router.GET("/network/:id", getnetwork)
    router.GET("/servers", getservers)
    router.GET("/server/:id", getserver)
    router.GET("/images", getimages)
    router.GET("/image/:id", getimage)
    router.GET("/containers", getcontainers)
    router.GET("/container/:id", getcontainer)
    router.GET("/loadbalancers", getloadbalancers)
    router.GET("/loadbalancer/:id", getloadbalancer)
    router.POST("/asset", createAsset)
    router.PUT("/asset", updateAsset)

    router.Run()

}

func getnetworks(c *gin.Context) {
    c.String(http.StatusOK, "OK")
}
func getnetwork(c *gin.Context) {
    c.String(http.StatusOK, c.Param("id"))
}
func getservers(c *gin.Context) {
    query := c.Query("query")
    if query == "" {
        c.String(http.StatusBadRequest, "Query is empty")
        return
    }
    skip := 0
    take := 10
    if i, err := strconv.Atoi(c.Query("skip")); err == nil {
        skip = i
    }
    if i, err := strconv.Atoi(c.Query("take")); err == nil {
        take = i
    }
    // Perform search
    esQuery := elastic.NewMultiMatchQuery(query, "title", "content").
        Fuzziness("2").
        MinimumShouldMatch("2")
    result, err := elasticClient.Search().
        Index(elasticIndexName).
        Query(esQuery).
        From(skip).Size(take).
        Do(c.Request.Context())
    if err != nil {
        return
    }
    c.JSON(http.StatusOK, result.Hits.Hits)
}
func getserver(c *gin.Context) {
    c.String(http.StatusOK, c.Param("id"))
}
func getimages(c *gin.Context) {
    c.String(http.StatusOK, "OK")
}
func getimage(c *gin.Context) {
    c.String(http.StatusOK, c.Param("id"))
}
func getcontainers(c *gin.Context) {
    c.String(http.StatusOK, "OK")
}
func getcontainer(c *gin.Context) {
    c.String(http.StatusOK, c.Param("id"))
}
func getloadbalancers(c *gin.Context) {
    c.String(http.StatusOK, "OK")
}
func getloadbalancer(c *gin.Context) {
    c.String(http.StatusOK, c.Param("id"))
}

func createAsset(c *gin.Context) {
    var json collectors.Asset 
    if err := c.ShouldBindJSON(&json); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
        return
    }
    bulk := elasticClient.Bulk().Index(elasticIndexName).Type(elasticTypeName)
    bulk.Add(elastic.NewBulkIndexRequest().Id(json.CloudID).Doc(json))
    if _, err := bulk.Do(c.Request.Context()); err != nil {
        c.JSON(http.StatusInternalServerError, "Failed to create assets")
        return
    }
    c.String(http.StatusOK, "OK")

}

func updateAsset(c *gin.Context) {
    c.String(http.StatusOK, "OK")
}
