# Bindless
Bindless is a puzzle game with a dystopian background story created as part of [Ebitengine's first game jam ever](https://itch.io/jam/ebiten-game-jam/results)! The jam's theme was "magnets".

![Bindless tutorial level](https://github.com/tinne26/bindless/blob/main/screenshots/01.png)

The puzzles are based on electromagnetic simulations where you use abilities to affect magnets and circuits while trying to reach a target. The game is now available in English, Spanish and Catalan, with a duration of at least half an hour if you are good at puzzles, with many people needing more than an hour to reach the end.

You can play from the browser or download at [itch.io](https://tinne26.itch.io/bindless), get static binaries from the releases here on Github, or if you have Golang 1.18+, you can also build and run directly with:
```
go run github.com/tinne26/bindless@v0.0.3
```
*(Notice that on linux Ebitengine has a [few dependencies](https://ebitengine.org/en/documents/install.html?os=linux#Installing_dependencies) that you may need to install if you have never used Ebitengine with Golang.)*

## Controls And Mechanics
The game requires a mouse or touchpad to play. Mechanics are explained in the tutorial within the game itself.

## Licenses
Code is licensed under the MIT License. Assets are licensed as described in the readme from the [assets folder](https://github.com/tinne26/bindless/tree/main/assets).

Special acknowledgments:
- [hajimehoshi/ebiten](https://github.com/hajimehoshi/ebiten) (Ebitengine game engine).
- [Vernon Adams](https://fonts.google.com/?query=Vernon+Adams) (Coda and many other great fonts).
- [jeandeaual/go-locale](https://github.com/jeandeaual/go-locale) (Cross-platform locale detection).

## Misc. Esoteric Knowledge
- Unusual program arguments: `--directx` to use DirectX (Windows only), `--windowed` to start the game in windowed mode, `--maxfps` to test performance and `--en`, `--es` or `--ca` to set the language from the start overriding locale autodetection.
- Large magnets have a wider magnetic field than small magnets (3 vs 2 tiles).
- In the cases where magnets could be pulled/pushed in different directions, closeness to other magnets and movement inertia are the deciding factors.

## TODO
I'd like to...
- Make different pixel art wireframe scenes for story sections, at least a couple more.
- Add two extra levels at the end instead of brushing it away with text.
- Fix the magnetic simulation to be more accurate.
