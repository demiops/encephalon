package collectors

import "google.golang.org/api/compute/v1"
import "golang.org/x/oauth2"
import "golang.org/x/oauth2/google"
import "fmt"
import "io/ioutil"
import "strconv"
import "net/url"
import "strings"


var kf string = "./google-key.json"
var pr string = "api-project-xxxxxxxxxxxx"

func GetLastUrlSegment(rawurl string) (string){
    uri, err := url.Parse(rawurl)
    if err != nil {
        return ""
    }
    path := uri.Path
    uriSegments := strings.Split(path,"/")
    res := uriSegments[len(uriSegments)-1]
    return res
}

func GoogleCloud(kf string, pr string) {
    keyfile, err := ioutil.ReadFile(kf)

    if err != nil {
        fmt.Printf("Error while reading google key json file %s", err.Error())
    }

    conf, err := google.JWTConfigFromJSON(keyfile, "https://www.googleapis.com/auth/compute")

    if err != nil {
        fmt.Printf("Error while authenticating with Compute API %s", err.Error())
    }

    client := conf.Client(oauth2.NoContext)

    srv, err := compute.New(client)

    if err != nil {
        fmt.Printf("Error while getting service compute %s", err.Error())
    }

    // Zones
    {
        r, err := srv.Zones.List("api-project-xxxxxxxxxxxxxx").Do()
        if err != nil {
            fmt.Printf("Error while getting Zone list %s", err.Error())
        }

        //fmt.Printf("%s\n", r.Id)
        for _, v := range r.Items {
            //fmt.Printf("%s is %s on %s\n", v.Name, v.Status, v.Region)
            var json Az
            json.Name = v.Name
            json.Cloud = "GoogleCloud"
            json.CloudID = v.Name
            json.AssetType = "availabilityzone"
            json.Status = v.Status
            json.Region = GetLastUrlSegment(v.Region)
            PushData("success", json)
        }
    }

    // Instances 
    {
        s, err := srv.Instances.List(pr,"us-central1-c").Do()
        if err != nil {
            fmt.Printf("Error while getting Instances List %s", err.Error())
        }

        //fmt.Printf("%s\n", s.Id)
        for _, i := range s.Items {
            //fmt.Printf("%s, %s, %s, %s, %s, %s\n", i.MachineType, i.Name, i.CpuPlatform, i.Status, i.NetworkInterfaces[0].Name, i.NetworkInterfaces[0].NetworkIP)
            var json VirtualMachine
            json.Name = i.Name
            json.Cloud = "GoogleCloud"
            json.CloudID = i.Name
            json.AssetType = "instance"
            json.Status = i.Status
            json.Disks = i.Disks[0].DeviceName
            json.Az = GetLastUrlSegment(i.Zone)
            json.Arch = i.CpuPlatform
            json.VMSize = GetLastUrlSegment(i.MachineType)
            json.Network = GetLastUrlSegment(i.NetworkInterfaces[0].Network)
            json.Subnet = GetLastUrlSegment(i.NetworkInterfaces[0].Subnetwork)
            json.PrivateIP = i.NetworkInterfaces[0].NetworkIP
            PushData("success", json) 
        }
    }

    //AcceleratorTypes
    //{
    //    s, err := srv.Instances.List(pr,"us-central1-c").Do()
    //    if err != nil {
    //        fmt.Printf("Error while getting Instances List %s", err.Error())
    //    }

    //    //fmt.Printf("%s\n", s.Id)
    //    for _, i := range s.Items {
    //        fmt.Printf("%s, %s, %s, %s, %s, %s\n", i.MachineType, i.Name, i.CpuPlatform, i.Status, i.NetworkInterfaces[0].Name, i.NetworkInterfaces[0].NetworkIP)
    //    }
    //}
    
    //addresses
    {
        s, err := srv.Addresses.List(pr,"us-central1").Do()
        if err != nil {
            fmt.Printf("Error while getting Addresses List %s", err.Error())
        }

        //fmt.Printf("%s\n", s.Items)
        for _, i := range s.Items {
            //fmt.Printf("%s, %s, %s, %s, %s, %s\n", i.Address, i.AddressType, i.Id, i.Name, i.IpVersion, i.Status)
            var json IpAddress
            json.Name = i.Name
            json.Cloud = "GoogleCloud"
            json.CloudID = strconv.FormatUint(i.Id,10)
            json.AssetType = "ipaddress"
            json.Status = i.Status          
            json.Address = i.Address
            PushData("success", json)
       }
    }

    //autoscalers
    //{
    //    s, err := srv.Instances.List(pr,"us-central1-c").Do()
    //    if err != nil {
    //        fmt.Printf("Error while getting Instances List %s", err.Error())
    //    }

    //    //fmt.Printf("%s\n", s.Id)
    //    for _, i := range s.Items {
    //        fmt.Printf("%s, %s, %s, %s, %s, %s\n", i.MachineType, i.Name, i.CpuPlatform, i.Status, i.NetworkInterfaces[0].Name, i.NetworkInterfaces[0].NetworkIP)
    //    }
    //}

    //backendBuckets
    //{
    //    s, err := srv.Instances.List(pr,"us-central1-c").Do()
    //    if err != nil {
    //        fmt.Printf("Error while getting Instances List %s", err.Error())
    //    }

        //fmt.Printf("%s\n", s.Id)
    //    for _, i := range s.Items {
    //        fmt.Printf("%s, %s, %s, %s, %s, %s\n", i.MachineType, i.Name, i.CpuPlatform, i.Status, i.NetworkInterfaces[0].Name, i.NetworkInterfaces[0].NetworkIP)
    //    }
    //}

    //backendServices
    //{
    //    s, err := srv.Instances.List(pr,"us-central1-c").Do()
    //    if err != nil {
    //        fmt.Printf("Error while getting Instances List %s", err.Error())
    //    }

        //fmt.Printf("%s\n", s.Id)
    //    for _, i := range s.Items {
    //        fmt.Printf("%s, %s, %s, %s, %s, %s\n", i.MachineType, i.Name, i.CpuPlatform, i.Status, i.NetworkInterfaces[0].Name, i.NetworkInterfaces[0].NetworkIP)
    //    }
    //}

    //disktypes
        {
        s, err := srv.DiskTypes.List(pr,"us-central1-c").Do()
        if err != nil {
            fmt.Printf("Error while getting Instances List %s", err.Error())
        }

        //fmt.Printf("%s\n", s)
        for _, i := range s.Items {
            fmt.Printf("%s, %s, %s, %s\n", i.Name, strconv.FormatInt(i.DefaultDiskSizeGb,10), i.Description, i.ValidDiskSize)
            
        }
    }
    //disks
    {
        s, err := srv.Disks.List(pr,"us-central1-c").Do()
        if err != nil {
            fmt.Printf("Error while getting Instances List %s", err.Error())
        }

        //fmt.Printf("%s\n", s)
        for _, i := range s.Items {
            //fmt.Printf("%s, %s, %s, %s, %s, %s, %s, %s, %s\n", i.Description, strconv.FormatUint(i.Id,10), i.Name, strconv.FormatInt(i.SizeGb,10), i.SourceImage, i.Status, i.Type, i.Region, i.Zone)
            var json Disk
            json.Name = i.Name
            json.Cloud = "GoogleCloud"
            json.CloudID = strconv.FormatUint(i.Id,10)
            json.AssetType = "disk"
            json.Status = i.Status
            json.Size = strconv.FormatInt(i.SizeGb,10)
            json.DiskType = GetLastUrlSegment(i.Type)
            json.Az = GetLastUrlSegment(i.Zone)
            PushData("success", json)
        }
    }

    fmt.Printf("\n")    
    fmt.Printf("\n")    
    fmt.Printf("\n")    
    fmt.Printf("\n")    
    //firewalls
        {
        s, err := srv.Firewalls.List(pr).Do()
        if err != nil {
            fmt.Printf("Error while getting Instances List %s", err.Error())
        }

        //fmt.Printf("%s\n", s)
        for _, i := range s.Items {
            //fmt.Printf("%s, %s, %s, %s, %s, %s, %s\n", i.Description, i.DestinationRanges, i.Direction, strconv.FormatUint(i.Id,10), i.Name, i.Network, i.SourceRanges)
            var json Firewall
            json.Name = i.Name
            json.Cloud = "GoogleCloud"
            json.CloudID = strconv.FormatUint(i.Id,10)
            json.AssetType = "firewall"
            json.Status = "available"
            PushData("success", json)
        }
    }

    //forwardingRules
    //    {
    //    s, err := srv.ForwardingRules.List(pr,"us-central1").Do()
    //    if err != nil {
    //        fmt.Printf("Error while getting Instances List %s", err.Error())
    //    }

    //    fmt.Printf("%s\n", s)
//        for _, i := range s.Items {
//            fmt.Printf("%s, %s, %s, %s, %s, %s\n", i.MachineType, i.Name, i.CpuPlatform, i.Status, i.NetworkInterfaces[0].Name, i.NetworkInterfaces[0].NetworkIP)
//        }
    //}

    //globalAddresses
    //    {
    //    s, err := srv.Instances.List(pr,"us-central1-c").Do()
    //    if err != nil {
    //        fmt.Printf("Error while getting Instances List %s", err.Error())
    //    }

        //fmt.Printf("%s\n", s.Id)
    //    for _, i := range s.Items {
    //        fmt.Printf("%s, %s, %s, %s, %s, %s\n", i.MachineType, i.Name, i.CpuPlatform, i.Status, i.NetworkInterfaces[0].Name, i.NetworkInterfaces[0].NetworkIP)
    //    }
    //}

    //globalForwardingrules
    //{
    //    s, err := srv.Instances.List(pr,"us-central1-c").Do()
    //    if err != nil {
    //        fmt.Printf("Error while getting Instances List %s", err.Error())
    //    }

        //fmt.Printf("%s\n", s.Id)
    //    for _, i := range s.Items {
    //        fmt.Printf("%s, %s, %s, %s, %s, %s\n", i.MachineType, i.Name, i.CpuPlatform, i.Status, i.NetworkInterfaces[0].Name, i.NetworkInterfaces[0].NetworkIP)
    //    }
    //}

    //globalOperations
    //{
    //    s, err := srv.Instances.List(pr,"us-central1-c").Do()
    //    if err != nil {
    //        fmt.Printf("Error while getting Instances List %s", err.Error())
    //    }

        //fmt.Printf("%s\n", s.Id)
    //    for _, i := range s.Items {
    //        fmt.Printf("%s, %s, %s, %s, %s, %s\n", i.MachineType, i.Name, i.CpuPlatform, i.Status, i.NetworkInterfaces[0].Name, i.NetworkInterfaces[0].NetworkIP)
    //    }
    //}

    //healthchecks
    //    {
    //    s, err := srv.Instances.List(pr,"us-central1-c").Do()
    //    if err != nil {
    //        fmt.Printf("Error while getting Instances List %s", err.Error())
    //    }

        //fmt.Printf("%s\n", s.Id)
    //    for _, i := range s.Items {
    //        fmt.Printf("%s, %s, %s, %s, %s, %s\n", i.MachineType, i.Name, i.CpuPlatform, i.Status, i.NetworkInterfaces[0].Name, i.NetworkInterfaces[0].NetworkIP)
    //    }
    //}

    //httphealthchecks
    //    {
    //    s, err := srv.Instances.List(pr,"us-central1-c").Do()
    //    if err != nil {
    //        fmt.Printf("Error while getting Instances List %s", err.Error())
    //    }

        //fmt.Printf("%s\n", s.Id)
    //    for _, i := range s.Items {
    //        fmt.Printf("%s, %s, %s, %s, %s, %s\n", i.MachineType, i.Name, i.CpuPlatform, i.Status, i.NetworkInterfaces[0].Name, i.NetworkInterfaces[0].NetworkIP)
    //    }
    //}

    //httpshealthchecks
    //{
    //    s, err := srv.Instances.List(pr,"us-central1-c").Do()
    //    if err != nil {
    //        fmt.Printf("Error while getting Instances List %s", err.Error())
    //    }

        //fmt.Printf("%s\n", s.Id)
    //    for _, i := range s.Items {
    //        fmt.Printf("%s, %s, %s, %s, %s, %s\n", i.MachineType, i.Name, i.CpuPlatform, i.Status, i.NetworkInterfaces[0].Name, i.NetworkInterfaces[0].NetworkIP)
    //    }
    //}

    //images
    {
        s, err := srv.Images.List(pr).Do()
        if err != nil {
            fmt.Printf("Error while getting Instances List %s", err.Error())
        }

        //fmt.Printf("%s\n", s)
        for _, i := range s.Items {
            //fmt.Printf("%s, %s, %s, %s, %s, %s\n", i.Description, strconv.FormatInt(i.DiskSizeGb,10), i.Family, strconv.FormatUint(i.Id,10), i.Name, i.Status)
            var json Image
            json.Name = i.Name
            json.Cloud = "GoogleCloud"
            json.CloudID = i.Name
            json.AssetType = "image"
            json.Status = i.Status
            json.Size = strconv.FormatInt(i.DiskSizeGb,10)
            json.ImageType = i.Family
            json.Distro = i.Description
            PushData("success", json)
        }
    }
    //instancegroupmanagers
    {
        s, err := srv.InstanceGroupManagers.List(pr,"us-central1-c").Do()
        if err != nil {
            fmt.Printf("Error while getting Instances List %s", err.Error())
        }

        //fmt.Printf("%s\n", s)
        for _, i := range s.Items {
            fmt.Printf("%s, %s, %s, %s, %s, %s\n", i.BaseInstanceName, i.Description, i.InstanceGroup, i.Name, i.Region, i.Zone)
        }
    }
    //instancegroups
        {
        s, err := srv.InstanceGroups.List(pr,"us-central1-c").Do()
        if err != nil {
            fmt.Printf("Error while getting Instances List %s", err.Error())
        }

        //fmt.Printf("%s\n", s)
        for _, i := range s.Items {
            fmt.Printf("%s, %s, %s, %s, %s, %s\n", i.Name, i.Network, i.Region, i.Subnetwork, strconv.FormatInt(i.Size,10), i.Zone)
        }
    }
    //instancetemplates
        {
        s, err := srv.InstanceTemplates.List(pr).Do()
        if err != nil {
            fmt.Printf("Error while getting Instances List %s", err.Error())
        }

        //fmt.Printf("%s\n", s)
        for _, i := range s.Items {
            fmt.Printf("%s, %s, %s, %s\n", i.Description, i.Name, i.SourceInstance, strconv.FormatUint(i.Id,10))
        }
    }

    //interconnectAttachments
    //{
    //    s, err := srv.Instances.List(pr,"us-central1-c").Do()
    //    if err != nil {
    //        fmt.Printf("Error while getting Instances List %s", err.Error())
    //    }

        //fmt.Printf("%s\n", s.Id)
    //    for _, i := range s.Items {
    //        fmt.Printf("%s, %s, %s, %s, %s, %s\n", i.MachineType, i.Name, i.CpuPlatform, i.Status, i.NetworkInterfaces[0].Name, i.NetworkInterfaces[0].NetworkIP)
    //    }
    //}

    //interconnectLocations
    //{
    //    s, err := srv.Instances.List(pr,"us-central1-c").Do()
    //    if err != nil {
    //        fmt.Printf("Error while getting Instances List %s", err.Error())
    //    }

        //fmt.Printf("%s\n", s.Id)
    //    for _, i := range s.Items {
    //        fmt.Printf("%s, %s, %s, %s, %s, %s\n", i.MachineType, i.Name, i.CpuPlatform, i.Status, i.NetworkInterfaces[0].Name, i.NetworkInterfaces[0].NetworkIP)
    //    }
    //}

    //interconnects
    //{
    //    s, err := srv.Instances.List(pr,"us-central1-c").Do()
    //    if err != nil {
    //        fmt.Printf("Error while getting Instances List %s", err.Error())
    //    }

        //fmt.Printf("%s\n", s.Id)
    //    for _, i := range s.Items {
    //        fmt.Printf("%s, %s, %s, %s, %s, %s\n", i.MachineType, i.Name, i.CpuPlatform, i.Status, i.NetworkInterfaces[0].Name, i.NetworkInterfaces[0].NetworkIP)
    //    }
    //}

    //licenses
    //{
    //    s, err := srv.Instances.List(pr,"us-central1-c").Do()
    //    if err != nil {
    //        fmt.Printf("Error while getting Instances List %s", err.Error())
    //    }

        //fmt.Printf("%s\n", s.Id)
    //    for _, i := range s.Items {
    //        fmt.Printf("%s, %s, %s, %s, %s, %s\n", i.MachineType, i.Name, i.CpuPlatform, i.Status, i.NetworkInterfaces[0].Name, i.NetworkInterfaces[0].NetworkIP)
    //    }
    //}

    //machineTypes
    {
        s, err := srv.MachineTypes.List(pr,"us-central1-c").Do()
        if err != nil {
            fmt.Printf("Error while getting Instances List %s", err.Error())
        }

        //fmt.Printf("%s\n", s)
        for _, i := range s.Items {
            //fmt.Printf("%s, %s, %s, %s, %s, %s\n", i.Description, strconv.FormatInt(i.GuestCpus,10), strconv.FormatInt(i.ImageSpaceGb,10), strconv.FormatInt(i.MemoryMb,10), i.Name, i.Zone)
            var json VMSize
            json.Name = i.Name
            json.Cloud = "GoogleCloud"
            json.CloudID = i.Name
            json.AssetType = "vmsize"
            json.Status = "available"
            json.Vcpu = strconv.FormatInt(i.GuestCpus,10)
            json.Memory = strconv.FormatInt(i.MemoryMb,10)
            json.Disk = strconv.FormatInt(i.ImageSpaceGb,10)
            json.Available = "available"
            PushData("success", json)
        }
    }

    //networks
    {
        s, err := srv.Networks.List(pr).Do()
        if err != nil {
            fmt.Printf("Error while getting Instances List %s", err.Error())
        }

        //fmt.Printf("%s\n", s)
        for _, i := range s.Items {
            //fmt.Printf("%s, %s, %s, %s, %s\n", i.IPv4Range, i.Description, i.GatewayIPv4, i.Name, i.Subnetworks)
            var json Network
            json.Name = i.Name
            json.Cloud = "GoogleCloud"
            json.CloudID = i.Name
            json.AssetType = "network"
            json.Status = "available"
            json.Range = i.IPv4Range
            PushData("success", json)
        }
    }

    //nodeGroups
    //{
    //    s, err := srv.Instances.List(pr,"us-central1-c").Do()
    //    if err != nil {
    //        fmt.Printf("Error while getting Instances List %s", err.Error())
    //    }

        //fmt.Printf("%s\n", s.Id)
    //    for _, i := range s.Items {
    //        fmt.Printf("%s, %s, %s, %s, %s, %s\n", i.MachineType, i.Name, i.CpuPlatform, i.Status, i.NetworkInterfaces[0].Name, i.NetworkInterfaces[0].NetworkIP)
    //    }
    //}
    
    //nodeTemplates
    //{
    //    s, err := srv.Instances.List(pr,"us-central1-c").Do()
    //    if err != nil {
    //        fmt.Printf("Error while getting Instances List %s", err.Error())
    //    }

        //fmt.Printf("%s\n", s.Id)
    //    for _, i := range s.Items {
    //        fmt.Printf("%s, %s, %s, %s, %s, %s\n", i.MachineType, i.Name, i.CpuPlatform, i.Status, i.NetworkInterfaces[0].Name, i.NetworkInterfaces[0].NetworkIP)
    //    }
    //}

    //nodeTypes
    //{
    //    s, err := srv.Instances.List(pr,"us-central1-c").Do()
    //    if err != nil {
    //        fmt.Printf("Error while getting Instances List %s", err.Error())
    //    }

        //fmt.Printf("%s\n", s.Id)
    //    for _, i := range s.Items {
    //        fmt.Printf("%s, %s, %s, %s, %s, %s\n", i.MachineType, i.Name, i.CpuPlatform, i.Status, i.NetworkInterfaces[0].Name, i.NetworkInterfaces[0].NetworkIP)
    //    }
    //}

    //regionAutoscalers
    //{
    //    s, err := srv.Instances.List(pr,"us-central1-c").Do()
    //    if err != nil {
    //        fmt.Printf("Error while getting Instances List %s", err.Error())
    //    }

        //fmt.Printf("%s\n", s.Id)
    //    for _, i := range s.Items {
    //        fmt.Printf("%s, %s, %s, %s, %s, %s\n", i.MachineType, i.Name, i.CpuPlatform, i.Status, i.NetworkInterfaces[0].Name, i.NetworkInterfaces[0].NetworkIP)
    //    }
    //}

    //regionBackendServices
    //{
    //    s, err := srv.Instances.List(pr,"us-central1-c").Do()
    //    if err != nil {
    //        fmt.Printf("Error while getting Instances List %s", err.Error())
    //    }

        //fmt.Printf("%s\n", s.Id)
    //    for _, i := range s.Items {
    //        fmt.Printf("%s, %s, %s, %s, %s, %s\n", i.MachineType, i.Name, i.CpuPlatform, i.Status, i.NetworkInterfaces[0].Name, i.NetworkInterfaces[0].NetworkIP)
    //    }
    //}

    //regionCommitments
    //{
    //    s, err := srv.Instances.List(pr,"us-central1-c").Do()
    //    if err != nil {
    //        fmt.Printf("Error while getting Instances List %s", err.Error())
    //    }

        //fmt.Printf("%s\n", s.Id)
    //    for _, i := range s.Items {
    //        fmt.Printf("%s, %s, %s, %s, %s, %s\n", i.MachineType, i.Name, i.CpuPlatform, i.Status, i.NetworkInterfaces[0].Name, i.NetworkInterfaces[0].NetworkIP)
    //    }
    //}

    //regionDisktypes
    //{
    //    s, err := srv.Instances.List(pr,"us-central1-c").Do()
    //    if err != nil {
    //        fmt.Printf("Error while getting Instances List %s", err.Error())
    //    }

        //fmt.Printf("%s\n", s.Id)
    //    for _, i := range s.Items {
    //        fmt.Printf("%s, %s, %s, %s, %s, %s\n", i.MachineType, i.Name, i.CpuPlatform, i.Status, i.NetworkInterfaces[0].Name, i.NetworkInterfaces[0].NetworkIP)
    //    }
    //}

    //regionDisks
    //{
    //    s, err := srv.Instances.List(pr,"us-central1-c").Do()
    //    if err != nil {
    //        fmt.Printf("Error while getting Instances List %s", err.Error())
    //    }

    //    //fmt.Printf("%s\n", s.Id)
    //    for _, i := range s.Items {
    //        fmt.Printf("%s, %s, %s, %s, %s, %s\n", i.MachineType, i.Name, i.CpuPlatform, i.Status, i.NetworkInterfaces[0].Name, i.NetworkInterfaces[0].NetworkIP)
    //    }
    //}

    //regionInstanceGroupManagers
    //{
    //    s, err := srv.Instances.List(pr,"us-central1-c").Do()
    //    if err != nil {
    //        fmt.Printf("Error while getting Instances List %s", err.Error())
    //    }

        //fmt.Printf("%s\n", s.Id)
    //    for _, i := range s.Items {
    //        fmt.Printf("%s, %s, %s, %s, %s, %s\n", i.MachineType, i.Name, i.CpuPlatform, i.Status, i.NetworkInterfaces[0].Name, i.NetworkInterfaces[0].NetworkIP)
    //    }
    //}

    //regionInstanceGroups
    //{
    //    s, err := srv.Instances.List(pr,"us-central1-c").Do()
    //    if err != nil {
    //        fmt.Printf("Error while getting Instances List %s", err.Error())
    //    }

        //fmt.Printf("%s\n", s.Id)
    //    for _, i := range s.Items {
    //        fmt.Printf("%s, %s, %s, %s, %s, %s\n", i.MachineType, i.Name, i.CpuPlatform, i.Status, i.NetworkInterfaces[0].Name, i.NetworkInterfaces[0].NetworkIP)
    //    }
    //}

    //RegionOperations
    //{
    //    s, err := srv.Instances.List(pr,"us-central1-c").Do()
    //    if err != nil {
    //        fmt.Printf("Error while getting Instances List %s", err.Error())
    //    }

        //fmt.Printf("%s\n", s.Id)
    //    for _, i := range s.Items {
    //        fmt.Printf("%s, %s, %s, %s, %s, %s\n", i.MachineType, i.Name, i.CpuPlatform, i.Status, i.NetworkInterfaces[0].Name, i.NetworkInterfaces[0].NetworkIP)
    //    }
    //}

    //Regions
    {
        s, err := srv.Regions.List(pr).Do()
        if err != nil {
            fmt.Printf("Error while getting Instances List %s", err.Error())
        }

        //fmt.Printf("%s\n", s)
        for _, i := range s.Items {
            //fmt.Printf("%s, %s, %s, %s\n", i.Description, i.Name, i.Status, i.Zones)
            var json Region
            json.Name = i.Name
            json.Cloud = "GoogleCloud"
            json.CloudID = i.Name
            json.AssetType = "region"
            json.Status = i.Status
            PushData("success", json)
        }
    }

    //Routers
    {
        s, err := srv.Routers.List(pr,"us-central1").Do()
        if err != nil {
            fmt.Printf("Error while getting Instances List %s", err.Error())
        }

        //fmt.Printf("%s\n", s)
        for _, i := range s.Items {
            //fmt.Printf("%s, %s, %s, %s\n", i.Description, i.Name, i.Network, i.Region)
            var json Router
            json.Name = i.Name
            json.Cloud = "GoogleCloud"
            json.CloudID = i.Name
            json.AssetType = "router"
            json.Status = "available"
            json.Description = i.Description
            json.Network = GetLastUrlSegment(i.Network)
            json.Region = GetLastUrlSegment(i.Region)
            PushData("success", json)
        }
    }

    //routes
    {
        s, err := srv.Routes.List(pr).Do()
        if err != nil {
            fmt.Printf("Error while getting Instances List %s", err.Error())
        }

        //fmt.Printf("%s\n", s)
        for _, i := range s.Items {
            //fmt.Printf("%s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s\n", i.Name, i.Description, i.DestRange, i.Network, i.NextHopGateway, i.NextHopInstance, i.NextHopIp, i.NextHopNetwork, i.NextHopPeering, i.NextHopVpnTunnel, strconv.FormatInt(i.Priority, 10))
            var json Routes
            json.Name = i.Name
            json.Cloud = "GoogleCloud"
            json.CloudID = i.Name
            json.AssetType = "route"
            json.Status = "available"
            json.Description = i.Description
            json.DstRange = i.DestRange
            json.Network = GetLastUrlSegment(i.Network)
            json.NextHopGW = GetLastUrlSegment(i.NextHopGateway)
            json.NextHopVM = GetLastUrlSegment(i.NextHopInstance)
            json.NextHopIP = i.NextHopIp
            json.NextHopNet = GetLastUrlSegment(i.NextHopNetwork)
            json.NextHopPeer = i.NextHopPeering
            json.NextHopVpn = GetLastUrlSegment(i.NextHopVpnTunnel)
            json.Priority = strconv.FormatInt(i.Priority, 10)
            PushData("success", json)
        }
    }

    //securityPolicies
    //{
    //    s, err := srv.Instances.List(pr,"us-central1-c").Do()
    //    if err != nil {
    //        fmt.Printf("Error while getting Instances List %s", err.Error())
    //    }

        //fmt.Printf("%s\n", s.Id)
    //    for _, i := range s.Items {
    //        fmt.Printf("%s, %s, %s, %s, %s, %s\n", i.MachineType, i.Name, i.CpuPlatform, i.Status, i.NetworkInterfaces[0].Name, i.NetworkInterfaces[0].NetworkIP)
    //    }
    //}

    fmt.Printf("\n")    
    fmt.Printf("\n")    
    fmt.Printf("\n")    
    fmt.Printf("\n")    
    //snapshots
    {
        s, err := srv.Snapshots.List(pr).Do()
        if err != nil {
            fmt.Printf("Error while getting Instances List %s", err.Error())
        }

        //fmt.Printf("%s\n", s)
        for _, i := range s.Items {
            //fmt.Printf("%s, %s, %s, %s, %s, %s, %s, %s\n", i.Description, strconv.FormatInt(i.DiskSizeGb,10), i.Name, i.SourceDisk, i.SourceDiskId, i.Status, strconv.FormatInt(i.StorageBytes,10), i.StorageBytesStatus)
            var json Snapshot
            json.Name = i.Name
            json.Cloud = "GoogleCloud"
            json.CloudID = i.Name
            json.AssetType = "snapshot"
            json.Status = i.Status
            json.Disk = i.SourceDisk
            PushData("success", json)
        }
    }

    //sslCertificates
    //{
    //    s, err := srv.Instances.List(pr,"us-central1-c").Do()
    //    if err != nil {
    //        fmt.Printf("Error while getting Instances List %s", err.Error())
    //    }

        //fmt.Printf("%s\n", s.Id)
    //    for _, i := range s.Items {
    //        fmt.Printf("%s, %s, %s, %s, %s, %s\n", i.MachineType, i.Name, i.CpuPlatform, i.Status, i.NetworkInterfaces[0].Name, i.NetworkInterfaces[0].NetworkIP)
    //    }
    //}

    //sslPolicies
    //{
    //    s, err := srv.Instances.List(pr,"us-central1-c").Do()
    //    if err != nil {
    //        fmt.Printf("Error while getting Instances List %s", err.Error())
    //    }

        //fmt.Printf("%s\n", s.Id)
    //    for _, i := range s.Items {
    //        fmt.Printf("%s, %s, %s, %s, %s, %s\n", i.MachineType, i.Name, i.CpuPlatform, i.Status, i.NetworkInterfaces[0].Name, i.NetworkInterfaces[0].NetworkIP)
    //    }
    //}

    fmt.Printf("\n")    
    fmt.Printf("\n")    
    fmt.Printf("\n")    
    fmt.Printf("\n")    
    //subnetworks
    {
        s, err := srv.Subnetworks.List(pr,"us-central1").Do()
        if err != nil {
            fmt.Printf("Error while getting Instances List %s", err.Error())
        }

        //fmt.Printf("%s\n", s)
        for _, i := range s.Items {
            //fmt.Printf("%s, %s, %s, %s, %s, %s\n", i.Description, i.GatewayAddress, i.IpCidrRange, i.Name, i.Network, i.Region)
            var json Subnet
            json.Name = i.Name
            json.Cloud = "GoogleCloud"
            json.CloudID = i.Name
            json.AssetType = "subnet"
            json.Status = "available"
            json.Range = i.IpCidrRange
            json.NetworkID = i.Network
            PushData("success", json)
        }
    }

    //targetHttpProxies
    //{
    //    s, err := srv.Instances.List(pr,"us-central1-c").Do()
    //    if err != nil {
    //        fmt.Printf("Error while getting Instances List %s", err.Error())
    //    }

        //fmt.Printf("%s\n", s.Id)
    //    for _, i := range s.Items {
    //        fmt.Printf("%s, %s, %s, %s, %s, %s\n", i.MachineType, i.Name, i.CpuPlatform, i.Status, i.NetworkInterfaces[0].Name, i.NetworkInterfaces[0].NetworkIP)
    //    }
    //}

    //targethttpsProxies
    //{
    //    s, err := srv.Instances.List(pr,"us-central1-c").Do()
    //    if err != nil {
    //        fmt.Printf("Error while getting Instances List %s", err.Error())
    //    }

        //fmt.Printf("%s\n", s.Id)
    //    for _, i := range s.Items {
    //        fmt.Printf("%s, %s, %s, %s, %s, %s\n", i.MachineType, i.Name, i.CpuPlatform, i.Status, i.NetworkInterfaces[0].Name, i.NetworkInterfaces[0].NetworkIP)
    //    }
    //}

    //targetinstances
    //{
    //    s, err := srv.Instances.List(pr,"us-central1-c").Do()
    //    if err != nil {
    //        fmt.Printf("Error while getting Instances List %s", err.Error())
    //    }

        //fmt.Printf("%s\n", s.Id)
    //    for _, i := range s.Items {
    //        fmt.Printf("%s, %s, %s, %s, %s, %s\n", i.MachineType, i.Name, i.CpuPlatform, i.Status, i.NetworkInterfaces[0].Name, i.NetworkInterfaces[0].NetworkIP)
    //    }
    //}

    //targetPools
    //{
    //    s, err := srv.Instances.List(pr,"us-central1-c").Do()
    //    if err != nil {
    //        fmt.Printf("Error while getting Instances List %s", err.Error())
    //    }

        //fmt.Printf("%s\n", s.Id)
    //    for _, i := range s.Items {
    //        fmt.Printf("%s, %s, %s, %s, %s, %s\n", i.MachineType, i.Name, i.CpuPlatform, i.Status, i.NetworkInterfaces[0].Name, i.NetworkInterfaces[0].NetworkIP)
    //    }
    //}

    //targetSslProxies
    //{
    //    s, err := srv.Instances.List(pr,"us-central1-c").Do()
    //    if err != nil {
    //        fmt.Printf("Error while getting Instances List %s", err.Error())
    //    }

        //fmt.Printf("%s\n", s.Id)
    //    for _, i := range s.Items {
    //        fmt.Printf("%s, %s, %s, %s, %s, %s\n", i.MachineType, i.Name, i.CpuPlatform, i.Status, i.NetworkInterfaces[0].Name, i.NetworkInterfaces[0].NetworkIP)
    //    }
    //}

    //targetTcpProxies
    //{
    //    s, err := srv.Instances.List(pr,"us-central1-c").Do()
    //    if err != nil {
    //        fmt.Printf("Error while getting Instances List %s", err.Error())
    //    }

        //fmt.Printf("%s\n", s.Id)
    //    for _, i := range s.Items {
    //        fmt.Printf("%s, %s, %s, %s, %s, %s\n", i.MachineType, i.Name, i.CpuPlatform, i.Status, i.NetworkInterfaces[0].Name, i.NetworkInterfaces[0].NetworkIP)
    //    }
    //}

    //targetVpnGateways
    //{
    //    s, err := srv.Instances.List(pr,"us-central1-c").Do()
    //    if err != nil {
    //        fmt.Printf("Error while getting Instances List %s", err.Error())
    //    }

        //fmt.Printf("%s\n", s.Id)
    //    for _, i := range s.Items {
    //        fmt.Printf("%s, %s, %s, %s, %s, %s\n", i.MachineType, i.Name, i.CpuPlatform, i.Status, i.NetworkInterfaces[0].Name, i.NetworkInterfaces[0].NetworkIP)
    //    }
    //}

    //urlMaps
    //{
    //    s, err := srv.Instances.List(pr,"us-central1-c").Do()
    //    if err != nil {
    //        fmt.Printf("Error while getting Instances List %s", err.Error())
    //    }

        //fmt.Printf("%s\n", s.Id)
    //    for _, i := range s.Items {
    //        fmt.Printf("%s, %s, %s, %s, %s, %s\n", i.MachineType, i.Name, i.CpuPlatform, i.Status, i.NetworkInterfaces[0].Name, i.NetworkInterfaces[0].NetworkIP)
    //    }
    //}

    //vpnTunnels
    //{
    //    s, err := srv.Instances.List(pr,"us-central1-c").Do()
    //    if err != nil {
    //        fmt.Printf("Error while getting Instances List %s", err.Error())
    //    }

        //fmt.Printf("%s\n", s.Id)
    //    for _, i := range s.Items {
    //        fmt.Printf("%s, %s, %s, %s, %s, %s\n", i.MachineType, i.Name, i.CpuPlatform, i.Status, i.NetworkInterfaces[0].Name, i.NetworkInterfaces[0].NetworkIP)
    //    }
    //}

    //ZoneOperations
    //{
    //    s, err := srv.Instances.List(pr,"us-central1-c").Do()
    //    if err != nil {
    //        fmt.Printf("Error while getting Instances List %s", err.Error())
    //    }

        //fmt.Printf("%s\n", s.Id)
    //    for _, i := range s.Items {
    //        fmt.Printf("%s, %s, %s, %s, %s, %s\n", i.MachineType, i.Name, i.CpuPlatform, i.Status, i.NetworkInterfaces[0].Name, i.NetworkInterfaces[0].NetworkIP)
    //    }
    //}

}

