use bevy::{
    prelude::*,
    core::FixedTimestep,
};

use bevy_ecs_ldtk::prelude::*;

mod player;
mod user_input;

pub const INTERNAL_TIMESTEP: f64 = 1.0 / 60.0;
pub const CAMERA_SCALE: f32 = 1.0 / 3.0;


fn startup(mut commands: Commands, asset_server: Res<AssetServer>) {
    let mut camera = OrthographicCameraBundle::new_2d();

    camera.orthographic_projection.scale = CAMERA_SCALE;

    commands.spawn_bundle(camera);
    commands.spawn_bundle(UiCameraBundle::default());

    commands.spawn_bundle(LdtkWorldBundle {
        ldtk_handle: asset_server.load("keydream.ldtk"),
        ..Default::default()
    });
}

#[derive(Bundle, LdtkEntity)]
pub struct MyBundle {
    #[sprite_sheet_bundle]
    #[bundle]
    sprite_bundle: SpriteSheetBundle,
}

fn main() {
    App::new()
        .insert_resource(ClearColor(Color::rgb(0.0, 0.0, 0.0)))
        .add_plugins(DefaultPlugins)
        .add_plugin(LdtkPlugin)
        .insert_resource(LevelSelection::Index(0))
        .register_ldtk_entity::<MyBundle>("MyEntityIdentifier")
        .add_startup_system(startup)
        .add_startup_system(player::player_startup)
        .add_system_set(
            SystemSet::new()
            .with_run_criteria(FixedTimestep::step(INTERNAL_TIMESTEP))
            .with_system(user_input::input_system)
            .with_system(player::player_animate_system))
        .run();
}
