package ipod

import "github.com/phlagg/infiniti-pod/iap"

// General
func OpenDataSession(data []byte) *iap.Command  { return nil }
func CloseDataSession(data []byte) *iap.Command { return nil }

func HandleDevDataTransfer(data []byte) *iap.Command  { return nil }
func HandleiPodDataTransfer(data []byte) *iap.Command { return nil }
