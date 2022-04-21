import pygame

from player import DEFAULT_PLAYER_SPEED
from objects import SpriteDirection, SpriteState

class NPC(pygame.sprite.Sprite):
    images: dict = {}

    image: pygame.Surface = None
    rect: pygame.Rect = None

    hitbox: pygame.Rect = None

    speed: int = DEFAULT_PLAYER_SPEED
    state: SpriteState = SpriteState.IDLE
    direction: SpriteDirection = SpriteDirection.DOWN

    def __init__(self):
        pass
