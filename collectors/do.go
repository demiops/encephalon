package collectors

import "github.com/digitalocean/godo"
import "context"
import "golang.org/x/oauth2"
import "fmt"
// import "reflect"
import "strconv"

const (
    dopat = "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
)

type TokenSource struct {
    AccessToken         string
}

func (t *TokenSource) Token() (*oauth2.Token, error) {
    token := &oauth2.Token{
        AccessToken: t.AccessToken,
    }
    return token, nil
}

func DigitalOcean(){
    tokenSource := &TokenSource{
        AccessToken: dopat,
    }
    var ctx = context.Background()
    oauthClient := oauth2.NewClient(ctx, tokenSource)
    client := godo.NewClient(oauthClient)

    {
        result, err := DropletList(ctx, client)
        if err != nil {
            fmt.Printf("Error getting list of droplets")
        } else {
            for _, a := range result {
                var json VirtualMachine
                    //fmt.Printf("%d\n", a.ID)
                    //fmt.Printf("%s\n", a.Name)
                    //fmt.Printf("%d\n", a.Memory)
                    //fmt.Printf("%d\n", a.Disk) 
                    //fmt.Printf("%d\n", a.Vcpus)
                    //fmt.Printf("%s\n", a.Region.Slug)
                    //fmt.Printf("%s\n", a.Size.Slug)
                    //fmt.Printf("%f\n", a.Size.PriceMonthly)
                    //fmt.Printf("%f\n", a.Size.PriceHourly)
                    //fmt.Printf("%s\n", *a.Networks)
                    //for _, v4 := range a.Networks.V4 {
                    //    fmt.Printf("%s %s\n", v4.IPAddress, v4.Type)
                    //}
                    //fmt.Printf("%s\n", a.Status)
                    json.CloudID = strconv.Itoa(a.ID)
                    json.Name = a.Name
                    json.Cloud = "DigitalOcean"
                    json.AssetType = "instance"
                    json.Memory = strconv.Itoa(a.Memory)
                    json.Vcpu = strconv.Itoa(a.Vcpus)
                    json.Status = a.Status
                    PushData("success", json)
            }
        }
    }
    // cdn
    {    
        result, err := CdnList(ctx, client)
        if err != nil {
            fmt.Printf("Error getting list of droplets")
        } else {
            for _, a := range result {
                var json Cdn
                    json.CloudID = a.ID
                    json.Name = "cdn-" + a.ID
                    json.Cloud = "DigitalOcean"
                    json.AssetType = "cdn"
                    json.Origin = a.Origin
                    json.Endpoint = a.Endpoint
                    json.TTL = strconv.FormatUint(uint64(a.TTL),10)
                    PushData("success", json)
            }
        }
    }
    // firewalls
    {    
        result, err := FirewallList(ctx, client)
        if err != nil {
            fmt.Printf("Error getting list of droplets")
        } else {
            for _, a := range result {
                var json Firewall
                    json.CloudID = a.ID
                    json.Name = a.Name
                    json.Cloud = "DigitalOcean"
                    json.AssetType = "firewall"
                    json.InRules = strconv.Itoa(len(a.InboundRules))
                    json.OutRules = strconv.Itoa(len(a.OutboundRules))
                    json.Status = a.Status
                    PushData("success", json)
            }
        }
    }
    // domains
    {    
        result, err := DomainList(ctx, client)
        if err != nil {
            fmt.Printf("Error getting list of droplets")
        } else {
            for _, a := range result {
                var json Domain
                    json.CloudID = a.Name
                    json.Name = a.Name
                    json.Cloud = "DigitalOcean"
                    json.AssetType = "domain"
                    json.TTL = strconv.Itoa(a.TTL)
                    json.ZoneFile = a.ZoneFile
                    PushData("success", json)
            }
        }
    }
    // floating_ips
    {    
        result, err := FloatingIpList(ctx, client)
        if err != nil {
            fmt.Printf("Error getting list of droplets")
        } else {
            for _, a := range result {
                var json IpAddress
                    json.CloudID = a.IP
                    json.Name = a.IP
                    json.Cloud = "DigitalOcean"
                    json.AssetType = "ipaddress"
                    json.Address = a.IP
                    PushData("success", json)
            }
        }
    }
    // images
    {    
        result, err := ImageList(ctx, client)
        if err != nil {
            fmt.Printf("Error getting list of droplets")
        } else {
            for _, a := range result {
                var json Image
                    json.CloudID = strconv.Itoa(a.ID)
                    json.Name = a.Name
                    json.Cloud = "DigitalOcean"
                    json.AssetType = "image"
                    json.ImageType = a.Type
                    json.Distro = a.Distribution
                    json.Size = strconv.Itoa(a.MinDiskSize)
                    PushData("success", json)
            }
        }
    }
    // keys
    {    
        result, err := KeyList(ctx, client)
        if err != nil {
            fmt.Printf("Error getting list of droplets")
        } else {
            for _, a := range result {
                var json Key
                    json.CloudID = strconv.Itoa(a.ID)
                    json.Name = a.Name
                    json.Cloud = "DigitalOcean"
                    json.AssetType = "key"
                    json.Fingerprint = a.Fingerprint
                    json.PublicKey = a.PublicKey
                    PushData("success", json)
            }
        }
    }
    // loadbalancers
    {    
        result, err := LoadBalancerList(ctx, client)
        if err != nil {
            fmt.Printf("Error getting list of droplets")
        } else {
            for _, a := range result {
                var json LoadBalancer
                    json.CloudID = a.ID
                    json.Name = a.Name
                    json.Cloud = "DigitalOcean"
                    json.AssetType = "loadbalancer"
                    json.IpAddress = a.IP
                    json.Algorithm = a.Algorithm
                    json.Status = a.Status
                    PushData("success", json)
            }
        }
    }
    // regions
    {    
        result, err := RegionList(ctx, client)
        if err != nil {
            fmt.Printf("Error getting list of droplets")
        } else {
            for _, a := range result {
                var json Region
                    json.CloudID = a.Slug
                    json.Name = a.Name
                    json.Cloud = "DigitalOcean"
                    json.AssetType = "region"
                    json.Available = strconv.FormatBool(a.Available)
                    PushData("success", json)
            }
        }
    }
    // sizes
    {    
        result, err := SizeList(ctx, client)
        if err != nil {
            fmt.Printf("Error getting list of droplets")
        } else {
            for _, a := range result {
                var json VMSize
                    json.CloudID = a.Slug
                    json.Name = a.Slug
                    json.Cloud = "DigitalOcean"
                    json.AssetType = "vmsize"
                    json.Memory = strconv.Itoa(a.Memory)
                    json.Vcpu = strconv.Itoa(a.Vcpus)
                    json.Disk = strconv.Itoa(a.Disk)
                    json.Available = strconv.FormatBool(a.Available)
                    PushData("success", json)
            }
        }
    }
    // snapshots
    {    
        result, err := SnapshotList(ctx, client)
        if err != nil {
            fmt.Printf("Error getting list of droplets")
        } else {
            for _, a := range result {
                var json Snapshot
                    json.CloudID = a.ID
                    json.Name = a.Name
                    json.Cloud = "DigitalOcean"
                    json.AssetType = "snapshot"
                    json.Disk = strconv.FormatFloat(a.SizeGigaBytes, 'E', -1, 64)
                    PushData("success", json)
            }
        }
    }
    // volumens -> listvolumes
/*    {    
        result, err := VolumeList(ctx, client)
        if err != nil {
            fmt.Printf("Error getting list of droplets")
        } else {
            for _, a := range result {
                var json Disk
                    json.CloudID = a.ID
                    json.Name = a.Name
                    json.Cloud = "DigitalOcean"
                    json.AssetType = "disk"
                    json.Size = strconv.Itoa(a.SizeGigaBytes)
                    json.FilesystemType = a.FilesystemType
                    json.FilesystemLabel = a.FilesystemLabel
                    PushData("success", json)
            }
        }
    }
*/
    fmt.Printf("I'm done\n")
}

