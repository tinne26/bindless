package misc

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
	_, _, err := fontLib.ParseEmbedDirFonts("assets/fonts", filesys)
	if err != nil { return nil, err }

	fontCache, err := ecache.NewDefaultCache(16*1024*1024) // 16MB cache
	if err != nil { return nil, err }

	return &Context {
		Filesys: filesys,
		FontLib: fontLib,
		FontCache: fontCache,
	}, nil
}
