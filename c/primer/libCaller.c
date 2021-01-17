#include <stdio.h>
#include "library/calc.h"

void main() {
    int a = 5;
    int b = 5;

    int c = add(a, b);
    printf("result: %d\n", c);
}