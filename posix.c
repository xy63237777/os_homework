#include <stdio.h>
#include <sys/types.h>
#include <unistd.h>
#include <signal.h>
#include <ctype.h>

#define MAX_CHILD_NUMBER 10
#define SLEEP_INTERVAL 2
int child_proc_number = 10;  
int proc_number=0;
void do_something();
pid_t pid[MAX_CHILD_NUMBER] = {0};
int book[MAX_CHILD_NUMBER] = {0};
void kill_flag_cmp(int iid, int index);

void kill_flag_cmp(int iid, int index) {
    if(book[index] == 0) {
        kill(pid[index], iid);
        book[index] = 1;
        printf("中断进程号%d\n",pid[index]);
    } else {
        printf("子进程%d已经中断，请勿再次中断\n", pid[index]);
        printf("--------------------------------------\n");
        printf("|           还 未 中 断 进 程         |\n");
        printf("--------------------------------------\n");
        for(int i = 0; i < child_proc_number; i++) {
            if(book[i] == 0)
                 printf("  进程下标%d, 进程号%d <-->\n",i,pid[i]);
        }
    }
}

int main(int argc, char* argv[]) {
    int i, ch;
    pid_t child_pid;
    if(argc > 1) {
        child_proc_number = atoi(argv[1]);
        child_proc_number = (child_proc_number > MAX_CHILD_NUMBER)?MAX_CHILD_NUMBER : child_proc_number;
    }
    for (int i = 0; i < child_proc_number; i++) {
        pid[i] = fork();
        proc_number = i;
        //printf("hello - %d\n",pid[i]);
        if(pid[i] == 0){
            do_something();    
        } 
    }
    while((ch = getchar()) != 'q') {
        if(isdigit(ch)) {
            kill_flag_cmp(SIGTERM,ch-'0');
        }
    }
 //   for(int i = 0; i < child_proc_number; i++) {
 //       kill_flag_cmp(SIGTERM,i);
 //   }
    kill(0,SIGTERM);
    return 0;
}

void do_something() {
    for(;;) {
        printf("This is process No.%d and its pid is %d \n", proc_number, getpid());
        sleep(2);
    }
}

