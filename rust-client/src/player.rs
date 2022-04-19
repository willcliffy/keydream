use std::time::Duration;

use bevy::prelude::*;


pub const PLAYER_MOVEMENT_SPEED: f32 = 2.5;
pub const PLAYER_ANIMATION_NUM_FRAMES: usize = 4;

const PLAYER_SPRITE_ATLAS_PATH: &str = "sprites/player_atlas.png";
const PLAYER_SPRITE_ATLAS_WIDTH: usize = 16;
const PLAYER_SPRITE_ATLAS_HEIGHT: usize = 4;

const PLAYER_SPRITE_DIMENSION: f32 = 32.0;
const PLAYER_SPRITE_SCALE: f32 = 1.0;

const PLAYER_ANIMATION_FRAME_DURATION: f32 = 0.15;
const PLAYER_ROLL_FRAME_DURATION: f32 = 0.15;
const PLAYER_ATTACK_FRAME_DURATION: f32 = 0.15;

const PLAYER_ANIMATION_DIRECTION_OFFSET: usize = PLAYER_ANIMATION_NUM_FRAMES;
const PLAYER_ANIMATION_STATE_OFFSET: usize = 4 * PLAYER_ANIMATION_NUM_FRAMES; // 4 player directions


#[repr(usize)]
#[derive(Copy, Clone, Debug, PartialEq)]
pub enum PlayerState {
    Idle = 0,
    Walk = 1,
    Attack = 2,
    Roll = 3,
}

#[repr(usize)]
#[derive(Copy, Clone, Debug)]
pub enum PlayerDirection {
    Down  = 0,
    Right = 1,
    Left  = 2,
    Up    = 3,
}

#[derive(Component)]
pub struct Player {
    pub direction: PlayerDirection,
    pub state: PlayerState,
    pub attacking_counter: usize,
    pub rolling_counter: usize,
}

impl Player {
    fn new() -> Self {
        Player {
            direction: PlayerDirection::Down,
            state: PlayerState::Idle,
            attacking_counter: 0,
            rolling_counter: 0,
        }
    }

    pub fn roll(&mut self, sprite: &mut TextureAtlasSprite, timer: &mut Timer) {
        self.state = PlayerState::Roll;
        self.rolling_counter = 0;
        sprite.index = next_sprite_index(&self, self.rolling_counter);
        timer.set_duration(Duration::from_secs_f32(PLAYER_ROLL_FRAME_DURATION));
    }

    pub fn attack(&mut self, sprite: &mut TextureAtlasSprite, timer: &mut Timer) {
        self.state = PlayerState::Attack;
        self.attacking_counter = 0;
        sprite.index = next_sprite_index(&self, self.attacking_counter);
        timer.set_duration(Duration::from_secs_f32(PLAYER_ATTACK_FRAME_DURATION));
    }
}

pub fn player_startup(
    mut commands: Commands,
    asset_server: Res<AssetServer>,
    mut texture_atlases: ResMut<Assets<TextureAtlas>>,
) {
    let texture_handle = asset_server.load(PLAYER_SPRITE_ATLAS_PATH);
    let texture_atlas = TextureAtlas::from_grid(
        texture_handle,
        Vec2::new(PLAYER_SPRITE_DIMENSION, PLAYER_SPRITE_DIMENSION),
        PLAYER_SPRITE_ATLAS_WIDTH,
        PLAYER_SPRITE_ATLAS_HEIGHT);
    let texture_atlas_handle = texture_atlases.add(texture_atlas);

    commands
        .spawn_bundle(SpriteSheetBundle {
            texture_atlas: texture_atlas_handle,
            transform: Transform::from_scale(Vec3::splat(PLAYER_SPRITE_SCALE))
                .with_translation(Vec3::new(0.0, 0.0, 5.0)),
            ..Default::default()
        })
        .insert(Timer::from_seconds(PLAYER_ANIMATION_FRAME_DURATION, true))
        .insert(Player::new());
}

pub fn player_animate_system(
    time: Res<Time>,
    mut query: Query<(&mut Player, &mut Timer, &mut TextureAtlasSprite)>,
) {
    for (mut player, mut timer, mut sprite) in query.iter_mut() {
        timer.tick(time.delta());
        if timer.finished() {
            match player.state {
                PlayerState::Idle => {
                    sprite.index = next_sprite_index(&player, sprite.index);
                },
                PlayerState::Walk => {
                    sprite.index = next_sprite_index(&player, sprite.index);
                },
                PlayerState::Attack => {
                    if player.attacking_counter < PLAYER_ANIMATION_NUM_FRAMES {
                        sprite.index = next_sprite_index(&player, player.attacking_counter);
                        player.attacking_counter += 1;
                        continue;
                    }

                    player.attacking_counter = 0;
                    player.state = PlayerState::Idle;
                    sprite.index = next_sprite_index(&player, sprite.index);
                },
                PlayerState::Roll => {
                    if player.rolling_counter < PLAYER_ANIMATION_NUM_FRAMES {
                        sprite.index = next_sprite_index(&player, player.rolling_counter);
                        player.rolling_counter += 1;
                        continue;
                    }

                    player.rolling_counter = 0;
                    player.state = PlayerState::Idle;
                    sprite.index = next_sprite_index(&player, sprite.index);
                },
            }
        }
    }
}

pub fn next_sprite_index(player: &Player, sprite_index: usize) -> usize {
    return (sprite_index + 1) % PLAYER_ANIMATION_NUM_FRAMES
        + player.direction as usize * PLAYER_ANIMATION_DIRECTION_OFFSET
        + player.state as usize * PLAYER_ANIMATION_STATE_OFFSET;
}
