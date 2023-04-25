package main

type binary struct {
	Pnpm string `default:"${BINARY_PNPM=pnpm}"`
	Yarn string `default:"${BINARY_YARN=yarn}"`
}
