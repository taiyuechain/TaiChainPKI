package main

import (
	"TaiChainPKI/bccsp/factory"
	"TaiChainPKI/msp"
	mspmgmt "TaiChainPKI/msp/mgmt"
	"fmt"
	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

var localMsp msp.MSP

func main() {
	mspMgrConfigDir := "/home/dwh/go/src/github.com/hyperledger/fabric/sampleconfig/msp"
	localMSPID := "SampleOrg"
	localMSPType := "bccsp"
	InitCrypto(mspMgrConfigDir, localMSPID, localMSPType)

	localMsp = mspmgmt.GetLocalMSP(factory.GetDefault())

	identityPath := "/home/dwh/go/src/github.com/hyperledger/fabric/sampleconfig/msp"

	identity := getIdentity(identityPath)

	err := localMsp.Validate(identity)
	if err != nil {
		fmt.Print(identity.GetIdentifier())

	}
	fmt.Println(localMsp)
}

// InitCrypto initializes crypto for this node
func InitCrypto(mspMgrConfigDir, localMSPID, localMSPType string) error {
	// Check whether msp folder exists
	fi, err := os.Stat(mspMgrConfigDir)
	if os.IsNotExist(err) || !fi.IsDir() {
		// No need to try to load MSP from folder which is not available
		return errors.Errorf("cannot init crypto, folder \"%s\" does not exist", mspMgrConfigDir)
	}
	// Check whether localMSPID exists
	if localMSPID == "" {
		return errors.New("the local MSP must have an ID")
	}

	// Init the BCCSP
	bccspConfig := factory.GetDefaultOpts()
	if config := viper.Get("peer.BCCSP"); config != nil {
		err = mapstructure.Decode(config, bccspConfig)
		if err != nil {
			return errors.WithMessage(err, "could not decode peer BCCSP configuration")
		}
	}

	err = mspmgmt.LoadLocalMspWithType(mspMgrConfigDir, bccspConfig, localMSPID, localMSPType)
	if err != nil {
		return errors.WithMessagef(err, "error when setting up MSP of type %s from directory %s", localMSPType, mspMgrConfigDir)
	}

	return nil
}

func getIdentity(path string) msp.Identity {
	mspDir := path
	pems, err := msp.GetPemMaterialFromDir(filepath.Join(mspDir, path))
	if err != nil {
		fmt.Print(err)
	}
	id, _, err := localMsp.GetIdentityFromConf(pems[0])
	if err != nil {
		fmt.Print(err)
	}
	return id
}
