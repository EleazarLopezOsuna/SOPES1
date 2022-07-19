use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Serialize, Deserialize)]
pub struct Game {
    pub game_id: String,
    pub players: String,
    pub game_name: String,
    pub winner: String,
    pub queue: String,
}