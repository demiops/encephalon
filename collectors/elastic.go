package collectors

import "github.com/olivere/elastic"
//import "net/http"
import "context"
import "fmt"
import "time"

const (
    elasticIndexName    = "assets"
    elasticTypeName     = "document"
)

var (
    elasticClient *elastic.Client
)

type PushValue struct {
    Status string
    CustomStruct        interface{}
}


func PushData(status string, class interface{}) (string, error){

    var err error
    var result = PushValue {Status: status, CustomStruct: class}
    
    ctx := context.Background()
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
    {
        data, ok := result.CustomStruct.(VirtualMachine)
        if ok {
            bulk := elasticClient.Bulk().Index(elasticIndexName).Type(elasticTypeName)
            bulk.Add(elastic.NewBulkIndexRequest().Id(data.CloudID).Doc(data))
            if _, err := bulk.Do(ctx); err != nil {
                fmt.Println("Failed to create assets")
                return "", err
            }
        }
    }
    {
        data, ok := result.CustomStruct.(Disk)
        if ok {
            bulk := elasticClient.Bulk().Index(elasticIndexName).Type(elasticTypeName)
            bulk.Add(elastic.NewBulkIndexRequest().Id(data.CloudID).Doc(data))
            if _, err := bulk.Do(ctx); err != nil {
                fmt.Println("Failed to create assets")
                return "", err
            }
        }
    }
    {
        data, ok := result.CustomStruct.(Image)
        if ok {
            bulk := elasticClient.Bulk().Index(elasticIndexName).Type(elasticTypeName)
            bulk.Add(elastic.NewBulkIndexRequest().Id(data.CloudID).Doc(data))
            if _, err := bulk.Do(ctx); err != nil {
                fmt.Println("Failed to create assets")
                return "", err
            }
        }
    }
    {
        data, ok := result.CustomStruct.(Region)
        if ok {
            bulk := elasticClient.Bulk().Index(elasticIndexName).Type(elasticTypeName)
            bulk.Add(elastic.NewBulkIndexRequest().Id(data.CloudID).Doc(data))
            if _, err := bulk.Do(ctx); err != nil {
                fmt.Println("Failed to create assets")
                return "", err
            }
        }
    }
    {
        data, ok := result.CustomStruct.(Az)
        if ok {
            bulk := elasticClient.Bulk().Index(elasticIndexName).Type(elasticTypeName)
            bulk.Add(elastic.NewBulkIndexRequest().Id(data.CloudID).Doc(data))
            if _, err := bulk.Do(ctx); err != nil {
                fmt.Println("Failed to create assets")
                return "", err
            }
        }
    }
    {
        data, ok := result.CustomStruct.(LoadBalancer)
        if ok {
            bulk := elasticClient.Bulk().Index(elasticIndexName).Type(elasticTypeName)
            bulk.Add(elastic.NewBulkIndexRequest().Id(data.CloudID).Doc(data))
            if _, err := bulk.Do(ctx); err != nil {
                fmt.Println("Failed to create assets")
                return "", err
            }
        }
    }
    {
        data, ok := result.CustomStruct.(Firewall)
        if ok {
            bulk := elasticClient.Bulk().Index(elasticIndexName).Type(elasticTypeName)
            bulk.Add(elastic.NewBulkIndexRequest().Id(data.CloudID).Doc(data))
            if _, err := bulk.Do(ctx); err != nil {
                fmt.Println("Failed to create assets")
                return "", err
            }
        }
    }
    {
        data, ok := result.CustomStruct.(VMSize)
        if ok {
            bulk := elasticClient.Bulk().Index(elasticIndexName).Type(elasticTypeName)
            bulk.Add(elastic.NewBulkIndexRequest().Id(data.CloudID).Doc(data))
            if _, err := bulk.Do(ctx); err != nil {
                fmt.Println("Failed to create assets")
                return "", err
            }
        }
    }
    {
        data, ok := result.CustomStruct.(IpAddress)
        if ok {
            bulk := elasticClient.Bulk().Index(elasticIndexName).Type(elasticTypeName)
            bulk.Add(elastic.NewBulkIndexRequest().Id(data.CloudID).Doc(data))
            if _, err := bulk.Do(ctx); err != nil {
                fmt.Println("Failed to create assets")
                return "", err
            }
        }
    }
    {
        data, ok := result.CustomStruct.(Network)
        if ok {
            bulk := elasticClient.Bulk().Index(elasticIndexName).Type(elasticTypeName)
            bulk.Add(elastic.NewBulkIndexRequest().Id(data.CloudID).Doc(data))
            if _, err := bulk.Do(ctx); err != nil {
                fmt.Println("Failed to create assets")
                return "", err
            }
        }
    }    
    {
        data, ok := result.CustomStruct.(InternetGateway)
        if ok {
            bulk := elasticClient.Bulk().Index(elasticIndexName).Type(elasticTypeName)
            bulk.Add(elastic.NewBulkIndexRequest().Id(data.CloudID).Doc(data))
            if _, err := bulk.Do(ctx); err != nil {
                fmt.Println("Failed to create assets")
                return "", err
            }
        }
    }    
    {
        data, ok := result.CustomStruct.(Subnet)
        if ok {
            bulk := elasticClient.Bulk().Index(elasticIndexName).Type(elasticTypeName)
            bulk.Add(elastic.NewBulkIndexRequest().Id(data.CloudID).Doc(data))
            if _, err := bulk.Do(ctx); err != nil {
                fmt.Println("Failed to create assets")
                return "", err
            }
        }
    }
    {
        data, ok := result.CustomStruct.(Router)
        if ok {
            bulk := elasticClient.Bulk().Index(elasticIndexName).Type(elasticTypeName)
            bulk.Add(elastic.NewBulkIndexRequest().Id(data.CloudID).Doc(data))
            if _, err := bulk.Do(ctx); err != nil {
                fmt.Println("Failed to create assets")
                return "", err
            }
        }
    }
    {
        data, ok := result.CustomStruct.(Routes)
        if ok {
            bulk := elasticClient.Bulk().Index(elasticIndexName).Type(elasticTypeName)
            bulk.Add(elastic.NewBulkIndexRequest().Id(data.CloudID).Doc(data))
            if _, err := bulk.Do(ctx); err != nil {
                fmt.Println("Failed to create assets")
                return "", err
            }
        }
    }
    return "OK", nil
}
