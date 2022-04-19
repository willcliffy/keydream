import enum
import pygame

from src.tiled_renderer import Level
from src.objects.direction import Direction


# todo - I'm really lazy. find a way to get config into every class
SCREEN_SCALE = 2 # TODO - this doesnt work, only affects player size

SCALED_SCREEN_WIDTH = 720 * SCREEN_SCALE
SCALED_SCREEN_HEIGHT = 480 * SCREEN_SCALE

DEFAULT_PLAYER_SPEED = 8
DEFAULT_PLAYER_ANIMATION_DELAY = 200

SCALED_DEFAULT_PLAYER_WIDTH = 128 * SCREEN_SCALE
SCALED_DEFAULT_PLAYER_HEIGHT = 128 * SCREEN_SCALE

SCALED_HITBOX_INFL_X = -SCALED_DEFAULT_PLAYER_WIDTH * 5.0 / 6.0
SCALED_HITBOX_INFL_Y = -SCALED_DEFAULT_PLAYER_HEIGHT * 6.0 / 7.0
SCALED_HITBOX_X_OFFSET = 0
SCALED_HITBOX_Y_OFFSET = SCALED_DEFAULT_PLAYER_HEIGHT / 12
SCALED_HALF_HITBOX_INFL_X = SCALED_HITBOX_INFL_X / 2.0
SCALED_HALF_HITBOX_INFL_Y = SCALED_HITBOX_INFL_Y / 2.0

MAP_COLLISION_LAYER = 3


class PlayerState(enum.Enum):
    IDLE = 0
    WALKING = 1
    ROLLING = 2
    ATTACKING = 3

