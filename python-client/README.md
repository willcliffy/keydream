# Keydream Python Client

This is version of the client that is being actively developed.

## Getting started

- Install Python 3.10 or later.
- Install project dependencies:

    ``` bash
    pip install pygame
    pip install pytmx
    ```

- Clone project:

    ``` bash
    git clone https://github.com/willcliffy/keydream.git
    ```

- Enter project directory:

    ``` bash
    cd keydream/python-client
    ```

- Run the game

    ``` bash
    python main.py
    ```

## Features

### Implemented

- Keyboard input for movement
- Player animations (idle and walking)
- Tiled map integration with PyTMX
  - Collision detection with Tiled objects
- Player-following camera
- y-sort camera

### Planned

- Multi-level support through Tiled
- story mode
- NPCs
  - chat
- combat system
  - PVM, bosses
- inventory system
  - equipment + weapons
  - "quest items"
  - currency & shops?
- SFX + music
- Camera improvements

### Reach

- Advanced player/npc animation
- Toggleable multiplayer mode
  - economy
  - chat
  - pvm combat
- Steam beta release
