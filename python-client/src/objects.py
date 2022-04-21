import enum


class SpriteState(enum.Enum):
    IDLE = 0
    WALK = 1
    ROLL = 2
    ATTACK = 3


class SpriteDirection(enum.Enum):
    UP = 0
    DOWN = 1
    LEFT = 2
    RIGHT = 3


class SpriteAction(enum.Enum):
    NONE = 0
    MOVE = 1
    ATTACK = 2
    ROLL = 3
    TALK = 4
