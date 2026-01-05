// Logic for Canopy Shield Validator Selection
// Priority: 60% Stake, 40% Reputation Score

pub fn select_active_set(mut candidates: Vec<Validator>) -> Vec<Validator> {
    // Filter by minimum reputation
    candidates.retain(|v| v.reputation_score >= 80); 

    // Sort by weighted score
    candidates.sort_by(|a, b| {
        let score_a = (a.stake_amount * 60) + (a.reputation_score as u128 * 40);
        let score_b = (b.stake_amount * 60) + (b.reputation_score as u128 * 40);
        score_b.cmp(&score_a)
    });

    candidates.truncate(100);
    candidates
}
