// Logic for Canopy Shield Reward Distribution
// 20% Security, 10% Dev Incentive, 70% Validators

pub fn distribute_rewards(total_reward: u128) {
    let to_shield = total_reward * 20 / 100; 
    let to_devs = total_reward * 10 / 100;
    let to_validators = total_reward - to_shield - to_devs;

    println!("Allocating {} to Canopy Shield Security", to_shield);
    println!("Allocating {} to Appchain Developers", to_devs);
    println!("Allocating {} to Active Validators", to_validators);
}
