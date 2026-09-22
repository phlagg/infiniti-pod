package ipod

func ResetDBSelection()          {}
func SelectDBRecord(data []byte) {}

func GetNumberCategorizedDBRecords(data []byte) []byte { return nil }
func RetrieveCategorizedDBRecord(data []byte) []byte   { return nil }

func SelectSortDBRecord(data []byte) {}
func ResetDBSelectionHierarchy()     {}

func GetDBiTunesInfo(data []byte) []byte { return nil }
func GetUIDTrackInfo(data []byte) []byte { return nil }
func GetDBTrackInfo(data []byte) []byte  { return nil }
func GetPBTrackInfo(data []byte) []byte  { return nil }
