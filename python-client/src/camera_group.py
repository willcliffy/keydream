# See: https://www.youtube.com/watch?v=u7LPRqrzry8

import pygame

from src.map import Map
from src.player import (
    SCALED_DEFAULT_PLAYER_HEIGHT,
    SCALED_DEFAULT_PLAYER_WIDTH,
    SCALED_SCREEN_WIDTH,
    SCALED_SCREEN_HEIGHT,
    SCREEN_SCALE,
    Player
)

PLAYER_X_OFFSET = SCALED_DEFAULT_PLAYER_HEIGHT / 2 - SCALED_SCREEN_WIDTH / 2
PLAYER_Y_OFFSET = SCALED_DEFAULT_PLAYER_WIDTH / 2 - SCALED_SCREEN_HEIGHT / 2

class CameraGroup(pygame.sprite.Group):
    display: pygame.Surface
    map: Map
    player: Player

    def __init__(self, screen, map):
        self.display = screen
        self.map = map

    def follow_player(self, player):
        self.player = player

    def render(self):
        self.display.fill((0, 0, 0))
        offset = (0, 0)
        ysort = 0 # todo - have no idea how this acts when unlocked

        if self.player:
            offset = (self.player.rect.x + PLAYER_X_OFFSET, self.player.rect.y + PLAYER_Y_OFFSET)
            ysort = self.player.hitbox.y
            # ysort will be somewhat broken because we're ysorting by the tile position, not the object position
            # todo - fix this

        self.map.render_ysort_1(self.display, offset, ysort)
        self.player.draw(self.display, offset)
        self.map.render_ysort_2(self.display, offset, ysort)
