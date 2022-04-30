import pygame

from src.player import (
    DEFAULT_PLAYER_SPEED,
    SCALED_DEFAULT_PLAYER_HEIGHT,
    SCALED_DEFAULT_PLAYER_WIDTH,
    SCALED_HITBOX_INFL_X,
    SCALED_HITBOX_INFL_Y,
    SCALED_HITBOX_X_OFFSET,
    SCALED_HITBOX_Y_OFFSET,
)
from src.objects import SpriteDirection, SpriteState
from src.keydream_sprite import KeydreamSprite


class NPC(KeydreamSprite):
    def __init__(self, starting_position: tuple[int, int], show_hitbox: bool = False):
        images = {}

        image = pygame.image.load("assets/_cyberrumor/player.png")

        for state in SpriteState:
            images[state] = {
                SpriteDirection.DOWN: [
                    pygame.transform.scale(image.subsurface(pygame.Rect(00, 0, 16, 32)), (64, 128)),
                    pygame.transform.scale(image.subsurface(pygame.Rect(16, 0, 16, 32)), (64, 128)),
                    pygame.transform.scale(image.subsurface(pygame.Rect(32, 0, 16, 32)), (64, 128)),
                    pygame.transform.scale(image.subsurface(pygame.Rect(16, 0, 16, 32)), (64, 128)),
                ],
                SpriteDirection.UP: [
                    pygame.transform.scale(image.subsurface(pygame.Rect(00, 32, 16, 32)), (64, 128)),
                    pygame.transform.scale(image.subsurface(pygame.Rect(16, 32, 16, 32)), (64, 128)),
                    pygame.transform.scale(image.subsurface(pygame.Rect(32, 32, 16, 32)), (64, 128)),
                    pygame.transform.scale(image.subsurface(pygame.Rect(16, 32, 16, 32)), (64, 128)),
                ],
                SpriteDirection.LEFT: [
                    pygame.transform.scale(image.subsurface(pygame.Rect(00, 64, 16, 32)), (64, 128)),
                    pygame.transform.scale(image.subsurface(pygame.Rect(16, 64, 16, 32)), (64, 128)),
                    pygame.transform.scale(image.subsurface(pygame.Rect(32, 64, 16, 32)), (64, 128)),
                    pygame.transform.scale(image.subsurface(pygame.Rect(16, 64, 16, 32)), (64, 128)),
                ],
                SpriteDirection.RIGHT: [
                    pygame.transform.flip(pygame.transform.scale(image.subsurface(pygame.Rect(00, 64, 16, 32)), (64, 128)), True, False),
                    pygame.transform.flip(pygame.transform.scale(image.subsurface(pygame.Rect(16, 64, 16, 32)), (64, 128)), True, False),
                    pygame.transform.flip(pygame.transform.scale(image.subsurface(pygame.Rect(32, 64, 16, 32)), (64, 128)), True, False),
                    pygame.transform.flip(pygame.transform.scale(image.subsurface(pygame.Rect(16, 64, 16, 32)), (64, 128)), True, False),
                ]
            }

        image = images[self.state][self.direction][self.current_frame]

        rect = image.get_rect()
        rect.x = starting_position[0] - rect.width
        rect.y = starting_position[1] - rect.height

        hitbox = rect.inflate(SCALED_HITBOX_INFL_X, SCALED_HITBOX_INFL_Y)
        hitbox.x += SCALED_HITBOX_X_OFFSET
        hitbox.y += SCALED_HITBOX_Y_OFFSET

        KeydreamSprite.__init__(self, images, image, rect, hitbox, DEFAULT_PLAYER_SPEED, 3, show_hitbox)

    def interact(self, player):
        pass

