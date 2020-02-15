package got

import "github.com/tennashi/got/linker"

func link() error {
	l, err := linker.New()
	if err != nil {
		return err
	}
	return l.Link()
}
