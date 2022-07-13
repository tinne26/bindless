# Bindless
My entry for [Ebitengine's first game jam ever](https://itch.io/jam/ebiten-game-jam/results)! The jam's theme was "magnets".

![Bindless tutorial level](https://github.com/tinne26/bindless/blob/main/screenshots/01.png)

Bindless is a puzzle game with a dystopian background story. The puzzles are mostly simulations of levels with magnets, where you can poke and prod a few things in order to solve them.

You can find static binaries in the releases, on itch.io, or if you have Golang 1.18+, you can also run directly with:
```
go run github.com/tinne26/bindless@v0.0.1
```

## Controls
- Use the mouse and left-click to select and use your abilities.
- Press TAB on levels to reset them if you are locked.

Additional controls:
- Press F to switch between fullscreen and windowed modes.
- Use 1-9 to make it easier to switch between abilities.
- Press TAB to skip the typewriter effect on story sections.

## Mechanics
If you already tried the game but are struggling to understand how the puzzles work, read this:
- You have limited "charges" for each ability on each level.
- Small magnets can move pushed or pulled by other magnets, while larger magnets are static, always stuck in place.
- Magnets have polarity (positive / negative) and can be powered or unpowered. Powered magnets will attract or repel each other. Unpowered magnets don't interact.
- Large magnets have a wider magnetic field than small magnets (3 vs 2 tiles).
- Circuits can share or transfer power from one magnet to another.
- In the cases where magnets could be pulled/pushed in different directions, closeness to other magnets and movement inertia are the deciding factors.

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
