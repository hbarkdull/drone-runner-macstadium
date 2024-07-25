// Copyright 2020 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE file.

package orka

// Config configures a virtual machine.
type Config struct {
	Name  string `json:"name"`
	Image string `json:"image"`
	CPU   int    `json:"cpu"`
}

type (
	// Response provides the API response.
	Response struct {
		Message string   `json:"message"`
		Errors  []*Error `json:"errors"`
	}

	// CreateRequest provides the create API request.
	//
	//     {
	// 	    "orka_vm_name": "myorkavm",
	// 	    "orka_base_image": "myStorage.img",
	// 	    "orka_image": "myorkavm",
	// 	    "orka_cpu_core": 6,
	// 	    "vcpu_count": 6,
	// 	    "iso_image": "Mojave.iso"
	//     }
	//
	// TODO(bradrydzewski) model the full json structure (above) and
	// replace and remove the Config struct with this struct.
	CreateRequest struct {
		Name  string `json:"name"`
		Image string `json:"image"`
		CPU   int    `json:"cpu"`
	}
	
	CreateResponse struct {
		Response
		RAM             string        `json:"memory"`
		IP             	string        `json:"ip"`
		SSHPort         int           `json:"ssh"`
		ScreenSharePort int           `json:"screenshare"`
		VncPort         int           `json:"vnc"`
	}

	// DeployResponse provides the deployment API response.
	DeployResponse struct {
		Response
		RAM             string        `json:"memory"`
		IP             	string        `json:"ip"`
		SSHPort         int           `json:"ssh"`
		ScreenSharePort int           `json:"screenshare"`
		VncPort         int           `json:"vnc"`
	}

	// StatusResponse provides the status API response.
	StatusResponse struct {
		Response
		VirtualMachineName string `json:"name"`
		VMDeploymentStatus string `json:"status"`
		Owner                 string `json:"owner"`
		NodeLocation          string `json:"node"`
		VirtualMachineIP      string `json:"ip"`
		VncPort               string `json:"vnc"`
		ScreenSharingPort     string `json:"screenshare"`
		SSHPort               string `json:"ssh"`
		CPU                   int    `json:"cpu"`
		RAM                   string `json:"memory"`
		Image                 string `json:"image"`
		VMStatus              string `json:"status"`
	}

	// TokenResponse provides the token API response.
	TokenResponse struct {
		Response
		Authenticated  bool   `json:"authenticated"`
		IsTokenRevoked bool   `json:"is_token_revoked"`
		Email          string `json:"email"`
	}
)

// Error represents an API error.
type Error struct {
	Message string `json:"message"`
}

// Error returns the error message.
func (e *Error) Error() string {
	return e.Message
}
