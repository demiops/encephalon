package collectors

type Asset struct {
    Name        string  `binding:"required"`
    Cloud       string  `binding:"required"`
    CloudID     string  `binding:"required"`
    AssetType   string  `json:"asset_type" binding:"required"`
    Status      string  `binding:"required"`
}

type Cdn struct {
    Name        string
    Cloud       string
    CloudID     string
    AssetType   string  `json:"asset_type" binding:"required"`
    Status      string
    Origin      string
    Endpoint    string
    TTL         string
}

type VirtualMachine struct {
    Name        string  `binding:"required"`
    Cloud       string  `binding:"required"`
    CloudID     string  `binding:"required"`
    AssetType   string  `json:"asset_type" binding:"required"`
    Status      string  `binding:"required"`
    Memory      string
    Vcpu        string
    Disks       string  `json: "omitempty"`
    Image       string  `json: "omitempty"`
    Region      string  `json: "omitempty"`
    Az          string  `json: "omitempty"`
    Arch        string  `json: "omitempty"`
    VMSize      string  `json: "omitempty"`
    Key         string  `json: "omitempty"`
    PrivateDNS  string  `json: "omitempty"`
    PrivateIP   string  `json: "omitempty"`
    PublicDNS   string  `json: "omitempty"`
    PublicIP    string  `json: "omitempty"`
    Firewall    string  `json: "omitempty"`
    Network     string  `json: "omitempty"`
    Subnet      string  `json: "omitempty"`
    Virt        string  
    Hyper       string
}

type Firewall struct {
    Name        string
    Cloud       string
    CloudID     string
    AssetType   string  `json:"asset_type" binding:"required"`
    Status      string
    InRules     string
    OutRules    string
}

type Domain struct {
    Name        string
    Cloud       string
    CloudID     string
    AssetType   string  `json:"asset_type" binding:"required"`
    Status      string
    TTL         string
    ZoneFile    string
}

type Disk struct {
    Name        string  `binding:"required"`
    Cloud       string  `binding:"required"`
    CloudID     string  `binding:"required"`
    AssetType   string  `json:"asset_type" binding:"required"`
    Status      string  `binding:"required"`
    Size        string
    DiskType    string
    FilesystemType      string
    FilesystemLabel     string
    Encrypted   string
    SnapshotId  string
    Az          string
}

type Image struct {
    Name        string  `binding:"required"`
    Cloud       string  `binding:"required"`
    CloudID     string  `binding:"required"`
    AssetType   string  `json:"asset_type" binding:"required"`
    Status      string  `binding:"required"`
    Size        string
    ImageType   string
    Distro      string
}

type Key struct {
    Name        string  `binding:"required"`
    Cloud       string  `binding:"required"`
    CloudID     string  `binding:"required"`
    AssetType   string  `json:"asset_type" binding:"required"`
    Status      string  `binding:"required"`
    Fingerprint string
    PublicKey   string
}

type Region struct {
    Name        string  `binding:"required"`
    Cloud       string  `binding:"required"`
    CloudID     string  `binding:"required"`
    AssetType   string  `json:"asset_type" binding:"required"`
    Status      string  `binding:"required"`
    Location    string
    Available   string
}

type Az struct {
    Name        string  `binding:"required"`
    Cloud       string  `binding:"required"`
    CloudID     string  `binding:"required"`
    AssetType   string  `json:"asset_type" binding:"required"`
    Status      string  `binding:"required"`
    Region      string
}

type LoadBalancer struct {
    Name        string  `binding:"required"`
    Cloud       string  `binding:"required"`
    CloudID     string  `binding:"required"`
    AssetType   string  `json:"asset_type" binding:"required"`
    Status      string  `binding:"required"`
    IpAddress   string
    Algorithm   string
}

type FirewallRule struct {
    From        string
    To          string
    Port        string
}

type VMSize struct {
    Name        string  `binding:"required"`
    Cloud       string  `binding:"required"`
    CloudID     string  `binding:"required"`
    AssetType   string  `json:"asset_type" binding:"required"`
    Status      string  `binding:"required"`
    Vcpu        string
    Memory      string
    Disk        string
    Available   string
}

type Snapshot struct {
    Name        string  `binding:"required"`
    Cloud       string  `binding:"required"`
    CloudID     string  `binding:"required"`
    AssetType   string  `json:"asset_type" binding:"required"`
    Status      string  `binding:"required"`
    Disk        string
}


type IpAddress struct {
    Name        string  `binding:"required"`
    Cloud       string  `binding:"required"`
    CloudID     string  `binding:"required"`
    AssetType   string  `json:"asset_type" binding:"required"`
    Status      string  `binding:"required"`
    Address     string  `binding:"required"`
    Public      bool
}

type Network struct {
    Name        string  `binding:"required"`
    Cloud       string  `binding:"required"`
    CloudID     string  `binding:"required"`
    AssetType   string  `json:"asset_type" binding:"required"`
    Status      string  `binding:"required"`
    Range       string  
}

type Subnet struct {
    Name        string  `binding:"required"`
    Cloud       string  `binding:"required"`
    CloudID     string  `binding:"required"`
    AssetType   string  `json:"asset_type" binding:"required"`
    Status      string  `binding:"required"`
    Az          string
    AvailableIP string
    Range       string
    NetworkID   string
    UsePublicIP string
    DefaultAz   string
}

type InternetGateway struct {
    Name        string  `binding:"required"`
    Cloud       string  `binding:"required"`
    CloudID     string  `binding:"required"`
    AssetType   string  `json:"asset_type" binding:"required"`
    Status      string  `binding:"required"`
}

type Router struct {
    Name        string  `binding:"required"`
    Cloud       string  `binding:"required"`
    CloudID     string  `binding:"required"`
    AssetType   string  `json:"asset_type" binding:"required"`
    Status      string  `binding:"required"`
    Description string
    Network     string
    Region      string
}

type Routes struct {
    Name        string  `binding:"required"`
    Cloud       string  `binding:"required"`
    CloudID     string  `binding:"required"`
    AssetType   string  `json:"asset_type" binding:"required"`
    Status      string  `binding:"required"`
    Description string
    DstRange    string
    Network     string
    NextHopGW   string
    NextHopVM   string
    NextHopIP   string
    NextHopNet  string
    NextHopPeer string
    NextHopVpn  string
    Priority    string
}

