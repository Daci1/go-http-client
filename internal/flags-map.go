package internal

type FlagsMap struct {
	flagsMap map[string]string
}

func (fm *FlagsMap) getValue(key string) string {
	return fm.flagsMap[key]
}
