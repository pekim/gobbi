// This is a generated file - DO NOT EDIT
// +build gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

package gio

import "C"

type NetworkConnectivity C.GNetworkConnectivity

const (
	NETWORK_CONNECTIVITY_LOCAL   NetworkConnectivity = 1
	NETWORK_CONNECTIVITY_LIMITED NetworkConnectivity = 2
	NETWORK_CONNECTIVITY_PORTAL  NetworkConnectivity = 3
	NETWORK_CONNECTIVITY_FULL    NetworkConnectivity = 4
)
