# See: https://www.youtube.com/watch?v=u7LPRqrzry8

import pygame

from src.map import Map
from src.npc_controller import NPCController
from src.player import (
    SCALED_DEFAULT_PLAYER_HEIGHT,
    SCALED_DEFAULT_PLAYER_WIDTH,
    SCALED_SCREEN_WIDTH,
    SCALED_SCREEN_HEIGHT,
    Player
)

from src.npc import NPC

PLAYER_X_OFFSET = SCALED_DEFAULT_PLAYER_HEIGHT / 2 - SCALED_SCREEN_WIDTH / 2
PLAYER_Y_OFFSET = SCALED_DEFAULT_PLAYER_WIDTH / 2 - SCALED_SCREEN_HEIGHT / 2

class CameraGroup(pygame.sprite.Group):
    display: pygame.Surface
    map: Map
    player: Player
    npc_controller: NPCController

    def __init__(self, screen, map, npc_controller):
        self.display = screen
        self.map = map
        self.npc_controller = npc_controller

    def follow_player(self, player):
        self.player = player

    def draw(self):
        self.display.fill((0, 0, 0))
        self.map.draw(
            self.display,
            (self.player.rect.x + PLAYER_X_OFFSET, self.player.rect.y + PLAYER_Y_OFFSET) if self.player else (0, 0),
            self.npc_controller.npcs + [self.player])
