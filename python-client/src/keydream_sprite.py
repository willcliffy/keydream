import pygame

from src.objects import SpriteDirection, SpriteState


class KeydreamSprite(pygame.sprite.Sprite):
    state: SpriteState = SpriteState.IDLE
    direction: SpriteDirection = SpriteDirection.DOWN

    current_frame: int = 0
    frames_per_animation: int = 4

    # TODO - this is measured in game ticks, do better
    frame_delay: int = 15

    # TODO - this is just horrible
    game_tick_counter: int = 0

    images: dict = {}
    image: pygame.Surface = None

    rect: pygame.Rect = None
    hitbox: pygame.Rect = None
    show_hitbox: bool = False

    speed: float = 8.0

    def __init__(self,
        images: dict,
        image: pygame.Surface,
        rect: pygame.Rect,
        hitbox: pygame.Rect,
        speed: float = 8.0,
        show_hitbox: bool = False):
        pygame.sprite.Sprite.__init__(self)
        self.images = images
        self.image = image
        self.rect = rect
        self.hitbox = hitbox
        self.speed = speed
        self.show_hitbox = show_hitbox

    def draw(self, screen: pygame.Surface, offset: tuple[int, int]):
        screen.blit(self.image, (self.rect.x - offset[0], self.rect.y - offset[1]))
        if self.show_hitbox:
            pygame.draw.rect(
                screen,
                (255, 0, 0),
                [
                    self.hitbox.x - offset[0],
                    self.hitbox.y - offset[1],
                    self.hitbox.width,
                    self.hitbox.height
                ],
                1)

    def update(self, directions: list[SpriteDirection], colliders):
        self.move(directions, colliders)
        self.game_tick_counter += 1
        if self.game_tick_counter > self.frame_delay:
            self.game_tick_counter = 0
            self.current_frame = (self.current_frame + 1) % self.frames_per_animation
            self.image = self.images[self.state][self.direction][self.current_frame]
            if self.state == SpriteState.ROLL:
                self.roll_counter += 1
                self.state = SpriteState.WALK
                if self.roll_counter >= 4:
                    self.roll_counter = 0
            elif self.state == SpriteState.ATTACK:
                self.attack_counter += 1
                self.state = SpriteState.IDLE
                if self.attack_counter >= 4:
                    self.attack_counter = 0
        
    def move(self, directions: list[SpriteDirection], colliders):
        dx, dy = directions_to_dx_dy(directions, self.speed)
        if dx == 0.0 and dy == 0.0:
            self.state = SpriteState.IDLE
        elif self.state != SpriteState.ROLL and self.state != SpriteState.ATTACK:
            self.state = SpriteState.WALK

        if dy != 0:
            self.hitbox.y += dy
            y_hit_list = pygame.sprite.spritecollide(
                self,
                colliders,
                False,
                lambda sprite, collider: sprite.hitbox.colliderect(collider))

            if len(y_hit_list) > 0:
                self.hitbox.y -= dy
            else:
                self.rect.y += dy
                self.change_direction(SpriteDirection.DOWN if dy > 0 else SpriteDirection.UP)

        if dx != 0:
            self.hitbox.x += dx
            x_hit_list = pygame.sprite.spritecollide(
                self,
                colliders,
                False,
                lambda sprite, collider: sprite.hitbox.colliderect(collider))

            if len(x_hit_list) > 0:
                self.hitbox.x -= dx
            else:
                self.rect.x += dx
                self.change_direction(SpriteDirection.RIGHT if dx > 0 else SpriteDirection.LEFT)

    def change_direction(self, direction: SpriteDirection):
        self.direction = direction
        self.current_frame = 0
        self.image = self.images[self.state][self.direction][self.current_frame]

    def change_state(self, state: SpriteState):
        self.state = state
        self.current_frame = 0
        self.image = self.images[self.state][self.direction][self.current_frame]


def directions_to_dx_dy(directions: list[SpriteDirection], speed):
    dx = 0.0

    if SpriteDirection.LEFT in directions:
        dx -= speed
    if SpriteDirection.RIGHT in directions:
        dx += speed

    dy = 0.0
    if SpriteDirection.UP in directions:
        dy -= speed
    if SpriteDirection.DOWN in directions:
        dy += speed

    return dx, dy
