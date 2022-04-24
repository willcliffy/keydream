import pygame
from pytmx.util_pygame import load_pygame
from pytmx import TiledMap, TiledTileLayer

from src.keydream_sprite import KeydreamSprite


SCREEN_SCALE = 2
TILE_DIMENSION = 32
YSORT_COLLISION_LAYER = 7


class Tile(pygame.sprite.Sprite):
    image: pygame.Surface
    rect: pygame.Rect
    hitbox: pygame.Rect = None

    def __init__(self, image, x, y, hitbox=None):
        pygame.sprite.Sprite.__init__(self)
        self.image = image
        self.rect = self.image.get_rect()
        self.rect.x = x
        self.rect.y = y
        if hitbox:
            self.hitbox = hitbox

    def draw(self, screen: pygame.Surface, offset: tuple[int, int]):
        screen.blit(self.image, (self.rect.x - offset[0], self.rect.y - offset[1]))


class Layer:
    tiles: list[Tile]
    tiles_with_hitboxes: list[Tile] = []

    def __init__(self, layer: TiledTileLayer, hitboxes):
        self.tiles = pygame.sprite.Group()

        hitbox_dict = {}
        for gid, object in hitboxes:
            hitbox_dict[gid] = object

        images = layer.parent.images
        for (x, y, gid) in layer.iter_data():
            image = images[gid]
            if image:
                tile = pygame.transform.scale(image,
                    (image.get_width() * SCREEN_SCALE, image.get_height() * SCREEN_SCALE))

                # TODO - document why this is the way that it is
                x = TILE_DIMENSION * (x - 1) - tile.get_width() / 2
                y = TILE_DIMENSION * (y - 1) - tile.get_height() / 2

                hitbox_new: pygame.Rect = None
                if gid in hitbox_dict:
                    for hb in hitbox_dict[gid]:
                        self.tiles_with_hitboxes.append(Tile(
                            image = tile,
                            x = SCREEN_SCALE * x,
                            y = SCREEN_SCALE * y,
                            hitbox = pygame.Rect(
                                SCREEN_SCALE * (x + hb.x),
                                SCREEN_SCALE * (y + hb.y),
                                SCREEN_SCALE * hb.width,
                                SCREEN_SCALE * hb.height)))
                else:
                    self.tiles.add(Tile(
                        image = tile,
                        x = SCREEN_SCALE * x,
                        y = SCREEN_SCALE * y))


class Level:
    map_object: TiledMap
    layers: list[Layer]

    def __init__(self, fileName: str):
        self.map_object = load_pygame(fileName)

        self.layers = []

        self.levelShift = 0
        for layer in self.map_object.layers:
            if isinstance(layer, TiledTileLayer):
                self.layers.append(Layer(layer, self.map_object.get_tile_colliders()))

    def get_starting_position(self):
        return (
            self.map_object.get_object_by_name("starting_position").x * SCREEN_SCALE,
            self.map_object.get_object_by_name("starting_position").y * SCREEN_SCALE)
    
    def get_npc_starting_postion(self, id):
        return (
            self.map_object.get_object_by_name(f"npc_{id}_spawn").x * SCREEN_SCALE,
            self.map_object.get_object_by_name(f"npc_{id}_spawn").y * SCREEN_SCALE)

    def shiftLevel(self, shiftX: int):
        self.levelShift += shiftX
        for layer in self.layers:
            for tile in layer.tiles:
                tile.rect.x += shiftX

    def get_hitboxes(self, layer=YSORT_COLLISION_LAYER):
        return self.layers[layer].tiles_with_hitboxes


class Map:
    levels: list
    current_level: Level
    show_hitboxes: bool = False

    def __init__(self, show_hitboxes: bool = False):
        self.show_hitboxes = show_hitboxes
        self.levels = [
            Level("assets/map/level_1a.tmx"),
            Level("assets/map/level_1b.tmx"),
            Level("assets/map/level_2.tmx"),
            Level("assets/map/level_2a.tmx"),
            Level("assets/map/level_2b.tmx"),
            Level("assets/map/level_2c.tmx"),
            Level("assets/map/level_3.tmx"),
        ]
        self.current_level = self.levels[0]

    def draw(self, screen: pygame.Surface, offset: tuple[int, int], sprites: list[KeydreamSprite] = None):
        sprites = sorted(sprites, key=lambda sprite: sprite.hitbox.y)

        for i in range(YSORT_COLLISION_LAYER):
            layer = self.current_level.layers[i]
            for tile in layer.tiles:
                tile.draw(screen, offset)
        for i in range(YSORT_COLLISION_LAYER, len(self.current_level.layers)):
            layer = self.current_level.layers[i]
            for tile in layer.tiles:
                tile.draw(screen, offset)
            sorted_tiles_with_hitboxes = sorted(layer.tiles_with_hitboxes, key=lambda tile: tile.hitbox.y)
            for tile in sorted_tiles_with_hitboxes:
                if len(sprites) != 0 and tile.hitbox.y > sprites[0].hitbox.y:
                    tile.draw(screen, offset)
                    if self.show_hitboxes:
                        pygame.draw.rect(screen, (255, 0, 0), [tile.hitbox.x - offset[0], tile.hitbox.y - offset[1], tile.hitbox.w, tile.hitbox.h], 1)
                    sprites[0].draw(screen, offset)
                    sprites = sprites[1:]
                else:
                    tile.draw(screen, offset)
                    if self.show_hitboxes:
                        pygame.draw.rect(screen, (255, 0, 0), [tile.hitbox.x - offset[0], tile.hitbox.y - offset[1], tile.hitbox.w, tile.hitbox.h], 1)

        # print(len(sprites))
        for sprite in sprites:
            sprite.draw(screen, offset)

    def draw_old_condensed(self, screen: pygame.Surface, offset: tuple[int, int], sprites: list[KeydreamSprite] = None):
        # def draw_ysort_1(self, screen: pygame.Surface, offset: tuple[int, int], ysort: int):
        for i in range(YSORT_COLLISION_LAYER):
            #self.layers[i].draw(screen, offset)
            for tile in self.tiles:
                tile.draw(screen, offset)
        for i in range(YSORT_COLLISION_LAYER, len(self.layers)):
            # self.layers[i].draw_ysort_1(screen, offset, ysort)
            for tile in self.tiles:
                if tile.hitbox:
                    if tile.hitbox.y < ysort:
                        tile.draw(screen, offset)
            else:
                tile.draw(screen, offset)

        # def draw_ysort_2(self, screen: pygame.Surface, offset: tuple[int, int], ysort: int):
        for i in range(YSORT_COLLISION_LAYER, len(self.layers)):
            # self.layers[i].draw_ysort_2(screen, offset, ysort)
            for tile in self.tiles:
                if tile.hitbox and tile.hitbox.y > ysort:
                    tile.draw(screen, offset)
            if self.show_colliders:
                for collider in self.colliders:
                    pygame.draw.rect(screen, (255, 0, 0), [collider.x - offset[0], collider.y - offset[1], collider.w, collider.h], 1)
