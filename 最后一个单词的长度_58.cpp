class Solution {
public:
    int lengthOfLastWord(string s) {
        int num = 0;
        for (int i = s.size() - 1; i >= 0; i--) {
            if (s[i] != ' ')
                num++;
            if (num != 0 && s[i] == ' ')
                break;
        }
        return num;
    }
};