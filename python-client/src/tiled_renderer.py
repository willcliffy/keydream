# See https://github.com/bitcraft/pytmx/blob/master/apps/pygame_demo.py

import pygame

from pytmx.util_pygame import load_pygame
from pytmx import TiledMap, TiledTileLayer, TiledObjectGroup, TiledObject


SCREEN_SCALE = 2

TILE_DIMENSION = 32

YSORT_COLLISION_LAYER = 4

class Tile(pygame.sprite.Sprite):
    image: pygame.Surface
    rect: pygame.Rect
    hitbox: pygame.Rect

    def __init__(self, image, x, y, collider=False):
        pygame.sprite.Sprite.__init__(self)
        self.image = image
        self.rect = self.image.get_rect()
        self.rect.x = x
        self.rect.y = y
        if collider:
            self.hitbox = collider


class Layer:
    tiles: pygame.sprite.Group
    colliders: list[Tile] = []
    show_colliders: bool = False

    def __init__(self, layer: TiledTileLayer, colliders, show_hitboxes: bool = False):
        self.tiles = pygame.sprite.Group()
        self.show_colliders = show_hitboxes

        collider_objects = {}
        for gid, object in colliders:
            collider_objects[gid] = object

        images = layer.parent.images
        for (x, y, gid) in layer.iter_data():
            image = images[gid]
            if image:
                tile = pygame.transform.scale(image,
                    (image.get_width() * SCREEN_SCALE, image.get_height() * SCREEN_SCALE))

                # TODO - document why this is the way that it is
                x = TILE_DIMENSION * (x - 2) - tile.get_width() / 2
                y = TILE_DIMENSION * (y - 2) - tile.get_height() / 2

                c_new: pygame.Rect = None
                if gid in collider_objects:
                    for collider in collider_objects[gid]:
                        c_new = pygame.Rect(
                            SCREEN_SCALE * (x + collider.x),
                            SCREEN_SCALE * (y + collider.y),
                            SCREEN_SCALE * collider.width,
                            SCREEN_SCALE * collider.height)
                        self.colliders.append(c_new)

                self.tiles.add(Tile(
                    image = tile,
                    x = SCREEN_SCALE * x,
                    y = SCREEN_SCALE * y,
                    collider = c_new))


    def draw(self, screen: pygame.Surface, offset: tuple[int, int]):
        for tile in self.tiles:
            screen.blit(tile.image, (tile.rect.x - offset[0], tile.rect.y - offset[1]))

    def draw_ysort_1(self, screen: pygame.Surface, offset: tuple[int, int], ysort: int):
        for tile in self.tiles:
            if tile.hitbox:
                if tile.hitbox.y < ysort:
                    screen.blit(tile.image, (tile.rect.x - offset[0], tile.rect.y - offset[1]))
            else:
                screen.blit(tile.image, (tile.rect.x - offset[0], tile.rect.y - offset[1]))

    def draw_ysort_2(self, screen: pygame.Surface, offset: tuple[int, int], ysort: int):
        for tile in self.tiles:
            if tile.hitbox and tile.hitbox.y > ysort:
                screen.blit(tile.image, (tile.rect.x - offset[0], tile.rect.y - offset[1]))
        if self.show_colliders:
            for collider in self.colliders:
                pygame.draw.rect(screen, (255, 0, 0), [collider.x - offset[0], collider.y - offset[1], collider.w, collider.h], 1)



class Level:
    map_object: TiledMap
    layers: list[Layer]

    def __init__(self, fileName: str, show_hitboxes: bool = False):
        self.map_object = load_pygame(fileName)

        self.layers = []

        self.levelShift = 0
        for layer in self.map_object.layers:
            if isinstance(layer, TiledTileLayer):
                self.layers.append(Layer(layer, self.map_object.get_tile_colliders(), show_hitboxes))

    def shiftLevel(self, shiftX: int):
        self.levelShift += shiftX
        for layer in self.layers:
            for tile in layer.tiles:
                tile.rect.x += shiftX

    def get_colliders(self, layer=YSORT_COLLISION_LAYER):
        return self.layers[layer].colliders

    def draw(self, screen: pygame.Surface, offset: tuple[int, int]):
        for layer in self.layers:
            layer.draw(screen, offset)

    def draw_ysort_1(self, screen: pygame.Surface, offset: tuple[int, int], ysort: int):
        for i in range(YSORT_COLLISION_LAYER):
            self.layers[i].draw(screen, offset)

        self.layers[YSORT_COLLISION_LAYER].draw_ysort_1(screen, offset, ysort)


    def draw_ysort_2(self, screen: pygame.Surface, offset: tuple[int, int], ysort: int):
        self.layers[YSORT_COLLISION_LAYER].draw_ysort_2(screen, offset, ysort)

        for i in range(YSORT_COLLISION_LAYER + 1, len(self.layers)):
            self.layers[i].draw(screen, offset)
