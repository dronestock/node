package main

func (p *plugin) install() (undo bool, err error) {
	switch p.Type {
	case typeNpm:
		err = p.npmInstall()
	case typeYarn:
		err = p.yarnInstall()
	}

	return
}
