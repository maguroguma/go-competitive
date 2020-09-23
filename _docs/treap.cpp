template<class T> class treap {
public:
    struct node {
        T val, sum;
        node *left, *right;
        int pri;
        unsigned sz;
        node(T val, int pri):val(val),sum(val),pri(pri),sz(1) {
            left = right = NULL;
        }
    };

    node *root;
    treap() : root(NULL) {
        srand(time(NULL));
    }
        
    unsigned size() { return size(root); }
    unsigned size(node *v) { return !v? 0: v->sz; }
    T sum(node *v) { return !v? 0: v->sum; }

    node *update(node *v) {
        v->sz = size(v->left)+size(v->right)+1;
        v->sum = sum(v->left)+sum(v->right)+v->val;
        return v;
    }

    node *merge(node *s, node *t) {
        if(!s or !t) return s? s: t;
        if(s->pri > t->pri) {
            s->right = merge(s->right, t);
            return update(s);
        }
        t->left = merge(s, t->left);
        return update(t);
    }

    pair<node*, node*> split(node *v, unsigned k) {
        if(!v) return pair<node*,node*>(NULL,NULL);
        if(k <= size(v->left)) {
            pair<node*,node*> s = split(v->left, k);
            v->left = s.second;
            return make_pair(s.first,update(v));
        }
        pair<node*, node*> s = split(v->right, k-size(v->left)-1);
        v->right = s.first;
        return make_pair(update(v),s.second);
    }

    node *find(unsigned k) {
        node *v = root;
        while(v) {
            unsigned s = size(v->left);
            if(s > k) v = v->left;
            else if(s == k) return v;
            else {
                v = v->right;
                k -= s+1;
            }
        }
        return v;
    }

    void insert(unsigned k, T val) { root = insert(root,k,val,rand()); }
    node *insert(node *t, unsigned k, T val, int pri) {
        pair<node*, node*> s = split(t,k);
        t = merge(s.first, new node(val,pri));
        t = merge(t, s.second);
        return update(t);
    }

    void erase(int k) { root = erase(root,k); }
    node *erase(node *t, unsigned k) {
        pair<node*, node*> u, v;
        u = split(t,k+1);
        v = split(u.first, k);
        t = merge(v.first, u.second);
        return update(t);
    }
};
