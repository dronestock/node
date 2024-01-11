package config

type Binary struct {
	Pnpm string `default:"${BINARY_PNPM=pnpm}"`
	Yarn string `default:"${BINARY_YARN=yarn}"`
	Npm  string `default:"${BINARY_PNPM=npm}"`
}
