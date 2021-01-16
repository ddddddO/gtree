#include <stdio.h>

// プロトタイプ宣言
// mainの後ろで関数を定義するとコンパイルするときにワーニングが出るため、
// どんな関数があるか？をmainの上に記載する、こと。
int data(void);
void conditions(void);
void wh(void);
void loop(void);
void st_val(void);
void pointer(void);
void swap(int *x, int *y);
void standard_input(void);
void person_struct(void);

// グローバル変数
int xx = 0;

void xxx(void) {
    xx++;
    printf("xxx func: xx = %d\n", xx);
}

int main(void) {
    data();
    conditions();
    wh();
    loop();

    xxx();
    xx++;
    printf("main func: xx = %d\n", xx);

    st_val();
    st_val();
    st_val();

    pointer();

    int num0 = 111;
    int num1 = 999;
    printf("before swap: num0 = %d, num1 = %d\n", num0, num1);
    swap(&num0, &num1);
    printf("after swap: num0 = %d, num1 = %d\n", num0, num1);

    standard_input();

    person_struct();

    return 0;
}

int data(void) {
    int i;
    float f;
    char c;
    printf("i = %d, f = %f, c = %c\n", i, f, c);

    i = 555;
    f = 3.344;
    c = 'C';
    printf("i = %d, f = %f, c = %c\n", i, f, c);

    int nums0[2];
    nums0[0] = 999;
    nums0[1] = 888;

    int nums1[] = {1, 2, 5};
    printf("array num1[2] = %d\n", nums1[2]);

    // C言語に文字列型は無い。charの配列で表わす(末尾の要素は「\0」)。
    char s[] = {'b', 'o', 'o', 'm', '\0'};
    // 上の書き方と、以下は同じ
    // char s[] = "boom";
    printf("%s\n", s);
}

void conditions(void) {
    int age = 28;
    if (age < 20) {
        printf("child\n");
    } else if (age >= 20) {
        printf("adult\n");
    } else {
        printf("dead code\n");
    }

    int num = 111;
    switch(num) {
        case 111:
            printf("num: 111\n");
            break;
        case 333:
            printf("num: 333\n");
            break;
        default:
            printf("default\n");
            break;
    }

    // 三項演算子
    int x = 4;
    (x > 5) ? printf("x > 5\n") : printf("x <= 5\n");
}

void wh(void) {
    int n = 0;
    // 条件判定後、処理
    while (n < 3) {
        printf("while n = %d\n", n);
        n++;
    }

    // 一回実行されてから条件判定
    do {
        printf("do while n = %d\n", n);
        n++;
    } while (n < 6);
}

void loop(void) {
    for (int n =0; n < 3; n++) {
        printf("for n = %d\n", n);
    }
    // continue
    // break
}

void st_val(void) {
    int non_static_val = 0;
    non_static_val++;
    printf("non_static_val: %d, ", non_static_val);

    // static ・・・プログラム終了時まで値を保持。静的変数
    static int val = 0;
    val++;
    printf("static_val: %d\n", val);
}

void pointer(void) {
    int x = 111;
    int *pointer_x;

    pointer_x = &x;
    printf("pointer_x_val: %d\n", *pointer_x);
}

// 参照渡し(引数にアドレスを渡す)。
// 返り値を一つしか渡せないけど、２値の値を交換する(swap)関数を作るには？以下関数
void swap(int *x, int *y) {
    int tmp = *x;
    *x = *y;
    *y = tmp;
}

// マクロ
#define N 10

void standard_input(void) {
    char buf[20];
    fgets(buf, sizeof(buf), stdin);

    char out[N];
    sscanf(buf, "%s", out);

    printf("std_out: %s\n", out);
}

// 構造体。「person」にあたる部分は「構造体タグ」という。
struct person {
    char name[10];  // 各々「メンバ」という。
    int age;
    char message[50];
};

// データ型に名前を付けることができる。
typedef struct person psn;

// typedefは構造体宣言と一緒に書ける。以下。
// typedef struct person {
//     char name[10];
//     int age;
//     char message[50];
// } psn;

// さらに、構造体タグも削除できる。以下。
// typedef struct {
//     char name[10];
//     int age;
//     char message[50];
// } psn;

#include <string.h>

// プロトタイプ宣言
void birthday(psn *p);

void person_struct(void) {
    struct person d = {
        "ddd",
        28,
        "Hello!"
    };

    printf("struct person d.name: %s\n", d.name);
    // d.message = "Bye"; これはダメ
    strcpy(d.message, "Bye");
    printf("struct person d.message: %s\n", d.message);
    d.age = 20;  // これはOK
    printf("struct person d.age: %d\n", d.age);

    birthday(&d);
    printf("struct person after birthday d.age: %d\n", d.age);

    psn x = {
        "XXX",
        99,
        "ffff"
    };

    printf("struct person x.name: %s\n", x.name);
}

// 構造体のポインタについて
void birthday(psn *p) {
    // アロー演算子でアクセス
    p->age+=1;
    // (*p).age+=1;  // こちらの書き方でも可
}