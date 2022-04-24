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
        rect.x = starting_position[0] - rect.width
        rect.y = starting_position[1] - rect.height

        hitbox = rect.inflate(SCALED_HITBOX_INFL_X, SCALED_HITBOX_INFL_Y)
        hitbox.x += SCALED_HITBOX_X_OFFSET
        hitbox.y += SCALED_HITBOX_Y_OFFSET

        KeydreamSprite.__init__(self, images, image, rect, hitbox, DEFAULT_PLAYER_SPEED, show_hitbox)