class Player(pygame.sprite.Sprite):
    images: dict = {
        PlayerState.IDLE: {
            Direction.UP: [
                pygame.image.load('assets/_rgs_dev/Character without weapon/idle/idle up1.png'),
                pygame.image.load('assets/_rgs_dev/Character without weapon/idle/idle up2.png'),
                pygame.image.load('assets/_rgs_dev/Character without weapon/idle/idle up3.png'),
                pygame.image.load('assets/_rgs_dev/Character without weapon/idle/idle up4.png'),
            ],
            Direction.DOWN: [
                pygame.image.load('assets/_rgs_dev/Character without weapon/idle/idle down1.png'),
                pygame.image.load('assets/_rgs_dev/Character without weapon/idle/idle down2.png'),
                pygame.image.load('assets/_rgs_dev/Character without weapon/idle/idle down3.png'),
                pygame.image.load('assets/_rgs_dev/Character without weapon/idle/idle down4.png'),
            ],
            Direction.LEFT: [
                pygame.image.load('assets/_rgs_dev/Character without weapon/idle/idle left1.png'),
                pygame.image.load('assets/_rgs_dev/Character without weapon/idle/idle left2.png'),
                pygame.image.load('assets/_rgs_dev/Character without weapon/idle/idle left3.png'),
                pygame.image.load('assets/_rgs_dev/Character without weapon/idle/idle left4.png'),
            ],
            Direction.RIGHT: [
                pygame.image.load('assets/_rgs_dev/Character without weapon/idle/idle right1.png'),
                pygame.image.load('assets/_rgs_dev/Character without weapon/idle/idle right2.png'),
                pygame.image.load('assets/_rgs_dev/Character without weapon/idle/idle right3.png'),
                pygame.image.load('assets/_rgs_dev/Character without weapon/idle/idle right4.png'),
            ],
        },
        PlayerState.WALKING: {
            Direction.UP: [
                pygame.image.load('assets/_rgs_dev/Character without weapon/walk/walk up1.png'),
                pygame.image.load('assets/_rgs_dev/Character without weapon/walk/walk up2.png'),
                pygame.image.load('assets/_rgs_dev/Character without weapon/walk/walk up3.png'),
                pygame.image.load('assets/_rgs_dev/Character without weapon/walk/walk up4.png'),
            ],
            Direction.DOWN: [
                pygame.image.load('assets/_rgs_dev/Character without weapon/walk/walk down1.png'),
                pygame.image.load('assets/_rgs_dev/Character without weapon/walk/walk down2.png'),
                pygame.image.load('assets/_rgs_dev/Character without weapon/walk/walk down3.png'),
                pygame.image.load('assets/_rgs_dev/Character without weapon/walk/walk down4.png'),
            ],
            Direction.LEFT: [
                pygame.image.load('assets/_rgs_dev/Character without weapon/walk/walk left1.png'),
                pygame.image.load('assets/_rgs_dev/Character without weapon/walk/walk left2.png'),
                pygame.image.load('assets/_rgs_dev/Character without weapon/walk/walk left3.png'),
                pygame.image.load('assets/_rgs_dev/Character without weapon/walk/walk left4.png'),
            ],
            Direction.RIGHT: [
                pygame.image.load('assets/_rgs_dev/Character without weapon/walk/walk right1.png'),
                pygame.image.load('assets/_rgs_dev/Character without weapon/walk/walk right2.png'),
                pygame.image.load('assets/_rgs_dev/Character without weapon/walk/walk right3.png'),
                pygame.image.load('assets/_rgs_dev/Character without weapon/walk/walk right4.png'),
            ],
        },
        PlayerState.ROLLING: {
            Direction.UP: [
                pygame.image.load('assets/_rgs_dev/Character without weapon/roll/roll up1.png'),
                pygame.image.load('assets/_rgs_dev/Character without weapon/roll/roll up2.png'),
                pygame.image.load('assets/_rgs_dev/Character without weapon/roll/roll up3.png'),
                pygame.image.load('assets/_rgs_dev/Character without weapon/roll/roll up4.png'),
            ],
            Direction.DOWN: [
                pygame.image.load('assets/_rgs_dev/Character without weapon/roll/roll down1.png'),
                pygame.image.load('assets/_rgs_dev/Character without weapon/roll/roll down2.png'),
                pygame.image.load('assets/_rgs_dev/Character without weapon/roll/roll down3.png'),
                pygame.image.load('assets/_rgs_dev/Character without weapon/roll/roll down4.png'),
            ],
            Direction.LEFT: [
                pygame.image.load('assets/_rgs_dev/Character without weapon/roll/roll left1.png'),
                pygame.image.load('assets/_rgs_dev/Character without weapon/roll/roll left2.png'),
                pygame.image.load('assets/_rgs_dev/Character without weapon/roll/roll left3.png'),
                pygame.image.load('assets/_rgs_dev/Character without weapon/roll/roll left4.png'),
            ],
            Direction.RIGHT: [
                pygame.image.load('assets/_rgs_dev/Character without weapon/roll/roll right1.png'),
                pygame.image.load('assets/_rgs_dev/Character without weapon/roll/roll right2.png'),
                pygame.image.load('assets/_rgs_dev/Character without weapon/roll/roll right3.png'),
                pygame.image.load('assets/_rgs_dev/Character without weapon/roll/roll right4.png'),
            ],
        },
        PlayerState.ATTACKING: {
            Direction.UP: [
                pygame.image.load('assets/_rgs_dev/Character without weapon/attack/attack up1.png'),
                pygame.image.load('assets/_rgs_dev/Character without weapon/attack/attack up2.png'),
                pygame.image.load('assets/_rgs_dev/Character without weapon/attack/attack up3.png'),
                pygame.image.load('assets/_rgs_dev/Character without weapon/attack/attack up4.png'),
            ],
            Direction.DOWN: [
                pygame.image.load('assets/_rgs_dev/Character without weapon/attack/attack down1.png'),
                pygame.image.load('assets/_rgs_dev/Character without weapon/attack/attack down2.png'),
                pygame.image.load('assets/_rgs_dev/Character without weapon/attack/attack down3.png'),
                pygame.image.load('assets/_rgs_dev/Character without weapon/attack/attack down4.png'),
            ],
            Direction.LEFT: [
                pygame.image.load('assets/_rgs_dev/Character without weapon/attack/attack left1.png'),
                pygame.image.load('assets/_rgs_dev/Character without weapon/attack/attack left2.png'),
                pygame.image.load('assets/_rgs_dev/Character without weapon/attack/attack left3.png'),
                pygame.image.load('assets/_rgs_dev/Character without weapon/attack/attack left4.png'),
            ],
            Direction.RIGHT: [
                pygame.image.load('assets/_rgs_dev/Character without weapon/attack/attack right1.png'),
                pygame.image.load('assets/_rgs_dev/Character without weapon/attack/attack right2.png'),
                pygame.image.load('assets/_rgs_dev/Character without weapon/attack/attack right3.png'),
                pygame.image.load('assets/_rgs_dev/Character without weapon/attack/attack right4.png'),
            ],
        },
    }

    image: pygame.Surface = None
    rect: pygame.Rect = None

    hitbox: pygame.Rect = None

    speed: int = DEFAULT_PLAYER_SPEED
    state: PlayerState = PlayerState.IDLE
    direction: Direction = Direction.DOWN

    current_level: Level = None

    current_frame: int = 0
    game_tick: int = 0

    def __init__(self, current_level: Level, show_hitbox: bool = False):
        pygame.sprite.Sprite.__init__(self)

        for state in self.images.keys():
            for direction in self.images[state].keys():
                for frame in range(4):
                    self.images[state][direction][frame] = pygame.transform.scale(
                        self.images[state][direction][frame],
                        (SCALED_DEFAULT_PLAYER_WIDTH, SCALED_DEFAULT_PLAYER_HEIGHT))

        self.show_hitbox = show_hitbox

        self.image = self.images[self.state][self.direction][self.current_frame]
        self.rect = self.image.get_rect()

        self.rect.x = SCALED_SCREEN_WIDTH / 2
        self.rect.y = SCALED_SCREEN_HEIGHT / 2

        self.hitbox = self.rect.inflate(SCALED_HITBOX_INFL_X, SCALED_HITBOX_INFL_Y)
        self.hitbox.x += SCALED_HITBOX_X_OFFSET
        self.hitbox.y += SCALED_HITBOX_Y_OFFSET

        self.current_level = current_level

    def update(self, directions: list[Direction]):
        self.move(directions)
        self.game_tick += 1
        if self.game_tick > 15:
            self.game_tick = 0
            self.current_frame = (self.current_frame + 1) % 4
            self.image = self.images[self.state][self.direction][self.current_frame]

    def draw(self, screen: pygame.Surface, offset: tuple[int, int]):
        screen.blit(self.image, (self.rect.x - offset[0], self.rect.y - offset[1]))
        if self.show_hitbox:
            pygame.draw.rect(screen, (255, 0, 0), [self.hitbox.x - offset[0], self.hitbox.y - offset[1], self.hitbox.width, self.hitbox.height], 1)

    def directions_to_dx_dy(self, directions: list[Direction]):
        dx = 0.0
        if Direction.LEFT in directions:
            dx -= self.speed
        if Direction.RIGHT in directions:
            dx += self.speed

        dy = 0.0
        if Direction.UP in directions:
            dy -= self.speed
        if Direction.DOWN in directions:
            dy += self.speed

        return dx, dy

    def change_direction(self, direction: Direction):
        self.direction = direction
        self.image = self.images[self.state][self.direction][self.current_frame]

    def move(self, directions: list[Direction]):
        dx, dy = self.directions_to_dx_dy(directions)
        if dx == 0.0 and dy == 0.0:
            self.state = PlayerState.IDLE
        else:
            self.state = PlayerState.WALKING

        if dy != 0:
            self.rect.y += dy
            self.hitbox.y += dy

            y_hit_list = pygame.sprite.spritecollide(
                self,
                self.current_level.get_colliders(),
                False,
                player_collided)
            if len(y_hit_list) > 0:
                self.rect.y -= dy
                self.hitbox.y -= dy
            else:
                self.change_direction(Direction.DOWN if dy > 0 else Direction.UP)

        if dx != 0:
            self.rect.x += dx
            self.hitbox.x += dx

            x_hit_list = pygame.sprite.spritecollide(
                self,
                self.current_level.get_colliders(),
                False,
                player_collided)

            if len(x_hit_list) > 0:
                self.rect.x -= dx
                self.hitbox.x -= dx
            else:
                self.change_direction(Direction.RIGHT if dx > 0 else Direction.LEFT)

def player_collided(player: Player, other_hitbox: pygame.Rect):
    return player.hitbox.colliderect(other_hitbox)