func DropletList(ctx context.Context, client *godo.Client) ([]godo.Droplet, error) {
    list := []godo.Droplet{}
    opt := &godo.ListOptions{}
    for {
        droplets, resp, err := client.Droplets.List(ctx, opt)
        if err != nil {
            return nil, err
        }
        for _, d := range droplets {
            list = append(list, d)
        }
        if resp.Links == nil || resp.Links.IsLastPage() {
            break
        }
        page, err := resp.Links.CurrentPage()
        if err != nil {
            return nil, err
        }
        opt.Page = page + 1
    }
    return list, nil
}

func CdnList(ctx context.Context, client *godo.Client) ([]godo.CDN, error) {
    list := []godo.CDN{}
    opt := &godo.ListOptions{}
    for {
        cdns, resp, err := client.CDNs.List(ctx, opt)
        if err != nil {
            return nil, err
        }
        for _, d := range cdns {
            list = append(list, d)
        }
        if resp.Links == nil || resp.Links.IsLastPage() {
            break
        }
        page, err := resp.Links.CurrentPage()
        if err != nil {
            return nil, err
        }
        opt.Page = page + 1
    }
    return list, nil
}

func FirewallList(ctx context.Context, client *godo.Client) ([]godo.Firewall, error) {
    list := []godo.Firewall{}
    opt := &godo.ListOptions{}
    for {
        fws, resp, err := client.Firewalls.List(ctx, opt)
        if err != nil {
            return nil, err
        }
        for _, d := range fws {
            list = append(list, d)
        }
        if resp.Links == nil || resp.Links.IsLastPage() {
            break
        }
        page, err := resp.Links.CurrentPage()
        if err != nil {
            return nil, err
        }
        opt.Page = page + 1
    }
    return list, nil
}

