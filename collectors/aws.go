package collectors

import "github.com/aws/aws-sdk-go/aws"
import "github.com/aws/aws-sdk-go/aws/session"
// import "github.com/aws/aws-sdk-go/aws/credentials"
import "github.com/aws/aws-sdk-go/service/ec2"
//import "github.com/aws/aws-sdk-go/service/elbv2"
//import "github.com/aws/aws-sdk-go/service/elb"
import "fmt"
import "strconv"

func Aws() {

    sess, err := session.NewSession(&aws.Config{
        Region:         aws.String("us-west-2"),
        // Credentials:    credentials.NewStaticCredentials("apikey", "apisecret",""),
})
    if err != nil {
        fmt.Println("there was an error ssession in", err.Error())
    }

    ec2svc := ec2.New(sess)
//    elbv2svc := elbv2.New(sess)
//    elbsvc := elb.New(sess)

    // instances
    {
        input := &ec2.DescribeInstancesInput{
            Filters: []*ec2.Filter{
                {
                    Name:   aws.String("instance-state-name"),
                    Values: []*string{aws.String("running"), aws.String("pending")},
                },
            },
        }
        resp, err := ec2svc.DescribeInstances(input)
        if err != nil {
            fmt.Println("there was an error listing instances in", err.Error())
        }

        for idx, _ := range resp.Reservations {
            for _, inst := range resp.Reservations[idx].Instances {
                var json VirtualMachine
                json.Name = *inst.InstanceId
                json.Cloud = "Aws"
                json.CloudID = *inst.InstanceId
                json.AssetType = "instance"
                json.Status = *inst.State.Name
                json.Memory = ""
                json.Vcpu = strconv.FormatInt(*inst.CpuOptions.CoreCount, 10)
                json.Image = *inst.ImageId
                json.Region = ""
                json.Az = *inst.Placement.AvailabilityZone
                json.Arch = *inst.Architecture
                json.VMSize = *inst.InstanceType
                json.Key = *inst.KeyName
                json.PrivateDNS = *inst.PrivateDnsName
                json.PrivateIP = *inst.PrivateIpAddress
                json.PublicDNS = *inst.PublicDnsName
                json.PublicIP = *inst.PublicIpAddress
                json.Firewall = *inst.SecurityGroups[0].GroupId
                json.Network = *inst.VpcId
                json.Subnet = *inst.SubnetId
                json.Virt = *inst.VirtualizationType
                json.Hyper = *inst.Hypervisor
                PushData("success", json)
            }
        }
    }

    // images
    {
        input := &ec2.DescribeImagesInput{
            Owners: []*string{
                aws.String("self"),
            },
        }
        resp, err := ec2svc.DescribeImages(input)
        if err != nil {
            fmt.Println("there was an error listing images in", err.Error())
        }
        for _, a := range resp.Images {
//            fmt.Printf("%s %s %s %t %s\n", *a.Name, *a.ImageId, *a.State, *a.Public, *a.RootDeviceType)
            var json Image
            json.CloudID = *a.ImageId
            json.Name = *a.Name
            json.Cloud = "Aws"
            json.AssetType = "image"
            json.ImageType = "EBS"
            json.Distro = *a.RootDeviceType
            json.Size = "0"
            PushData("success", json)
        }
    }

    // keypair
    {
        input := &ec2.DescribeKeyPairsInput{}
        resp, err := ec2svc.DescribeKeyPairs(input)
        if err != nil {
            fmt.Println("there was an error listing images in", err.Error())
        }
        for _, a := range resp.KeyPairs {
//            fmt.Printf("%s %s\n", *a.KeyName, *a.KeyFingerprint)
            var json Key
            json.CloudID = *a.KeyName
            json.Name = *a.KeyName
            json.Cloud = "Aws"
            json.AssetType = "key"
            json.Fingerprint = *a.KeyFingerprint
            json.PublicKey = ""
            PushData("success", json)
        }
    }

    // Availablity zones
    {
        input := &ec2.DescribeAvailabilityZonesInput{}
        resp, err := ec2svc.DescribeAvailabilityZones(input)
        if err != nil {
            fmt.Println("there was an error listing az in", err.Error())
        }
        for _, a := range resp.AvailabilityZones {
            var json Az
            json.CloudID = *a.ZoneName
            json.Cloud = "Aws"
            json.Name = *a.ZoneName
            json.AssetType = "availabilityzone"
            json.Status = *a.State
            json.Region = *a.RegionName
            PushData("success", json)
        }
    }

    // customer gateways
/*    {
        input := &ec2.DescribeCustomerGatewaysInput{}
        resp, err := ec2svc.DescribeCustomerGateways(input)
        if err != nil {
            fmt.Println("there was an error listing customer gateways in", err.Error())
        }
        fmt.Printf("%s\n", resp)
    }
*/
    // egress only internet gateways
/*    {
        input := &ec2.DescribeEgressOnlyInternetGatewaysInput{}
        resp, err := ec2svc.DescribeEgressOnlyInternetGateways(input)
        if err != nil {
            fmt.Println("there was an error listing egress only internet gateways in", err.Error())
        }
        fmt.Printf("%s\n", resp)
    }
*/
    // fleet instances
    // TODO only run if fleetid is available
    //{
    //    input := &ec2.DescribeFleetInstancesInput{
    //        FleetId: aws.String("name"),
    //    }
    //    resp, err := ec2svc.DescribeFleetInstances(input)
    //    if err != nil {
    //        fmt.Println("there was an error listing fleet instances in", err.Error())
    //    }
    //    fmt.Printf("%s\n", resp)
    //}

    // Fleets
/*    {
        input := &ec2.DescribeFleetsInput{}
        resp, err := ec2svc.DescribeFleets(input)
        if err != nil {
            fmt.Println("there was an error listing fleets in", err.Error())
        }
        fmt.Printf("%s\n", resp)
    }
*/
    // Hosts
/*    {
        input := &ec2.DescribeHostsInput{}
        resp, err := ec2svc.DescribeHosts(input)
        if err != nil {
            fmt.Println("there was an error listing hosts in", err.Error())
        }
        fmt.Printf("%s\n", resp)
    }
*/
    // Internet Gateways
    {
        input := &ec2.DescribeInternetGatewaysInput{}
        resp, err := ec2svc.DescribeInternetGateways(input)
        if err != nil {
            fmt.Println("there was an error listing internet gateways in", err.Error())
        }
        for _, a := range resp.InternetGateways {
            var json Az
            json.CloudID = *a.InternetGatewayId
            json.Cloud = "Aws"
            json.Name = *a.InternetGatewayId
            json.AssetType = "internetgateway"
            json.Status = "available"
            PushData("success", json)
        }
    }
    
    // Nat Gateways
    {
        input := &ec2.DescribeNatGatewaysInput{}
        resp, err := ec2svc.DescribeNatGateways(input)
        if err != nil {
            fmt.Println("there was an error listing nat gateways in", err.Error())
        }
        fmt.Printf("%s\n", resp)
    }

    // Network Acls
/*    {
        input := &ec2.DescribeNetworkAclsInput{}
        resp, err := ec2svc.DescribeNetworkAcls(input)
        if err != nil {
            fmt.Println("there was an error listing network acls in", err.Error())
        }
        fmt.Printf("%s\n", resp)
    }
*/
    // Network interfaces
/*    {
        input := &ec2.DescribeNetworkInterfacesInput{}
        resp, err := ec2svc.DescribeNetworkInterfaces(input)
        if err != nil {
            fmt.Println("there was an error listing network interfaces in", err.Error())
        }
        fmt.Printf("%s\n", resp)
    }
*/
    // Placement groups
/*    {
        input := &ec2.DescribePlacementGroupsInput{}
        resp, err := ec2svc.DescribePlacementGroups(input)
        if err != nil {
            fmt.Println("there was an error listing placement groups in", err.Error())
        }
        fmt.Printf("%s\n", resp)
    }
*/
    // Regions
    {
        input := &ec2.DescribeRegionsInput{}
        resp, err := ec2svc.DescribeRegions(input)
        if err != nil {
            fmt.Println("there was an error listing regions in", err.Error())
        }
        for _, a := range resp.Regions {
            var json Region
            json.CloudID = *a.Endpoint
            json.Cloud = "Aws"
            json.Name = *a.RegionName
            json.AssetType = "region"
            json.Status = ""
            json.Location = *a.RegionName
            json.Available = ""
            PushData("success", json)
        }
    }    

    // Reserved instances
/*    {
        input := &ec2.DescribeReservedInstancesInput{}
        resp, err := ec2svc.DescribeReservedInstances(input)
        if err != nil {
            fmt.Println("there was an error listing reserved instances in", err.Error())
        }
        fmt.Printf("%s\n", resp)
    }
*/
    // Routing TAbles
/*    {
        input := &ec2.DescribeRouteTablesInput{}
        resp, err := ec2svc.DescribeRouteTables(input)
        if err != nil {
            fmt.Println("there was an error listing route tables in", err.Error())
        }
        fmt.Printf("%s\n", resp)
    }
*/
    // scheduled instances
/*    {
        input := &ec2.DescribeScheduledInstancesInput{}
        resp, err := ec2svc.DescribeScheduledInstances(input)
        if err != nil {
            fmt.Println("there was an error listing scheduled instances in", err.Error())
        }
        fmt.Printf("%s\n", resp)
    }
*/
    // security groups
    {
        input := &ec2.DescribeSecurityGroupsInput{}
        resp, err := ec2svc.DescribeSecurityGroups(input)
        if err != nil {
            fmt.Println("there was an error listing security groups in", err.Error())
        }
        for _, a := range resp.SecurityGroups {
            var json Firewall
            json.Name = *a.GroupName
            json.CloudID = *a.GroupId
            json.Cloud = "Aws"
            json.AssetType = "firewall"
            json.Status = "Available"
            json.InRules = strconv.Itoa(len(a.IpPermissions))
            json.OutRules = strconv.Itoa(len(a.IpPermissionsEgress))
            PushData("success", json)
        }
    }

    // snapshots        
    // TODO: get owner id and pass it to the input 
    //{
    //    var max int64 = 100
    //    input := &ec2.DescribeSnapshotsInput{
    //        MaxResults: &max,
    //    }
    //    resp, err := ec2svc.DescribeSnapshots(input)
    //    if err != nil {
    //        fmt.Println("there was an error listing snapshots in", err.Error())
    //    }
    //    fmt.Printf("%s\n", resp)
    //}


    // subnets
    {
        input := &ec2.DescribeSubnetsInput{}
        resp, err := ec2svc.DescribeSubnets(input)
        if err != nil {
            fmt.Println("there was an error listing subnets in", err.Error())
        }
        for _, a := range resp.Subnets {
            var json Subnet
            json.Name = *a.SubnetId
            json.CloudID = *a.SubnetId
            json.Cloud = "Aws"
            json.AssetType = "subnet"
            json.Status = *a.State
            json.Az = *a.AvailabilityZone
            json.AvailableIP = strconv.FormatInt(*a.AvailableIpAddressCount,10)
            json.Range = *a.CidrBlock
            json.NetworkID = *a.VpcId
            json.UsePublicIP = strconv.FormatBool(*a.MapPublicIpOnLaunch)
            json.DefaultAz = strconv.FormatBool(*a.DefaultForAz)
            PushData("success", json)
        }
    }

    // volumes
    {
        input := &ec2.DescribeVolumesInput{}
        resp, err := ec2svc.DescribeVolumes(input)
        if err != nil {
            fmt.Println("there was an error listing volumes in", err.Error())
        }
        for _, a := range resp.Volumes {
            var json Disk
            json.Name = *a.VolumeId
            json.Cloud = "Aws"
            json.CloudID = *a.VolumeId
            json.AssetType = "disk"
            json.Status = *a.State
            json.Size = strconv.FormatInt(*a.Size,10)
            json.DiskType = *a.VolumeType
            json.Encrypted = strconv.FormatBool(*a.Encrypted)
            json.SnapshotId = *a.SnapshotId
            json.Az = *a.AvailabilityZone
            PushData("success", json)
        }
    }

    // vpc
    {
        input := &ec2.DescribeVpcsInput{}
        resp, err := ec2svc.DescribeVpcs(input)
        if err != nil {
            fmt.Println("there was an error listing vpc in", err.Error())
        }
        for _, a := range resp.Vpcs {
            var json Network
            json.Name = *a.VpcId
            json.Cloud = "Aws"
            json.CloudID = *a.VpcId
            json.AssetType = "network"
            json.Status = *a.State
            json.Range = *a.CidrBlock
            PushData("success", json)    
        }
    }

    // vpn connections
/*    {
        input := &ec2.DescribeVpnConnectionsInput{}
        resp, err := ec2svc.DescribeVpnConnections(input)
        if err != nil {
            fmt.Println("there was an error listing vpn connections in", err.Error())
        }
        fmt.Printf("%s\n", resp)
    }
*/
    // vpn gateways
/*    {
        input := &ec2.DescribeVpnGatewaysInput{}
        resp, err := ec2svc.DescribeVpnGateways(input)
        if err != nil {
            fmt.Println("there was an error listing vpn gateways in", err.Error())
        }
        fmt.Printf("%s\n", resp)
    }
*/
    // elbv2
/*    {
        input := &elbv2.DescribeLoadBalancersInput{}
        resp, err := elbv2svc.DescribeLoadBalancers(input)
        if err != nil {
            fmt.Println("there was an error listing elbv2 loadbalacners", err.Error())
        }
        fmt.Printf("%s\n", resp)

    }
*/
    // elb
/*    {
        input := &elb.DescribeLoadBalancersInput{}
        resp, err := elbsvc.DescribeLoadBalancers(input)
        if err != nil {
            fmt.Println("there was an error listing elb loadbalancers", err.Error())
        }
      fmt.Printf("%s\n", resp)
    }
*/

}
