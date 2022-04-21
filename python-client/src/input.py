import pygame

from src.objects import SpriteDirection


class KeyboardInput:
    directions: list
    actions: list

    def __init__(self):
        self.directions = []

    # Returns false when the game should end
    def input_system(self):
        for event in pygame.event.get():
            if event.type == pygame.QUIT:
                return True
            if event.type == pygame.KEYDOWN:
                if event.key == pygame.K_ESCAPE:
                    return True
                elif event.key == pygame.K_UP or event.key == pygame.K_w:
                    print("Up")
                    self.directions.append(SpriteDirection.UP)
                elif event.key == pygame.K_DOWN or event.key == pygame.K_s:
                    print("Down")
                    self.directions.append(SpriteDirection.DOWN)
                elif event.key == pygame.K_LEFT or event.key == pygame.K_a:
                    print("Left")
                    self.directions.append(SpriteDirection.LEFT)
                elif event.key == pygame.K_RIGHT or event.key == pygame.K_d:
                    print("Right")
                    self.directions.append(SpriteDirection.RIGHT)
                elif event.key == pygame.K_SPACE:
                    print("Space")
                    self.actions.append(SpriteDirection.ATTACK)
                elif event.key == pygame.K_LSHIFT:
                    print("Shift")
                    self.actions.append(SpriteDirection.ROLL)
            elif event.type == pygame.KEYUP:
                if event.key == pygame.K_UP or event.key == pygame.K_w:
                    self.directions.remove(SpriteDirection.UP)
                elif event.key == pygame.K_DOWN or event.key == pygame.K_s:
                    self.directions.remove(SpriteDirection.DOWN)
                elif event.key == pygame.K_LEFT or event.key == pygame.K_a:
                    self.directions.remove(SpriteDirection.LEFT)
                elif event.key == pygame.K_RIGHT or event.key == pygame.K_d:
                    self.directions.remove(SpriteDirection.RIGHT)
                elif event.key == pygame.K_SPACE:
                    self.actions.remove(SpriteDirection.ATTACK)
                elif event.key == pygame.K_LSHIFT:
                    self.actions.remove(SpriteDirection.ROLL)

        return False
