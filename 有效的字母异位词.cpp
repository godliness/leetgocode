class Solution {
public:
    bool isAnagram(string s, string t) {
        if(s.size()!=t.size()) return false;
        unordered_map<char , int > mymap_s;
        unordered_map<char , int > mymap_t;
        for(auto e : s)
        {
            mymap_s[e]++;
        }
        for(auto a : t)
        {
            mymap_t[a]++;
        }
        for(auto b : mymap_s)
        {
            if(b.second!=mymap_t[b.first])
            return false;
        }
        return true;
    }
};