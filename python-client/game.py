import pygame

from src.input import KeyboardInput
from src.player import Player
from src.map import Map
from src.camera_group import CameraGroup


SCREEN_SCALE = 2

SCALED_SCREEN_WIDTH = 720 * SCREEN_SCALE
SCALED_SCREEN_HEIGHT = 480 * SCREEN_SCALE


class Game:
    screen: pygame.Surface = None
    input: KeyboardInput = None
    map: Map = None
    player: pygame.sprite.Sprite = None

    def __init__(self, show_hitboxes: bool = False):
        self.screen = pygame.display.set_mode([SCALED_SCREEN_WIDTH, SCALED_SCREEN_HEIGHT])
        self.input = KeyboardInput()
        self.map = Map(show_hitboxes=show_hitboxes)
        self.player = Player(self.map.current_level, show_hitbox=show_hitboxes)
        self.camera = CameraGroup(self.screen, self.map)
        self.camera.follow_player(self.player)

    def run(self):
        pygame_clock = pygame.time.Clock()

        pygame.init()
        while self.input.input_system():
            self.player.update(self.input.directions)
            self.camera.render()
            pygame.display.update()
            pygame_clock.tick(60)
        pygame.quit()
