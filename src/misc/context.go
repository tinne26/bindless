package misc

import "io/ioutil"
import "embed"

import "github.com/tinne26/etxt"
import "github.com/tinne26/etxt/ecache"

// Shared context for scene initializations.
type Context struct {
	Filesys *embed.FS
	FontLib *etxt.FontLibrary
	FontCache *ecache.DefaultCache
	// ...
}

func NewContext(filesys *embed.FS) (*Context, error) {
	fontLib := etxt.NewFontLibrary()
	file, err := filesys.Open("assets/fonts/Coda-Regular.ttf")
	if err != nil { return nil, err }
	bytes, err := ioutil.ReadAll(file)
	if err != nil { return nil, err }
	_, err = fontLib.ParseFontBytes(bytes)
	if err != nil { return nil, err }

	// _, _, err := fontLib.ParseEmbedDirFonts("assets/fonts", filesys)
	// if err != nil { return nil, err }

	fontCache, err := ecache.NewDefaultCache(32*1024*1024) // 32MB cache
	if err != nil { return nil, err }

	return &Context {
		Filesys: filesys,
		FontLib: fontLib,
		FontCache: fontCache,
	}, nil
}
