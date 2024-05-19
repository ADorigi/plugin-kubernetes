package preferences

import (
	"github.com/kaytu-io/kaytu/pkg/plugin/proto/src/golang"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

var DefaultEC2Preferences = []*golang.PreferenceItem{
	{Service: "EC2Instance", Key: "Tenancy", Pinned: true, PossibleValues: []string{"", "Host", "Shared", "Dedicated"}},
	{Service: "EC2Instance", Key: "EBSOptimized", PossibleValues: []string{"", "Yes", "No"}},
	{Service: "EC2Instance", Key: "LicenseModel", PossibleValues: []string{"", "Bring your own license", "No License required"}},
	{Service: "EC2Instance", Key: "Region", Pinned: true},
	{Service: "EC2Instance", Key: "CurrentGeneration", PossibleValues: []string{"", "Yes", "No"}},
	{Service: "EC2Instance", Key: "PhysicalProcessor"},
	{Service: "EC2Instance", Key: "ClockSpeed"},
	{Service: "EC2Instance", Key: "OperatingSystem", PreventPinning: true, PossibleValues: []string{"", "Windows", "Linux/UNIX"}},
	{Service: "EC2Instance", Key: "ProcessorArchitecture", Pinned: true, PossibleValues: []string{"", "x86_64", "arm64", "arm64_mac"}},
	{Service: "EC2Instance", Key: "InstanceFamily", PossibleValues: []string{"", "General purpose", "Compute optimized", "Memory optimized", "Storage optimized", "FPGA Instances", "GPU instance", "Machine Learning ASIC Instances", "Media Accelerator Instances"}},
	{Service: "EC2Instance", Key: "UsageOperation", Pinned: true, PossibleValues: []string{"", "Linux/UNIX", "Red Hat BYOL Linux", "Red Hat Enterprise Linux", "Red Hat Enterprise Linux with HA", "Red Hat Enterprise Linux with SQL Server Standard and HA", "Red Hat Enterprise Linux with SQL Server Enterprise and HA", "Red Hat Enterprise Linux with SQL Server Standard", "Red Hat Enterprise Linux with SQL Server Web", "Red Hat Enterprise Linux with SQL Server Enterprise", "SQL Server Enterprise", "SQL Server Standard", "SQL Server Web", "SUSE Linux", "Ubuntu Pro", "Windows", "Windows BYOL", "Windows with SQL Server Enterprise", "Windows with SQL Server Standard", "Windows with SQL Server Web"}},
	{Service: "EC2Instance", Key: "ENASupported"},
	{Service: "EC2Instance", Key: "SupportedRootDeviceTypes", Value: wrapperspb.String("EBSOnly"), PreventPinning: true, PossibleValues: []string{"EBSOnly"}},
	{Service: "EC2Instance", Key: "vCPU", IsNumber: true},
	{Service: "EC2Instance", Key: "MemoryGB", Alias: "Memory", IsNumber: true, Pinned: true, Unit: "GiB"},
	{Service: "EC2Instance", Key: "CPUBreathingRoom", IsNumber: true, Value: wrapperspb.String("10"), PreventPinning: true, Unit: "%"},
	{Service: "EC2Instance", Key: "MemoryBreathingRoom", IsNumber: true, Value: wrapperspb.String("5"), PreventPinning: true, Unit: "%"},
	{Service: "EC2Instance", Key: "NetworkBreathingRoom", IsNumber: true, Value: wrapperspb.String("10"), PreventPinning: true, Unit: "%"},
	{Service: "EC2Instance", Key: "ObservabilityTimePeriod", Value: wrapperspb.String("7"), PreventPinning: true, Unit: "days", PossibleValues: []string{"7"}},
	{Service: "EC2Instance", Key: "RuntimeInterval", Value: wrapperspb.String("730"), PreventPinning: true, Unit: "hours", PossibleValues: []string{"730"}},
	{Service: "EC2Instance", Key: "ExcludeBurstableInstances", Value: wrapperspb.String("if current resource is burstable"), PreventPinning: true, PossibleValues: []string{"No", "Yes", "if current resource is burstable"}},
	{Service: "EBSVolume", Key: "IOPS", IsNumber: true},
	{Service: "EBSVolume", Key: "Throughput", IsNumber: true, Unit: "Mbps"},
	{Service: "EBSVolume", Key: "Size", IsNumber: true, Pinned: true, Unit: "GB"},
	{Service: "EBSVolume", Key: "VolumeFamily", PossibleValues: []string{"", "General Purpose", "Solid State Drive", "IO Optimized", "Hard Disk Drive"}},
	{Service: "EBSVolume", Key: "VolumeType", PossibleValues: []string{"", "standard", "io1", "io2", "gp2", "gp3", "sc1", "st1"}},
	{Service: "EBSVolume", Key: "IOPSBreathingRoom", IsNumber: true, Value: wrapperspb.String("10"), PreventPinning: true, Unit: "%"},
	{Service: "EBSVolume", Key: "ThroughputBreathingRoom", IsNumber: true, Value: wrapperspb.String("10"), PreventPinning: true, Unit: "%"},
}
