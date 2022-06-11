// << File: src/main.cpp  Line: 100  Function: main>> helloworld
// addition:
// DBG_PRINTF("Compile Time: %s  %s\n", __DATE__, __TIME__);
#define DEBUG 1

#if DEBUG
#define DBG_PRINTF(fmt, args...)                                     \
  do {                                                               \
    printf("<<File:%s  Line:%d  Function:%s>> ", __FILE__, __LINE__, \
           __FUNCTION__);                                            \
    printf(fmt, ##args);                                             \
  } while (0)
#else
#define DBG_PRINTF(fmt, args...)
#endif
