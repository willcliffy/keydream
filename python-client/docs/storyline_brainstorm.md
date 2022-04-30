# Keydream Storyline

## Plot structure

- Keydream will follow an established narrative structure. After some consideration, I've decided on Harmon's Plot
- Here is a basic flowchart of the Harmon's Plot:

``` mermaid
flowchart LR
    1[Status Quo]
    2[Goal Established]
    3[The Journey]
    4[The Search]
    5[The Fight]
    6[Attain Goal...]
    7[... at a cost]
    8[The Return]
    1 --> 2
    2 --> 3
    3 --> 4
    4 --> 5
    5 --> 6
    6 --> 7
    7 --> 8
    8 --> 1
```

## First pass notes

### 1. STATUS QUO

| character is in zone of comfort

- Imagine something like a simple Stardew Valley or Animal Crossing
- Very simple inventory/resource management
- Introduce the mundane world and some NPCs

### 2. GOAL ESTABLISHED

| character has a desire

- Understand their past?

### 3. JOURNEY BEGINS

| character journeys towards desire

- Enters the cave where they were found?

### 4. CHARACTER ADAPTS

| character fights to achieve their goal

- Finds equipment
- Learns skills
- Defeat enemies
- Solve puzzles

- This will be the core part of gameplay

### 5. GOAL ATTAINED

| character achieves the desire

- Finds out the real backstory - TBD
- Would be interesting if player character has convoluted backstory, or was previously a villain.

### 6. ... AT A PRICE

| character pays the price for their achievements

- NPC death
- (Temporary) loss of equipment or abilities due to "injury"
- Community from starting area ostracizes player character

### 7. JOURNEY ENDS

| character returns to the start, changed by the journey

### 8. CLOSURE

| TBD

## Map design

The game map will need to be split into multiple levels.

The level for plot sections 1 and 2 are the same, we will call this Level 1a.
The level for plot section 3 is a corridor or path to Level 3, we will call this level 1b.

The level for plot section 4 will be the largest section, and will likely be split into 3 subsections.
We will call these subsections Level 2a, Level 2b, and Level 2c.

The level for plot sections 5 and 6 will be a boss chamber of some sort, we will call this Level 3.

The level for plot section 7 is a modified version of Level 1b, we will call this Level 1bM
The level for plot section 8 is modified version of Level 1a, we will call this Level 1aM.

So the intended path through the game is:

- Level 1a
- Level 1b
- Level 2 (portal room/fork in the road)
- Level 2a, 2b, and 2c, in any order
- Level 3 (boss)
- Level 2M
- Level 1bM
- Level 1aM

In order to not overscope, the modified versions of levels 1 and 2 will be restricted to use the same tmx file.

So there will be 7 pyTMX files:

- level_1a.tmx - Plot sections 1, 2, 8
- level_1b.tmx - Plot sections 3, 7
- level_2.tmx  - Plot section 4
- level_2a.tmx - Plot section 4
- level_2b.tmx - Plot section 4
- level_2c.tmx - Plot section 4
- level_3.tmx  - Plot section 5, 6

## Gameplay considerations

Rough time outline:

- 20 mins for intro
  - movement, npc interaction, picking up items and tools.
  - introduce characters
- 5 mins transition
- 15 mins dungeon I
- 15 mins dungeon II
- 15 mins dungeon III
- 20 mins conlusion
