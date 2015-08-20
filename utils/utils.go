package utils

func now() time.Time {
	return time.Now().UTC().Format(Iso8061Format)
}
