class Solution {
public:
    bool isPalindrome(int x) {
        string s = to_string(x);
        return s == string(s.rbegin(),s.rend());
    }
};


class Solution {
public:
    bool judge(string s,int left, int right) {
        while (left >= 0 && right < s.size() && s[left]==s[right]) {
            left--;
            right++;
        }
        if (left < 0) {
            return true;
        }
        return false;
    }
    bool isPalindrome(int x) {
        string str = to_string(x);
        int len = str.size();
        int mid = len / 2;
        if (len % 2 == 0) {
            return judge(str, mid-1, mid);
        }
        else {
            return judge(str, mid, mid);
        }
    }
}
