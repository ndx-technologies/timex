package timex

import "time"

func TimeParser(v *time.Time) func(string) error {
	return func(s string) error {
		if s == "" {
			return nil
		}

		if s == "now" {
			*v = time.Now()
			return nil
		}

		t, err := time.Parse(time.RFC3339, s)
		if err != nil {
			return err
		}

		*v = t
		return nil
	}
}
