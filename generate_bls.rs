
use blst::{SecretKey, PublicKey};

fn main() {
    // Generate a random secret key
    let sk = SecretKey::key_gen(&mut rand::thread_rng(), &[]).expect("Failed to generate key");
    
    // Generate the corresponding public key
    let pk = sk.sk_to_pk();

    println!("BLS Public Key: {:?}", pk.to_bytes());
    println!("BLS Private Key: {:?}", sk.to_bytes());
}
