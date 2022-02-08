package main

func (p *plugin) scripts() (undo bool, err error) {
	for _, script := range p.Scripts {
		switch p.Type {
		case typeNpm:
			err = p.npmScript(script)
		case typeYarn:
			err = p.yarnScript(script)
		}

		if nil != err {
			return
		}
	}

	return
}
