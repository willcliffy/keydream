import pygame
from src.npc import NPC


class NPCController():
    # TODO this should be a spritegroup for performance reasons
    npcs: list[NPC] = []
    hitboxes: list[pygame.Rect] = []

    def __init__(self, npcs):
        self.npcs = npcs
        for npc in self.npcs:
            self.hitboxes.append(npc.hitbox)

    def add(self, npc):
        self.npcs.append(npc)
    
    def remove(self, npc):
        self.npcs.remove(npc)

    def update(self):
        for npc in self.npcs:
            npc.update([], [])

    def draw(self, display, offset):
        for npc in self.npcs:
            npc.draw(display, offset)

    def get_hitboxes(self):
        return self.hitboxes
