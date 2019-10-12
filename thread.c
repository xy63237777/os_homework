 #include <stdio.h>
#include <sys/types.h>
#include <unistd.h>
#include <ctype.h>
#include <pthread.h>

#define MAX_THREAD 3

unsigned long long main_counter,counter[MAX_THREAD];

void* thread_worker(void*);

int main(int argc, char* argv[]) {
    int ch, i, rth;
    pthread_t pthread_id[MAX_THREAD] = {0};
    for (int i = 0; i < MAX_THREAD; i++) {
          rth = pthread_create(&pthread_id[i],0,thread_worker,&i);
    }
    do {
		unsigned long long sum = 0;
		for (i = 0; i < MAX_THREAD; i++) {
			sum += counter[i];
			printf("%llu ", counter[i]);
		}
	 	printf("%llu/%llu", main_counter, sum);
	} while ((ch = getchar()) != 'q');

	return 0;
}
   
void* thread_worker(void* p) {
    printf("worker invoke...\n");

    int thread_num;
	thread_num = *((int *)p);
    for(;;) {
          counter[thread_num]++;
          main_counter++;
     }
}            
