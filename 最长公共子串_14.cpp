class Solution {
public:
    string longestCommonPrefix(vector<string>& strs) {
        if(strs.size() == 0) {
            return "";
        }
        string res = strs[0];
        for(string& str: strs) {
            while(str.size() < res.size() || str.substr(0, res.size()) != res)        {
                res = res.substr(0, res.size() - 1);
            }
        }
        return res; // could be ""
    }
};