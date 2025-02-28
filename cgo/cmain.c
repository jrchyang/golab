#include <stdio.h>
#include "libsdk.h"

int main() {
	printf("C calling Go ...\n");
	hello_from_go("Gopher");
	return 0;
}
