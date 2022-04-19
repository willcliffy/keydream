import pygame
from src.tiled_renderer import Level

class Map:
    levels: list
    current_level: Level

    def __init__(self, show_hitboxes: bool = False):
        self.levels = [Level("assets/map_0_collisions.tmx", show_hitboxes=show_hitboxes)]
        self.current_level = self.levels[0]

    def render(self, screen: pygame.Surface, offset: tuple[int, int]):
        self.current_level.draw(screen, offset)

    def render_ysort_1(self, screen: pygame.Surface, offset: tuple[int, int], ysort: int):
        self.current_level.draw_ysort_1(screen, offset, ysort)

    def render_ysort_2(self, screen: pygame.Surface, offset: tuple[int, int], ysort: int):
        self.current_level.draw_ysort_2(screen, offset, ysort)
