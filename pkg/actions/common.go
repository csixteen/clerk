package actions

const dateLayout = "2006-01-02 15:04:05"

func getIdFieldAndValue(id string) (string, string) {
	if id[0] == '#' {
		return "id", id[1:]
	}

	return "name", id
}
