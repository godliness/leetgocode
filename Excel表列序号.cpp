class Solution {
public:
    int titleToNumber(string columnTitle) {
       int n = columnTitle.size();
       int ans = 0;
       for (int i = 0; i < n; i++) { // 从先往后
        ans = ans*26+(columnTitle.at(i)-'A'+1);
       }
       return ans;
    }
};