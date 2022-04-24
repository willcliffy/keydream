use bevy::prelude::*;

use crate::player::{
    Player,
    PlayerState,
    PlayerDirection,
    PLAYER_MOVEMENT_SPEED
};

pub fn input(
    keyboard_input: Res<Input<KeyCode>>,
    mut query: Query<(&mut Player, &mut Timer, &mut Transform, &mut TextureAtlasSprite, Without<Camera>)>,
    mut camera: Query<&mut Transform, With<Camera>>
) {
    let (mut player, mut timer, mut transform, mut sprite, _) = query.single_mut();

    if keyboard_input.just_pressed(KeyCode::Space) {
        player.attack(&mut sprite, &mut timer);
    }

    if player.state == PlayerState::Attack {
        return;
    }

    if !keyboard_input.any_pressed(vec![
        KeyCode::W, KeyCode::A, KeyCode::S, KeyCode::D,
        KeyCode::Up, KeyCode::Left, KeyCode::Down, KeyCode::Right,
    ]) {
        player.state = PlayerState::Idle;
        return;
    }

    let mut x = 0.0;
    let mut y = 0.0;

    let up = keyboard_input.any_pressed(vec![KeyCode::Up, KeyCode::W]);
    let down = keyboard_input.any_pressed(vec![KeyCode::Down, KeyCode::S]);
    let left = keyboard_input.any_pressed(vec![KeyCode::Left, KeyCode::A]);
    let right = keyboard_input.any_pressed(vec![KeyCode::Right, KeyCode::D]);

    if up && !down {
        player.direction = PlayerDirection::Up;
        y += PLAYER_MOVEMENT_SPEED;
    } else if down && !up {
        player.direction = PlayerDirection::Down;
        y -= PLAYER_MOVEMENT_SPEED;
    }

    if left && !right {
        player.direction = PlayerDirection::Left;
        x -= PLAYER_MOVEMENT_SPEED;
    } else if right && !left {
        player.direction = PlayerDirection::Right;
        x += PLAYER_MOVEMENT_SPEED;
    }

    if x == 0.0 && y == 0.0 {
        player.state = PlayerState::Idle;
        return;
    }

    if keyboard_input.just_pressed(KeyCode::LShift) {
        player.roll(&mut sprite, &mut timer);
    }

    if player.state != PlayerState::Roll {
        player.state = PlayerState::Walk;
    }

    let translation = &mut transform.translation;

    translation.x += x;
    translation.y += y;

    for mut camera_transform in camera.iter_mut() {
        camera_transform.translation.x = translation.x;
        camera_transform.translation.y = translation.y;
    }
}
