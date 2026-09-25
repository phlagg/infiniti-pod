package ipod

import "github.com/phlagg/infiniti-pod/iap"

// General
func StartIDPS()                                 {}
func SetFIDTokenValues(data []byte) *iap.Command { return nil }
func EndIDPS()                                   {}
func GetIDPSStatus() *iap.Command                { return nil }
