package main

import (
	"html/template"
	"time"
)

var temp *template.Template

var settings struct {
	importPath string
	exportPath string
}

type LvsFromJSON struct {
	SizeG string `json:"size_g"`
	Vg    string `json:"vg"`
}

type PvsFromJSON struct {
	FreeG string `json:"free_g"`
	SizeG string `json:"size_g"`
	Vg    string `json:"v_g"`
}
type VgsFromJSON struct {
	FreeG  string `json:"free_g"`
	NumLvs string `json:"num_lvs"`
	NumPvs string `json:"num_pvs"`
	SizeG  string `json:"size_g"`
}

type AnsibleFactFull struct {
	AnsibleFacts struct {
		AnsibleAllIpv4Addresses []string `json:"ansible_all_ipv4_addresses"`
		AnsibleAllIpv6Addresses []string `json:"ansible_all_ipv6_addresses"`
		AnsibleApparmor         struct {
			Status string `json:"status"`
		} `json:"ansible_apparmor"`
		AnsibleArchitecture    string `json:"ansible_architecture"`
		AnsibleBiosDate        string `json:"ansible_bios_date"`
		AnsibleBiosVendor      string `json:"ansible_bios_vendor"`
		AnsibleBiosVersion     string `json:"ansible_bios_version"`
		AnsibleBoardAssetTag   string `json:"ansible_board_asset_tag"`
		AnsibleBoardName       string `json:"ansible_board_name"`
		AnsibleBoardSerial     string `json:"ansible_board_serial"`
		AnsibleBoardVendor     string `json:"ansible_board_vendor"`
		AnsibleBoardVersion    string `json:"ansible_board_version"`
		AnsibleChassisAssetTag string `json:"ansible_chassis_asset_tag"`
		AnsibleChassisSerial   string `json:"ansible_chassis_serial"`
		AnsibleChassisVendor   string `json:"ansible_chassis_vendor"`
		AnsibleChassisVersion  string `json:"ansible_chassis_version"`
		AnsibleCmdline         struct {
			BOOTIMAGE string `json:"BOOT_IMAGE"`
			Elevator  string `json:"elevator"`
			Quiet     bool   `json:"quiet"`
			Ro        bool   `json:"ro"`
			Root      string `json:"root"`
		} `json:"ansible_cmdline"`
		AnsibleDateTime struct {
			Date              string    `json:"date"`
			Day               string    `json:"day"`
			Epoch             string    `json:"epoch"`
			EpochInt          string    `json:"epoch_int"`
			Hour              string    `json:"hour"`
			Iso8601           time.Time `json:"iso8601"`
			Iso8601Basic      string    `json:"iso8601_basic"`
			Iso8601BasicShort string    `json:"iso8601_basic_short"`
			Iso8601Micro      time.Time `json:"iso8601_micro"`
			Minute            string    `json:"minute"`
			Month             string    `json:"month"`
			Second            string    `json:"second"`
			Time              string    `json:"time"`
			Tz                string    `json:"tz"`
			TzDst             string    `json:"tz_dst"`
			TzOffset          string    `json:"tz_offset"`
			Weekday           string    `json:"weekday"`
			WeekdayNumber     string    `json:"weekday_number"`
			Weeknumber        string    `json:"weeknumber"`
			Year              string    `json:"year"`
		} `json:"ansible_date_time"`
		AnsibleDefaultIpv4 struct {
			Address    string `json:"address"`
			Alias      string `json:"alias"`
			Broadcast  string `json:"broadcast"`
			Gateway    string `json:"gateway"`
			Interface  string `json:"interface"`
			Macaddress string `json:"macaddress"`
			Mtu        int    `json:"mtu"`
			Netmask    string `json:"netmask"`
			Network    string `json:"network"`
			Prefix     string `json:"prefix"`
			Type       string `json:"type"`
		} `json:"ansible_default_ipv4"`
		AnsibleDefaultIpv6 struct {
			Address    string `json:"address"`
			Gateway    string `json:"gateway"`
			Interface  string `json:"interface"`
			Macaddress string `json:"macaddress"`
			Mtu        int    `json:"mtu"`
			Prefix     string `json:"prefix"`
			Scope      string `json:"scope"`
			Type       string `json:"type"`
		} `json:"ansible_default_ipv6"`
		AnsibleDeviceLinks struct {
			Ids struct {
				Dm0  []string `json:"dm-0"`
				Dm1  []string `json:"dm-1"`
				Dm10 []string `json:"dm-10"`
				Dm11 []string `json:"dm-11"`
				Dm12 []string `json:"dm-12"`
				Dm13 []string `json:"dm-13"`
				Dm14 []string `json:"dm-14"`
				Dm2  []string `json:"dm-2"`
				Dm3  []string `json:"dm-3"`
				Dm4  []string `json:"dm-4"`
				Dm5  []string `json:"dm-5"`
				Dm6  []string `json:"dm-6"`
				Dm7  []string `json:"dm-7"`
				Dm8  []string `json:"dm-8"`
				Dm9  []string `json:"dm-9"`
				Sda  []string `json:"sda"`
				Sda1 []string `json:"sda1"`
				Sda2 []string `json:"sda2"`
				Sda3 []string `json:"sda3"`
				Sr0  []string `json:"sr0"`
			} `json:"ids"`
			Labels struct {
			} `json:"labels"`
			Masters struct {
				Sda3 []string `json:"sda3"`
			} `json:"masters"`
			Uuids struct {
				Dm0  []string `json:"dm-0"`
				Dm1  []string `json:"dm-1"`
				Dm10 []string `json:"dm-10"`
				Dm11 []string `json:"dm-11"`
				Dm12 []string `json:"dm-12"`
				Dm13 []string `json:"dm-13"`
				Dm14 []string `json:"dm-14"`
				Dm2  []string `json:"dm-2"`
				Dm3  []string `json:"dm-3"`
				Dm4  []string `json:"dm-4"`
				Dm5  []string `json:"dm-5"`
				Dm6  []string `json:"dm-6"`
				Dm7  []string `json:"dm-7"`
				Dm8  []string `json:"dm-8"`
				Dm9  []string `json:"dm-9"`
				Sda2 []string `json:"sda2"`
			} `json:"uuids"`
		} `json:"ansible_device_links"`
		AnsibleDevices struct {
			Dm0 struct {
				Holders []interface{} `json:"holders"`
				Host    string        `json:"host"`
				Links   struct {
					Ids     []string      `json:"ids"`
					Labels  []interface{} `json:"labels"`
					Masters []interface{} `json:"masters"`
					Uuids   []string      `json:"uuids"`
				} `json:"links"`
				Model      interface{} `json:"model"`
				Partitions struct {
				} `json:"partitions"`
				Removable       string      `json:"removable"`
				Rotational      string      `json:"rotational"`
				SasAddress      interface{} `json:"sas_address"`
				SasDeviceHandle interface{} `json:"sas_device_handle"`
				SchedulerMode   string      `json:"scheduler_mode"`
				Sectors         string      `json:"sectors"`
				Sectorsize      string      `json:"sectorsize"`
				Size            string      `json:"size"`
				SupportDiscard  string      `json:"support_discard"`
				Vendor          interface{} `json:"vendor"`
				Virtual         int         `json:"virtual"`
			} `json:"dm-0"`
			Dm1 struct {
				Holders []interface{} `json:"holders"`
				Host    string        `json:"host"`
				Links   struct {
					Ids     []string      `json:"ids"`
					Labels  []interface{} `json:"labels"`
					Masters []interface{} `json:"masters"`
					Uuids   []string      `json:"uuids"`
				} `json:"links"`
				Model      interface{} `json:"model"`
				Partitions struct {
				} `json:"partitions"`
				Removable       string      `json:"removable"`
				Rotational      string      `json:"rotational"`
				SasAddress      interface{} `json:"sas_address"`
				SasDeviceHandle interface{} `json:"sas_device_handle"`
				SchedulerMode   string      `json:"scheduler_mode"`
				Sectors         string      `json:"sectors"`
				Sectorsize      string      `json:"sectorsize"`
				Size            string      `json:"size"`
				SupportDiscard  string      `json:"support_discard"`
				Vendor          interface{} `json:"vendor"`
				Virtual         int         `json:"virtual"`
			} `json:"dm-1"`
			Dm10 struct {
				Holders []interface{} `json:"holders"`
				Host    string        `json:"host"`
				Links   struct {
					Ids     []string      `json:"ids"`
					Labels  []interface{} `json:"labels"`
					Masters []interface{} `json:"masters"`
					Uuids   []string      `json:"uuids"`
				} `json:"links"`
				Model      interface{} `json:"model"`
				Partitions struct {
				} `json:"partitions"`
				Removable       string      `json:"removable"`
				Rotational      string      `json:"rotational"`
				SasAddress      interface{} `json:"sas_address"`
				SasDeviceHandle interface{} `json:"sas_device_handle"`
				SchedulerMode   string      `json:"scheduler_mode"`
				Sectors         string      `json:"sectors"`
				Sectorsize      string      `json:"sectorsize"`
				Size            string      `json:"size"`
				SupportDiscard  string      `json:"support_discard"`
				Vendor          interface{} `json:"vendor"`
				Virtual         int         `json:"virtual"`
			} `json:"dm-10"`
			Dm11 struct {
				Holders []interface{} `json:"holders"`
				Host    string        `json:"host"`
				Links   struct {
					Ids     []string      `json:"ids"`
					Labels  []interface{} `json:"labels"`
					Masters []interface{} `json:"masters"`
					Uuids   []string      `json:"uuids"`
				} `json:"links"`
				Model      interface{} `json:"model"`
				Partitions struct {
				} `json:"partitions"`
				Removable       string      `json:"removable"`
				Rotational      string      `json:"rotational"`
				SasAddress      interface{} `json:"sas_address"`
				SasDeviceHandle interface{} `json:"sas_device_handle"`
				SchedulerMode   string      `json:"scheduler_mode"`
				Sectors         string      `json:"sectors"`
				Sectorsize      string      `json:"sectorsize"`
				Size            string      `json:"size"`
				SupportDiscard  string      `json:"support_discard"`
				Vendor          interface{} `json:"vendor"`
				Virtual         int         `json:"virtual"`
			} `json:"dm-11"`
			Dm12 struct {
				Holders []interface{} `json:"holders"`
				Host    string        `json:"host"`
				Links   struct {
					Ids     []string      `json:"ids"`
					Labels  []interface{} `json:"labels"`
					Masters []interface{} `json:"masters"`
					Uuids   []string      `json:"uuids"`
				} `json:"links"`
				Model      interface{} `json:"model"`
				Partitions struct {
				} `json:"partitions"`
				Removable       string      `json:"removable"`
				Rotational      string      `json:"rotational"`
				SasAddress      interface{} `json:"sas_address"`
				SasDeviceHandle interface{} `json:"sas_device_handle"`
				SchedulerMode   string      `json:"scheduler_mode"`
				Sectors         string      `json:"sectors"`
				Sectorsize      string      `json:"sectorsize"`
				Size            string      `json:"size"`
				SupportDiscard  string      `json:"support_discard"`
				Vendor          interface{} `json:"vendor"`
				Virtual         int         `json:"virtual"`
			} `json:"dm-12"`
			Dm13 struct {
				Holders []interface{} `json:"holders"`
				Host    string        `json:"host"`
				Links   struct {
					Ids     []string      `json:"ids"`
					Labels  []interface{} `json:"labels"`
					Masters []interface{} `json:"masters"`
					Uuids   []string      `json:"uuids"`
				} `json:"links"`
				Model      interface{} `json:"model"`
				Partitions struct {
				} `json:"partitions"`
				Removable       string      `json:"removable"`
				Rotational      string      `json:"rotational"`
				SasAddress      interface{} `json:"sas_address"`
				SasDeviceHandle interface{} `json:"sas_device_handle"`
				SchedulerMode   string      `json:"scheduler_mode"`
				Sectors         string      `json:"sectors"`
				Sectorsize      string      `json:"sectorsize"`
				Size            string      `json:"size"`
				SupportDiscard  string      `json:"support_discard"`
				Vendor          interface{} `json:"vendor"`
				Virtual         int         `json:"virtual"`
			} `json:"dm-13"`
			Dm14 struct {
				Holders []interface{} `json:"holders"`
				Host    string        `json:"host"`
				Links   struct {
					Ids     []string      `json:"ids"`
					Labels  []interface{} `json:"labels"`
					Masters []interface{} `json:"masters"`
					Uuids   []string      `json:"uuids"`
				} `json:"links"`
				Model      interface{} `json:"model"`
				Partitions struct {
				} `json:"partitions"`
				Removable       string      `json:"removable"`
				Rotational      string      `json:"rotational"`
				SasAddress      interface{} `json:"sas_address"`
				SasDeviceHandle interface{} `json:"sas_device_handle"`
				SchedulerMode   string      `json:"scheduler_mode"`
				Sectors         string      `json:"sectors"`
				Sectorsize      string      `json:"sectorsize"`
				Size            string      `json:"size"`
				SupportDiscard  string      `json:"support_discard"`
				Vendor          interface{} `json:"vendor"`
				Virtual         int         `json:"virtual"`
			} `json:"dm-14"`
			Dm2 struct {
				Holders []interface{} `json:"holders"`
				Host    string        `json:"host"`
				Links   struct {
					Ids     []string      `json:"ids"`
					Labels  []interface{} `json:"labels"`
					Masters []interface{} `json:"masters"`
					Uuids   []string      `json:"uuids"`
				} `json:"links"`
				Model      interface{} `json:"model"`
				Partitions struct {
				} `json:"partitions"`
				Removable       string      `json:"removable"`
				Rotational      string      `json:"rotational"`
				SasAddress      interface{} `json:"sas_address"`
				SasDeviceHandle interface{} `json:"sas_device_handle"`
				SchedulerMode   string      `json:"scheduler_mode"`
				Sectors         string      `json:"sectors"`
				Sectorsize      string      `json:"sectorsize"`
				Size            string      `json:"size"`
				SupportDiscard  string      `json:"support_discard"`
				Vendor          interface{} `json:"vendor"`
				Virtual         int         `json:"virtual"`
			} `json:"dm-2"`
			Dm3 struct {
				Holders []interface{} `json:"holders"`
				Host    string        `json:"host"`
				Links   struct {
					Ids     []string      `json:"ids"`
					Labels  []interface{} `json:"labels"`
					Masters []interface{} `json:"masters"`
					Uuids   []string      `json:"uuids"`
				} `json:"links"`
				Model      interface{} `json:"model"`
				Partitions struct {
				} `json:"partitions"`
				Removable       string      `json:"removable"`
				Rotational      string      `json:"rotational"`
				SasAddress      interface{} `json:"sas_address"`
				SasDeviceHandle interface{} `json:"sas_device_handle"`
				SchedulerMode   string      `json:"scheduler_mode"`
				Sectors         string      `json:"sectors"`
				Sectorsize      string      `json:"sectorsize"`
				Size            string      `json:"size"`
				SupportDiscard  string      `json:"support_discard"`
				Vendor          interface{} `json:"vendor"`
				Virtual         int         `json:"virtual"`
			} `json:"dm-3"`
			Dm4 struct {
				Holders []interface{} `json:"holders"`
				Host    string        `json:"host"`
				Links   struct {
					Ids     []string      `json:"ids"`
					Labels  []interface{} `json:"labels"`
					Masters []interface{} `json:"masters"`
					Uuids   []string      `json:"uuids"`
				} `json:"links"`
				Model      interface{} `json:"model"`
				Partitions struct {
				} `json:"partitions"`
				Removable       string      `json:"removable"`
				Rotational      string      `json:"rotational"`
				SasAddress      interface{} `json:"sas_address"`
				SasDeviceHandle interface{} `json:"sas_device_handle"`
				SchedulerMode   string      `json:"scheduler_mode"`
				Sectors         string      `json:"sectors"`
				Sectorsize      string      `json:"sectorsize"`
				Size            string      `json:"size"`
				SupportDiscard  string      `json:"support_discard"`
				Vendor          interface{} `json:"vendor"`
				Virtual         int         `json:"virtual"`
			} `json:"dm-4"`
			Dm5 struct {
				Holders []interface{} `json:"holders"`
				Host    string        `json:"host"`
				Links   struct {
					Ids     []string      `json:"ids"`
					Labels  []interface{} `json:"labels"`
					Masters []interface{} `json:"masters"`
					Uuids   []string      `json:"uuids"`
				} `json:"links"`
				Model      interface{} `json:"model"`
				Partitions struct {
				} `json:"partitions"`
				Removable       string      `json:"removable"`
				Rotational      string      `json:"rotational"`
				SasAddress      interface{} `json:"sas_address"`
				SasDeviceHandle interface{} `json:"sas_device_handle"`
				SchedulerMode   string      `json:"scheduler_mode"`
				Sectors         string      `json:"sectors"`
				Sectorsize      string      `json:"sectorsize"`
				Size            string      `json:"size"`
				SupportDiscard  string      `json:"support_discard"`
				Vendor          interface{} `json:"vendor"`
				Virtual         int         `json:"virtual"`
			} `json:"dm-5"`
			Dm6 struct {
				Holders []interface{} `json:"holders"`
				Host    string        `json:"host"`
				Links   struct {
					Ids     []string      `json:"ids"`
					Labels  []interface{} `json:"labels"`
					Masters []interface{} `json:"masters"`
					Uuids   []string      `json:"uuids"`
				} `json:"links"`
				Model      interface{} `json:"model"`
				Partitions struct {
				} `json:"partitions"`
				Removable       string      `json:"removable"`
				Rotational      string      `json:"rotational"`
				SasAddress      interface{} `json:"sas_address"`
				SasDeviceHandle interface{} `json:"sas_device_handle"`
				SchedulerMode   string      `json:"scheduler_mode"`
				Sectors         string      `json:"sectors"`
				Sectorsize      string      `json:"sectorsize"`
				Size            string      `json:"size"`
				SupportDiscard  string      `json:"support_discard"`
				Vendor          interface{} `json:"vendor"`
				Virtual         int         `json:"virtual"`
			} `json:"dm-6"`
			Dm7 struct {
				Holders []interface{} `json:"holders"`
				Host    string        `json:"host"`
				Links   struct {
					Ids     []string      `json:"ids"`
					Labels  []interface{} `json:"labels"`
					Masters []interface{} `json:"masters"`
					Uuids   []string      `json:"uuids"`
				} `json:"links"`
				Model      interface{} `json:"model"`
				Partitions struct {
				} `json:"partitions"`
				Removable       string      `json:"removable"`
				Rotational      string      `json:"rotational"`
				SasAddress      interface{} `json:"sas_address"`
				SasDeviceHandle interface{} `json:"sas_device_handle"`
				SchedulerMode   string      `json:"scheduler_mode"`
				Sectors         string      `json:"sectors"`
				Sectorsize      string      `json:"sectorsize"`
				Size            string      `json:"size"`
				SupportDiscard  string      `json:"support_discard"`
				Vendor          interface{} `json:"vendor"`
				Virtual         int         `json:"virtual"`
			} `json:"dm-7"`
			Dm8 struct {
				Holders []interface{} `json:"holders"`
				Host    string        `json:"host"`
				Links   struct {
					Ids     []string      `json:"ids"`
					Labels  []interface{} `json:"labels"`
					Masters []interface{} `json:"masters"`
					Uuids   []string      `json:"uuids"`
				} `json:"links"`
				Model      interface{} `json:"model"`
				Partitions struct {
				} `json:"partitions"`
				Removable       string      `json:"removable"`
				Rotational      string      `json:"rotational"`
				SasAddress      interface{} `json:"sas_address"`
				SasDeviceHandle interface{} `json:"sas_device_handle"`
				SchedulerMode   string      `json:"scheduler_mode"`
				Sectors         string      `json:"sectors"`
				Sectorsize      string      `json:"sectorsize"`
				Size            string      `json:"size"`
				SupportDiscard  string      `json:"support_discard"`
				Vendor          interface{} `json:"vendor"`
				Virtual         int         `json:"virtual"`
			} `json:"dm-8"`
			Dm9 struct {
				Holders []interface{} `json:"holders"`
				Host    string        `json:"host"`
				Links   struct {
					Ids     []string      `json:"ids"`
					Labels  []interface{} `json:"labels"`
					Masters []interface{} `json:"masters"`
					Uuids   []string      `json:"uuids"`
				} `json:"links"`
				Model      interface{} `json:"model"`
				Partitions struct {
				} `json:"partitions"`
				Removable       string      `json:"removable"`
				Rotational      string      `json:"rotational"`
				SasAddress      interface{} `json:"sas_address"`
				SasDeviceHandle interface{} `json:"sas_device_handle"`
				SchedulerMode   string      `json:"scheduler_mode"`
				Sectors         string      `json:"sectors"`
				Sectorsize      string      `json:"sectorsize"`
				Size            string      `json:"size"`
				SupportDiscard  string      `json:"support_discard"`
				Vendor          interface{} `json:"vendor"`
				Virtual         int         `json:"virtual"`
			} `json:"dm-9"`
			Loop0 struct {
				Holders []interface{} `json:"holders"`
				Host    string        `json:"host"`
				Links   struct {
					Ids     []interface{} `json:"ids"`
					Labels  []interface{} `json:"labels"`
					Masters []interface{} `json:"masters"`
					Uuids   []interface{} `json:"uuids"`
				} `json:"links"`
				Model      interface{} `json:"model"`
				Partitions struct {
				} `json:"partitions"`
				Removable       string      `json:"removable"`
				Rotational      string      `json:"rotational"`
				SasAddress      interface{} `json:"sas_address"`
				SasDeviceHandle interface{} `json:"sas_device_handle"`
				SchedulerMode   string      `json:"scheduler_mode"`
				Sectors         string      `json:"sectors"`
				Sectorsize      string      `json:"sectorsize"`
				Size            string      `json:"size"`
				SupportDiscard  string      `json:"support_discard"`
				Vendor          interface{} `json:"vendor"`
				Virtual         int         `json:"virtual"`
			} `json:"loop0"`
			Loop1 struct {
				Holders []interface{} `json:"holders"`
				Host    string        `json:"host"`
				Links   struct {
					Ids     []interface{} `json:"ids"`
					Labels  []interface{} `json:"labels"`
					Masters []interface{} `json:"masters"`
					Uuids   []interface{} `json:"uuids"`
				} `json:"links"`
				Model      interface{} `json:"model"`
				Partitions struct {
				} `json:"partitions"`
				Removable       string      `json:"removable"`
				Rotational      string      `json:"rotational"`
				SasAddress      interface{} `json:"sas_address"`
				SasDeviceHandle interface{} `json:"sas_device_handle"`
				SchedulerMode   string      `json:"scheduler_mode"`
				Sectors         string      `json:"sectors"`
				Sectorsize      string      `json:"sectorsize"`
				Size            string      `json:"size"`
				SupportDiscard  string      `json:"support_discard"`
				Vendor          interface{} `json:"vendor"`
				Virtual         int         `json:"virtual"`
			} `json:"loop1"`
			Loop2 struct {
				Holders []interface{} `json:"holders"`
				Host    string        `json:"host"`
				Links   struct {
					Ids     []interface{} `json:"ids"`
					Labels  []interface{} `json:"labels"`
					Masters []interface{} `json:"masters"`
					Uuids   []interface{} `json:"uuids"`
				} `json:"links"`
				Model      interface{} `json:"model"`
				Partitions struct {
				} `json:"partitions"`
				Removable       string      `json:"removable"`
				Rotational      string      `json:"rotational"`
				SasAddress      interface{} `json:"sas_address"`
				SasDeviceHandle interface{} `json:"sas_device_handle"`
				SchedulerMode   string      `json:"scheduler_mode"`
				Sectors         string      `json:"sectors"`
				Sectorsize      string      `json:"sectorsize"`
				Size            string      `json:"size"`
				SupportDiscard  string      `json:"support_discard"`
				Vendor          interface{} `json:"vendor"`
				Virtual         int         `json:"virtual"`
			} `json:"loop2"`
			Loop3 struct {
				Holders []interface{} `json:"holders"`
				Host    string        `json:"host"`
				Links   struct {
					Ids     []interface{} `json:"ids"`
					Labels  []interface{} `json:"labels"`
					Masters []interface{} `json:"masters"`
					Uuids   []interface{} `json:"uuids"`
				} `json:"links"`
				Model      interface{} `json:"model"`
				Partitions struct {
				} `json:"partitions"`
				Removable       string      `json:"removable"`
				Rotational      string      `json:"rotational"`
				SasAddress      interface{} `json:"sas_address"`
				SasDeviceHandle interface{} `json:"sas_device_handle"`
				SchedulerMode   string      `json:"scheduler_mode"`
				Sectors         string      `json:"sectors"`
				Sectorsize      string      `json:"sectorsize"`
				Size            string      `json:"size"`
				SupportDiscard  string      `json:"support_discard"`
				Vendor          interface{} `json:"vendor"`
				Virtual         int         `json:"virtual"`
			} `json:"loop3"`
			Loop4 struct {
				Holders []interface{} `json:"holders"`
				Host    string        `json:"host"`
				Links   struct {
					Ids     []interface{} `json:"ids"`
					Labels  []interface{} `json:"labels"`
					Masters []interface{} `json:"masters"`
					Uuids   []interface{} `json:"uuids"`
				} `json:"links"`
				Model      interface{} `json:"model"`
				Partitions struct {
				} `json:"partitions"`
				Removable       string      `json:"removable"`
				Rotational      string      `json:"rotational"`
				SasAddress      interface{} `json:"sas_address"`
				SasDeviceHandle interface{} `json:"sas_device_handle"`
				SchedulerMode   string      `json:"scheduler_mode"`
				Sectors         string      `json:"sectors"`
				Sectorsize      string      `json:"sectorsize"`
				Size            string      `json:"size"`
				SupportDiscard  string      `json:"support_discard"`
				Vendor          interface{} `json:"vendor"`
				Virtual         int         `json:"virtual"`
			} `json:"loop4"`
			Loop5 struct {
				Holders []interface{} `json:"holders"`
				Host    string        `json:"host"`
				Links   struct {
					Ids     []interface{} `json:"ids"`
					Labels  []interface{} `json:"labels"`
					Masters []interface{} `json:"masters"`
					Uuids   []interface{} `json:"uuids"`
				} `json:"links"`
				Model      interface{} `json:"model"`
				Partitions struct {
				} `json:"partitions"`
				Removable       string      `json:"removable"`
				Rotational      string      `json:"rotational"`
				SasAddress      interface{} `json:"sas_address"`
				SasDeviceHandle interface{} `json:"sas_device_handle"`
				SchedulerMode   string      `json:"scheduler_mode"`
				Sectors         string      `json:"sectors"`
				Sectorsize      string      `json:"sectorsize"`
				Size            string      `json:"size"`
				SupportDiscard  string      `json:"support_discard"`
				Vendor          interface{} `json:"vendor"`
				Virtual         int         `json:"virtual"`
			} `json:"loop5"`
			Loop6 struct {
				Holders []interface{} `json:"holders"`
				Host    string        `json:"host"`
				Links   struct {
					Ids     []interface{} `json:"ids"`
					Labels  []interface{} `json:"labels"`
					Masters []interface{} `json:"masters"`
					Uuids   []interface{} `json:"uuids"`
				} `json:"links"`
				Model      interface{} `json:"model"`
				Partitions struct {
				} `json:"partitions"`
				Removable       string      `json:"removable"`
				Rotational      string      `json:"rotational"`
				SasAddress      interface{} `json:"sas_address"`
				SasDeviceHandle interface{} `json:"sas_device_handle"`
				SchedulerMode   string      `json:"scheduler_mode"`
				Sectors         string      `json:"sectors"`
				Sectorsize      string      `json:"sectorsize"`
				Size            string      `json:"size"`
				SupportDiscard  string      `json:"support_discard"`
				Vendor          interface{} `json:"vendor"`
				Virtual         int         `json:"virtual"`
			} `json:"loop6"`
			Loop7 struct {
				Holders []interface{} `json:"holders"`
				Host    string        `json:"host"`
				Links   struct {
					Ids     []interface{} `json:"ids"`
					Labels  []interface{} `json:"labels"`
					Masters []interface{} `json:"masters"`
					Uuids   []interface{} `json:"uuids"`
				} `json:"links"`
				Model      interface{} `json:"model"`
				Partitions struct {
				} `json:"partitions"`
				Removable       string      `json:"removable"`
				Rotational      string      `json:"rotational"`
				SasAddress      interface{} `json:"sas_address"`
				SasDeviceHandle interface{} `json:"sas_device_handle"`
				SchedulerMode   string      `json:"scheduler_mode"`
				Sectors         string      `json:"sectors"`
				Sectorsize      string      `json:"sectorsize"`
				Size            string      `json:"size"`
				SupportDiscard  string      `json:"support_discard"`
				Vendor          interface{} `json:"vendor"`
				Virtual         int         `json:"virtual"`
			} `json:"loop7"`
			Sda struct {
				Holders []interface{} `json:"holders"`
				Host    string        `json:"host"`
				Links   struct {
					Ids     []string      `json:"ids"`
					Labels  []interface{} `json:"labels"`
					Masters []interface{} `json:"masters"`
					Uuids   []interface{} `json:"uuids"`
				} `json:"links"`
				Model      string `json:"model"`
				Partitions struct {
					Sda1 struct {
						Holders []interface{} `json:"holders"`
						Links   struct {
							Ids     []string      `json:"ids"`
							Labels  []interface{} `json:"labels"`
							Masters []interface{} `json:"masters"`
							Uuids   []interface{} `json:"uuids"`
						} `json:"links"`
						Sectors    string      `json:"sectors"`
						Sectorsize int         `json:"sectorsize"`
						Size       string      `json:"size"`
						Start      string      `json:"start"`
						UUID       interface{} `json:"uuid"`
					} `json:"sda1"`
					Sda2 struct {
						Holders []interface{} `json:"holders"`
						Links   struct {
							Ids     []string      `json:"ids"`
							Labels  []interface{} `json:"labels"`
							Masters []interface{} `json:"masters"`
							Uuids   []string      `json:"uuids"`
						} `json:"links"`
						Sectors    string `json:"sectors"`
						Sectorsize int    `json:"sectorsize"`
						Size       string `json:"size"`
						Start      string `json:"start"`
						UUID       string `json:"uuid"`
					} `json:"sda2"`
					Sda3 struct {
						Holders []string `json:"holders"`
						Links   struct {
							Ids     []string      `json:"ids"`
							Labels  []interface{} `json:"labels"`
							Masters []string      `json:"masters"`
							Uuids   []interface{} `json:"uuids"`
						} `json:"links"`
						Sectors    string      `json:"sectors"`
						Sectorsize int         `json:"sectorsize"`
						Size       string      `json:"size"`
						Start      string      `json:"start"`
						UUID       interface{} `json:"uuid"`
					} `json:"sda3"`
				} `json:"partitions"`
				Removable       string      `json:"removable"`
				Rotational      string      `json:"rotational"`
				SasAddress      interface{} `json:"sas_address"`
				SasDeviceHandle interface{} `json:"sas_device_handle"`
				SchedulerMode   string      `json:"scheduler_mode"`
				Sectors         string      `json:"sectors"`
				Sectorsize      string      `json:"sectorsize"`
				Size            string      `json:"size"`
				SupportDiscard  string      `json:"support_discard"`
				Vendor          string      `json:"vendor"`
				Virtual         int         `json:"virtual"`
			} `json:"sda"`
			Sr0 struct {
				Holders []interface{} `json:"holders"`
				Host    string        `json:"host"`
				Links   struct {
					Ids     []string      `json:"ids"`
					Labels  []interface{} `json:"labels"`
					Masters []interface{} `json:"masters"`
					Uuids   []interface{} `json:"uuids"`
				} `json:"links"`
				Model      string `json:"model"`
				Partitions struct {
				} `json:"partitions"`
				Removable       string      `json:"removable"`
				Rotational      string      `json:"rotational"`
				SasAddress      interface{} `json:"sas_address"`
				SasDeviceHandle interface{} `json:"sas_device_handle"`
				SchedulerMode   string      `json:"scheduler_mode"`
				Sectors         string      `json:"sectors"`
				Sectorsize      string      `json:"sectorsize"`
				Size            string      `json:"size"`
				SupportDiscard  string      `json:"support_discard"`
				Vendor          string      `json:"vendor"`
				Virtual         int         `json:"virtual"`
			} `json:"sr0"`
		} `json:"ansible_devices"`
		AnsibleDistribution             string `json:"ansible_distribution"`
		AnsibleDistributionFileParsed   bool   `json:"ansible_distribution_file_parsed"`
		AnsibleDistributionFilePath     string `json:"ansible_distribution_file_path"`
		AnsibleDistributionFileVariety  string `json:"ansible_distribution_file_variety"`
		AnsibleDistributionMajorVersion string `json:"ansible_distribution_major_version"`
		AnsibleDistributionRelease      string `json:"ansible_distribution_release"`
		AnsibleDistributionVersion      string `json:"ansible_distribution_version"`
		AnsibleDNS                      struct {
			Nameservers []string `json:"nameservers"`
			Search      []string `json:"search"`
		} `json:"ansible_dns"`
		AnsibleDomain           string `json:"ansible_domain"`
		AnsibleEffectiveGroupID int    `json:"ansible_effective_group_id"`
		AnsibleEffectiveUserID  int    `json:"ansible_effective_user_id"`
		AnsibleEnv              struct {
			HOME        string `json:"HOME"`
			LANG        string `json:"LANG"`
			LOGNAME     string `json:"LOGNAME"`
			MAIL        string `json:"MAIL"`
			PATH        string `json:"PATH"`
			PWD         string `json:"PWD"`
			SHELL       string `json:"SHELL"`
			SUDOCOMMAND string `json:"SUDO_COMMAND"`
			SUDOGID     string `json:"SUDO_GID"`
			SUDOUID     string `json:"SUDO_UID"`
			SUDOUSER    string `json:"SUDO_USER"`
			TERM        string `json:"TERM"`
			USER        string `json:"USER"`
		} `json:"ansible_env"`
		AnsibleFibreChannelWwn []interface{} `json:"ansible_fibre_channel_wwn"`
		AnsibleFips            bool          `json:"ansible_fips"`
		AnsibleFormFactor      string        `json:"ansible_form_factor"`
		AnsibleFqdn            string        `json:"ansible_fqdn"`
		AnsibleHostname        string        `json:"ansible_hostname"`
		AnsibleHostnqn         string        `json:"ansible_hostnqn"`
		AnsibleInterfaces      []string      `json:"ansible_interfaces"`
		AnsibleIsChroot        bool          `json:"ansible_is_chroot"`
		AnsibleIscsiIqn        string        `json:"ansible_iscsi_iqn"`
		AnsibleKernel          string        `json:"ansible_kernel"`
		AnsibleKernelVersion   string        `json:"ansible_kernel_version"`
		AnsibleLo              struct {
			Active bool   `json:"active"`
			Device string `json:"device"`
			Ipv4   struct {
				Address   string `json:"address"`
				Broadcast string `json:"broadcast"`
				Netmask   string `json:"netmask"`
				Network   string `json:"network"`
				Prefix    string `json:"prefix"`
			} `json:"ipv4"`
			Ipv6 []struct {
				Address string `json:"address"`
				Prefix  string `json:"prefix"`
				Scope   string `json:"scope"`
			} `json:"ipv6"`
			Mtu     int    `json:"mtu"`
			Promisc bool   `json:"promisc"`
			Type    string `json:"type"`
		} `json:"ansible_lo"`
		AnsibleLocal struct {
		} `json:"ansible_local"`
		AnsibleLsb struct {
			Codename     string `json:"codename"`
			Description  string `json:"description"`
			ID           string `json:"id"`
			MajorRelease string `json:"major_release"`
			Release      string `json:"release"`
		} `json:"ansible_lsb"`
		AnsibleLvm struct {
			Lvs map[string]LvsFromJSON `json:"lvs"`
			Pvs map[string]PvsFromJSON `json:"pvs"`
			Vgs map[string]VgsFromJSON `json:"vgs"`
		} `json:"ansible_lvm"`
		AnsibleMachine   string `json:"ansible_machine"`
		AnsibleMachineID string `json:"ansible_machine_id"`
		AnsibleMemfreeMb int    `json:"ansible_memfree_mb"`
		AnsibleMemoryMb  struct {
			Nocache struct {
				Free int `json:"free"`
				Used int `json:"used"`
			} `json:"nocache"`
			Real struct {
				Free  int `json:"free"`
				Total int `json:"total"`
				Used  int `json:"used"`
			} `json:"real"`
			Swap struct {
				Cached int `json:"cached"`
				Free   int `json:"free"`
				Total  int `json:"total"`
				Used   int `json:"used"`
			} `json:"swap"`
		} `json:"ansible_memory_mb"`
		AnsibleMemtotalMb int `json:"ansible_memtotal_mb"`
		AnsibleMounts     []struct {
			BlockAvailable int64  `json:"block_available"`
			BlockSize      int64  `json:"block_size"`
			BlockTotal     int64  `json:"block_total"`
			BlockUsed      int64  `json:"block_used"`
			Device         string `json:"device"`
			Fstype         string `json:"fstype"`
			InodeAvailable int64  `json:"inode_available"`
			InodeTotal     int64  `json:"inode_total"`
			InodeUsed      int64  `json:"inode_used"`
			Mount          string `json:"mount"`
			Options        string `json:"options"`
			SizeAvailable  int64  `json:"size_available"`
			SizeTotal      int64  `json:"size_total"`
			UUID           string `json:"uuid"`
		} `json:"ansible_mounts"`
		AnsibleNodename    string `json:"ansible_nodename"`
		AnsibleOsFamily    string `json:"ansible_os_family"`
		AnsiblePkgMgr      string `json:"ansible_pkg_mgr"`
		AnsibleProcCmdline struct {
			BOOTIMAGE string `json:"BOOT_IMAGE"`
			Elevator  string `json:"elevator"`
			Quiet     bool   `json:"quiet"`
			Ro        bool   `json:"ro"`
			Root      string `json:"root"`
		} `json:"ansible_proc_cmdline"`
		AnsibleProcessor               []string `json:"ansible_processor"`
		AnsibleProcessorCores          int      `json:"ansible_processor_cores"`
		AnsibleProcessorCount          int      `json:"ansible_processor_count"`
		AnsibleProcessorNproc          int      `json:"ansible_processor_nproc"`
		AnsibleProcessorThreadsPerCore int      `json:"ansible_processor_threads_per_core"`
		AnsibleProcessorVcpus          int      `json:"ansible_processor_vcpus"`
		AnsibleProductName             string   `json:"ansible_product_name"`
		AnsibleProductSerial           string   `json:"ansible_product_serial"`
		AnsibleProductUUID             string   `json:"ansible_product_uuid"`
		AnsibleProductVersion          string   `json:"ansible_product_version"`
		AnsiblePython                  struct {
			Executable    string `json:"executable"`
			HasSslcontext bool   `json:"has_sslcontext"`
			Type          string `json:"type"`
			Version       struct {
				Major        int    `json:"major"`
				Micro        int    `json:"micro"`
				Minor        int    `json:"minor"`
				Releaselevel string `json:"releaselevel"`
				Serial       int    `json:"serial"`
			} `json:"version"`
			VersionInfo []interface{} `json:"version_info"`
		} `json:"ansible_python"`
		AnsiblePythonVersion string `json:"ansible_python_version"`
		AnsibleRealGroupID   int    `json:"ansible_real_group_id"`
		AnsibleRealUserID    int    `json:"ansible_real_user_id"`
		AnsibleSelinux       struct {
			Status string `json:"status"`
		} `json:"ansible_selinux"`
		AnsibleSelinuxPythonPresent           bool     `json:"ansible_selinux_python_present"`
		AnsibleServiceMgr                     string   `json:"ansible_service_mgr"`
		AnsibleSSHHostKeyDsaPublic            string   `json:"ansible_ssh_host_key_dsa_public"`
		AnsibleSSHHostKeyDsaPublicKeytype     string   `json:"ansible_ssh_host_key_dsa_public_keytype"`
		AnsibleSSHHostKeyEcdsaPublic          string   `json:"ansible_ssh_host_key_ecdsa_public"`
		AnsibleSSHHostKeyEcdsaPublicKeytype   string   `json:"ansible_ssh_host_key_ecdsa_public_keytype"`
		AnsibleSSHHostKeyEd25519Public        string   `json:"ansible_ssh_host_key_ed25519_public"`
		AnsibleSSHHostKeyEd25519PublicKeytype string   `json:"ansible_ssh_host_key_ed25519_public_keytype"`
		AnsibleSSHHostKeyRsaPublic            string   `json:"ansible_ssh_host_key_rsa_public"`
		AnsibleSSHHostKeyRsaPublicKeytype     string   `json:"ansible_ssh_host_key_rsa_public_keytype"`
		AnsibleSwapfreeMb                     int      `json:"ansible_swapfree_mb"`
		AnsibleSwaptotalMb                    int      `json:"ansible_swaptotal_mb"`
		AnsibleSystem                         string   `json:"ansible_system"`
		AnsibleSystemCapabilities             []string `json:"ansible_system_capabilities"`
		AnsibleSystemCapabilitiesEnforced     string   `json:"ansible_system_capabilities_enforced"`
		AnsibleSystemVendor                   string   `json:"ansible_system_vendor"`
		AnsibleUptimeSeconds                  int      `json:"ansible_uptime_seconds"`
		AnsibleUserDir                        string   `json:"ansible_user_dir"`
		AnsibleUserGecos                      string   `json:"ansible_user_gecos"`
		AnsibleUserGid                        int      `json:"ansible_user_gid"`
		AnsibleUserID                         string   `json:"ansible_user_id"`
		AnsibleUserShell                      string   `json:"ansible_user_shell"`
		AnsibleUserUID                        int      `json:"ansible_user_uid"`
		AnsibleUserspaceArchitecture          string   `json:"ansible_userspace_architecture"`
		AnsibleUserspaceBits                  string   `json:"ansible_userspace_bits"`
		AnsibleVirtualizationRole             string   `json:"ansible_virtualization_role"`
		AnsibleVirtualizationTechGuest        []string `json:"ansible_virtualization_tech_guest"`
		AnsibleVirtualizationTechHost         []string `json:"ansible_virtualization_tech_host"`
		AnsibleVirtualizationType             string   `json:"ansible_virtualization_type"`
		GatherSubset                          []string `json:"gather_subset"`
		ModuleSetup                           bool     `json:"module_setup"`
	} `json:"ansible_facts"`
	Changed bool `json:"changed"`
}

var allServer []AnsibleFactFull

var ServerFacts struct {
	PageTitle    string
	Version      string
	CreationDate time.Time
	Facts        []AnsibleFactFull
}
