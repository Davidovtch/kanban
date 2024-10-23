package forms

type errors map[string][]string

func (e errors) Add(field, message string) {
	e[field] = append(e[field], message)
}

func (e errors) List(field string) []string {
	errors := e[field]

	if len(errors) == 0 {
		return nil
	}

	return errors
}
