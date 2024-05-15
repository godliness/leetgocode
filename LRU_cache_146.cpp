class LRUCache {
    unordered_map<int, list<pair<int, int>>::iterator> mp;
    list<pair<int, int>> li;
    const int capacity;

public:
    LRUCache(int capacity) : capacity(capacity) {}

    int get(int key) {
        if (!mp.count(key))
            return -1;
        li.push_front(*mp[key]);
        li.erase(mp[key]);
        mp[key] = li.begin();
        return li.begin()->second;
    }

    void put(int key, int value) {
        if (mp.count(key)) {
            li.erase(mp[key]);
        } else if (mp.size() >= capacity) {
            mp.erase(li.back().first);
            li.pop_back();
        }

        li.push_front({key, value});
        mp[key] = li.begin();
    }
};