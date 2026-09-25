package ipod

import "github.com/phlagg/infiniti-pod/iap"

func ResetDBSelection() *iap.Command          { return nil }
func SelectDBRecord(data []byte) *iap.Command { return nil }

func GetNumberCategorizedDBRecords(data []byte) *iap.Command { return nil }
func RetrieveCategorizedDBRecord(data []byte) *iap.Command   { return nil }

func SelectSortDBRecord(data []byte) *iap.Command { return nil }
func ResetDBSelectionHierarchy() *iap.Command     { return nil }

func GetDBiTunesInfo(data []byte) *iap.Command { return nil }
func GetUIDTrackInfo(data []byte) *iap.Command { return nil }
func GetDBTrackInfo(data []byte) *iap.Command  { return nil }
func GetPBTrackInfo(data []byte) *iap.Command  { return nil }
