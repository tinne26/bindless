package misc

import "fmt"
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
	loadedCount, _, err := fontLib.ParseEmbedDirFonts("assets/fonts", filesys)
	if err != nil { return nil, err }
	if loadedCount != 1 {
		return nil, fmt.Errorf("expected to load 1 font, got %d instead", loadedCount)
	}

	fontCache, err := ecache.NewDefaultCache(32*1024*1024) // 32MB cache
	if err != nil { return nil, err }

	return &Context {
		Filesys: filesys,
		FontLib: fontLib,
		FontCache: fontCache,
	}, nil
}
