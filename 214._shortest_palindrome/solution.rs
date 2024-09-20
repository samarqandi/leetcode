impl Solution {
    pub fn shortest_palindrome(s: String) -> String {
        let rev_s = s.chars().rev().collect::<String>();

        for i in 0..rev_s.len() {
            if s.starts_with(&rev_s[i..]) {
                return format!("{}{}", &rev_s[..i], s)
            }
        }
        format!("{}{}", &rev_s, s)
    }
}
