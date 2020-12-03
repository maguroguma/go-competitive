// https://atcoder.jp/contests/joisc2014/tasks/joisc2014_d

#include<bits/stdc++.h>
#include"ramen.h"
using namespace std;

// quoted from: https://lattemalta.hatenablog.jp/entry/2015/09/30/182630

void Ramen(int N){
    if(N==1){
        Answer(0,0);
        return;
    }

    vector<int>win,lose;
    for(int i=0;i<N/2;i++){
        if(Compare(i*2,i*2+1)==1){
            win.push_back(i*2);
            lose.push_back(i*2+1);
        }
        else{
            win.push_back(i*2+1);
            lose.push_back(i*2);
        }
    }
    if(N&1){
        win.push_back(N-1);
        lose.push_back(N-1);
    }

    int ma=0,mi=0;
    for(int i=1;i<win.size();i++){
        assert(win[i]!=win[ma]);
        if(Compare(win[i],win[ma])==1)ma=i;
    }

    for(int i=1;i<lose.size();i++){
        assert(lose[mi]!=lose[i]);
        if(Compare(lose[mi],lose[i])==1)mi=i;
    }

    Answer(lose[mi],win[ma]);
}

