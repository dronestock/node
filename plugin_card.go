package main

func (p *plugin) Scheme() string {
	return `https://raw.githubusercontent.com/dronestock/node/master/scheme.json`
}

func (p *plugin) Card() (card any, err error) {
	card = p.card

	return
}
