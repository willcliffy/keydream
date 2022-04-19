import pygame

from src.objects.direction import Direction
from src.player import Player

class KeyboardInput:
    directions: list

    def __init__(self):
        self.directions = []

    # Returns false when the game should end
    def input_system(self) -> bool:
        for event in pygame.event.get():
            if event.type == pygame.QUIT:
                return False
            if event.type == pygame.KEYDOWN:
                if event.key == pygame.K_ESCAPE:
                    return False
                elif event.key == pygame.K_UP or event.key == pygame.K_w:
                    self.directions.append(Direction.UP)
                elif event.key == pygame.K_DOWN or event.key == pygame.K_s:
                    self.directions.append(Direction.DOWN)
                elif event.key == pygame.K_LEFT or event.key == pygame.K_a:
                    self.directions.append(Direction.LEFT)
                elif event.key == pygame.K_RIGHT or event.key == pygame.K_d:
                    self.directions.append(Direction.RIGHT)
            elif event.type == pygame.KEYUP:
                if event.key == pygame.K_UP or event.key == pygame.K_w:
                    self.directions.remove(Direction.UP)
                elif event.key == pygame.K_DOWN or event.key == pygame.K_s:
                    self.directions.remove(Direction.DOWN)
                elif event.key == pygame.K_LEFT or event.key == pygame.K_a:
                    self.directions.remove(Direction.LEFT)
                elif event.key == pygame.K_RIGHT or event.key == pygame.K_d:
                    self.directions.remove(Direction.RIGHT)

        return True
