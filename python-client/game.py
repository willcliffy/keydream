import collections
import pygame

from src.input import KeyboardInput
from src.player import Player
from src.map import Map
from src.camera import CameraGroup
from src.npc_controller import NPCController
from src.npc import NPC
from src.keydream_sprite import KeydreamSprite


SCREEN_SCALE = 2

SCALED_SCREEN_WIDTH = 720 * SCREEN_SCALE
SCALED_SCREEN_HEIGHT = 480 * SCREEN_SCALE


class Game:
    screen: pygame.Surface = None
    input: KeyboardInput = None
    map: Map = None
    player: KeydreamSprite = None
    npcs: list[KeydreamSprite] = []

    def __init__(self, resizeable: bool = False, show_hitboxes: bool = False):
        self.screen = pygame.display.set_mode([SCALED_SCREEN_WIDTH, SCALED_SCREEN_HEIGHT], pygame.RESIZABLE if resizeable else 0)
        self.input = KeyboardInput()

        self.map = Map(show_hitboxes=show_hitboxes)
        self.player = Player(self.map.current_level.get_starting_position(), show_hitbox=show_hitboxes)
        self.npc_controller = NPCController([
            NPC(self.map.current_level.get_npc_starting_postion(1),
            show_hitbox=show_hitboxes)
        ])

        self.camera = CameraGroup(self.screen, self.map, self.npc_controller)
        self.camera.follow_player(self.player)

    def run(self):
        pygame_clock = pygame.time.Clock()
        pygame.init()

        done = self.input.input_system()

        while not done:
            hitboxes = self.map.current_level.get_hitboxes() + self.npc_controller.npcs
            self.player.update(self.input.directions, hitboxes=hitboxes)
            self.npc_controller.update()
            self.camera.draw()
            pygame.display.update()
            pygame_clock.tick(60)
            done = self.input.input_system()

        pygame.quit()
