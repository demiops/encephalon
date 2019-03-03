package main

import (
    "context"
    "strconv"
//    "reflect"
    "github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2018-10-01/compute"
    "github.com/Azure/azure-sdk-for-go/services/network/mgmt/2018-08-01/network"
    //"github.com/Azure/azure-sdk-for-go/services/storage/mgmt/2018-07-01/storage"
    "fmt"
    "github.com/Azure/go-autorest/autorest/azure/auth"
    "os"
)


func main() {
    os.Setenv("AZURE_TENANT_ID","111111-1111-1111-111111111")
    os.Setenv("AZURE_CLIENT_ID","222222-2222-2222-222222222")
    os.Setenv("AZURE_CLIENT_SECRET","secret")

    computeClient := compute.NewVirtualMachinesClient("3333333-3333-3333-3333333333")
    diskClient := compute.NewDisksClient("33333333-3333-3333-333333333")
    imageClient := compute.NewImagesClient("3333333-33333-3333-333333333-33333")
    containerClient := compute.NewContainerServicesClient("3333333-33333-333333-333333333")
    asetClient := compute.NewAvailabilitySetsClient("33333333-33333-333333-33333333")
    snapshotClient := compute.NewSnapshotsClient("33333-333333-333333-3333333")

    vnetClient := network.NewVirtualNetworksClient("3333333333-33333333-3333333333-333333333")
    pubipClient := network.NewPublicIPAddressesClient("333333333333-3333333-3333333-333333333333")
    inetClient := network.NewInterfacesClient("333333333-33333333-33333333333-33333333333333")
    sgClient := network.NewSecurityGroupsClient("33333333333-333333333333-3333333333-3333333333")

    authorizer, err := auth.NewAuthorizerFromEnvironment()
    if err == nil {
        computeClient.Authorizer = authorizer
        diskClient.Authorizer = authorizer
        imageClient.Authorizer = authorizer
        containerClient.Authorizer = authorizer
        asetClient.Authorizer = authorizer
        snapshotClient.Authorizer = authorizer   
        vnetClient.Authorizer = authorizer
        pubipClient.Authorizer = authorizer             
        inetClient.Authorizer = authorizer
        sgClient.Authorizer = authorizer
    }

    {
        resp, err := computeClient.ListAll(context.Background())
        if err != nil {
            fmt.Printf("Having some problems here %s\n", err.Error())
            return
        }
        //fmt.Printf("%s\n", resp)
        vmlist := resp.Values()
        for  _, v := range vmlist {
            fmt.Printf("%s\n", v.Identity)
            fmt.Printf("%s\n", *v.Name) 
            fmt.Printf("%s\n", *v.ID) 
            fmt.Printf("%s\n", *v.Type) 
            fmt.Printf("%s\n", *v.Location)
            fmt.Printf("%s\n", *v.VirtualMachineProperties.VMID) 
            fmt.Printf("%s\n", *v.VirtualMachineProperties.ProvisioningState) 
            fmt.Printf("%s\n", *v.VirtualMachineProperties.HardwareProfile)
            fmt.Printf("%s\n", *v.VirtualMachineProperties.StorageProfile.ImageReference.Publisher)
            fmt.Printf("%s\n", *v.VirtualMachineProperties.StorageProfile.ImageReference.Offer)
            fmt.Printf("%s\n", *v.VirtualMachineProperties.StorageProfile.ImageReference.Sku)
            fmt.Printf("%s\n", *v.VirtualMachineProperties.StorageProfile.ImageReference.Version)
            fmt.Printf("%s\n", *v.VirtualMachineProperties.StorageProfile.OsDisk.Name)
            fmt.Printf("%s\n", strconv.FormatInt(int64(*v.VirtualMachineProperties.StorageProfile.OsDisk.DiskSizeGB),10))
            fmt.Printf("%s\n", *v.VirtualMachineProperties.OsProfile.ComputerName)
            fmt.Printf("%s\n", *v.VirtualMachineProperties.OsProfile.AdminUsername)
        }
    }

    {
        resp, err := diskClient.List(context.Background())
        if err != nil {
            fmt.Printf("Having some problems here with disks %s\n", err.Error())
            return
        }
        disklist := resp.Values()
        for _, v := range disklist {
            fmt.Printf("%s\n", *v.Name)
            fmt.Printf("%s\n", *v.ManagedBy)
            fmt.Printf("%s\n", *v.ID)
            fmt.Printf("%s\n", *v.Type)
            fmt.Printf("%s\n", *v.Location)
        }
    }
    {
        resp, err := imageClient.List(context.Background())
        if err != nil {
            fmt.Printf("Having some problems here with disks %s\n", err.Error())
            return
        }
        imagelist := resp.Values()
        for _, v := range imagelist {
            fmt.Printf("%s\n", *v.Name)
            fmt.Printf("%s\n", *v.ID)
            fmt.Printf("%s\n", *v.Type)
            fmt.Printf("%s\n", *v.Location)
            fmt.Printf("%s\n", *v.ImageProperties.ProvisioningState)
        }
    }

    {
        resp, err := containerClient.List(context.Background())
        if err != nil {
            fmt.Printf("Having some problems here with disks %s\n", err.Error())
            return
        }
        containerlist := resp.Values()
        for _, v := range containerlist {
            fmt.Printf("%s\n", *v.Name)
            fmt.Printf("%s\n", *v.ID)
            fmt.Printf("%s\n", *v.Type)
            fmt.Printf("%s\n", *v.Location)
            fmt.Printf("%s\n", *v.ContainerServiceProperties.ProvisioningState)
            //fmt.Printf("%s\n", *v.ContainerServiceProperties.OrchestratorProfile.OrchestratorType)
            //fmt.Printf("%s\n", *v.ContainerServiceProperties.OrchestratorProfile.OrchestratorType)
        }
    }

    {
        resp, err := asetClient.List(context.Background(),"demiops")
        if err != nil {
            fmt.Printf("Having some problems here with disks %s\n", err.Error())
            return
        }
        asetlist := resp.Values()
        for _, v := range asetlist {
            fmt.Printf("%s\n", *v.Name)
            fmt.Printf("%s\n", *v.ID)
            fmt.Printf("%s\n", *v.Type)
            fmt.Printf("%s\n", *v.Location)
        }
    }

    {
        resp, err := snapshotClient.List(context.Background())
        if err != nil {
            fmt.Printf("Having some problems here with disks %s\n", err.Error())
            return
        }
        snaplist := resp.Values()
        for _, v := range snaplist {
            fmt.Printf("%s\n", *v.Name)
            fmt.Printf("%s\n", *v.ManagedBy)
            fmt.Printf("%s\n", *v.ID)
            fmt.Printf("%s\n", *v.Type)
            fmt.Printf("%s\n", *v.Location)
        }
    }

    {
        resp, err := vnetClient.List(context.Background(), "demiops")
        if err != nil {
            fmt.Printf("Having some problems here with disks %s\n", err.Error())
            return
        }
        vnetlist := resp.Values()
        for _, v := range vnetlist {
            fmt.Printf("%s\n", *v.Name)
            fmt.Printf("%s\n", *v.Etag)
            fmt.Printf("%s\n", *v.ID)
            fmt.Printf("%s\n", *v.Type)
            fmt.Printf("%s\n", *v.Location)
            fmt.Printf("%s\n", *v.VirtualNetworkPropertiesFormat.ProvisioningState)
            fmt.Printf("%s\n", *v.VirtualNetworkPropertiesFormat.AddressSpace.AddressPrefixes)
        }
    }

    {
        resp, err := pubipClient.ListAll(context.Background())
        if err != nil {
            fmt.Printf("Having some problems here with disks %s\n", err.Error())
            return
        }
        pubiplist := resp.Values()
        for _, v := range pubiplist {
            fmt.Printf("%s\n", *v.Name)
            fmt.Printf("%s\n", *v.Etag)
            fmt.Printf("%s\n", *v.ID)
            fmt.Printf("%s\n", *v.Type)
            fmt.Printf("%s\n", *v.Location)
            fmt.Printf("%s\n", *v.PublicIPAddressPropertiesFormat.IPAddress)
            fmt.Printf("%s\n", *v.PublicIPAddressPropertiesFormat.ProvisioningState)
        }
    }

    {
        resp, err := inetClient.ListAll(context.Background())
        if err != nil {
            fmt.Printf("Having some problems here with disks %s\n", err.Error())
            return
        }
        inetlist := resp.Values()
        for _, v := range inetlist {
            fmt.Printf("%s\n", *v.Name)
            fmt.Printf("%s\n", *v.Etag)
            fmt.Printf("%s\n", *v.ID)
            fmt.Printf("%s\n", *v.Type)
            fmt.Printf("%s\n", *v.Location)
            fmt.Printf("%s\n", *v.InterfacePropertiesFormat.MacAddress)
            fmt.Printf("%s\n", *v.InterfacePropertiesFormat.ProvisioningState)
            for _, y := range *v.InterfacePropertiesFormat.IPConfigurations {
                fmt.Printf("%s\n", *y.Name)
                fmt.Printf("%s\n", *y.InterfaceIPConfigurationPropertiesFormat.PrivateIPAddress)
            }
            //fmt.Printf("%s\n", *v.InterfacePropertiesFormat.IPConfigurations)

        }
    }

    {
        resp, err := sgClient.ListAll(context.Background())
        if err != nil {
            fmt.Printf("Having some problems here with disks %s\n", err.Error())
            return
        }
        sglist := resp.Values()
        for _, v := range sglist {
            fmt.Printf("%s\n", *v.Name)
            fmt.Printf("%s\n", *v.Etag)
            fmt.Printf("%s\n", *v.ID)
            fmt.Printf("%s\n", *v.Type)
            fmt.Printf("%s\n", *v.Location)
            for _, y := range *v.SecurityGroupPropertiesFormat.SecurityRules {
                fmt.Printf("%s\n", *y.ID)
                fmt.Printf("%s ", *y.Name)
                //fmt.Printf("%s ", *y.SecurityRulePropertiesFormat.Protocol)
                fmt.Printf("%s ", *y.SecurityRulePropertiesFormat.SourcePortRange)
                fmt.Printf("%s ", *y.SecurityRulePropertiesFormat.DestinationPortRange)
                fmt.Printf("%s ", *y.SecurityRulePropertiesFormat.SourceAddressPrefix)
                fmt.Printf("%s\n", *y.SecurityRulePropertiesFormat.DestinationAddressPrefix)
                //fmt.Printf("%s\n", *y.SecurityRulePropertiesFormat.Access)
                //fmt.Printf("%s\n", *y.SecurityRulePropertiesFormat.Direction)
                
            }
            for _, y := range *v.SecurityGroupPropertiesFormat.DefaultSecurityRules {
                fmt.Printf("%s\n", *y.ID)
                fmt.Printf("%s ", *y.Name)
                //fmt.Printf("%s ", *y.SecurityRulePropertiesFormat.Protocol)
                fmt.Printf("%s ", *y.SecurityRulePropertiesFormat.SourcePortRange)
                fmt.Printf("%s ", *y.SecurityRulePropertiesFormat.DestinationPortRange)
                fmt.Printf("%s ", *y.SecurityRulePropertiesFormat.SourceAddressPrefix)
                fmt.Printf("%s\n", *y.SecurityRulePropertiesFormat.DestinationAddressPrefix)
                //fmt.Printf("%s\n", *y.SecurityRulePropertiesFormat.Access)
                //fmt.Printf("%s\n", *y.SecurityRulePropertiesFormat.Direction)

            }
        }
    }

}