func DomainList(ctx context.Context, client *godo.Client) ([]godo.Domain, error) {
    list := []godo.Domain{}
    opt := &godo.ListOptions{}
    for {
        domains, resp, err := client.Domains.List(ctx, opt)
        if err != nil {
            return nil, err
        }
        for _, d := range domains {
            list = append(list, d)
        }
        if resp.Links == nil || resp.Links.IsLastPage() {
            break
        }
        page, err := resp.Links.CurrentPage()
        if err != nil {
            return nil, err
        }
        opt.Page = page + 1
    }
    return list, nil
}

func FloatingIpList(ctx context.Context, client *godo.Client) ([]godo.FloatingIP, error) {
    list := []godo.FloatingIP{}
    opt := &godo.ListOptions{}
    for {
        floatingips, resp, err := client.FloatingIPs.List(ctx, opt)
        if err != nil {
            return nil, err
        }
        for _, d := range floatingips {
            list = append(list, d)
        }
        if resp.Links == nil || resp.Links.IsLastPage() {
            break
        }
        page, err := resp.Links.CurrentPage()
        if err != nil {
            return nil, err
        }
        opt.Page = page + 1
    }
    return list, nil
}

func ImageList(ctx context.Context, client *godo.Client) ([]godo.Image, error) {
    list := []godo.Image{}
    opt := &godo.ListOptions{}
    for {
        images, resp, err := client.Images.List(ctx, opt)
        if err != nil {
            return nil, err
        }
        for _, d := range images {
            list = append(list, d)
        }
        if resp.Links == nil || resp.Links.IsLastPage() {
            break
        }
        page, err := resp.Links.CurrentPage()
        if err != nil {
            return nil, err
        }
        opt.Page = page + 1
    }
    return list, nil
}

func KeyList(ctx context.Context, client *godo.Client) ([]godo.Key, error) {
    list := []godo.Key{}
    opt := &godo.ListOptions{}
    for {
        keys, resp, err := client.Keys.List(ctx, opt)
        if err != nil {
            return nil, err
        }
        for _, d := range keys {
            list = append(list, d)
        }
        if resp.Links == nil || resp.Links.IsLastPage() {
            break
        }
        page, err := resp.Links.CurrentPage()
        if err != nil {
            return nil, err
        }
        opt.Page = page + 1
    }
    return list, nil
}

func LoadBalancerList(ctx context.Context, client *godo.Client) ([]godo.LoadBalancer, error) {
    list := []godo.LoadBalancer{}
    opt := &godo.ListOptions{}
    for {
        loadbalancers, resp, err := client.LoadBalancers.List(ctx, opt)
        if err != nil {
            return nil, err
        }
        for _, d := range loadbalancers {
            list = append(list, d)
        }
        if resp.Links == nil || resp.Links.IsLastPage() {
            break
        }
        page, err := resp.Links.CurrentPage()
        if err != nil {
            return nil, err
        }
        opt.Page = page + 1
    }
    return list, nil
}

func RegionList(ctx context.Context, client *godo.Client) ([]godo.Region, error) {
    list := []godo.Region{}
    opt := &godo.ListOptions{}
    for {
        regions, resp, err := client.Regions.List(ctx, opt)
        if err != nil {
            return nil, err
        }
        for _, d := range regions {
            list = append(list, d)
        }
        if resp.Links == nil || resp.Links.IsLastPage() {
            break
        }
        page, err := resp.Links.CurrentPage()
        if err != nil {
            return nil, err
        }
        opt.Page = page + 1
    }
    return list, nil
}

func SizeList(ctx context.Context, client *godo.Client) ([]godo.Size, error) {
    list := []godo.Size{}
    opt := &godo.ListOptions{}
    for {
        sizes, resp, err := client.Sizes.List(ctx, opt)
        if err != nil {
            return nil, err
        }
        for _, d := range sizes {
            list = append(list, d)
        }
        if resp.Links == nil || resp.Links.IsLastPage() {
            break
        }
        page, err := resp.Links.CurrentPage()
        if err != nil {
            return nil, err
        }
        opt.Page = page + 1
    }
    return list, nil
}

func SnapshotList(ctx context.Context, client *godo.Client) ([]godo.Snapshot, error) {
    list := []godo.Snapshot{}
    opt := &godo.ListOptions{}
    for {
        snapshots, resp, err := client.Snapshots.List(ctx, opt)
        if err != nil {
            return nil, err
        }
        for _, d := range snapshots {
            list = append(list, d)
        }
        if resp.Links == nil || resp.Links.IsLastPage() {
            break
        }
        page, err := resp.Links.CurrentPage()
        if err != nil {
            return nil, err
        }
        opt.Page = page + 1
    }
    return list, nil
}
/*
func VolumeList(ctx context.Context, client *godo.Client) ([]godo.Volume, error) {
    list := []godo.Volume{}
    opt := &godo.ListVolumeParams{}
    for {
        volumes, resp, err := client.Storage.ListVolumes(ctx, opt)
        if err != nil {
            return nil, err
        }
        for _, d := range volumes {
            list = append(list, d)
        }
        if resp.Links == nil || resp.Links.IsLastPage() {
            break
        }
        page, err := resp.Links.CurrentPage()
        if err != nil {
            return nil, err
        }
        opt.Page = page + 1
    }
    return list, nil
}
*/
