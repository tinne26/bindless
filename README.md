# Bindless
My entry for [Ebitengine's first game jam ever](https://itch.io/jam/ebiten-game-jam)! Theme was "magnets".

You can find static binaries in the releases, or if you have Golang 1.18+, run directly with:
```
go run github.com/tinne26/bindless@v0.0.1
```

Bindless is a puzzle game with a dystopian background story. The puzzles are mostly simulations of levels with magnets, where you can poke and prod a few things in order to solve them.

## Controls
- You use the mouse and left-click to select and use your abilities.
- You press ESC on levels to reset them if you are locked.

Additional controls:
- You can press F to switch between fullscreen/windowed.
- You can also use 1-4 to make it easier to switch between abilities.
- You can also press ESC to skip the typewriter effect on story sections.


## Licenses
Code is licensed under the MIT License. Assets are licensed as described in the readme from the [assets folder](https://github.com/tinne26/bindless/tree/main/assets).

## TODO
I'd like to...
- Add a more fully fledged ¿optional? tutorial stage after the preamble, explaining magnetism and giving some tips to solve levels.
- Make different pixel art wireframe scenes for story sections, at least a couple more.
- Figure out what's the problem with audio loops and Ebitengine streams.
- Add two extra levels at the end instead of brushing it away with text.
- Better management of window size / fullscreen on desktop, or explicitly mention F to fullscreen.
- Spanish and catalan translations.
