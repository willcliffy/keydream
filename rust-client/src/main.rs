use bevy::{
    prelude::*,
    core::FixedTimestep,
};

mod player;
mod user_input;

pub const INTERNAL_TIMESTEP: f64 = 1.0 / 60.0;
pub const CAMERA_SCALE: f32 = 1.0 / 2.0;


fn startup(mut commands: Commands) {
    let mut camera = OrthographicCameraBundle::new_2d();

    camera.orthographic_projection.scale = CAMERA_SCALE;

    commands.spawn_bundle(camera);
    commands.spawn_bundle(UiCameraBundle::default());
}

fn main() {
    App::new()
        .insert_resource(ClearColor(Color::rgb(0.0, 0.0, 0.0)))
        .add_system(bevy::input::system::exit_on_esc_system.system())
        .add_plugins(DefaultPlugins)
        .add_startup_system(startup)
        .add_startup_system(player::startup.system())
        .add_system_set(
            SystemSet::new()
            .with_run_criteria(FixedTimestep::step(INTERNAL_TIMESTEP))
            .with_system(user_input::input.system())
            .with_system(player::animate.system()))
        .run();
}
