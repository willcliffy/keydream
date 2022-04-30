import pygame

from src.input import SpriteDirection
from src.objects import SpriteState

from src.keydream_sprite import KeydreamSprite


# todo - I'm really lazy. find a way to get config into every class
SCREEN_SCALE = 2 # TODO - this doesnt work, only affects player size

SCALED_SCREEN_WIDTH = 720 * SCREEN_SCALE
SCALED_SCREEN_HEIGHT = 480 * SCREEN_SCALE

DEFAULT_PLAYER_SPEED = 8

SCALED_DEFAULT_PLAYER_WIDTH = 128 * SCREEN_SCALE
SCALED_DEFAULT_PLAYER_HEIGHT = 128 * SCREEN_SCALE

SCALED_HITBOX_INFL_X = -20
SCALED_HITBOX_INFL_Y = -100
SCALED_HITBOX_X_OFFSET = 0
SCALED_HITBOX_Y_OFFSET = 5
SCALED_HALF_HITBOX_INFL_X = SCALED_HITBOX_INFL_X / 2.0
SCALED_HALF_HITBOX_INFL_Y = SCALED_HITBOX_INFL_Y / 2.0


class Player(KeydreamSprite):
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
        rect.x = starting_position[0] + rect.width
        rect.y = starting_position[1] + rect.height

        hitbox = rect.inflate(SCALED_HITBOX_INFL_X, SCALED_HITBOX_INFL_Y)
        hitbox.x += SCALED_HITBOX_X_OFFSET
        hitbox.y += SCALED_HITBOX_Y_OFFSET

        KeydreamSprite.__init__(self, images, image, rect, hitbox, DEFAULT_PLAYER_SPEED, 4, show_hitbox)
