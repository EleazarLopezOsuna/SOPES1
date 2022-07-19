//! Example code for using MongoDB with Actix.

mod model;
#[cfg(test)]
mod test;

use actix_cors::Cors;
use actix_web::{http, get, web, App, HttpResponse, HttpServer};
use mongodb::{Client};
use model::Game;
use futures::StreamExt;
use std::env;
use dotenv::dotenv;

const DB_NAME: &str = "sopes";
const COLL_NAME: &str = "fdos";




#[get("/get-all")]
async fn get_all_logs(client: web::Data<Client>) -> HttpResponse {
    let collection = client.database(DB_NAME).collection(COLL_NAME);
    
    let mut cursor = collection
        .find(None, None).await.ok().expect("Failed to execute find");
        
    let mut docs: Vec<Game> = Vec::new();
    
    while let Some(doc) = cursor.next().await{
        docs.push(doc.unwrap())
    }
    let result = web::block(move || docs).await;
    

    match result {
        Ok(result) => HttpResponse::Ok().json(result),
        Err(err) => HttpResponse::InternalServerError().body(err.to_string()),
    }
}



#[actix_web::main]
async fn main() -> std::io::Result<()> {

    dotenv().ok();
    env::set_var("RUST_LOG", "actix_web=debug,actix_server=info");
    env_logger::init();

    let uri = std::env::var("MONGODB_URI").unwrap_or_else(|_| "mongodb://localhost:27017".into());

    let client = Client::with_uri_str(uri).await.expect("failed to connect");

    /*let options = ClientOptions::parse_with_resolver_config(&uri, ResolverConfig::cloudflare())
        .await?;
    let client = Client::with_options(options)?;*/
    

    HttpServer::new(move || {
        App::new()
            .app_data(web::Data::new(client.clone()))
            
            .service(get_all_logs)
    })
    .bind(("0.0.0.0", 8080))?
    .run()
    .await
}