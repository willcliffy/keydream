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

SCALED_HITBOX_INFL_X = -SCALED_DEFAULT_PLAYER_WIDTH * 5.0 / 6.0
SCALED_HITBOX_INFL_Y = -SCALED_DEFAULT_PLAYER_HEIGHT * 6.0 / 7.0
SCALED_HITBOX_X_OFFSET = 0
SCALED_HITBOX_Y_OFFSET = SCALED_DEFAULT_PLAYER_HEIGHT / 12
SCALED_HALF_HITBOX_INFL_X = SCALED_HITBOX_INFL_X / 2.0
SCALED_HALF_HITBOX_INFL_Y = SCALED_HITBOX_INFL_Y / 2.0


class Player(KeydreamSprite):
    def __init__(self, starting_position: tuple[int, int], show_hitbox: bool = False):
        images = {}

        for state in SpriteState:
            images[state] = {}
            for direction in SpriteDirection:
                images[state][direction] = []
                for frame in range(4):
                    url = f"assets/_rgs_dev/Character without weapon/{state.name.lower()}/{state.name.lower()} {direction.name.lower()}{frame + 1}.png"
                    image = pygame.image.load(url)
                    images[state][direction].append(pygame.transform.scale(
                        image,
                        (SCALED_DEFAULT_PLAYER_WIDTH, SCALED_DEFAULT_PLAYER_HEIGHT)))

        image = images[self.state][self.direction][self.current_frame]

        rect = image.get_rect()
        rect.x = starting_position[0]
        rect.y = starting_position[1]

        hitbox = rect.inflate(SCALED_HITBOX_INFL_X, SCALED_HITBOX_INFL_Y)
        hitbox.x += SCALED_HITBOX_X_OFFSET
        hitbox.y += SCALED_HITBOX_Y_OFFSET

        KeydreamSprite.__init__(self, images, image, rect, hitbox, DEFAULT_PLAYER_SPEED, show_hitbox)
