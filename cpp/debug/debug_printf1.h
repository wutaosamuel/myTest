// << File: src/main.cpp  Line: 100  Function: main>> helloworld
// addition:
// DBG_PRINTF("Compile Time: %s  %s\n", __DATE__, __TIME__);
#define DEBUG 1

#define DBG_ERROR 0x00   /* error */
#define DBG_NONE 0x01    /*  do nothing*/
#define DBG_WARNING 0x02 /* warning */

#if DEBUG
#define DBG_PRINTF(flag, fmt, args...)                               \
  do {                                                               \
    printf("<<File:%s  Line:%d  Function:%s>> ", __FILE__, __LINE__, \
           __FUNCTION__);                                            \
    printf(fmt, ##args);                                             \
  } while (0)
#else
#define DBG_PRINTF(flag, fmt, args...)                   \
  do {                                                    \
    switch (flag) {                                       \
      case DBG_ERROR:                                   \
        printf("<<Error   %s:%d>> ", __FILE__, __LINE__);   \
        printf(fmt, ##args);                              \
        break;                                            \
      case DBG_WARNING:                                 \
        printf("<<Warning %s:%d>> ", __FILE__, __LINE__); \
        printf(fmt, ##args);                              \
    }                                                     \
  } while (0)
#endif